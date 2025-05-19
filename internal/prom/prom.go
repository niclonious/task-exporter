package prom

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusServer struct {
	Registry  *prometheus.Registry
	Handler   gin.HandlerFunc
	GaugeVecs map[string]*prometheus.GaugeVec
}

func NewPrometheusServer() *PrometheusServer {
	promServer := &PrometheusServer{}
	promServer.Registry = prometheus.NewRegistry()
	promServer.Handler = gin.WrapH(promhttp.HandlerFor(promServer.Registry, promhttp.HandlerOpts{}))
	promServer.GaugeVecs = map[string]*prometheus.GaugeVec{}

	return promServer
}

// Creates and registers new gauge in the registry
func (ps *PrometheusServer) NewGaugeVec(opts prometheus.GaugeOpts, labels []string) *prometheus.GaugeVec {
	gaugevec := prometheus.NewGaugeVec(opts, labels)
	ps.Registry.MustRegister(gaugevec)
	ps.GaugeVecs[opts.Name] = gaugevec
	return gaugevec
}
