package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/murat96k/kitaptar.kz/internal/kitaptar/docs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       kitaptar.kz Kitaptar service
// @description Service for books, authors...
// @version     1.0
// @host        localhost:8082
// @BasePath    /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @contact.name   Meiirzhan Uristemov
// @contact.email  admin@kitaptar.kz
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(HTTPMetrics())

	book := router.Group("/books", h.authMiddleware())
	{
		book.GET("/:id", h.getBookById)
		book.GET("", h.getAllBooks)
		book.GET("/search/:name", h.getBookByName)
		book.POST("", h.createBook)
		book.DELETE("/:id", h.deleteBook)
		book.PUT("/:id", h.updateBook)
		//book.GET("/user/books", h.userBooks)

		favorites := book.Group("/favorites")
		{
			favorites.GET("/:id", h.getFromFavorites)
			favorites.POST("/:id", h.addToFavorites)
			favorites.DELETE("/:id", h.deleteFromFavorites)
		}
	}

	author := router.Group("/authors", h.authMiddleware())
	{
		author.GET("/:id", h.getAuthorById)
		author.GET("", h.getAllAuthors)
		author.GET("/search/:name", h.getAuthorByName)
		author.POST("", h.createAuthor)
		author.DELETE("/:id", h.deleteAuthor)
		author.PUT("/:id", h.updateAuthor)
	}

	filePath := router.Group("/file_paths", h.authMiddleware())
	{
		filePath.GET("/:id", h.getFilePathById)
		filePath.GET("", h.getAllFilePaths)
		filePath.POST("", h.createFilePath)
		filePath.DELETE("/:id", h.deleteFilePath)
		filePath.PUT("/:id", h.updateFilePath)
	}

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return router
}
