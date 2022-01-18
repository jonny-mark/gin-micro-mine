/**
 * @author jiangshangfang
 * @date 2021/10/28 8:01 PM
 **/
package trace

import (
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	appConfig "gin/pkg/config"
	"log"
	"go.opentelemetry.io/otel"
	jaegerExporter "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	jaegerprop "go.opentelemetry.io/contrib/propagators/jaeger"
	"errors"
	"strings"
)

func Init() {
	var traceConfig Config
	if err := appConfig.Load("trace", &traceConfig); err != nil {
		log.Panicf("trace config load %+v1", err)
	}
	_, err := initTracerProvider(traceConfig.ServiceName, traceConfig.CollectorEndpoint)
	if err != nil {
		log.Panicf("trace init err %+v1", err)
	}
}

func initTracerProvider(serviceName, endpoint string, options ...Option) (*tracesdk.TracerProvider, error) {
	var endpointOption jaegerExporter.EndpointOption
	if serviceName == "" {
		return nil, errors.New("no service name provided")
	}

	if strings.HasPrefix(endpoint, "http") {
		endpointOption = jaegerExporter.WithCollectorEndpoint(jaegerExporter.WithEndpoint(endpoint))
	} else {
		endpointOption = jaegerExporter.WithAgentEndpoint(jaegerExporter.WithAgentHost(endpoint))
	}

	//初始化exporter
	exporter, err := jaegerExporter.New(endpointOption)
	if err != nil {
		return nil, err
	}

	opts := applyOptions(options...)
	tp := tracesdk.NewTracerProvider(
		//设置Sampler
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(opts.SamplingRatio)),
		//允许批量生产
		tracesdk.WithBatcher(exporter),
		//记录Resource信息
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	//全局注册
	otel.SetTracerProvider(tp)
	//设置全局TextMapPropagator
	otel.SetTextMapPropagator(jaegerprop.Jaeger{})

	return tp, nil
}
