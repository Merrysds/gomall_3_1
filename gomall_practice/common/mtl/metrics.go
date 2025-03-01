package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry // 相当于注册中心，用来提供指标

func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())                                       // 初始化go运行时相关的指标
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})) // 初始化进程相关的指标
	r, _ := consul.NewConsulRegister(registryAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registryInfo)
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	// 异步启动一个sever给PROMETHUES
	go http.ListenAndServe(metricsPort, nil)
	return r, registryInfo
}
