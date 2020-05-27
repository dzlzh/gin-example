package utils

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func InitRedis() {

	host := Config.Section("redis").Key("host").String()
	database, _ := Config.Section("redis").Key("database").Int()
	max_idle, _ := Config.Section("redis").Key("max_idle").Int()
	max_active, _ := Config.Section("redis").Key("max_active").Int()

	idle_timeout_int, _ := Config.Section("redis").Key("idle_timeout").Int64()
	idle_timeout := time.Duration(idle_timeout_int) * time.Second

	timeout_int, _ := Config.Section("redis").Key("timeout").Int64()
	timeout := time.Duration(timeout_int) * time.Second

	// 建立连接池
	RedisPool = &redis.Pool{
		MaxIdle:     max_idle,
		MaxActive:   max_active,
		IdleTimeout: idle_timeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", host,
				//redis.DialPassword(conf["Password"].(string)),
				redis.DialDatabase(database),
				redis.DialConnectTimeout(timeout),
				redis.DialReadTimeout(timeout),
				redis.DialWriteTimeout(timeout))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}
