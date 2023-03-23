package main

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"mabang-arch-demo-go/common/config/cache"
	"mabang-arch-demo-go/common/config/db"
	"mabang-arch-demo-go/common/config/gin"
	"mabang-arch-demo-go/common/config/http"
	"mabang-arch-demo-go/common/config/log"
	"mabang-arch-demo-go/common/config/redis"
	"mabang-arch-demo-go/common/config/token"
	vc "mabang-arch-demo-go/common/config/viper"
	"mabang-arch-demo-go/controller"
	"mabang-arch-demo-go/dao"
	_ "net/url"
)

func main() {
	initComponents()
}

// 初始化服务所有组件
func initComponents() {
	// 初始化日志
	log.InitLogConfig()
	logger.Info("===================================================================================")
	logger.Info("Starting Application")
	// 读取本地配置文件
	vc.InitLocalConfigFile()
	// 初始化Mysql
	db.InitDbConfig()
	// 自动生成表
	// autoMigrate()
	// 初始化缓存
	cache.InitBigCacheConfig()
	// 初始化Redis
	redis.InitRedisConfig()

	// 初始化HttpClient连接池
	http.InitHttpClientConfig()

	// 初始化token
	token.InitTokenConfig()

	// 初始化Gin
	router := gin.InitGinConfig()

	// 注册Api
	// 用户api
	controller.UserApi(router)

	// 启动Gin
	gin.RunGin(router)
}

// 自动生成表
func autoMigrate() {
	err := db.DB.AutoMigrate(dao.User{})
	if err != nil {
		_ = fmt.Errorf("自动生成user表失败")
		panic(err)
	}

}
