package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"templateRestAPI/internal/api/middleware"
)

func EnrichLogger(logger *zap.Logger, c *gin.Context) *zap.Logger {
	return logger.With(
		zap.String("trackingID", middleware.GetTrackingId(c)),
		zap.String("operation", middleware.GetOperationName(c)),
		zap.String("sessionID", middleware.GetSessionId(c)),
	)
}
