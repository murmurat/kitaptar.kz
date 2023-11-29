package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
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

func (h *Handler) addToFavorites(ctx *gin.Context) {
	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "book id is empty"})
		return
	}

	userId, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: fmt.Sprintf("user id getting error: %v", err)})
		return
	}

	favoriteId, err := h.srvs.AddToFavorites(ctx, userId, bookId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: fmt.Sprintf("Successfully added to favorited, id: %s", favoriteId)})
}

func (h *Handler) getFromFavorites(ctx *gin.Context) {
	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "book id is empty"})
		return
	}

	userId, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: fmt.Sprintf("user id getting error: %v", err)})
		return
	}

	favorite, err := h.srvs.GetFromFavorites(ctx, userId, bookId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, favorite)
}

func (h *Handler) deleteFromFavorites(ctx *gin.Context) {
	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "book id is empty"})
		return
	}

	userId, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: fmt.Sprintf("user id getting error: %v", err)})
		return
	}

	err = h.srvs.DeleteFromFavorites(ctx, userId, bookId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "Book deleted from favorites"})
}
