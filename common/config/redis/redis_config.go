package redis

//import (
//	"errors"
//	"github.com/gomodule/redigo/redis"
//	log "github.com/sirupsen/logrus"
//	"github.com/spf13/viper"
//	"go-web-demo/utils"
//	"time"
//)
//
//var redisPool *redis.Pool
//
//// InitRedisConfig 初始化redis
//func InitRedisConfig() {
//	log.Info("Initializing Redis")
//
//	var (
//		address     = viper.GetString("redis.address")
//		index       = viper.GetInt("redis.index")
//		password    = viper.GetString("redis.password")
//		idleTimeout = time.Duration(viper.GetInt("redis.idleTimeout")) * time.Second
//		maxIdle     = viper.GetInt("redis.maxIdle")
//		maxActive   = viper.GetInt("redis.maxActive")
//	)
//	redisPool = &redis.Pool{
//		MaxIdle:     maxIdle,     // 最大的空闲连接数
//		MaxActive:   maxActive,   // 最大的激活连接数
//		IdleTimeout: idleTimeout, // 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
//		Dial: func() (redis.Conn, error) {
//			conn, err := redis.Dial("tcp", address)
//			if err != nil {
//				return nil, err
//			}
//			if password != "" {
//				// 验证密码
//				if _, err := conn.Do("AUTH", password); err != nil {
//					err := conn.Close()
//					if err != nil {
//						return nil, err
//					}
//					return nil, err
//				}
//			}
//			// 选择db
//			if _, err := conn.Do("SELECT", index); err != nil {
//				return nil, err
//			}
//			return conn, nil
//		},
//		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
//			_, err := conn.Do("PING")
//			return err
//		},
//	}
//	log.Info("Redis: initialization completed")
//}
//
//// existedObject 根据key是否存在该对象
//func existedObject(key string) (isExisted bool) {
//	conn := redisPool.Get()
//	if conn == nil {
//		panic(errors.New("redis connection is nil"))
//	}
//	defer func(conn redis.Conn) {
//		err := conn.Close()
//		if err != nil {
//
//		}
//	}(conn)
//	reply, err := conn.Do("EXISTS", key)
//	if err != nil {
//		panic(err)
//	}
//	existed, _ := redis.Int(reply, nil)
//	return existed == 1
//}
//
//// SetObject 以Key、Value将对象存入Redis
//// 并设置Key有效时间
//func SetObject(key string, value interface{}, duration int) {
//	conn := redisPool.Get()
//	if conn == nil {
//		panic(errors.New("redis connection is nil"))
//	}
//	defer func(conn redis.Conn) {
//		err := conn.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(conn)
//	valueBytes := utils.Serialize(value)
//	_, err := conn.Do("SET", key, valueBytes, "EX", duration)
//	if err != nil {
//		panic(err)
//	}
//}
//
//// GetObject 根据Key从Redis获取对象
//func GetObject(key string) (value interface{}) {
//	// 存档是否存在
//	isExisted := existedObject(key)
//	if !isExisted {
//		return nil
//	}
//	conn := redisPool.Get()
//	if conn == nil {
//		panic(errors.New("redis connection is nil"))
//	}
//	defer func(conn redis.Conn) {
//		err := conn.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(conn)
//	valueBytes, err := redis.Bytes(conn.Do("GET", key))
//	if err != nil {
//		panic(err)
//	}
//	value = utils.Deserialize(valueBytes)
//	return
//}
