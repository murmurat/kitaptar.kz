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
	router.GET("/user/books", h.userBooks)          // not realised
	router.GET("/book/:id", h.getBookById)          // checked
	router.GET("/books", h.getAllBooks)             // checked
	router.POST("/book/create", h.createBook)       // checked
	router.DELETE("/book/delete/:id", h.deleteBook) // checked
	router.PUT("/book/update/:id", h.updateBook)    // checked

	router.GET("/author/:id", h.getAuthorById)          //checked
	router.GET("/authors", h.getAllAuthors)             //checked
	router.POST("/author/create", h.createAuthor)       //checked
	router.DELETE("/author/delete/:id", h.deleteAuthor) //checked
	router.PUT("/author/update/:id", h.updateAuthor)    //checked

	router.GET("/file_path/:id", h.getFilePathById)          //checked
	router.GET("/file_paths", h.getAllFilePaths)             //checked
	router.POST("/file_path/create", h.createFilePath)       //checked
	router.DELETE("/file_path/delete/:id", h.deleteFilePath) //checked
	router.PUT("/file_path/update/:id", h.updateFilePath)    //checked

	router.PUT("/user/update", h.updateUser)
	router.DELETE("/user/delete", h.deleteUser)
	router.GET("/user/data", h.getUser)

	return router
}
