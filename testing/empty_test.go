package testing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	metrics_lib "github.com/TempMee/go-metrics-lib"
)

func Test_EmptyMetrics(t *testing.T) {
	metrics := NewEmptyMetrics()

	assert.NoError(t, metrics.HistogramMetric("test", 100, map[string]string{"test": "test"}))

	assert.NoError(t, metrics.CountMetric("test", map[string]string{"test": "test"}))

	assert.NoError(t, metrics.GaugeMetric("test", 100, map[string]string{"test": "test"}))

	assert.NoError(t, metrics.SummaryMetric("test", 100, map[string]string{"test": "test"}))

	assert.NoError(t, metrics.ResolverMetric(100, metrics_lib.ResolverMetricLabels{
		Resolver: "test",
		Result:   metrics_lib.Success,
	}))

	assert.Nil(t, metrics.HttpMiddlewareMetric("test"))

	assert.Nil(t, metrics.PlainHttpMiddlewareMetric(metrics_lib.HttpMiddlewareMetricConfig{
		Service: "test",
	}))

	assert.NoError(t, metrics.ApiMetric(100, metrics_lib.ApiMetricLabels{
		Service: "test",
		Vendor:  "test",
		Call:    "test",
		Result:  metrics_lib.Success,
	}))

	assert.NoError(t, metrics.ApiMetricDuration(time.Now(), metrics_lib.ApiMetricLabels{
		Service: "test",
		Vendor:  "test",
		Call:    "test",
		Result:  metrics_lib.Success,
	}, nil))

	assert.NoError(t, metrics.DatabaseMetric(100, metrics_lib.DatabaseMetricLabels{
		Service: "test",
		Table:   "test",
		Method:  metrics_lib.DatabaseMetricMethodSelect,
		Result:  metrics_lib.Success,
	}))

	assert.NoError(t, metrics.CallMetric(100, metrics_lib.CallMetricLabels{
		Service:  "test",
		Function: "test",
		Result:   metrics_lib.Success,
	}))

	assert.NoError(t, metrics.EventProcessMetric(100, metrics_lib.EventProcessMetricLabels{
		Service: "test",
		Queue:   "test",
		Result:  metrics_lib.Success,
	}))

	assert.NoError(t, metrics.EventPublishMetric(metrics_lib.EventPublishMetricLabels{
		Service: "test",
		Queue:   "test",
		Result:  metrics_lib.Success,
	}))
}
