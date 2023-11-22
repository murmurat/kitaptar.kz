package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
	"net/http"
)

// func (h *Handler) userBooks(ctx *gin.Context) {

// TODO Implement all books which liked by user(Saved books)

// }

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

func (h *Handler) getBookByName(ctx *gin.Context) {

	bookName := ctx.Param("name")

	books, err := h.srvs.GetBookByName(ctx, bookName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (h *Handler) createBook(ctx *gin.Context) {

	var req api.BookRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	bookId, err := h.srvs.CreateBook(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, api.Response{Message: bookId})
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

	if req == (api.BookRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update data not provided"})
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
