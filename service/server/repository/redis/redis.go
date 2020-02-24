package redis

import (
	"github.com/garyburd/redigo/redis"
)

// redis repo implement CacheRepo
type Repo struct {
	// 由service层传入
	Entity interface{}

	// 数据库实例
	DB *redis.Conn
}
