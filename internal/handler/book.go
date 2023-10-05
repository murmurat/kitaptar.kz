package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
	"net/http"
)

// not ready
//func (h *Handler) userBooks(ctx *gin.Context) {
//	email, ok := ctx.MustGet(authUserID).(string)
//	if !ok {
//		log.Printf("can't get user email on userBooks")
//		ctx.Status(http.StatusBadRequest)
//		return
//	}
//	_, err := h.srvs.GetUserBooks(email)
//	if err != nil {
//		log.Printf("get User books error %s", err.Error())
//	}
//	// logic
//	//fmt.Println("Email of book owner user: ", email)
//	//ctx.Status(http.StatusOK)
//}

func (h *Handler) getAllBooks(ctx *gin.Context) {

	books, err := h.srvs.GetAllBooks(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (h *Handler) getBookById(ctx *gin.Context) {

	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "book id is empty"})
		return
	}

	book, err := h.srvs.GetBookById(ctx, bookId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *Handler) createBook(ctx *gin.Context) {
	var req api.BookRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	err := h.srvs.CreateBook(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, req)
}

func (h *Handler) updateBook(ctx *gin.Context) {
	var req api.BookRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "book id is empty"})
		return
	}

	err := h.srvs.UpdateBook(ctx, bookId, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteBook(ctx *gin.Context) {

	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "book id is empty"})
		return
	}

	err := h.srvs.DeleteBook(ctx, bookId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "Book deleted"})
}
