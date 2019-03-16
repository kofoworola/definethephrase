package redis

import (
	"github.com/gomodule/redigo/redis"
)

var activePool *redis.Pool

func GetPool() *redis.Pool {
	if activePool != nil {
		return activePool
	}
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
