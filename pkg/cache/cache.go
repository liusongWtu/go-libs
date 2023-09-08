package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	// SharedCache
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	// warning cache key conflict 注意key冲突
	SharedCache = cache.New(5*time.Minute, 10*time.Minute)
)
