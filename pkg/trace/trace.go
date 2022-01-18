/**
 * @author jiangshangfang
 * @date 2021/10/28 8:01 PM
 **/
package trace

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

func GetTracer(cfg *Config) opentracing.Tracer {
	// 判断是否已注册了，单例模式
	if opentracing.IsGlobalTracerRegistered() {
		return opentracing.GlobalTracer()
	}
	c := &config.Configuration{
		ServiceName: cfg.traceAgent,
		Sampler: &config.SamplerConfig{
			SamplingServerURL: cfg.jaeger.samplingServerURL,
			Type:              "const",
			Param:             1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort: cfg.jaeger.localAgentHostPort,
			CollectorEndpoint:  cfg.jaeger.collectorEndpoint,
			User:               cfg.jaeger.collectorUser,
			Password:           cfg.jaeger.collectorPassword,
		},
		Headers: &jaeger.HeadersConfig{
			JaegerDebugHeader:        "x-debug-id",
			JaegerBaggageHeader:      "x-baggage",
			TraceContextHeaderName:   "x-trace-id",
			TraceBaggageHeaderPrefix: "x-ctx",
		},
	}

	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, _, err := c.NewTracer(
		config.Logger(jaeger.StdLogger),
		config.Injector(opentracing.HTTPHeaders, propagator),
		config.Extractor(opentracing.HTTPHeaders, propagator),
		config.ZipkinSharedRPCSpan(true),
		config.MaxTagValueLength(256),
		config.PoolSpans(true),
	)
	if err != nil {
		panic(fmt.Sprintf("init trace failed: %v1\n", err))
	}
	// 设置到全局里
	opentracing.SetGlobalTracer(tracer)

	return tracer
}
