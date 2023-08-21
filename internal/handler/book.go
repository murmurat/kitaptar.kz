package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"one-lab/api"
)

// not ready
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

// ready to test
func (h *Handler) getAllBooks(ctx *gin.Context) {

	books, err := h.srvs.GetAllBooks(ctx)

	if err != nil {
		log.Printf("Handler all book getting error %w", err)
	}

	ctx.JSON(http.StatusOK, books)
}

// ready to test
func (h *Handler) getBookById(ctx *gin.Context) {

	bookId := ctx.Param("id")
	book, err := h.srvs.GetBookById(ctx, bookId)

	if err != nil {
		log.Printf("Handler book getting by id error %w", err)
	}

	ctx.JSON(http.StatusOK, book)
}

// ready to test
func (h *Handler) createBook(ctx *gin.Context) {
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

// ready to test
func (h *Handler) updateBook(ctx *gin.Context) {
	var req api.BookRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Println("Bind json error ", err.Error())
		ctx.JSON(http.StatusBadRequest, err)
	}
	bookId := ctx.Param("id")
	//_, err := h.srvs.GetBookById(ctx, bookId)
	//
	//if err != nil {
	//	log.Printf("Handler book getting by id error %w", err)
	//}
	err := h.srvs.UpdateBook(ctx, bookId, &req)
	if err != nil {
		log.Printf("Error %w", err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, req)
}

// ready to test
func (h *Handler) deleteBook(ctx *gin.Context) {

	bookId := ctx.Param("id")

	err := h.srvs.DeleteBook(ctx, bookId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, "Book deleted")
}
