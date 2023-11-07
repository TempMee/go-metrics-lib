mocks:
	go get go.uber.org/mock/mockgen/model
	go install go.uber.org/mock/mockgen@latest
	$GOBIN/mockgen -destination=./mocks/metrics_impl.go -package=mocks github.com/tempmee/go-metrics-lib MetricsImpl
	$GOBIN/mockgen -destination=./mocks/metrics_client.go -package=mocks github.com/tempmee/go-metrics-lib Client
