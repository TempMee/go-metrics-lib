package metrics_lib

import (
	"net/http"
)

type StandardMetrics interface {
	ResolverMetric(name string, value float64, labels ResolverMetricLabels) error
	HttpMiddlewareMetric(config HttpMiddlewareMetricConfig) func(http.Handler) http.Handler
}
