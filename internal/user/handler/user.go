package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
	"net/http"
)

func (h *Handler) updateUser(ctx *gin.Context) {

	var req api.UpdateUserRequest

	userID, err := getUserId(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	if req == (api.UpdateUserRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{"Update user data not provided"})
		return
	}
	err = h.srvs.UpdateUser(ctx, userID, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{"User data updated!"})
}

func (h *Handler) getUser(ctx *gin.Context) {

	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	user, err := h.srvs.GetUserById(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, user)
}

func (h *Handler) deleteUser(ctx *gin.Context) {

	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	err = h.srvs.DeleteUser(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{"User deleted"})
}
