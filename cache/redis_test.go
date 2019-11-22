package cache

import (
	"strconv"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

func TestRedisCache(t *testing.T) {
	conf := map[string]string{"conn": "127.0.0.1:6379"}
	bm := NewRedisCache()
	err := bm.StartAndGC(conf)
	if err != nil {
		t.Error("init err")
	}
	timeoutDuration := 10 * time.Second
	if err = bm.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie") {
		t.Error("check err")
	}

	time.Sleep(11 * time.Second)

	if bm.IsExist("astaxie") {
		t.Error("check err")
	}
	if err = bm.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 1 {
		t.Error("get err")
	}

	if _, err = bm.Incr("astaxie"); err != nil {
		t.Error("Incr Error", err)
	}

	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 2 {
		t.Error("get err")
	}

	if _, err = bm.Decr("astaxie"); err != nil {
		t.Error("Decr Error", err)
	}

	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 1 {
		t.Error("get err")
	}
	bm.Delete("astaxie")
	if bm.IsExist("astaxie") {
		t.Error("delete err")
	}

	//test string
	if err = bm.Put("astaxie", "author", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie") {
		t.Error("check err")
	}

	if v, _ := redis.String(bm.Get("astaxie"), err); v != "author" {
		t.Error("get err")
	}

	//test GetMulti
	if err = bm.Put("astaxie1", "author1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie1") {
		t.Error("check err")
	}

	vv := bm.GetMulti([]string{"astaxie", "astaxie1"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if v, _ := redis.String(vv[0], nil); v != "author" {
		t.Error("GetMulti ERROR")
	}
	if v, _ := redis.String(vv[1], nil); v != "author1" {
		t.Error("GetMulti ERROR")
	}

	// test clear all
	if err = bm.ClearAll(); err != nil {
		t.Error("clear all err")
	}
}

func TestRedisCache_DeleteLargeSet(t *testing.T) {
	conf := map[string]string{
		"conn":     "39.96.187.72:6379",
		"dbnum":    "3",
		"password": "Hjd123!@#",
	}
	bm := NewRedisCache()
	err := bm.StartAndGC(conf)
	if err != nil {
		t.Error("init err")
	}

	setKey := "test:set"
	bm.Delete(setKey)
	for i := 0; i < 10000; i++ {
		bm.do("SADD", setKey, i)
	}

	bm.DeleteLargeSet(setKey)
}

func TestRedisCache_DeleteLargeHash(t *testing.T) {
	conf := map[string]string{
		"conn":     "39.96.187.72:6379",
		"dbnum":    "3",
		"password": "Hjd123!@#",
	}
	bm := NewRedisCache()
	err := bm.StartAndGC(conf)
	if err != nil {
		t.Error("init err")
	}

	setKey := "test:hash"
	bm.Delete(setKey)
	for i := 0; i < 1000; i++ {
		key := strconv.Itoa(i)
		bm.do("HSET", setKey, key, "val"+key)
	}

	bm.DeleteLargeHash(setKey)
}
