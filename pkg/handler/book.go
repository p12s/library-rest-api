package handler

import (
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p12s/library-rest-api/pkg/models"
	"github.com/sirupsen/logrus"
)

// @Summary Create
// @Security ApiKeyAuth
// @Tags Book
// @Description Creating book
// @ID createBook
// @Accept  json
// @Produce  json
// @Param input body models.Book true "Book created info"
// @Success 200 {string} string "id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/ [post]
func (h *Handler) createBook(c *gin.Context) {
	userId, err := getUserId(c)
	logrus.Info("User with id ", userId, " created book")

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Book
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateBook(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all
// @Security ApiKeyAuth
// @Tags Books
// @Description Getting all books
// @ID getAllBooks
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Book
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/ [get]
func (h *Handler) getAllBooks(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO тут вопрос - стоит ли кешировать всю библиотеку. я думаю - не стоит:
	// 1 - книг может быть много
	// 2 - запрос может быть нечастым, а память будет занята всегда

	items, err := h.services.GetAllBook()
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get by id
// @Security ApiKeyAuth
// @Tags Book
// @Description Getting book by {id}
// @ID getBookById
// @Accept  json
// @Produce  json
// @Param id path integer true "Book id"
// @Success 200 {object} []models.Book
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/{id} [get]
func (h *Handler) getBookById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	book := h.cache.Get(strconv.Itoa(bookId))
	if book == nil {
		book, err = h.services.GetBookById(bookId)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		h.cache.SetWithExpire(strconv.Itoa(bookId), book, time.Hour*24)
	}

	c.JSON(http.StatusOK, book)
}

// @Summary Update by id
// @Security ApiKeyAuth
// @Tags Book
// @Description Update book by {id}
// @ID updateBook
// @Accept  json
// @Produce  json
// @Param id path integer true "Book id"
// @Param input body models.Book true "Book updated info"
// @Success 200 {string} string 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/{id} [put]
func (h *Handler) updateBook(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	h.cache.Delete(strconv.Itoa(bookId))

	var input models.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateBook(bookId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"OK"})
}

// @Summary Delete
// @Security ApiKeyAuth
// @Tags Book
// @Description Deleting book by {id}
// @ID deleteBook
// @Accept  json
// @Produce  json
// @Param id path integer true "Book id"
// @Success 200 {string} string 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/{id} [delete]
func (h *Handler) deleteBook(c *gin.Context) {
	_, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	err = h.services.DeleteBook(bookId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.cache.Delete(strconv.Itoa(bookId))

	c.JSON(http.StatusOK, statusResponse{"OK"})
}
