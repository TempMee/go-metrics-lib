package datadog_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tempmee/go-metrics-lib/clients/datadog"
	"testing"
)

func TestHistogram_GenerateMetric(t *testing.T) {
	t.Run("Should generate metric", func(t *testing.T) {
		a := assert.New(t)

		metric := datadog.NewHistogram("test", []float64{10, 20, 30}, map[string]string{"test": "test"}, 1.0)
		val, err := metric.GenerateMetric(1, map[string]string{"test": "test"}, 1.0)

		a.NoError(err)
		a.Equal(val.Labels["le"], "10")

	})
	t.Run("Should generate metric", func(t *testing.T) {
		a := assert.New(t)

		metric := datadog.NewHistogram("test", []float64{10, 20, 30}, map[string]string{"test": "test"}, 1.0)
		val, err := metric.GenerateMetric(10, map[string]string{"test": "test"}, 1.0)

		a.NoError(err)
		a.Equal(val.Labels["le"], "10")

	})

	t.Run("Should generate metric", func(t *testing.T) {
		a := assert.New(t)

		metric := datadog.NewHistogram("test", []float64{10, 20, 30}, map[string]string{"test": "test"}, 1.0)
		val, err := metric.GenerateMetric(25, map[string]string{"test": "test"}, 1.0)

		a.NoError(err)
		a.Equal(val.Labels["le"], "30")

	})

}
