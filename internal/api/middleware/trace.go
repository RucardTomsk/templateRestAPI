package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const trackingIdHeader = "X-Tracking-Id"
const sessionIdHeader = "X-Session-Id"

func GetTrackingId(c *gin.Context) string {
	return c.GetString(trackingIdHeader)
}

func GetTrackingIdHeader() string {
	return trackingIdHeader
}

func GetSessionId(c *gin.Context) string {
	return c.GetString(sessionIdHeader)
}

func GetSessionIdHeader() string {
	return sessionIdHeader
}

func SetTracingContext(logger zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		tId := c.GetHeader(trackingIdHeader)
		sId := c.GetHeader(sessionIdHeader)

		if tId == "" {
			tId = uuid.New().String()
			c.Header(trackingIdHeader, tId)
			logger.Debug(fmt.Sprintf("trackingID %s generated", tId))
		}

		if sId == "" {
			sId = uuid.New().String()
			c.Header(sessionIdHeader, sId)
			logger.Debug(fmt.Sprintf("sessionID %s generated", sId))
		}

		c.Set(trackingIdHeader, tId)
		c.Set(sessionIdHeader, sId)

		c.Next()
	}
}
