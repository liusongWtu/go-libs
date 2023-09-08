package clickhousex

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// BaseModel 注意T必须为指针类型
type BaseModel[T any] struct {
	Client driver.Conn
	Table  string
}

// BatchInsert 注意添加字段时，先发布代码，再往数据库添加字段。不然先加字段会出现插不进去
func (m *BaseModel[T]) BatchInsert(ctx context.Context, items []T) error {
	batch, err := m.Client.PrepareBatch(ctx, "INSERT INTO "+m.Table)
	if err != nil {
		return err
	}
	for i := range items {
		err := batch.AppendStruct(items[i])
		if err != nil {
			return err
		}
	}
	err = batch.Send()
	return err
}
