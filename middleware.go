package metrics_lib

import (
	"fmt"
	"time"

	"github.com/TempMee/x/logging"
	"github.com/gin-gonic/gin"
)

func httpMetricMiddleware(client Client, serviceName string, rate float64) gin.HandlerFunc {
	return func(c *gin.Context) {
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
					"service": serviceName,      // service name
					"method":  c.Request.Method, // method
					"path":    c.FullPath(),     // path
					"result":  result,           // result
				},
				rate, // rate
			); err != nil {
				logger := logging.FromContext(c)
				logger.Warn().Err(err).Msg("failed to record http request duration")
			}

			if err := client.Count(
				"http_request_count", // metric name
				map[string]string{
					"service":     serviceName,                          // service name
					"method":      c.Request.Method,                     // method
					"path":        c.FullPath(),                         // path
					"status_code": fmt.Sprintf("%d", c.Writer.Status()), // status code
					"result":      result,                               // result
				},
				rate, // rate
			); err != nil {
				logger := logging.FromContext(c)
				logger.Warn().Err(err).Msg("failed to record http request count")
			}
		}()

		c.Next()
	}
}
