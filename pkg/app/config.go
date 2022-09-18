/**
 * @author jiangshangfang
 * @date 2021/12/12 8:11 PM
 **/
package app

import (
	"gin/internal/constant"
	"gin-micro-mine/pkg/config"
	"gin-micro-mine/pkg/load/nacos"
	"gopkg.in/yaml.v3"
	"log"
	"time"
)

var Conf *Config

type Config struct {
	Name              string         `yaml:"Name"`
	Version           string         `yaml:"Version"`
	Mode              string         `yaml:"Mode"`
	PprofPort         string         `yaml:"PprofPort"`
	JwtSecret         string         `yaml:"JwtSecret"`
	JwtTimeout        int            `yaml:"JwtTimeout"`
	SSL               bool           `yaml:"SSL"`
	CtxDefaultTimeout int64          `yaml:"CtxDefaultTimeout"`
	CSRF              bool           `yaml:"CSRF"`
	Debug             bool           `yaml:"Debug"`
	EnableTrace       bool           `yaml:"EnableTrace"`
	EnablePprof       bool           `yaml:"EnablePprof"`
	HTTP              ServerConfig   `yaml:"HTTP"`
	GRPC              ServerConfig   `yaml:"GRPC"`
	Registry          RegistryConfig `yaml:"Registry"`
}

type ServerConfig struct {
	Network      string        `yaml:"Network"`
	Addr         string        `yaml:"Addr"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

type RegistryConfig struct {
	Endpoints string `yaml:"Endpoints"`
}

//初始化配置项
func Init() {
	var cfg Config
	if nacos.NacosClient.Enable {
		context, err := nacos.NacosClient.LoadConfiguration(constant.NacosAppKey)
		if err != nil {
			log.Panicf("load app conf err: %v", err)
		}
		if err := yaml.Unmarshal([]byte(context), &cfg); err != nil {
			log.Panicf("load app conf unmarshal err: %v", err)
		}
		listenConfiguration(constant.NacosAppKey)
	} else {
		if err := config.Load(constant.AppKey, &cfg); err != nil {
			panic(err)
		}
	}
	// set global
	Conf = &cfg
}

//监听nacos的变化
func listenConfiguration(name string) {
	context, err := nacos.NacosClient.ListenConfiguration(name)
	if err != nil {
		log.Panicf("load app conf err: %v", err)
	}
	if err := yaml.Unmarshal([]byte(context), Conf); err != nil {
		log.Panicf("load app conf unmarshal err: %v", err)
	}
}
