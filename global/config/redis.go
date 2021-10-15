/**
 * @author jiangshangfang
 * @date 2021/7/30 10:24 AM
 **/
package config

type Redis struct {
Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
Pass         string `mapstructure:"password" json:"password" yaml:"password"`
Db           int    `mapstructure:"db" json:"db" yaml:"db"`
MaxRetries   int    `mapstructure:"maxRetries" json:"maxRetries" yaml:"maxRetries"`
PoolSize     int    `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize"`
MinIdleConns int    `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize" toml:"minIdleConns"`
}
