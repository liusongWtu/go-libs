package goredis

import (
	"fmt"
	"testing"
	"time"
)

func TestRedisCache_Ping(t *testing.T) {

	client := NewRedisCache(map[string]string{
		"conn":     "127.0.0.1:63798",
		"dbnum":    "1",
		"password": "XiaoZi527",
	})
	err := client.Ping()
	if err != nil {
		t.Errorf("Ping error:%s\n", err.Error())
	}
}

func TestRedisCache_Do(t *testing.T) {
	redisCache := NewRedisCache(map[string]string{
		"conn":     "127.0.0.1:63798",
		"dbnum":    "1",
		"password": "XiaoZi527",
	})
	for {
		reply, err := redisCache.Do("GET", "signin:1")
		if err != nil {
			t.Errorf("do get error:%s\n", reply)
			return
		}
		fmt.Println(reply)
		t.Log(reply)
		time.Sleep(3 * time.Second)
	}

}
