package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
)

//func (h *Handler) getBookFilePaths(ctx *gin.Context) {
//		filePathId := ctx.Param("id")
//
//		filePath, err := h.srvs.GetBookFilePaths(ctx, filePathId)
//		if !err {
//			log.Printf("getBookFilePaths err %w", err)
//			ctx.JSON(http.StatusInternalServerError, err)
//			return
//		}
//		ctx.JSON(http.StatusOK, filePath)
//}

func (h *Handler) getAllFilePaths(ctx *gin.Context) {

	filePaths, err := h.srvs.GetAllFilePaths(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, filePaths)
}

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
