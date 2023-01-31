package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"templateRestAPI/internal/common"
)

func SetAuthorizationCheck(cfg common.ServerConfig, logger zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.With(zap.String("trackingID", GetTrackingId(c))).Debug("auth is not implemented")
		c.Next()
	}
}
