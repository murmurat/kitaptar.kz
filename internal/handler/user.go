package handler

import (
	"fmt"
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
	log.Println("Access Token: ", accessToken)
	ctx.Status(http.StatusOK)
}

func (h *Handler) userBooks(ctx *gin.Context) {
	email, ok := ctx.MustGet(authUserID).(string)
	if !ok {
		log.Printf("can't get user email on userBooks")
		ctx.Status(http.StatusBadRequest)
		return
	}

	// logic
	fmt.Println("Email of book owner user: ", email)
	ctx.Status(http.StatusOK)
}
