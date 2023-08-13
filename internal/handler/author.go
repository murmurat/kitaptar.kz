package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"one-lab/api"
)

func (h *Handler) getAllAuthors(ctx *gin.Context) {

	authors, err := h.srvs.GetAllAuthors(ctx)

	if err != nil {
		log.Printf("Handler all authors getting error %w", err)
	}

	ctx.JSON(http.StatusOK, authors)
}

func (h *Handler) getAuthorById(ctx *gin.Context) {

	bookId := ctx.Param("id")
	book, err := h.srvs.GetBookById(ctx, bookId)

	if err != nil {
		log.Printf("Handler book getting by id error %w", err)
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *Handler) createAuthor(ctx *gin.Context) {
	var req api.BookRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Println("Bind json error ", err.Error())
		ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.srvs.CreateBook(ctx, &req)
	if err != nil {
		log.Printf("Error %w", err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusCreated, req)
}

func (h *Handler) updateAuthor(ctx *gin.Context) {
	var req api.BookRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Println("Bind json error ", err.Error())
		ctx.JSON(http.StatusBadRequest, err)
	}
	bookId := ctx.Param("id")
	_, err := h.srvs.GetBookById(ctx, bookId)

	if err != nil {
		log.Printf("Handler book getting by id error %w", err)
	}
	err = h.srvs.UpdateBook(ctx, bookId, &req)
	if err != nil {
		log.Printf("Error %w", err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteAuthor(ctx *gin.Context) {

	bookId := ctx.Param("id")

	err := h.srvs.DeleteBook(ctx, bookId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, "Book deleted")
}
