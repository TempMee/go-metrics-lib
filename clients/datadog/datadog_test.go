package datadog

import (
	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataDogClient_CountWithValue(t *testing.T) {
	client := NewDatadogClient(DataDogConfig{
		DD_AGENT_HOST: "localhost",
		DD_AGENT_PORT: 8125,
	})

	t.Run("Should count with custom value", func(t *testing.T) {
		a := assert.New(t)

		err := client.CountWithValue("test_count", 5, map[string]string{"test": "test"}, 1.0)
		a.NoError(err)
	})
}

func TestDataDogClient_CreateHistogram(t *testing.T) {
	statsdClient, _ := statsd.New("localhost:8125")
	datadogClient := &DataDogClient{
		Client:     statsdClient,
		Histograms: map[string]*Histogram{},
	}

	datadogClient.CreateHistogram("test", []float64{10, 20, 30}, map[string]string{"test": "test"}, 1.0)

	t.Run("Should create histogram", func(t *testing.T) {
		a := assert.New(t)

		histogram, err := datadogClient.Histograms["test"].GenerateMetric(1, map[string]string{"test": "test"}, 1.0)
		a.NoError(err)
		a.Equal(histogram.Labels["le"], "10")
	})

	t.Run("Should use latest tags", func(t *testing.T) {
		a := assert.New(t)

		histogram, err := datadogClient.Histograms["test"].GenerateMetric(1, map[string]string{"test": "test2"}, 1.0)
		a.NoError(err)
		a.Equal(histogram.Labels["le"], "10")
		a.Equal(histogram.Labels["test"], "test2")

		client := NewDatadogClient(DataDogConfig{
			DD_AGENT_HOST: "localhost",
			DD_AGENT_PORT: 8125,
		})

		client.CreateHistogram("test", []float64{10, 20, 30}, map[string]string{"test": "test"}, 1.0)

		err = client.Histogram("test", 1, map[string]string{"test": "test2"}, 1.0)

		a.NoError(err)

	})

}
