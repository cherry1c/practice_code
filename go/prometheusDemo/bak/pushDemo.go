package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	pusher *push.Pusher

	httpRequestCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "service",
			Name:      "http_request_total",
			Help:      "Total number of http_request",
		},
	)

	// 统计请求数量
	httpRequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "service",
			Name:      "http_request_total",
			Help:      "Total number of http_request",
		},
		[]string{"kind"},
	)

	// 监控实时并发量（处理中的请求）
	concurrentHttpRequestsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Subsystem: "sdk",
			Name:      "http_handle_concurrent",
			Help:      "Number of incoming HTTP Requests handling concurrently now.",
		},
	)

	// 监控请求量，请求耗时等
	concurrentHttpRequestsGauge = prometheus.NewHistogramVec(
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

	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
)

func main() {

	pusher = push.New("http://pushgateway:9091", "job_name") // 初始化一个pusher
	// 为pusher添加一些grouping key
	pusher.Grouping("service", "live_backend_go").Grouping("host", "localhost")

	// 向pusher中注册一个metric收集器
	pusher.Collector(completionTime)

	// 向puser中注册多个meterics
	registry := prometheus.NewRegistry()                                                                         // 向创建一个自定义的register
	registry.MustRegister(httpRequestCounter, concurrentHttpRequestsGauge, concurrentHttpRequestsGauge, summary) // 向register中注册多个meterics

	// 将register添加进pusher
	pusher.Gatherer(registry)

	// 将各metrics中的指标推送至push gateway
	pusher.Push() // 使用http的PUT方法
	pusher.Add()  // 使用http的POST方法
}
