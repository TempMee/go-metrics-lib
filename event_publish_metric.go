package metrics_lib

type EventPublishMetricLabels struct {
	Service string
	Queue   string
	Result  Result
}

func EventPublishMetric(client Client, labels EventProcessMetricLabels) error {
	err := client.Count("event_publish_count", map[string]string{
		"service": labels.Service,
		"queue":   labels.Queue,
		"result":  labels.Result,
	}, 1)

	return err
}
