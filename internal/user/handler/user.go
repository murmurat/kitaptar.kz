package handler

import (
	"fmt"
	"github.com/murat96k/kitaptar.kz/internal/user/handler/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
)

// updateUser godoc
// @Summary      Update user
// @Security	 ApiKeyAuth
// @Description  Update existing user data
// @Tags         user
// @Accept       json
// @Produce      json
// @Param req body api.UpdateUserRequest true "req body"
// @Success      200
// @Failure      400  {object}  api.Error
// @Failure      401  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /users [put]
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update user data not provided"})
		return
	}
	err = h.srvs.UpdateUser(ctx, userID, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "User data updated!"})
}

// getUser godoc
// @Summary      Get user
// @Security	 ApiKeyAuth
// @Description  Get existing user data
// @Tags         user
// @Produce      json
// @Success      302
// @Failure      401  {object}  error
// @Failure      500  {object}  api.Error
// @Router       /users [get]
func (h *Handler) getUser(ctx *gin.Context) {

	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	user, err := h.srvs.GetUserById(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, user)
}

// deleteUser godoc
// @Summary      Delete user
// @Security	 ApiKeyAuth
// @Description  Delete existing user data
// @Tags         user
// @Produce      json
// @Success      200  {object}  api.Response
// @Failure      401  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /users [delete]
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

	ctx.JSON(http.StatusOK, api.Response{Message: "User deleted"})
}

// confirmUser godoc
// @Summary      Confirm user by email
// @Security	 ApiKeyAuth
// @Description  Confirm registered user by code which send to email
// @Tags         user
// @Accept       json
// @Produce      json
// @Param req body dto.CodeInput true "req body"
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      401  {object}  error
// @Failure      500  {object}  api.Error
// @Router       /users/confirm [post]
func (h *Handler) confirmUser(ctx *gin.Context) {

	var code dto.CodeInput

	if err := ctx.BindJSON(&code); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "invalid input body"})
		return
	}

	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	err = h.srvs.ConfirmUser(ctx, userID, code.Code)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: fmt.Sprintf("confirm user error: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "User confirmed successfully!"})
}

// getAllUsers godoc
// @Summary      Get all users
// @Security	 ApiKeyAuth
// @Description  Get all users
// @Tags         admin
// @Produce      json
// @Success      200  {object}  []entity.User
// @Failure      401  {object}  error
// @Failure      500  {object}  api.Error
// @Router       /admin/users [get]
func (h *Handler) getAllUsers(ctx *gin.Context) {

	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	users, err := h.srvs.GetAllUsers(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// setUserRole godoc
// @Summary      Set user role by admin
// @Security	 ApiKeyAuth
// @Description  Set user role by admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param req body dto.RoleInput true "request"
// @Param id path string true "id"
// @Success      200  {object}  api.Response
// @Failure      401  {object}  api.Error
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /admin/users/{id} [post]
func (h *Handler) setUserRole(ctx *gin.Context) {

	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	targetUserId := ctx.Param("id")
	if targetUserId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "user id is empty"})
		return
	}

	var req dto.RoleInput

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	err = h.srvs.SetUserRoleById(ctx, userID, targetUserId, req.Role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: fmt.Sprintf("User role updated to %s successfully", req.Role)})
}
