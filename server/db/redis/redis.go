package redis

import (
	"github.com/gomodule/redigo/redis"
	"server/config"
	"time"
)

var rds *redis.Pool

func Init() {
	rds = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				config.Configs.RedisHost,
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(config.Configs.RedisDb),
				redis.DialPassword(config.Configs.RedisPwd),
			)
		},
	}
}

func GetConn() *redis.Pool {
	return rds
}
