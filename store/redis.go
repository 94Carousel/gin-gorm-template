package store

import (
	"github.com/yingce/gin-gorm-template/config"

	redis "gopkg.in/redis.v5"
)

// Redis defined redis global client
var Redis *redis.Client

// InitRedis config
func InitRedis() {
	redisCfg := config.GetSection("redis")
	addr := redisCfg.Key("addr").Value()
	password := redisCfg.Key("password").Value()
	db, _ := redisCfg.Key("port").Int()
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}
