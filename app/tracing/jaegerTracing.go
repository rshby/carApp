package tracing

import (
	"carApp/app/config"
	"carApp/app/logging"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
)

func ConnectJaeger(config config.IConfig, log logging.ILogger, serviceName string) (opentracing.Tracer, io.Closer) {
	cfg := jaegerConfig.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%v:%v", config.Config().Jaeger.Host, config.Config().Jaeger.Port),
		},
	}

	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		log.LogConsole().Fatalf("cant connect jaeger : %v", err)
	}

	log.LogConsole().Info("success connect jaeger")
	return tracer, closer
}
