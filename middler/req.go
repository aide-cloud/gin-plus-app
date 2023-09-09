package middler

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// MetricSeconds 请求耗时
	MetricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "server requests duration(ms).",
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	}, []string{"kind", "operation"})

	// MetricRequests 已处理请求的总数
	MetricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})

	// IpMetricCounter 出现的IP频次
	IpMetricCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "ip_total",
		Help:      "The total number of processed requests",
	}, []string{"ip"})
)

func init() {
	prometheus.MustRegister(MetricSeconds, MetricRequests, IpMetricCounter)
}

// MetricSecondsMiddleware MetricSeconds中间件, 用于记录请求耗时
func MetricSecondsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()
		ctx.Next()
		elapsed := time.Since(now)
		MetricSeconds.WithLabelValues(ctx.Request.Method, ctx.Request.URL.Path).Observe(float64(elapsed.Nanoseconds()))
		log.Println("[MetricSecondsMiddleware] ", ctx.Request.Method, ctx.Request.URL.Path, elapsed)
	}
}

// MetricRequestsMiddleware MetricRequests中间件, 用于记录已处理请求的总数
func MetricRequestsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		MetricRequests.WithLabelValues(ctx.Request.Method, ctx.Request.URL.Path, strconv.Itoa(ctx.Writer.Status()), "").Inc()
		log.Println("[MetricRequestsMiddleware] ", ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status())
	}
}

// IpMetricMiddleware IpMetric中间件, 用于记录出现的IP频次
func IpMetricMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		IpMetricCounter.WithLabelValues(ctx.Request.RemoteAddr).Inc()
		log.Println("[IpMetricMiddleware] ", ctx.Request.RemoteAddr)
	}
}
