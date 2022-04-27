/**
 * @author jiangshangfang
 * @date 2021/10/30 9:48 PM
 **/
package trace

type Config struct {
	ServiceName        string `yaml:"ServiceName"`
	LocalAgentHostPort string `yaml:"LocalAgentHostPort"`
	CollectorEndpoint  string `yaml:"CollectorEndpoint"`
	CollectorUser      string `yaml:"CollectorUser"`
	CollectorPassword  string `yaml:"CollectorPassword"`
}
