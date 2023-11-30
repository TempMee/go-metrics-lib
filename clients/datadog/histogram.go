package datadog

import (
	"fmt"
	"strings"
)

type Histogram struct {
	MetricName string
	Buckets    []float64
	labels     map[string]string
	rate       float64
}

func NewHistogram(metricName string, buckets []float64, labels map[string]string, rate float64) *Histogram {
	return &Histogram{
		MetricName: metricName,
		Buckets:    buckets,
		labels:     labels,
		rate:       rate,
	}
}

func (h *Histogram) GenerateMetric(value float64, labels map[string]string, rate float64) (Histogram, error) {
	// set le label related to value in buckets
	le := ""
	for _, bucket := range h.Buckets {
		if value <= bucket {
			// remove 0s from float
			le = strings.TrimRight(fmt.Sprintf("%f", bucket), "0")
			le = strings.TrimRight(le, ".")
			break
		}
	}
	if le == "" {
		le = "+Inf"
	}

	// set labels
	h.labels = labels
	h.labels["le"] = le
	h.rate = rate

	return *h, nil

}
