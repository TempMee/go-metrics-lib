package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type PrometheusClient struct {
	HistogramVecs map[string]*prometheus.HistogramVec
	CounterVecs   map[string]*prometheus.CounterVec
	GaugeVecs     map[string]*prometheus.GaugeVec
	SummaryVecs   map[string]*prometheus.SummaryVec
}

func NewPrometheusClient() *PrometheusClient {
	return &PrometheusClient{}
}

func (p *PrometheusClient) ServerHandler() {
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(":2112", nil)
}

func (p *PrometheusClient) CreateHistogramVec(name string, help string, labelNames []string, buckets []float64) error {
	if p.HistogramVecs == nil {
		p.HistogramVecs = make(map[string]*prometheus.HistogramVec)
	}

	if _, ok := p.HistogramVecs[name]; ok {
		return nil
	}
	p.HistogramVecs[name] = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    name,
		Help:    help,
		Buckets: buckets,
	}, labelNames)
	return nil
}

func (p *PrometheusClient) CreateCounterVec(name string, help string, labelNames []string) error {
	if p.CounterVecs == nil {
		p.CounterVecs = make(map[string]*prometheus.CounterVec)
	}

	if _, ok := p.CounterVecs[name]; ok {
		return nil
	}
	p.CounterVecs[name] = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: name,
		Help: help,
	}, labelNames)
	return nil
}

func (p *PrometheusClient) CreateGaugeVec(name string, help string, labelNames []string) error {
	if p.GaugeVecs == nil {
		p.GaugeVecs = make(map[string]*prometheus.GaugeVec)
	}

	if _, ok := p.GaugeVecs[name]; ok {
		return nil
	}
	p.GaugeVecs[name] = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	}, labelNames)
	return nil
}

func (p *PrometheusClient) CreateSummaryVec(name string, help string, labelNames []string) error {
	if p.SummaryVecs == nil {
		p.SummaryVecs = make(map[string]*prometheus.SummaryVec)
	}

	if _, ok := p.SummaryVecs[name]; ok {
		return nil
	}
	p.SummaryVecs[name] = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: name,
		Help: help,
	}, labelNames)
	return nil
}

func (p *PrometheusClient) Histogram(name string, value float64, labels map[string]string, rate float64) error {
	labelNames := make([]string, 0, len(labels))
	labelValues := make([]string, 0, len(labels))
	for k, v := range labels {
		labelNames = append(labelNames, k)
		labelValues = append(labelValues, v)
	}
	if _, ok := p.HistogramVecs[name]; !ok {
		_ = p.CreateHistogramVec(name, "", labelNames, nil)
	}
	p.HistogramVecs[name].WithLabelValues(labelValues...).Observe(value)
	return nil
}

func (p *PrometheusClient) Count(name string, labels map[string]string, rate float64) error {
	labelNames := make([]string, 0, len(labels))
	labelValues := make([]string, 0, len(labels))
	for k, v := range labels {
		labelNames = append(labelNames, k)
		labelValues = append(labelValues, v)
	}
	if _, ok := p.CounterVecs[name]; !ok {
		_ = p.CreateCounterVec(name, "", labelNames)
	}
	p.CounterVecs[name].WithLabelValues(labelValues...).Inc()
	return nil
}

func (p *PrometheusClient) Gauge(name string, value float64, labels map[string]string, rate float64) error {
	labelNames := make([]string, 0, len(labels))
	labelValues := make([]string, 0, len(labels))
	for k, v := range labels {
		labelNames = append(labelNames, k)
		labelValues = append(labelValues, v)
	}
	if _, ok := p.GaugeVecs[name]; !ok {
		_ = p.CreateGaugeVec(name, "", labelNames)
	}
	p.GaugeVecs[name].WithLabelValues(labelValues...).Set(value)
	return nil
}

func (p *PrometheusClient) Summary(name string, value float64, labels map[string]string, rate float64) error {
	labelNames := make([]string, 0, len(labels))
	labelValues := make([]string, 0, len(labels))
	for k, v := range labels {
		labelNames = append(labelNames, k)
		labelValues = append(labelValues, v)
	}
	if _, ok := p.SummaryVecs[name]; !ok {
		_ = p.CreateSummaryVec(name, "", labelNames)
	}
	p.SummaryVecs[name].WithLabelValues(labelValues...).Observe(value)
	return nil
}
