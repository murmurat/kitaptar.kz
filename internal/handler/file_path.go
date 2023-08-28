package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
	"log"
	"net/http"
)

//	func (h *Handler) getBookFilePaths(ctx *gin.Context) {
//		filePathId := ctx.Param("id")
//
//		filePath, err := h.srvs.GetBookFilePaths(ctx, filePathId)
//		if !err {
//			log.Printf("getBookFilePaths err %w", err)
//			ctx.JSON(http.StatusInternalServerError, err)
//			return
//		}
//		ctx.JSON(http.StatusOK, filePath)
//	}
func (h *Handler) getAllFilePaths(ctx *gin.Context) {

	filePaths, err := h.srvs.GetAllFilePaths(ctx)

	if err != nil {
		log.Printf("Handler all file paths getting error %w", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, filePaths)
}

func (h *Handler) getFilePathById(ctx *gin.Context) {

	filePathId := ctx.Param("id")
	filePath, err := h.srvs.GetFilePathById(ctx, filePathId)

	if err != nil {
		log.Printf("Handler filePath getting by id error %w", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, filePath)
}

func (h *Handler) createFilePath(ctx *gin.Context) {
	var req api.FilePathRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Println("Bind json error ", err.Error())
		ctx.JSON(http.StatusBadRequest, err)
	}
	log.Println("Input path request: ", *req.Mobi)
	err := h.srvs.CreateFilePath(ctx, &req)
	if err != nil {
		log.Printf("Error %w", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, req)
}

func (h *Handler) updateFilePath(ctx *gin.Context) {
	var req api.FilePathRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Println("Bind json error ", err.Error())
		ctx.JSON(http.StatusBadRequest, err)
	}
	filePathId := ctx.Param("id")
	//_, err := h.srvs.GetBookById(ctx, filePathId)
	//
	//if err != nil {
	//	log.Printf("Handler filePath getting by id error %w", err)
	//}
	err := h.srvs.UpdateFilePath(ctx, filePathId, &req)
	if err != nil {
		log.Printf("Error %w", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *Handler) deleteFilePath(ctx *gin.Context) {

	filePathId := ctx.Param("id")

	err := h.srvs.DeleteFilePath(ctx, filePathId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, "File Path deleted")
}
