package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/murat96k/kitaptar.kz/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := router.Group("/user", h.authMiddleware())
	{
		user.PUT("/update", h.updateUser)
		user.DELETE("/delete", h.deleteUser)
		user.GET("/data", h.getUser)
		user.POST("/confirm", h.confirmUser)
	}

	return router
}
