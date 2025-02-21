package config

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

var Client *redis.Client

func InitRedis() {
	redisCfg := Cfg.Redis
	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Ports),
		Username: redisCfg.User,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	}
	Client = redis.NewClient(&opt)
	return
}
