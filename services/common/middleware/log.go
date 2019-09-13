package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mewil/portal/common/logger"
)

func LogMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				log.Error(e)
			}
		} else {
			log.Infof(
				"%s %s %s %d %s %s %s %s",
				end.UTC().Format(time.RFC3339Nano),
				c.Request.Method,
				path,
				c.Writer.Status(),
				latency,
				query,
				c.ClientIP(),
				c.Request.UserAgent(),
			)
		}
	}
}
