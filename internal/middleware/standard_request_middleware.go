package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"otp-auth/internal/constants"
	"otp-auth/internal/utils"
)

func StandardRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract request-id from request header
		requestId := c.GetHeader(constants.XRequestId)

		// generate new request-id if not found
		if requestId == "" {
			requestId = uuid.New().String()
		}

		// set the request-id in context so that it can be used later
		c.Set(constants.XRequestId, requestId)

		// Before calling next, set the response header with the request-id
		c.Writer.Header().Set(constants.XRequestId, requestId)

		// log the beginning of the request
		utils.LogInfo("Starting request with X-Request-Id", logrus.Fields{
			"request_id": requestId,
			"path":       c.Request.URL.Path,
			"method":     c.Request.Method,
		})

		// proceed with the request handling
		c.Next()

		// log after the request is processed
		utils.LogInfo("Completed request with X-Request-Id", logrus.Fields{
			"request_id": requestId,
			"status":     c.Writer.Status(),
		})
	}
}
