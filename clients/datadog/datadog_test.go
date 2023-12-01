package datadog

import (
	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataDogClient_CreateHistogram(t *testing.T) {
	t.Run("Should create histogram", func(t *testing.T) {
		a := assert.New(t)
		statsdClient, _ := statsd.New("localhost:8125")
		datadogClient := &DataDogClient{
			Client:     statsdClient,
			Histograms: map[string]*Histogram{},
		}

		err := datadogClient.CreateHistogram("test", []float64{10, 20, 30}, map[string]string{"test": "test"}, 1.0)
		a.NoError(err)
		histogram, err := datadogClient.Histograms["test"].GenerateMetric(1, map[string]string{"test": "test"}, 1.0)
		a.NoError(err)
		a.Equal(histogram.labels["le"], "10")
	})

	t.Run("Should use latest tags", func(t *testing.T) {
		a := assert.New(t)
		statsdClient, _ := statsd.New("localhost:8125")
		datadogClient := &DataDogClient{
			Client:     statsdClient,
			Histograms: map[string]*Histogram{},
		}

		err := datadogClient.CreateHistogram("test", []float64{10, 20, 30}, map[string]string{"test": "test"}, 1.0)
		a.NoError(err)

		histogram, err := datadogClient.Histograms["test"].GenerateMetric(1, map[string]string{"test": "test2"}, 1.0)
		a.NoError(err)
		a.Equal(histogram.labels["le"], "10")
		a.Equal(histogram.labels["test"], "test2")
	})

}
