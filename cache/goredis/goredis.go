package goredis

import (
	"errors"
	"github.com/go-redis/redis"
	"strconv"
)

var (
	// DefaultKey the collection name of redis for cache adapter.
	DefaultKey = ""
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(conf map[string]string) *RedisCache {
	client := redis.NewClient(getOptions(conf))
	return &RedisCache{client: client}
}

func getOptions(conf map[string]string) *redis.Options {
	opts := &redis.Options{}
	if _, ok := conf["key"]; !ok {
		conf["key"] = DefaultKey
	}
	if _, ok := conf["conn"]; !ok {
		panic("config has no conn key")
	}
	opts.Addr = conf["conn"]
	if _, ok := conf["dbnum"]; !ok {
		conf["dbnum"] = "0"
	}
	opts.DB, _ = strconv.Atoi(conf["dbnum"])
	if _, ok := conf["password"]; !ok {
		conf["password"] = ""
	}
	opts.Password = conf["password"]
	if _, ok := conf["maxactive"]; !ok {
		conf["maxactive"] = "0"
	}
	opts.PoolSize, _ = strconv.Atoi(conf["maxactive"])
	return opts
}

// Ping tests connectivity for redis (PONG should be returned)
func (r *RedisCache) Ping() error {
	pong, err := r.client.Ping().Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return errors.New(pong)
	}
	// Output: PONG <nil>

	return nil
}

func (r *RedisCache) Do(args ...interface{}) (reply interface{}, err error) {
	return r.client.Do(args...).Result()
}
