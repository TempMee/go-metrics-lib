package metrics_lib

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type HttpMiddlewareMetricConfig struct {
	Service string
}

func HttpMiddlewareMetric(client Client, config HttpMiddlewareMetricConfig, rate float64) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return measurementHandler(client, config, rate, h)
	}
}

func measurementHandler(client Client, config HttpMiddlewareMetricConfig, rate float64, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()

		defer func() {
			elasped := time.Since(startTime).Milliseconds()
			err := client.Histogram("http_request_duration_histogram_milliseconds", float64(elasped),
				map[string]string{
					"service": config.Service,
					"method":  r.Method,
					"result":  "success",
				}, rate)
			if err != nil {
				log.Fatal(err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func ginHttpMetricMiddleware(client Client, serviceName string, rate float64) gin.HandlerFunc {
	return func(c *gin.Context) {
		// allow health endpoint to skip metrics
		if c.Request.URL.Path == "/health" {
			c.Next()
			return
		}

		start := time.Now()

		defer func() {
			elapsed := time.Since(start).Milliseconds()
			result := ResultSuccess
			if c.Writer.Status() < 200 || c.Writer.Status() >= 300 {
				result = ResultError
			}

			if err := client.Histogram(
				"http_request_duration_histogram_milliseconds", // metric name
				float64(elapsed), // elapsed time in milliseconds
				map[string]string{
					"service":     serviceName,      // service name
					"method":      c.Request.Method, // method
					"path":        c.FullPath(),     // path
					"result":      result,           // result
					"status_code": strconv.Itoa(c.Writer.Status()),
				},
				rate, // rate
			); err != nil {
				zerolog.Ctx(c).Warn().Err(err).Msg("failed to record http request duration")
			}
		}()

		c.Next()
	}
}
