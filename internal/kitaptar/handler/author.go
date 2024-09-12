package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
)

func (h *Handler) getAllAuthors(ctx *gin.Context) {

	authors, err := h.srvs.GetAllAuthors(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func (h *Handler) getAuthorById(ctx *gin.Context) {

	authorId := ctx.Param("id")
	if authorId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "author id is empty"})
		return
	}

	author, err := h.srvs.GetAuthorById(ctx, authorId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (h *Handler) getAuthorByName(ctx *gin.Context) {

	authorName := ctx.Param("name")

	authors, err := h.srvs.GetAuthorByName(ctx, authorName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func (h *Handler) createAuthor(ctx *gin.Context) {

	var req *api.AuthorRequest
	var reqDto AuthorHandlerDto

	if err := ctx.BindJSON(&reqDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	builder := NewAuthorBuilder()

	director := NewDirector(builder)

	req = director.ConstructAuthor(reqDto.Firstname, reqDto.Lastname, reqDto.AboutAuthor, reqDto.ImagePath)

	authorId, err := h.srvs.CreateAuthor(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, api.Response{Message: authorId})
}

func (h *Handler) updateAuthor(ctx *gin.Context) {

	var req api.AuthorRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	authorId := ctx.Param("id")
	if authorId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "author id is empty"})
		return
	}
	if req == (api.AuthorRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update data not provided"})
		return
	}

	err := h.srvs.UpdateAuthor(ctx, authorId, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteAuthor(ctx *gin.Context) {

	authorId := ctx.Param("id")
	if authorId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "author id is empty"})
		return
	}

	err := h.srvs.DeleteAuthor(ctx, authorId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "Author deleted"})
}
