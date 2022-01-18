/**
 * @author jiangshangfang
 * @date 2021/10/23 10:34 PM
 **/
package grpc

import "time"

type Grpc struct {
	Addr         string        `mapstructure:"addr" json:"addr" yaml:"addr"`                         // 服务器地址:端口
	ReadTimeout  time.Duration `mapstructure:"readTimeout" json:"readTimeout" yaml:"readTimeout"`    //套接字读超时时间
	WriteTimeout time.Duration `mapstructure:"writeTimeout" json:"writeTimeout" yaml:"writeTimeout"` //套接字写超时时间
}
