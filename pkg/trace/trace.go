package trace

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

// InitJaeger 将jaeger tracer设置为全局tracer
func InitJaeger(service string) io.Closer {
	// 根据配置初始化Tracer 返回Closer
	tracer, closer, err := (&config.Configuration{
		ServiceName: service,
		Disabled:    false,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			// param的值在0到1之间，设置为1则将所有的Operation输出到Reporter
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
			// 将span发往jaeger-collector的服务地址，需要拉一个jaeger起来（详情看docker-compose）
			// CollectorEndpoint: "http://localhost:14268/api/traces",
		},
	}).NewTracer()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	// 设置全局Tracer - 如果不设置将会导致上下文无法生成正确的Span
	opentracing.SetGlobalTracer(tracer)
	return closer
}
