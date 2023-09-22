package viper

import (
	"app/src/common/utils"
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const fileName = "application_"

// InitLocalConfigFile 加载本地配置文件
func InitLocalConfigFile() {

	//config.GetConfig()
	log.Info("初始化本地配置文件 " + fileName + os.Getenv("ACTIVE_PROFILE") + "……")
	profile := "local"
	if os.Getenv("ACTIVE_PROFILE") == "" {
		viper.SetConfigName(fileName + profile)
	} else {
		viper.SetConfigName(fileName + os.Getenv("ACTIVE_PROFILE"))
	}

	viper.SetConfigType("properties")
	viper.AddConfigPath("./resources")

	err := viper.ReadInConfig()

	if viper.GetString("cloud.nacos.server-addr") != "" {

		log.Info("nacos……" + viper.GetString("cloud.nacos.server-addr"))
		config := utils.NacosConfigClient{
			IpAddrs:         viper.GetString("cloud.nacos.server-addr"),
			Port:            443,
			NaocesNameSpace: viper.GetString("cloud.nacos.config.namespace"),
			DataId:          viper.GetString("cloud.nacos.config.dataId"),
			Group:           viper.GetString("cloud.nacos.config.group"),
			NacosLogDir:     viper.GetString("logging.config.local-path"),
			NacosCacheDir:   viper.GetString("logging.config.local-path"),
		}
		client := utils.NewNacosClient(config)
		configStr, _ := client.GetConfig()
		err := viper.ReadConfig(bytes.NewReader([]byte(configStr)))
		if err != nil {
			panic(fmt.Errorf("读取配置文件失败: %s \n", err))
		}

	}

	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s \n", err))
	}

}
