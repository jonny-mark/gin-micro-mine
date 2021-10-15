/**
 * @author jiangshangfang
 * @date 2021/7/30 10:23 AM
 **/
package config

type Server struct {
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//log
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	//System
	System System `mapstructure:"system" json:"system" yaml:"system"`
	//redis
	Redis System `mapstructure:"system" json:"system" yaml:"system"`
}
