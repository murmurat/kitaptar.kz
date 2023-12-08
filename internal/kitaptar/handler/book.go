package handler

import (
	"fmt"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/metrics"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
)

// func (h *Handler) userBooks(ctx *gin.Context) {

// TODO Implement all books which liked by user(Saved books)

// }

// getAllBooks godoc
// @Summary      Get all books
// @Security	 ApiKeyAuth
// @Description  Get all books
// @Tags         books
// @Produce      json
// @Param sort query string false "asc or desc by created_at"
// @Success      200  {object}  []entity.Book
// @Failure      500  {object}  api.Error
// @Router       /books [get]
func (h *Handler) getAllBooks(ctx *gin.Context) {

	sortBy := ctx.Query("sort")

	books, err := h.srvs.GetAllBooks(ctx, sortBy)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

// getBookById godoc
// @Summary      Get book by id
// @Security	 ApiKeyAuth
// @Description  Get book by id
// @Tags         books
// @Param id path  string true "id"
// @Produce      json
// @Success      200  {object}  entity.Book
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /books/{id} [get]
func (h *Handler) getBookById(ctx *gin.Context) {

	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "book id is empty"})
		return
	}
	metrics.BookRequests.WithLabelValues(time.Now().Format("2006-01-02")).Add(1)
	book, err := h.srvs.GetBookById(ctx, bookId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// getBookByName godoc
// @Summary      Get book by name (Search)
// @Security	 ApiKeyAuth
// @Description  Get book by name and annotation
// @Tags         books
// @Param name path string true "name"
// @Produce      json
// @Success      200  {object}  []entity.Book
// @Failure      500  {object}  api.Error
// @Router       /books/search/{name} [get]
func (h *Handler) getBookByName(ctx *gin.Context) {

	bookName := ctx.Param("name")

	books, err := h.srvs.GetBookByName(ctx, bookName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

// createBook godoc
// @Summary      Create Books
// @Security	 ApiKeyAuth
// @Description  Create new book
// @Tags         books
// @Param req body api.BookRequest true "req"
// @Accept       json
// @Produce      json
// @Success      201  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /books [post]
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

// updateBook godoc
// @Summary      Update book
// @Security	 ApiKeyAuth
// @Description  Update existing book
// @Tags         books
// @Param req body api.BookRequest true "request"
// @Param id path string true "id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  api.BookRequest
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /books/{id} [put]
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

// deleteBook godoc
// @Summary      Delete book
// @Security	 ApiKeyAuth
// @Description  Delete existing book
// @Tags         books
// @Param id path string true "id"
// @Produce      json
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /books/{id} [delete]
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

// addToFavorites godoc
// @Summary      Add book to favorites
// @Security	 ApiKeyAuth
// @Description  Add book to favorites
// @Tags         books/favorites
// @Param id path string true "book_id"
// @Produce      json
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /books/favorites/{id} [post]
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

// getFromFavorites godoc
// @Summary      Get book from user favorites
// @Security	 ApiKeyAuth
// @Description  Get book from user favorites
// @Tags         books/favorites
// @Param id path string true "book_id"
// @Produce      json
// @Success      200  {object}  entity.FavoriteBook
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /books/favorites/{id} [get]
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

// deleteFromFavorites godoc
// @Summary      Delete book from user favorites
// @Security	 ApiKeyAuth
// @Description  Delete book from user favorites
// @Tags         books/favorites
// @Param id path string true "book_id"
// @Produce      json
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /books/favorites/{id} [delete]
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
