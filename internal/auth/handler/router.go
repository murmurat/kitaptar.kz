package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/murat96k/kitaptar.kz/docs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(HTTPMetrics())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", h.createUser)
	router.POST("/refresh", h.refresh)
	router.POST("/login", h.loginUser)

	return router
}
