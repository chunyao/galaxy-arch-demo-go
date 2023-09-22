package cos

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var Client *cos.Client

func InitCosClientConfig() {
	log.Info("初始化 CosClient 连接池……" + viper.GetString("tencent.baseUrl"))

	u, _ := url.Parse(viper.GetString("tencent.baseUrl"))
	b := &cos.BaseURL{BucketURL: u}
	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("tencent.secretId"),
			SecretKey: viper.GetString("tencent.secretKey"),
		},
	})
	log.Info("初始化 CosClient 完成……")
}
