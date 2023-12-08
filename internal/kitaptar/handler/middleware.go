package handler

import (
	"errors"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/metrics"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
)

const (
	authUserID = "auth_user_id"
)

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")

		if authorizationHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "authorization header is not set"})
			return
		}
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "invalid header value"})
			return
		}
		if len(headerParts[1]) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "empty token"})
			return
		}
		userId, err := h.srvs.VerifyToken(headerParts[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "invalid token"})
			return
		}
		ctx.Set(authUserID, userId)
		ctx.Next()
	}

}

//nolint:all
func getUserId(c *gin.Context) (string, error) {
	idDirty, ok := c.Get(authUserID)
	if !ok {
		return "", errors.New("user id not found")
	}

	id, ok := idDirty.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return id, nil
}

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
