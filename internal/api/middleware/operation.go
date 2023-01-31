package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"templateRestAPI/internal/common"
)

const operationKey = "operation"

func GetOperationName(c *gin.Context) string {
	return c.GetString(operationKey)
}

func SetOperationName(cfg common.ServerConfig, logger zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := fmt.Sprintf("%s %s", strings.ToLower(c.Request.Method), c.Request.RequestURI)
		operation := cfg.Operations[method]

		if operation == "" {
			operation = "unknown"
		}

		c.Set(operationKey, operation)
		logger.With(zap.String("trackingID", GetTrackingId(c))).Debug("operation is " + operation)

		c.Next()
	}
}
