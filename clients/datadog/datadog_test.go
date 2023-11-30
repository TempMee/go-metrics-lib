package datadog

import (
	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataDogClient_CreateHistogram(t *testing.T) {
	a := assert.New(t)
	statsdClient, _ := statsd.New("localhost:8125")
	datadogClient := &DataDogClient{
		Client:     statsdClient,
		Histograms: map[string]*Histogram{},
	}

	_ = datadogClient.CreateHistogram("test", []float64{10, 20, 30}, map[string]string{"test": "test"}, 1.0)
	histogram, err := datadogClient.Histograms["test"].GenerateMetric(1, map[string]string{"test": "test"}, 1.0)
	a.NoError(err)
	a.Equal(histogram.labels["le"], "10")

}
