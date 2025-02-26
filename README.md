# Go Metrics Library

The **Go Metrics Library** simplifies the implementation of metrics and provides a standardized approach to metric usage
in Go applications. This library supports a range of metrics, offering flexibility and consistency in tracking
application performance and behavior.

---

## **Supported Standard Metrics**

| Metric                                                   | Labels                                                                                                                                                                                                  | Description                                                                                 |
|----------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| **resolver\_request\_duration\_histogram\_milliseconds** | `resolver` (function name),<br/> `result` (success/fail), <br/>`service` (service name), <br/>`protocol` (http/grpc/graphql) <br/> `env`: production\|staging\|development                              | Tracks the success/failure, duration, and duration distribution of resolver requests.       |
| **http\_request\_duration\_histogram\_milliseconds**     | `result` (success/fail), <br/>`service` (service name), <br/>`method` (HTTP method: POST, GET, PATCH, etc.) <br/> env: production\|staging\|development                                                 | Captures metrics for all HTTP requests to the service. (Datadog provides this by default.)  |
| **api\_request\_duration\_histogram\_milliseconds**      | `service` (current service), <br/>`vendor` (internal/external), <br/>`call` (query/function name), <br/>`result` (success/fail) <br/> `env`: production\|staging\|development                           | Tracks communication duration between services/vendors, source, destination, and result.    |
| **database\_query\_duration\_histogram\_milliseconds**   | `service` (service name), <br/>`result` (success/fail), <br/>`table` (table name), <br/>`method` (insert/delete/find), <br/>`database` (mongodb/postgres) <br/> `env`: production\|staging\|development | Measures query durations by service, database, and operation type.                          |
| **call\_duration\_histogram\_milliseconds**              | `service` (service name), <br/>`result` (success/fail), <br/>`function` (function name) <br/> `env`: production\|staging\|development                                                                   | Monitors the duration of specific function calls (used selectively for targeted observation). |
| **event\_process\_histogram\_duration\_milliseconds**    | `service` (service name), <br/>`queue` (queue name), <br/>`result` (success/fail) <br/> `env`: production\|staging\|development                                                                         | Tracks the duration of event processing by service, event, and result.                        |

---

## **Client Integrations**

### **Datadog Integration**

The library includes a Datadog client for integrating with Datadog’s monitoring system.

#### Key Features

- Histogram creation and value submission
- Support for default and custom buckets
- Label customization

#### Example Usage

```go
import (
"github.com/TempMee/go-metrics-lib/clients/datadog"
)

config := datadog.DataDogConfig{
    DD_AGENT_PORT: 8125,
    DD_AGENT_HOST: "localhost",
}

datadogClient := datadog.NewDatadogClient(config)

datadogClient.CreateHistogram("example.histogram", []float64{10, 20, 30}, map[string]string{"example": "label"}, 1.0)

datadogClient.Histogram("example.histogram", 15, map[string]string{"example": "label"}, 1.0)
```

---

### **Prometheus Integration**

A Prometheus client is also included for users leveraging Prometheus for metrics.

#### Key Features

- Support for Prometheus’s `HistogramVec`, `CounterVec`, `GaugeVec`, and `SummaryVec`
- Label-based metric tracking
- Integrated HTTP handler for exposing `/metrics`

#### Example Usage

```go
import (
"github.com/TempMee/go-metrics-lib/clients/prometheus"
)

prometheusClient := prometheus.NewPrometheusClient()

// Create a histogram vector
prometheusClient.CreateHistogramVec("example_histogram", "Example help text", []string{"label1"}, []float64{10, 20, 30})

// Record a value
prometheusClient.Histogram("example_histogram", 15, map[string]string{"label1": "value"}, 1.0)

// Start the Prometheus HTTP handler
go prometheusClient.ServerHandler()
```

---

## **Setting Up Metrics**

### **Create a Metric**

Define a new metric using the `CreateHistogram` method:

```go
datadogClient.CreateHistogram("graphql.resolver.millisecond", []float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, map[string]string{
    "resolver": "resolver",
    "service":  "graphql",
    "result":   "success",
}, 1)
```

### **Use a Metric**

Report a metric using the `HistogramMetric` method:

```go
metrics := MetricsLib.NewMetrics(datadogClient, 1)

err := metrics.HistogramMetric("graphql.resolver.millisecond", 100, map[string]string{
    "resolver": "resolver",
    "service":  "graphql",
    "result":   "success",
})
```

### **Using HTTP Metrics in Middleware**

You can track HTTP request durations and counts for both success and error scenarios using middleware:

```go
import (
    "net/http"
    "time"
    "github.com/TempMee/go-metrics-lib"
    "github.com/TempMee/go-metrics-lib/clients/datadog"
)

func MetricsMiddleware(metrics metrics_lib.MetricsImpl, serviceName string) func (http.Handler) http.Handler {
    return func (next http.Handler) http.Handler {
        return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            
            defer func() {
                duration := float64(time.Since(start).Milliseconds())
                result := "success"
                if w.WriteHeader != http.StatusOK {
                    result = "error"
                }

                metrics.HistogramMetric("http_request_duration_histogram_milliseconds", duration, map[string]string{
                    "service": serviceName,
                    "method":  r.Method,
                    "result":  result,
                })
            }()

            next.ServeHTTP(w, r)
        })
    }
}

func main() {
    // Configure Datadog client
    config := datadog.DataDogConfig{
        DD_AGENT_HOST: "localhost",
        DD_AGENT_PORT: 8125,
    }
    datadogClient := datadog.NewDatadogClient(config)
    metrics := metrics_lib.NewMetrics(datadogClient, 1)
    
    // Use middleware
    http.Handle("/example", MetricsMiddleware(metrics, "example_service")(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Success"))
    })))
    
    http.Handle("/example-error", MetricsMiddleware(metrics, "example_service")(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Error"))
    })))
    
    http.ListenAndServe(":8080", nil)
}
```

This middleware automatically tracks the duration of each HTTP request and records metrics for success and error
scenarios.

---

## **References**

### Prometheus Documentation

- [Metric Types](https://prometheus.io/docs/concepts/metric_types/)
- [Best Practices for Naming](https://prometheus.io/docs/practices/naming/)
- [Best Practices for Instrumentation](https://prometheus.io/docs/practices/instrumentation/)
- [Histograms and Summaries](https://prometheus.io/docs/practices/histograms/)
- [Alerting Practices](https://prometheus.io/docs/practices/alerting/)
- [Rule Practices](https://prometheus.io/docs/practices/rules/)

For detailed examples, refer to the `examples` folder in the repository.

