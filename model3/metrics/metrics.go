package metrics

import (
	"time"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	MetricsNamespace = "httpserver"
)

type ExecutionTimer struct {
	histogram *prometheus.HistogramVec
	start     time.Time
	last      time.Time
}

var (
	functionLatency = CreateExecutionTimeMetric(MetricsNamespace, "Time spent.")
)

func Register() {
	err := prometheus.Register(functionLatency)
	if err != nil {
		glog.Error("register metrics err")
	}
}

func NewTime() *ExecutionTimer {
	return NewExecutionTimer(functionLatency)
}

func NewExecutionTimer(histogram *prometheus.HistogramVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histogram: histogram,
		start:     now,
		last:      now,
	}
}

func (t *ExecutionTimer) ObserveTotal() {
	(*t.histogram).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "execution_latency_seconds",
			Help:      help,
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"step"},
	)
}
