/**
 * @author jiangshangfang
 * @date 2021/10/21 4:35 PM
 **/
package orm

import "time"

type Config struct {
	Name            string
	Addr            string
	UserName        string
	Password        string
	MaxIdleConn     int
	ConnMaxLifeTime time.Duration
	MaxOpenConn     int
	ShowLog         bool
	SlowThreshold   time.Duration
}
