/**
 * @author jiangshangfang
 * @date 2021/10/17 9:14 PM
 **/
package nacos

type Config struct {
	Enable       bool
	ClientConfig clientConfig
	ServerConfig serverConfig
}

type clientConfig struct {
	Env                 string
	TimeoutMs           uint64
	NotLoadCacheAtStart bool
	AppName             string
	LogDir              string
	CacheDir            string
	MaxAge              uint64
	LogLevel            string
	Namespace           string
	FileExtension       string
	ExtensionConfigs    []extensionConfigs
}

type serverConfig struct {
	ServerScheme  string
	ServerAddr    string
	ServerContext string
	ServerPort    uint64
}

type extensionConfigs struct {
	DataId string
	Group  string
}
