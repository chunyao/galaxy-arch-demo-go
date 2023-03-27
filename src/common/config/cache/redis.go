package cache

import (
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var Redis *redis.Client

// InitRedisConfig 初始化redis
func InitRedisConfig() {
	log.Info("Initializing Redis")

	var (
		address     = viper.GetString("redis.address")
		db          = viper.GetInt("redis.db")
		password    = viper.GetString("redis.password")
		idleTimeout = time.Duration(viper.GetInt("redis.idleTimeout")) * time.Second
		maxIdle     = viper.GetInt("redis.maxIdle")
		poolSize    = viper.GetInt("redis.poolSize")
	)
	Redis = redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     password, // no password set
		DB:           db,       // use default DB
		PoolSize:     poolSize,
		MaxIdleConns: maxIdle,
		PoolTimeout:  idleTimeout,
	})

	log.Info("Redis: initialization completed")
}
