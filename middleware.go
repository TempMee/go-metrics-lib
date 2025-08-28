package metrics_lib

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
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
				zerolog.Ctx(c).Warn().Err(err).Msg("failed to record http request duration")
			}
		}()

		c.Next()
	}
}
