package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"net/http"
)

// createUser godoc
// @Summary      Create user
// @Description  Create new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param req body entity.User true "req body"
// @Success      201
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /register [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var req entity.User

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{"invalid input body"})
		return
	}

	userId, err := h.srvs.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, api.Response{Message: userId})
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{"invalid input body"})
		return
	}

	accessToken, err := h.srvs.Login(ctx, req.Email, req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, accessToken)
}

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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{"something went wrong"})
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
	user, err := h.srvs.GetUser(ctx, userID)
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
