package serversuite

import (
	"github.com/cloudwego/gomall07/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler), // 服务器元数据处理器
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ // 服务器基本信息
			ServiceName: s.CurrentServiceName,
		}),
		server.WithTracer(prometheus.NewServerTracer("", // 追踪器
			"",
			prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry)),
		),
		server.WithSuite(tracing.NewServerSuite()),
	}
	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))
	return opts
}
