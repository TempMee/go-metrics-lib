package metrics_lib

type Result = string

const (
	Success Result = "success"
	Error   Result = "error"
)

type ResolverMetricLabels struct {
	Resolver string
	Result   Result
}

func ResolverMetric(client Client, name string, value float64, labels ResolverMetricLabels) error {
	err := client.Histogram(name, value, map[string]string{
		"resolver": labels.Resolver,
		"result":   string(labels.Result),
	}, 1)

	return err
}
