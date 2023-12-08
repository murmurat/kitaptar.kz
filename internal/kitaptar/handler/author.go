package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
)

// getAllAuthors godoc
// @Summary      Get all authors
// @Security	 ApiKeyAuth
// @Description  Get all authors
// @Tags         authors
// @Produce      json
// @Success      200  {object}  []entity.Author
// @Failure      500  {object}  api.Error
// @Router       /authors [get]
func (h *Handler) getAllAuthors(ctx *gin.Context) {

	authors, err := h.srvs.GetAllAuthors(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

// getAuthorById godoc
// @Summary      Get author by id
// @Security	 ApiKeyAuth
// @Description  Get author by id
// @Tags         authors
// @Param id path  string true "id"
// @Produce      json
// @Success      200  {object}  entity.Author
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /authors/{id} [get]
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

// getAuthorByName godoc
// @Summary      Get author by name (Search)
// @Security	 ApiKeyAuth
// @Description  Get author by name
// @Tags         authors
// @Param name path string true "name"
// @Produce      json
// @Success      200  {object}  []entity.Author
// @Failure      500  {object}  api.Error
// @Router       /authors/search/{name} [get]
func (h *Handler) getAuthorByName(ctx *gin.Context) {

	authorName := ctx.Param("name")

	authors, err := h.srvs.GetAuthorByName(ctx, authorName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

// createAuthor godoc
// @Summary      Create Authors
// @Security	 ApiKeyAuth
// @Description  Create new author
// @Tags         authors
// @Param req body AuthorHandlerDto true "req"
// @Accept       json
// @Produce      json
// @Success      201  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /authors [post]
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

// updateAuthor godoc
// @Summary      Update author
// @Security	 ApiKeyAuth
// @Description  Update existing author
// @Tags         authors
// @Param req body api.AuthorRequest true "request"
// @Param id path string true "id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  api.AuthorRequest
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /authors/{id} [put]
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

// deleteAuthor godoc
// @Summary      Delete author
// @Security	 ApiKeyAuth
// @Description  Delete existing author
// @Tags         authors
// @Param id path string true "id"
// @Produce      json
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /authors/{id} [delete]
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
