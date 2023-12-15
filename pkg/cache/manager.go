package cache

import (
	"context"
	"errors"
	"fmt"
	"libs/pkg/cron"
	"reflect"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
)

type DBFetcher[U int64 | string, T any] interface {
	FindOne(ctx context.Context, id U) (*T, error)
}

type Manager[U int64 | string, T any] struct {
	fetcher        DBFetcher[U, T]
	values         sync.Map
	updateInterval time.Duration
	singleflightG  singleflight.Group
}

// NewManager 创建一个新的 Manager 实例
// U: 数据类型，可以是 int64 或 string
// T: 数据结构类型
// fetcher: DBFetcher 接口的实现，用于获取数据
// updateInterval: 更新数据的间隔时间,0时从数据库获取一次后不更新
func NewManager[U int64 | string, T any](fetcher DBFetcher[U, T], updateInterval time.Duration) *Manager[U, T] {
	m := &Manager[U, T]{
		fetcher:        fetcher,
		values:         sync.Map{}, // 使用 sync.Map 存储数据
		updateInterval: updateInterval,
	}
	if updateInterval > 0 {
		duration := time.Duration(updateInterval) * time.Second
		pattern := fmt.Sprintf("@every %s", duration.String())
		_, err := cron.AddJob(pattern, func() {
			m.UpdateAll()
		}, cron.SkipIfStillRunning)
		if err != nil {
			logx.Errorf("NewManager init cron error:%v %v", err, reflect.TypeOf(fetcher))
			panic(err)
		}
	}
	return m
}

func (m *Manager[U, T]) Get(ctx context.Context, id U) (*T, error) {
	if value, ok := m.values.Load(id); ok {
		return value.(*T), nil
	}

	key := fmt.Sprintf("%v", id)
	info, err, _ := m.singleflightG.Do(key, func() (interface{}, error) {
		data, err := m.fetchItem(id)
		if err != nil {
			return nil, err
		}
		m.values.Store(id, data)
		return data, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*T), nil
}

func (m *Manager[U, T]) UpdateAll() {
	m.values.Range(func(key, value interface{}) bool {
		id := key.(U)
		data, err := m.fetchItem(id)
		if err != nil {
			logx.Errorf("cache manager update fetchItem error:%v id:%v", err, id)
			return true
		}
		if data != nil {
			m.values.Store(id, data)
		}
		return true
	})
}

func (m *Manager[U, T]) fetchItem(id U) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	info, err := m.fetcher.FindOne(ctx, id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, sqlx.ErrNotFound
		}
		return nil, err
	}
	return info, nil
}
