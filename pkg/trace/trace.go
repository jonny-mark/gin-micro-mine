package trace

import (
	"errors"
	"github.com/jonny-mark/gin-micro-mine/internal/constant"
	appConfig "github.com/jonny-mark/gin-micro-mine/pkg/config"
	"github.com/jonny-mark/gin-micro-mine/pkg/load/nacos"
	jaegerprop "go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel"
	jaegerExporter "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

func Init() {
	var cfg Config
	if nacos.NacosClient.Enable {
		context, err := nacos.NacosClient.LoadConfiguration(constant.NacosTraceKey)
		if err != nil {
			log.Panicf("load trace conf err: %v", err)
		}
		if err := yaml.Unmarshal([]byte(context), &cfg); err != nil {
			log.Panicf("load trace conf unmarshal err: %v", err)
		}
		listenConfiguration(constant.NacosTraceKey, &cfg)
	} else {
		if err := appConfig.Load(constant.TraceKey, &cfg); err != nil {
			log.Panicf("trace config load %+v", err)
		}
		_, err := initTracerProvider(cfg.ServiceName, cfg.CollectorEndpoint)
		if err != nil {
			log.Panicf("trace init err %+v", err)
		}
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

// 监听nacos的变化
func listenConfiguration(name string, cfg *Config) {
	ctx, err := nacos.NacosClient.ListenConfiguration(name)
	if err != nil {
		log.Panicf("load trace conf err: %v", err)
	}
	if err := yaml.Unmarshal([]byte(ctx), cfg); err != nil {
		log.Panicf("load trace conf unmarshal err: %v", err)
	}
}
