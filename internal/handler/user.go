package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"one-lab/api"
	"one-lab/internal/entity"
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
	//var req api.RegisterRequest
	var req entity.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, api.Error{err.Error()})
		return
	}

	//u := &entity.User{
	//	Email:     req.Email,
	//	FirstName: req.FirstName,
	//	LastName:  req.LastName,
	//	Password:  req.Password,
	//}

	err = h.srvs.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}
	//ctx.JSON(http.StatusCreated, req)
	ctx.Status(http.StatusCreated)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		return
	}

	accessToken, err := h.srvs.Login(ctx, req.Email, req.Password)
	if err != nil {
		log.Printf("srvs login send err: %s", err.Error())
		return
	}
	//log.Println("Access Token: ", accessToken)
	//ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, accessToken)
}

func (h *Handler) updateUser(ctx *gin.Context) {
	var req api.UpdateUserRequest
	//id := ctx.Params.ByName("id")
	userID, err := getUserId(ctx)
	//email, err := getUserEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	err = h.srvs.UpdateUser(ctx, userID, &req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
	}
	//var person Person
	//id := c.Params.ByName(“id”)
	//if err := db.Where(“id = ?”, id).First(&person).Error; err != nil {
	//	c.AbortWithStatus(404)
	//	fmt.Println(err)
	//}
	//c.BindJSON(&person)
	//db.Save(&person)
	//c.JSON(200, person)
}

func (h *Handler) getUser(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	user, err := h.srvs.GetUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusFound, user)

}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	err = h.srvs.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, "User deleted")
}
