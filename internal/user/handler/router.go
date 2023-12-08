package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/murat96k/kitaptar.kz/internal/user/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       kitaptar.kz User service
// @description Service for users (Personal cabinet)
// @version     1.0
// @host        localhost:8081
// @BasePath    /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @contact.name   Meiirzhan Uristemov
// @contact.email  admin@kitaptar.kz
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := router.Group("/users", h.authMiddleware())
	{
		user.PUT("/", h.updateUser)
		user.DELETE("/", h.deleteUser)
		user.GET("/", h.getUser)
		user.POST("/confirm", h.confirmUser)
	}

	admin := router.Group("/admin", h.authMiddleware())
	{
		admin.GET("/users", h.getAllUsers)
		admin.POST("/users/:id", h.setUserRole)
	}

	return router
}
