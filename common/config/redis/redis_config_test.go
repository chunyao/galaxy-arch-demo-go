package redis

//import (
//	"github.com/gomodule/redigo/redis"
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"time"
//)
//
//func init() {
//	var (
//		address     = "127.0.0.1:6379"
//		index       = 0
//		password    = ""
//		idleTimeout = time.Duration(30) * time.Second
//		maxIdle     = 50
//		maxActive   = 1000
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
//}
//
//func TestSetObject(t *testing.T) {
//	SetObject("484ac281d6e1473a81824db648ada376", 18, 30)
//}
//
//func TestGetObject(t *testing.T) {
//	result := GetObject("484ac281d6e1473a81824db648ada376").(int)
//	respect := 18
//	assert.Equal(t, result, respect)
//}
