package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/murat96k/kitaptar.kz/internal/auth/docs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       kitaptar.kz Auth service
// @description Service for authorization for users
// @version     1.0
// @host        localhost:8080
// @BasePath    /
// @in header
// @name Authorization
// @contact.name   Meiirzhan Uristemov
// @contact.email  admin@kitaptar.kz
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//router.Use(HTTPMetrics())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.POST("/register", h.createUser)
	router.POST("/refresh", h.refresh)
	router.POST("/login", h.loginUser)

	return router
}
