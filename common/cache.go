package common

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var c *cache.Cache

func InitCache() *cache.Cache {
	// 默认为两天过期，每10分钟清理一次
	cache := cache.New(2*24*time.Hour, 10*time.Minute)
	return cache
}

func GetCache() *cache.Cache {
	if c == nil {
		c = InitCache()
	}
	return c
}
