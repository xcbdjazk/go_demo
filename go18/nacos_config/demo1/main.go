package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// NACOS 2.2.2
func main() {

	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           500,
		NotLoadCacheAtStart: true,
		LogDir:              ".",
		CacheDir:            ".",
		LogLevel:            "debug",
		Username:            "nacos",
		Password:            "123456",
	}

	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "172.25.89.138",
			ContextPath: "/nacos",
			Port:        8847,
			Scheme:      "http",
		},
		//{
		//	IpAddr:      "console2.nacos.io",
		//	ContextPath: "/nacos",
		//	Port:        80,
		//	Scheme:      "http",
		//},
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
	ok, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  "config_pre1.yaml",
		Group:   "DEFAULT_GROUP",
		Content: `ddd:1234`,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
	data, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "config_pre1.yaml",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "config_pre1.yaml",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		panic(err)
	}
	select {}
}
