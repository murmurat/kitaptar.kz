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
	router.GET("/user/books", h.userBooks)
	router.GET("/book/:id", h.getBookById)
	router.GET("/books", h.getAllBooks)
	router.POST("/book/create", h.createBook)
	router.DELETE("/book/delete", h.deleteBook)
	router.PUT("/book/update/:id", h.updateBook)

	router.GET("/author/:id", h.getAuthorById)
	router.GET("/authors", h.getAllAuthors)
	router.POST("/author/create", h.createAuthor)
	router.DELETE("/author/delete", h.deleteAuthor)
	router.PUT("/author/update/:id", h.updateAuthor)

	router.PUT("/user/update", h.updateUser)
	router.DELETE("/user/delete", h.deleteUser)
	router.GET("/user/data", h.getUser)

	return router
}
