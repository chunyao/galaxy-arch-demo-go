package log

import (
	"app/src/common/config/log/lumberjack"
	"app/src/common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"time"
)

func InitLogConfig() {
	// 设置日志输出路径和名称
	logFilePath := path.Join(viper.GetString("logging.config.local-path"), viper.GetString("server.appName")+".log")
	// 日志输出滚动设置
	fileOut := &lumberjack.Logger{
		Filename:   logFilePath, // 日志文件位置
		MaxSize:    100,         // 单文件最大容量,单位是MB
		MaxBackups: 500,         // 最大保留过期文件个数
		MaxAge:     15,          // 保留过期文件的最大时间间隔,单位是天
		LocalTime:  true,        // 启用当地时区计时
	}
	// 文件和控制台日志输出
	writers := []io.Writer{
		fileOut,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdoutWriter)
	// 设置日志格式为Text格式
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置日志级别为Info以上
	log.SetLevel(log.InfoLevel)
}

// LoggerAccess 入口日志打印
func LoggerAccess(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	// 处理请求
	c.Next()
	// 请求方式
	reqMethod := c.Request.Method
	// 请求路由
	reqUri := c.Request.RequestURI
	// 状态码
	statusCode := c.Writer.Status()
	// 服务器IP
	serverIP := utils.GetLocalIP()
	// 客户端IP
	clientIP := c.ClientIP()
	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))
	//日志格式
	log.WithFields(log.Fields{
		"server-ip": serverIP,
		"duration":  latencyTime,
		"status":    statusCode,
		"method":    reqMethod,
		"uri":       reqUri,
		"client-ip": clientIP,
	}).Info("Api accessing")
}
