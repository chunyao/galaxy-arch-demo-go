package db

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

const filePath = "../../application.yaml"

func TestInitDb(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s \n", err))
	}
	InitDbConfig()
	assert.NotNil(t, DB)
}
