package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/murat96k/kitaptar.kz/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", h.createUser)
	router.POST("/login", h.loginUser)

	user := router.Group("/user", h.authMiddleware())
	{
		user.PUT("/update", h.updateUser)
		user.DELETE("/delete", h.deleteUser)
		user.GET("/data", h.getUser)
	}

	book := router.Group("/book", h.authMiddleware())
	{
		book.GET("/:id", h.getBookById)
		book.GET("/all", h.getAllBooks)
		book.POST("/create", h.createBook)
		book.DELETE("/delete/:id", h.deleteBook)
		book.PUT("/update/:id", h.updateBook)
		//book.GET("/user/books", h.userBooks)
	}

	author := router.Group("/author", h.authMiddleware())
	{
		author.GET("/:id", h.getAuthorById)
		author.GET("/all", h.getAllAuthors)
		author.POST("/create", h.createAuthor)
		author.DELETE("/delete/:id", h.deleteAuthor)
		author.PUT("/update/:id", h.updateAuthor)
	}

	filePath := router.Group("/file_path", h.authMiddleware())
	{
		filePath.GET("/:id", h.getFilePathById)
		filePath.GET("/all", h.getAllFilePaths)
		filePath.POST("/create", h.createFilePath)
		filePath.DELETE("/delete/:id", h.deleteFilePath)
		filePath.PUT("/update/:id", h.updateFilePath)
	}

	return router
}
