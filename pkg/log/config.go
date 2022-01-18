/**
 * @author jiangshangfang
 * @date 2021/10/17 9:14 PM
 **/
package log

import "time"

type Config struct {
	Name             string
	Development      bool
	Level            string
	Format           string
	Stacktrace       bool
	LinkName         string
	Prefix           string
	Director         string
	LogRollingPolicy string
	LoggerInfoFile   string
	LoggerWarnFile   string
	LoggerErrorFile  string
	MaxAge           time.Duration
}
