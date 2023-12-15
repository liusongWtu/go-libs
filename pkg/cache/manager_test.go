package cache

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type UserMock struct {
	Id         int64
	UpdateTime int64
}

type UserMockModel struct {
}

func (m UserMockModel) FindOne(ctx context.Context, id int64) (*UserMock, error) {
	return &UserMock{Id: id, UpdateTime: time.Now().Unix()}, nil
}

func TestNewManager(t *testing.T) {
	ctx := context.Background()
	feacher := UserMockModel{}

	m := NewManager[int64, UserMock](feacher, 10)
	// m := NewManager[int64, UserMock](feacher, 0)
	info, err := m.Get(ctx, 1)
	assert.NoError(t, err)

	t.Log(info)

	time.Sleep(time.Second * 12)

	info, err = m.Get(ctx, 1)
	assert.NoError(t, err)

	t.Log(info)
}
