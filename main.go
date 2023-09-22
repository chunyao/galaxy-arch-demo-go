package main

import (
	"app/src/common/config/cache"
	"app/src/common/config/cos"
	"app/src/common/config/db"
	"app/src/common/config/gin"
	http2 "app/src/common/config/http"
	"app/src/common/config/log"
	"app/src/common/config/mongo"
	"app/src/common/config/token"
	vc "app/src/common/config/viper"
	"app/src/common/config/xxlJob"
	"app/src/controller"
	"runtime"

	logger "github.com/sirupsen/logrus"
	_ "net/url"
)

func main() {
	var MULTICORE int = runtime.NumCPU() //number of core

	runtime.GOMAXPROCS(MULTICORE) //running in multicore
	initComponents()
}

// 初始化服务所有组件
func initComponents() {
	// 读取本地配置文件
	vc.InitLocalConfigFile()
	// 初始化日志
	log.InitLogConfig()
	logger.Info("===================================================================================")
	logger.Info("Starting Application")

	// 初始化Mysql
	db.InitDbConfig()

	// 初始化缓存
	cache.InitBigCacheConfig()
	// 初始化Redis
	cache.InitRedisConfig()
	// 初始化mongo
	mongo.InitMongoDB()
	// 初始化HttpClient连接池
	http2.InitHttpClientConfig()
	// 初始化Xxl-Job
	xxlJob.InitJob()
	// 初始化token
	token.InitTokenConfig()
	cos.InitCosClientConfig()
	// 初始化Gin
	router := gin.InitGinConfig()

	// 注册Api
	// 用户api
	controller.UserApi(router)
	controller.ImageApi(router)
	controller.ScreenshotApi(router)
	// 启动Gin
	gin.RunGin(router)
}
