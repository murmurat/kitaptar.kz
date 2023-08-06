package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "one-lab/docs"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", h.createUser)
	router.POST("/login", h.loginUser)

	router.Use(h.authMiddleware())
	router.GET("/user/posts", h.userBooks)

	return router
}
