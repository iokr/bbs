package redis

import (
	"github.com/iokr/bbs/core/cache/redis"
)

var Client *redis.Client

func Init(c *redis.Config) {
	if Client == nil {
		Client = redis.NewRedis(c)
	}
}
