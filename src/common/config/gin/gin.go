package gin

import (
	"app/src/common/config/log"
	err "app/src/common/exception"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitGinConfig 初始化Gin
func InitGinConfig() *gin.Engine {
	logger.Info("初始化 gin……")
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	// 入口日志打印
	router.Use(log.LoggerAccess)
	// 统一异常处理
	router.Use(err.ErrHandle)
	// 跨域处理
	router.Use(cors.Default())

	// token校验
	//router.Use(token.TokenVerify)
	// 健康检测
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	logger.Info("Gin: 初始化完成……")
	return router
}

// RunGin 启动Gin
func RunGin(router *gin.Engine) {
	port := viper.GetString("server.port")
	logger.Info(fmt.Sprintf("Service started on port(s): %s", port))
	_ = router.Run(":" + port)
}
