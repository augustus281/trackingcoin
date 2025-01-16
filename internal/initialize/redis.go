package initialize

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
)

var (
	_ctx = context.Background()
)

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		//Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Addr:     "localhost:6379",
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(_ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error: %v", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Connecting redis successfully!")
	global.Rdb = rdb
}
