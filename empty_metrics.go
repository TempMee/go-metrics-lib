package metrics_lib

import (
	"time"

	"github.com/gin-gonic/gin"
)

type emptyMetrics struct{}

func NewEmptyMetrics() MetricsImpl {
	return &emptyMetrics{}
}

func (m *emptyMetrics) HistogramMetric(name string, value float64, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) CountMetric(name string, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) GaugeMetric(name string, value float64, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) SummaryMetric(name string, value float64, labels map[string]string) error {
	return nil
}

func (m *emptyMetrics) ResolverMetric(value float64, labels ResolverMetricLabels) error {
	return nil
}

func (m *emptyMetrics) HttpMiddlewareMetric(serviceName string) gin.HandlerFunc {
	return nil
}

func (m *emptyMetrics) ApiMetric(value float64, labels ApiMetricLabels) error {
	return nil
}

func (m *emptyMetrics) ApiMetricDuration(startTime time.Time, labels ApiMetricLabels, err error) error {
	return nil
}

func (m *emptyMetrics) DatabaseMetric(value float64, labels DatabaseMetricLabels) error {
	return nil
}

func (m *emptyMetrics) CallMetric(value float64, labels CallMetricLabels) error {
	return nil
}

func (m *emptyMetrics) EventProcessMetric(value float64, labels EventProcessMetricLabels) error {
	return nil
}

func (m *emptyMetrics) EventPublishMetric(labels EventPublishMetricLabels) error {
	return nil
}
