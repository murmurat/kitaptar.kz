package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/auth/entity"
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
