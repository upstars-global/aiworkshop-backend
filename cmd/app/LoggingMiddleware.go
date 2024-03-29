package main

// Потрібно згенерувати middleware для логування швидкості виконання запитів

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// TimingMiddleware ...
func TimingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		start := time.Now()
		c.Next()
		// after request
		logger.Info("request processed",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.Duration("duration", time.Since(start)),
		)
	}
}
