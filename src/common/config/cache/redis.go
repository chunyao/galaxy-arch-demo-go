package cache

import (
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var (
	RDs = map[string]*RDBManager{} // 初始化时加载数据源到集合
)

// 连接管理器
type RDBManager struct {
	Redis *redis.Client // redis
}

// InitRedisConfig 初始化redis
func InitRedisConfig() {
	log.Info("初始化 Redis")
	for k, _ := range viper.GetStringMap("redis") {
		log.Printf("初始化Redis数据源 %s ", k)
		db := RDsSetup(k)
		rdb := &RDBManager{
			Redis: db,
		}
		RDs[k] = rdb
	}

	log.Info("Redis: 初始化完成")
}
func RDsSetup(name string) *redis.Client {
	var (
		address     = viper.GetString("redis." + name + ".address")
		db          = viper.GetInt("redis." + name + ".db")
		password    = viper.GetString("redis." + name + ".password")
		idleTimeout = time.Duration(viper.GetInt("redis."+name+".idleTimeout")) * time.Second
		maxIdle     = viper.GetInt("redis." + name + ".maxIdle")
		poolSize    = viper.GetInt("redis." + name + ".poolSize")
	)
	Redis := redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     password, // no password set
		DB:           db,       // use default DB
		PoolSize:     poolSize,
		MaxIdleConns: maxIdle,
		PoolTimeout:  idleTimeout,
	})
	return Redis
}
