/**
 * @author jiangshangfang
 * @date 2021/8/8 2:30 PM
 **/
package config

type Zap struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`                     // 级别
	Format      string `mapstructure:"format" json:"format" yaml:"format"`                  // 输出
	ShowLine    bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`           // 显示行
	LinkName    string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`          // 软链接名称
	Prefix      string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                  // 日志前缀
	EncodeLevel string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"` // 编码级
	Director    string `mapstructure:"director" json:"director" yaml:"director"`            // 日志文件夹
}
