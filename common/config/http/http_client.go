package http

import (
	"github.com/spf13/viper"
	"net"
	"net/http"
	"time"
)

var HttpClient *http.Client // HttpClient

// InitHttpClientConfig 初始化HttpClient
func InitHttpClientConfig() {
	maxIdleConns := viper.GetInt("http-client.maxIdleConns")
	maxIdleConnsPerHost := viper.GetInt("http-client.maxIdleConnsPerHost")
	maxConnsPerHost := viper.GetInt("http-client.maxConnsPerHost")
	idleConnTimeout := viper.GetInt("http-client.idleConnTimeout")
	HttpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			DisableKeepAlives:     false,
			TLSHandshakeTimeout:   10 * time.Second,
			MaxIdleConns:          maxIdleConns,
			MaxIdleConnsPerHost:   maxIdleConnsPerHost,
			MaxConnsPerHost:       maxConnsPerHost,
			IdleConnTimeout:       time.Duration(idleConnTimeout) * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 0,
		},
	}
}
