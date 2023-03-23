package viper

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const fileName = "application"

// InitLocalConfigFile 加载本地配置文件
func InitLocalConfigFile() {
	log.Info("初始化本地配置文件……")
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s \n", err))
	}
	log.Info("本地配置文件初始化完成……")
}
