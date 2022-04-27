/**
 * @author jiangshangfang
 * @date 2021/7/30 10:23 AM
 **/
package config

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sync"
)

//全局Config
var Conf *Config

type Config struct {
	env        string
	configDir  string
	configType string
	val        map[string]*viper.Viper
	mu         sync.Mutex
}

func New(dir string, opts ...Option) *Config {
	if dir == "" {
		panic("config dir is not set")
	}
	c := &Config{
		configDir:  dir,
		configType: "yaml",
		val:        make(map[string]*viper.Viper),
	}
	for _, opt := range opts {
		opt(c)
	}
	Conf = c
	return c
}

//Load函数
func Load(filename string, val interface{}) error {
	return Conf.Load(filename, val)
}

//Load方法
func (c *Config) Load(filename string, val interface{}) error {
	v, err := c.LoadWithType(filename, c.configType)
	if err != nil {
		return err
	}
	err = v.Unmarshal(&val)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) LoadWithType(filename string, configType string) (*viper.Viper, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	vip, ok := c.val[filename]
	if ok {
		return vip, nil
	}
	vip, err := c.load(filename, configType)
	if err != nil {
		return nil, err
	}
	c.val[filename] = vip
	return vip, nil
}

func (c *Config) load(filename string, configType string) (*viper.Viper, error) {
	var path string
	if c.env != "" {
		path = filepath.Join(c.configDir, c.env)
	} else {
		env := GetEnv("APP_ENV", "")
		path = filepath.Join(c.configDir, env)
	}
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(filename)
	v.SetConfigType(c.configType)
	if configType != "" {
		v.SetConfigType(configType)
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
	})
	return v, nil
}

func GetEnv(key string, defaultKey string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultKey
	}
	return val
}
