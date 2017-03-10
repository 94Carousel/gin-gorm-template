package store

import (
	"encoding/json"
	"time"

	"github.com/yingce/gin-gorm-template/config"

	redis "gopkg.in/redis.v5"
)

// Redis defined redis global client
var Redis *redis.Client

// InitRedis global initialize
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

func CacheSet(key string, value interface{}, expired ...time.Duration) {
	cacheName := config.Get("app", "name").Value()
	keyName := cacheName + "_cache:" + key
	j, _ := json.Marshal(value)
	var expiredAt time.Duration
	if len(expired) > 0 {
		expiredAt = expired[0]
	}
	Redis.Set(keyName+key, j, expiredAt)
}

func CacheGet(key string) (map[string]interface{}, bool) {
	cacheName := config.Get("app", "name").Value()
	keyName := cacheName + "_cache:" + key
	var data map[string]interface{}
	result, err := Redis.Get(keyName + key).Result()
	if err != nil || len(result) == 0 {
		return nil, false
	}
	byteDate := []byte(result)
	json.Unmarshal(byteDate, &data)
	return data, true
}
