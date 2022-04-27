package initialize

import (
	"github.com/go-redis/redis/v8"
	"mall/global"
)

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password, // no password set
		DB:       global.Config.Redis.DB,       // use default DB
	})
	global.Rdb = rdb
}
