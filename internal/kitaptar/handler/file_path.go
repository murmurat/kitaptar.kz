package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
)

// getAllFilePaths godoc
// @Summary      Get all file paths
// @Security	 ApiKeyAuth
// @Description  Get all file paths
// @Tags         file_paths
// @Produce      json
// @Success      200  {object}  []entity.FilePath
// @Failure      500  {object}  api.Error
// @Router       /file_paths [get]
func (h *Handler) getAllFilePaths(ctx *gin.Context) {

	filePaths, err := h.srvs.GetAllFilePaths(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, filePaths)
}

// getFilePathById godoc
// @Summary      Get file path by id
// @Security	 ApiKeyAuth
// @Description  Get file path by id
// @Tags         file_paths
// @Param id path  string true "id"
// @Produce      json
// @Success      200  {object}  entity.FilePath
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /file_paths/{id} [get]
func (h *Handler) getFilePathById(ctx *gin.Context) {

	filePathId := ctx.Param("id")
	if filePathId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Empty file path id"})
		return
	}

	filePath, err := h.srvs.GetFilePathById(ctx, filePathId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, filePath)
}

// createFilePath godoc
// @Summary      Create File Paths
// @Security	 ApiKeyAuth
// @Description  Create new file paths
// @Tags         file_paths
// @Param req body api.FilePathRequest true "req"
// @Accept       json
// @Produce      json
// @Success      201  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /file_paths [post]
func (h *Handler) createFilePath(ctx *gin.Context) {
	var req api.FilePathRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	filePathId, err := h.srvs.CreateFilePath(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, api.Response{Message: filePathId})
}

// updateFilePath godoc
// @Summary      Update file path
// @Security	 ApiKeyAuth
// @Description  Update existing file path
// @Tags         file_paths
// @Param req body api.FilePathRequest true "request"
// @Param id path string true "id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  api.FilePathRequest
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /file_paths/{id} [put]
func (h *Handler) updateFilePath(ctx *gin.Context) {
	var req api.FilePathRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}
	filePathId := ctx.Param("id")
	if filePathId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Empty file path id"})
		return
	}

	if req == (api.FilePathRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Update data not provided"})
		return
	}

	err := h.srvs.UpdateFilePath(ctx, filePathId, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

// deleteFilePath godoc
// @Summary      Delete file paths
// @Security	 ApiKeyAuth
// @Description  Delete existing file paths
// @Tags         file_paths
// @Param id path string true "id"
// @Produce      json
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Error
// @Failure      500  {object}  api.Error
// @Router       /file_paths/{id} [delete]
func (h *Handler) deleteFilePath(ctx *gin.Context) {

	filePathId := ctx.Param("id")
	if filePathId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "Empty file path id"})
		return
	}

	err := h.srvs.DeleteFilePath(ctx, filePathId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "File path deleted"})
}
