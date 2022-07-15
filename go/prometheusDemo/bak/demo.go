package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

// 定义指标
var (
	// 统计请求数量
	httpRequestCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "service",
			Name:      "http_request_total",
			Help:      "Total number of http_request",
		},
	)

	//prometheus.NewCounter与prometheus.NewCounterVec的区别
	//httpRequestCounter = prometheus.NewCounterVec(
	//   prometheus.CounterOpts{
	//      Subsystem: "service",
	//      Name:      "http_request_total",
	//      Help:      "Total number of http_request",
	//   },
	//   []string{"kind"}
	//)

	// 监控实时并发量（处理中的请求）
	concurrentHttpRequestsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Subsystem: "sdk",
			Name:      "http_handle_concurrent",
			Help:      "Number of incoming HTTP Requests handling concurrently now.",
		},
	)

	// 监控请求量，请求耗时等
	httpRequestsHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "sdk",
			Name:      "http_handle_requests",
			Help:      "Histogram statistics of http requests handle by elete http. Buckets by latency",
			Buckets:   []float64{0.001, 0.002, 0.005, 0.01, 0.05, 0.1, 0.2, 0.3, 0.4, 0.5, 0.8, 1, 2, 5, 10},
		},
		[]string{"code"},
	)

	summary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "test_summary",
			Help: "test summary",
			Objectives: map[float64]float64{
				0.5:  0.05,
				0.9:  0.01,
				0.99: 0.001,
			}, // 计算的分位数和对应的允许误差值
		},
		[]string{"name"},
	)
)

// 注册指标收集器
func init() {
	prometheus.MustRegister(dropRequestCounter)
	// prometheus.Register(dropRequestCounter)
	prometheus.MustRegister(concurrentHttpRequestsGauge)
	prometheus.MustRegister(httpRequestsHistogram)
	prometheus.MustRegister(summary)
}

func GinMetricsMid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 统计接口请求数量
		httpRequestCounter.Inc()

		// 监控并发量，进入接口前 +1
		concurrentHttpRequestsGauge.Inc()

		startTime := time.Now()

		// 处理后续逻辑
		ctx.Next()

		// after request
		finishTime := time.Now()

		// 监控计算接口耗时，请求数量等
		httpRequestsHistogram.With(prometheus.Labels{"code": strconv.Itoa(w2.StatusCode)}).Observe(float64(finishTime.Sub(startTime)) / (1000 * 1000 * 1000))

		// 监控并发量，离开接口 -1
		concurrentHttpRequestsGauge.Dec()
	}
}
