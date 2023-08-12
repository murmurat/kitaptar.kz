package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) userBooks(ctx *gin.Context) {
	email, ok := ctx.MustGet(authUserID).(string)
	if !ok {
		log.Printf("can't get user email on userBooks")
		ctx.Status(http.StatusBadRequest)
		return
	}
	_, err := h.srvs.GetUserBooks(email)
	if err != nil {
		log.Printf("get User books error %s", err.Error())
	}
	// logic
	//fmt.Println("Email of book owner user: ", email)
	//ctx.Status(http.StatusOK)
}

func (h *Handler) getAllBooks(ctx *gin.Context) {

	books, err := h.srvs.GetAllBooks(ctx)

	if err != nil {
		log.Printf("Handler all book getting error %w", err)
	}

	ctx.JSON(http.StatusOK, books)
}

func (h *Handler) getBookById(ctx *gin.Context) {

	bookId := ctx.Param("id")
	book, err := h.srvs.GetBookById(ctx, bookId)

	if err != nil {
		log.Printf("Handler book getting by id error %w", err)
	}

	ctx.JSON(http.StatusOK, book)
}
