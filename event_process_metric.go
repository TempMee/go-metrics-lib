package metrics_lib

type EventProcessMetricLabels struct {
	Service string
	Queue   string
	Result  Result
}

func EventProcessMetric(client Client, value float64, labels EventProcessMetricLabels) error {
	err := client.Histogram("call_duration_histogram_milliseconds", value, map[string]string{
		"service": labels.Service,
		"queue":   labels.Queue,
		"result":  labels.Result,
	}, 1)

	return err
}
