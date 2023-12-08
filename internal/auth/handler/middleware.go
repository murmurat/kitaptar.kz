package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/internal/auth/metrics"
	"strconv"
	"time"
)

func HTTPMetrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()

		ctx.Next()

		elapsedSeconds := time.Since(now).Seconds()
		pattern := ctx.FullPath()
		method := ctx.Request.Method
		status := ctx.Writer.Status()

		metrics.HttpRequestsDurationHistorgram.WithLabelValues(pattern, method).Observe(elapsedSeconds)
		metrics.HttpRequestsDurationSummary.WithLabelValues(pattern, method).Observe(elapsedSeconds)
		metrics.HttpRequestsTotal.WithLabelValues(pattern, method, strconv.Itoa(status)).Inc()
	}
}
