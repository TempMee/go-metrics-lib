package datadog

import (
	"fmt"
	"log"
	"sync"

	"github.com/DataDog/datadog-go/v5/statsd"
)

type DataDogClient struct {
	Client     *statsd.Client
	Histograms map[string]*Histogram

	mu sync.Mutex
}

type DataDogConfig struct {
	DD_AGENT_HOST string `env:"DD_AGENT_HOST" default:"localhost"`
	DD_AGENT_PORT int    `env:"DD_AGENT_PORT" default:"8125"`
}

func NewDatadogClient(datadogConfig DataDogConfig) *DataDogClient {
	dogstatsd_client, err := statsd.New(fmt.Sprintf("%s:%d", datadogConfig.DD_AGENT_HOST, datadogConfig.DD_AGENT_PORT))
	if err != nil {
		log.Fatal(err)
	}

	if dogstatsd_client == nil {
		log.Fatal("dogstatsd_client is nil")
	}
	return &DataDogClient{
		dogstatsd_client,
		make(map[string]*Histogram),
		sync.Mutex{},
	}
}

// CreateHistogram creates a new histogram metric
// If the metric already exists, it will be ignored
func (d *DataDogClient) CreateHistogram(metric string, buckets []float64, labels map[string]string, rate float64) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.Histograms[metric]; ok {
		return
	}

	histogram := NewHistogram(metric, buckets, labels, rate)
	d.Histograms[metric] = histogram
}

// Histogram pushes a value to a histogram metric
// If the metric does not exist, it will be created with default buckets (0.0, 1.0)
func (d *DataDogClient) Histogram(metric string, value float64, labels map[string]string, rate float64) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.Histograms[metric]; !ok {
		histogram := NewHistogram(metric, []float64{0.0, 1.0}, labels, rate)
		d.Histograms[metric] = histogram
	}

	histogram, err := d.Histograms[metric].GenerateMetric(value, d.Histograms[metric].Labels, rate)
	if err != nil {
		return err
	}
	// merge labels
	for k, v := range labels {
		histogram.Labels[k] = v
	}
	tags := labelsToStringArray(histogram.Labels)
	err = d.Client.Histogram(histogram.MetricName, value, tags, rate)
	if err != nil {
		return err
	}

	err = d.Client.Distribution(metric, value, tags, rate)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataDogClient) Count(metric string, labels map[string]string, rate float64) error {
	tags := labelsToStringArray(labels)
	err := d.Client.Count(metric, 1, tags, rate)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataDogClient) Gauge(metric string, value float64, labels map[string]string, rate float64) error {
	tags := labelsToStringArray(labels)
	err := d.Client.Gauge(metric, value, tags, rate)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataDogClient) Summary(metric string, value float64, labels map[string]string, rate float64) error {
	log.Println("[Datadog] Summary is unsupported")
	return nil
}

func labelsToStringArray(labels map[string]string) []string {
	var tags []string
	for k, v := range labels {
		tags = append(tags, fmt.Sprintf("%s:%s", k, v))
	}
	return tags
}
