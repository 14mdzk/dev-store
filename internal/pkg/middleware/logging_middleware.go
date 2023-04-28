package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		requestMethod := ctx.Request.Method
		reqURL := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		log.WithFields(log.Fields{
			"latency_time":   latencyTime,
			"request_method": requestMethod,
			"req_url":        reqURL,
			"status_code":    statusCode,
			"client_ip":      clientIP,
		}).Info("http request")

		ctx.Next()
	}
}
