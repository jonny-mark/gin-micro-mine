/**
 * @author jiangshangfang
 * @date 2022/3/4 11:25 PM
 **/
package nacos

import (
	"errors"
	"fmt"
	"github.com/jonny-mark/gin-micro-mine/pkg/config"
	"github.com/jonny-mark/gin-micro-mine/pkg/load"
	"github.com/jonny-mark/gin-micro-mine/pkg/utils"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
	"strings"
)

var _ load.Load = &Nacos{}

var NacosClient *Nacos

type Nacos struct {
	Enable       bool
	configParam  map[string]vo.ConfigParam
	configClient config_client.IConfigClient
}

func Init() *Nacos {
	var c Config
	if err := config.Load("nacos", &c); err != nil {
		log.Panicf("nacos config load %+v", err)
	}

	nacosIps := strings.Split(c.ServerConfig.ServerAddr, ",")
	if 0 == len(nacosIps) {
		fmt.Println(fmt.Sprintf(`{"log_level":"ERROR","time":"%s","msg":"%s","server_name":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", c.ClientConfig.AppName, "nacos地址不能为空"))
		os.Exit(1)
	}

	var serverCs []constant.ServerConfig
	for _, ip := range nacosIps {
		serverCs = append(serverCs, constant.ServerConfig{
			Scheme:      c.ServerConfig.ServerScheme,
			IpAddr:      ip,
			ContextPath: c.ServerConfig.ServerContext,
			Port:        c.ServerConfig.ServerPort,
		})
	}
	// 初始化nacos的client服务
	if 0 == len(c.ClientConfig.Namespace) {
		fmt.Println(fmt.Sprintf(`{"log_level":"ERROR","time":"%s","msg":"%s","server_name":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", c.ClientConfig.AppName, "nacos命名空间不能为空"))
		os.Exit(1)
	}

	clientC := constant.ClientConfig{
		TimeoutMs:           c.ClientConfig.TimeoutMs,
		NamespaceId:         c.ClientConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		NotLoadCacheAtStart: c.ClientConfig.NotLoadCacheAtStart,
		AppName:             c.ClientConfig.AppName,
		LogDir:              c.ClientConfig.LogDir,
		CacheDir:            c.ClientConfig.CacheDir,
		LogLevel:            c.ClientConfig.LogLevel,
	}

	configMap := make(map[string]vo.ConfigParam)
	for _, configParam := range c.ClientConfig.ExtensionConfigs {
		key := fmt.Sprintf("%s_%s", configParam.Group, configParam.DataId)
		if _, ok := configMap[key]; ok {
			fmt.Println(fmt.Sprintf(`{"log_level":"ERROR","time":"%s","msg":"%s","server_name":"%s","key":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", c.ClientConfig.AppName, key, "key exists"))
			os.Exit(1)
		}
		configMap[key] = vo.ConfigParam{
			DataId: configParam.DataId,
			Group:  configParam.Group,
		}
	}

	configClient, configClientErr := clients.CreateConfigClient(map[string]interface{}{
		constant.KEY_SERVER_CONFIGS: serverCs,
		constant.KEY_CLIENT_CONFIG:  clientC,
	})
	if nil != configClientErr {
		fmt.Println(fmt.Sprintf(`{"log_level":"ERROR","time":"%s","msg":"%s","server_name":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", "连接失败", "nacos配置中心连接错误"))
		os.Exit(1)
	}
	NacosClient = &Nacos{
		Enable:       c.Enable,
		configClient: configClient,
		configParam:  configMap,
	}
	return NacosClient
}

func (n *Nacos) LoadConfiguration(name string) (content string, err error) {
	v, ok := n.configParam[name]
	if !ok {
		fmt.Sprintf(`{"log_level":"ERROR","time":"%s","msg":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", "nacos本地配置中没有该key")
		err = errors.New(fmt.Sprintf("nacos key:%v not exist", name))
		return
	}
	content, contentErr := n.configClient.GetConfig(v)
	if contentErr != nil {
		fmt.Sprintf(`{"log_level":"ERROR","time":"%s","msg":"%s","group":"%s","data-id":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", v.Group, v.DataId, "nacos配置中心获取配置错误")
		err = errors.New(fmt.Sprintf("nacos key:%v get config failed,err:%v", name, contentErr))
		return
	}
	return
}

func (n *Nacos) ListenConfiguration(name string) (content string, err error) {
	v, ok := n.configParam[name]
	if !ok {
		fmt.Sprintf(`{"log_level":"ERROR","time":"%s","msg":"%s","key":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", name, "nacos本地配置中没有该key")
		err = errors.New(fmt.Sprintf("nacos key:%v not exist", name))
		return
	}
	err = n.configClient.ListenConfig(vo.ConfigParam{
		DataId: v.DataId,
		Group:  v.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Sprintf(`{"time":"%s","msg":"%s","group":"%s","data-id":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", group, dataId, "nacos监听到配置的变化")
			content = data
		},
	})
	if nil != err {
		fmt.Println(fmt.Sprintf(`{"log_level":"ERROR","time":"%s","msg":"%s","group":"%s","data-id":"%s","desc":"%s"}`, utils.TimeNow(), "配置中心", v.Group, v.DataId, "nacos对该配置监听失败"))
	}
	return
}
