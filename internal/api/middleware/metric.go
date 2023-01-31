package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetMetrics(logger zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.With(zap.String("trackingID", GetTrackingId(c))).Debug("metrics are not implemented")
		c.Next()
	}
}
