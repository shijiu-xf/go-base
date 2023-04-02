package jeagersj

import (
	"github.com/opentracing/opentracing-go"
	jeager "github.com/uber/jaeger-client-go"
	jeagerConf "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	conf := &jeagerConf.Configuration{
		ServiceName: serviceName,
		Sampler: &jeagerConf.SamplerConfig{
			Type:  jeager.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jeagerConf.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
	}
	return conf.NewTracer()
}
