package testing

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	metrics_lib "github.com/TempMee/go-metrics-lib"
)

type emptyMetrics struct{}

func NewEmptyMetrics() metrics_lib.MetricsImpl {
	return &emptyMetrics{}
}

func (m *emptyMetrics) HistogramMetric(name string, value float64, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) CountMetric(name string, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) CountMetricWithValue(name string, value int64, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) GaugeMetric(name string, value float64, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) SummaryMetric(name string, value float64, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) ResolverMetric(value float64, labels metrics_lib.ResolverMetricLabels) error {
	return nil
}

func (m *emptyMetrics) HttpMiddlewareMetric(config metrics_lib.HttpMiddlewareMetricConfig) func(http.Handler) http.Handler {
	return nil
}

func (m *emptyMetrics) GinHttpMiddlewareMetric(serviceName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func (m *emptyMetrics) ApiMetric(value float64, labels metrics_lib.ApiMetricLabels) error {
	return nil
}

func (m *emptyMetrics) ApiMetricDuration(startTime time.Time, labels metrics_lib.ApiMetricLabels, err error) error {
	return nil
}

func (m *emptyMetrics) DatabaseMetric(value float64, labels metrics_lib.DatabaseMetricLabels) error {
	return nil
}

func (m *emptyMetrics) CallMetric(value float64, labels metrics_lib.CallMetricLabels) error {
	return nil
}

func (m *emptyMetrics) EventProcessMetric(value float64, labels metrics_lib.EventProcessMetricLabels) error {
	return nil
}

func (m *emptyMetrics) EventPublishMetric(labels metrics_lib.EventPublishMetricLabels) error {
	return nil
}
