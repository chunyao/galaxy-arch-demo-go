package utils

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

type NacosConfigClient struct {
	//ConfigInput
	//"ip1,ip2,ip3"
	IpAddrs         string
	Port            int
	NaocesNameSpace string
	DataId          string
	Group           string
	NacosLogDir     string
	NacosCacheDir   string
	//NamingInput
	NamingServiceName string
	NamingServiceIP   string
	NamingOwner       string
	NamingPort        int
	//Output
	ConfigClient config_client.IConfigClient
	NamingClient naming_client.INamingClient
}

func NewNacosClient(nacosConfigClient NacosConfigClient) NacosConfigClient {

	serverConfigs := make([]constant.ServerConfig, 0)
	address := nacosConfigClient.IpAddrs
	serverConfigs = append(serverConfigs, constant.ServerConfig{
		IpAddr: address,
		Port:   uint64(nacosConfigClient.Port),
	})
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfigClient.NaocesNameSpace, //
		NotLoadCacheAtStart: false,
		TimeoutMs:           5 * 1000,
		LogDir:              nacosConfigClient.NacosLogDir,
		CacheDir:            nacosConfigClient.NacosCacheDir,
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {

		panic(err)
	}
	nacosConfigClient.ConfigClient = configClient

	// Naming client
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}
	nacosConfigClient.NamingClient = namingClient
	return nacosConfigClient
}
func (client *NacosConfigClient) GetConfig() (string, error) {
	content, err := client.ConfigClient.GetConfig(vo.ConfigParam{
		DataId: client.DataId,
		Group:  client.Group,
	})

	return content, err
}
func (client *NacosConfigClient) Register() (bool, error) {
	param := vo.RegisterInstanceParam{
		Ip:          client.NamingServiceIP,
		Port:        uint64(client.NamingPort),
		ServiceName: client.NamingServiceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"owenr": client.NamingOwner, "uptime": time.Now().Format("2006-01-02 15:04:05")},
	}
	success, err := client.NamingClient.RegisterInstance(param)
	return success, err
}
