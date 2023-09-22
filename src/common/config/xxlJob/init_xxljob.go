package xxlJob

import (
	"app/src/common/config/log"
	err "app/src/common/exception"
	"app/src/job"
	"fmt"
	"github.com/gin-gonic/gin"
	xxl_job_executor_gin "github.com/gin-middleware/xxl-job-executor"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xxl-job/xxl-job-executor-go"
)

const Port = "9999"

func InitJob() {
	if viper.GetString("xxl.job.enable") == "true" {
		//初始化执行器
		logger.Info("初始化执行器..." + viper.GetString("xxl.job.admin.addresses"))
		exec := xxl.NewExecutor(
			xxl.ServerAddr(viper.GetString("xxl.job.admin.addresses")),
			xxl.AccessToken(viper.GetString("xxl.job.accessToken")), //请求令牌(默认为空)
			//	xxl.ExecutorIp("127.0.0.1"),                                  //可自动获取
			xxl.ExecutorPort(Port), //默认9999（此处要与gin服务启动port必需一至）
			xxl.RegistryKey(viper.GetString("xxl.job.executor.appname")), //执行器名称

		)
		exec.Init()
		defer exec.Stop()
		//添加到gin路由
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.Use(log.LoggerAccess)
		// 统一异常处理
		r.Use(err.ErrHandle)
		xxl_job_executor_gin.XxlJobMux(r, exec)

		//注册gin的handler
		r.GET("ping", func(cxt *gin.Context) {
			cxt.JSON(200, "pong")
		})

		//注册任务handler

		task := job.Task{}

		dispatch := task.DoTask()
		for k, v := range dispatch {
			logger.Info("XXL-Job: 注册任务: " + k)
			exec.RegTask(k, v)
		}

		go runGin(r)
		logger.Info("XXL-Job: 初始化完成……")

	}

}
func runGin(router *gin.Engine) {
	logger.Info(fmt.Sprintf("XXL-Job Service started on port(s): %s", Port))
	_ = router.Run(":" + Port)
}
