/**
 * @author jiangshangfang
 * @date 2021/12/12 8:11 PM
 **/
package app

import "time"

var Conf *Config

type Config struct {
	Name              string
	Version           string
	Mode              string
	PprofPort         string
	URL               string
	JwtSecret         string
	JwtTimeout        int
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
	EnableTrace       bool
	EnablePprof       bool
	HTTP              ServerConfig
	GRPC              ServerConfig
	Registry          RegistryConfig
}

type ServerConfig struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type RegistryConfig struct {
	Endpoints string
}
