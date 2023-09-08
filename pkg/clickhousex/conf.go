package clickhousex

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Conf struct {
	Addr     []string
	Db       string
	User     string
	Password string
	Debug    bool `json:",optional"`
}

func NewConn(conf Conf) (driver.Conn, error) {
	return clickhouse.Open(&clickhouse.Options{
		Addr: conf.Addr,
		Auth: clickhouse.Auth{
			Database: conf.Db,
			Username: conf.User,
			Password: conf.Password,
		},
		Debug: conf.Debug,
		Settings: clickhouse.Settings{
			"max_execution_time": 120,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:      time.Duration(10) * time.Second,
		ConnMaxLifetime:  time.Duration(10) * time.Minute,
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		MaxOpenConns:     20,
		MaxIdleConns:     5,
	})
}
