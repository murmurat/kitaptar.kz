package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const authUserID = "auth_user_id"

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorationHeader := ctx.GetHeader("authorization")
		if authorationHeader == "" {
			err := errors.New("authorization header is not set")
			fmt.Println("get auth header err %w", err)
			ctx.Status(http.StatusUnauthorized)
			panic(err)
			return
		}
		fields := strings.Fields(authorationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header incorrect format")
			fmt.Println("get auth header err %w", err)
			ctx.Status(http.StatusUnauthorized)
			panic(err)
			return
		}
		userId, err := h.srvs.VerifyToken(fields[1])
		if err != nil {
			fmt.Println("get auth header err %w", err)
			ctx.Status(http.StatusUnauthorized)
			panic(err)
			return
		}
		ctx.Set(authUserID, userId)
		ctx.Next()
	}

}

func getUserEmail(c *gin.Context) (string, error) {
	emailDirty, ok := c.Get(authUserID)
	if !ok {
		return "", errors.New("user email not found")
	}

	email, ok := emailDirty.(string)
	if !ok {
		return "", errors.New("user email is of invalid type")
	}

	return email, nil
}

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

func setNewEmail(c *gin.Context, email string) error {
	c.Set(authUserID, email)
	return nil
}
