package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
	"log"
	"net/http"
)

func (h *Handler) getAllAuthors(ctx *gin.Context) {

	authors, err := h.srvs.GetAllAuthors(ctx)

	if err != nil {
		log.Printf("Handler all authors getting error %w", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, authors)
}

func (h *Handler) getAuthorById(ctx *gin.Context) {

	authorId := ctx.Param("id")
	author, err := h.srvs.GetAuthorById(ctx, authorId)

	if err != nil {
		log.Printf("Handler author getting by id error %w", err)
		ctx.JSON(http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, author)
}

func (h *Handler) createAuthor(ctx *gin.Context) {
	var req api.AuthorRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Println("Bind json error ", err.Error())
		ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.srvs.CreateAuthor(ctx, &req)
	if err != nil {
		log.Printf("Error %w", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, req)
}

func (h *Handler) updateAuthor(ctx *gin.Context) {
	var req api.AuthorRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Println("Bind json error ", err.Error())
		ctx.JSON(http.StatusBadRequest, err)
	}
	authorId := ctx.Param("id")
	//_, err := h.srvs.GetBookById(ctx, authorId)
	//
	//if err != nil {
	//	log.Printf("Handler author getting by id error %w", err)
	//}
	err := h.srvs.UpdateAuthor(ctx, authorId, &req)
	if err != nil {
		log.Printf("Error %w", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteAuthor(ctx *gin.Context) {

	authorId := ctx.Param("id")

	err := h.srvs.DeleteAuthor(ctx, authorId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, "Author deleted")
}
