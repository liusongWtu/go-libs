package transactionex

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Model struct {
	sqlc.CachedConn
}

type TableTransactionFunc func(session sqlx.Session) error

func NewModel(conn sqlx.SqlConn, c cache.CacheConf) *Model {
	return &Model{
		CachedConn: sqlc.NewConn(conn, c),
	}
}

// Transaction 事务方式插入多个表的数据
func (m *Model) Transactions(transactions []TableTransactionFunc) error {
	err := m.CachedConn.Transact(func(session sqlx.Session) error {
		for _, transaction := range transactions {
			err := transaction(session)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
