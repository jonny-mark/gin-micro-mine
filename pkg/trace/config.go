/**
 * @author jiangshangfang
 * @date 2021/10/30 9:48 PM
 **/
package trace

type Config struct {
	ServiceName        string // The name of this service
	LocalAgentHostPort string //Set jaeger-agent's host:port that the reporter will used
	CollectorEndpoint  string //Instructs reporter to send spans to jaeger-collector at this URL
	CollectorUser      string // CollectorUser for basic http authentication when sending spans to jaeger-collector
	CollectorPassword  string // CollectorPassword for basic http authentication when sending spans to jaeger-collector
}
