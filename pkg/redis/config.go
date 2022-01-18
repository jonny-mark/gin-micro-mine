/**
 * @author jiangshangfang
 * @date 2021/10/23 4:33 PM
 **/
package redis

import "time"

type Config struct {
	Addr         string
	Password     string
	Db           int
	MaxRetries   int
	PoolSize     int
	PoolTimeout  time.Duration
	MinIdleConns int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	EnableTrace  bool
}
