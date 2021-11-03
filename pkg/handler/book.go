package handler

import (
	"github.com/gin-gonic/gin"
	//"github.com/p12s/library-rest-api"
	"net/http"
	//"strconv"
)

func (h *Handler) createBook(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// listId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	// 	return
	// }

	// var input todo.TodoItem
	// if err := c.BindJSON(&input); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// id, err := h.services.TodoItem.Create(userId, listId, input)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": 1,
	})
}

func (h *Handler) getAllBooks(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// listId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	// 	return
	// }

	// items, err := h.services.TodoItem.GetAll(userId, listId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getBookById(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// itemId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
	// 	return
	// }

	// item, err := h.services.TodoItem.GetById(userId, itemId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) updateBook(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
	// 	return
	// }

	// var input todo.UpdateItemInput
	// if err := c.BindJSON(&input); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// if err := h.services.TodoItem.Update(userId, id, input); err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteBook(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// itemId, err := strconv.Atoi(c.Param("id"))

	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
	// 	return
	// }

	// err = h.services.TodoItem.Delete(userId, itemId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
