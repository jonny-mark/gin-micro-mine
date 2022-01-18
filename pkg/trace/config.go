/**
 * @author jiangshangfang
 * @date 2021/10/30 9:48 PM
 **/
package trace

type Config struct {
	serviceName string // The name of this service
	traceAgent  string // The type of trace agent: zipkin, jaeger or elastic

	jaeger Jaeger
}

type Jaeger struct {
	samplingServerURL string // provide sampling strategy to this service
	localAgentHostPort string // Set jaeger-agent's host:port that the reporter will used
	collectorEndpoint  string // Instructs reporter to send spans to jaeger-collector at this URL
	collectorUser      string // CollectorUser for basic http authentication when sending spans to jaeger-collector
	collectorPassword  string // CollectorPassword for basic http authentication when sending spans to jaeger-collector
}
