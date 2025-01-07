package handler

import (
	"net/http"
	"strconv"
	todoapp "todo-app"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context){
	userId ,err := getUserId(c)
	if err != nil {
		return
	}
	var input todoapp.ToDoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err := h.services.ToDoList.Create(userId,input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})



}

type getAllListsResponse struct {
	Data []todoapp.ToDoList `json:"data"`
}

func (h *Handler) getAlllists(c *gin.Context){
	userId, err :=getUserId(c)
	if err != nil{
		return
	}

	lists, err := h.services.ToDoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllListsResponse{
		Data: lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {
	userId, err :=getUserId(c)
	if err != nil{
		return
	}

	id , err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest,"invalid id params")
		return
	}

	list, err := h.services.ToDoList.GetById(userId,id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}


	id , err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}

	var input todoapp.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	if err := h.services.ToDoList.Update(userId,id,input);err !=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})




}

func (h *Handler) deleteList (c *gin.Context) {
	userId, err :=getUserId(c)
	if err != nil{
		return
	}

	id , err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest,"invalid id params")
		return
	}

	err = h.services.ToDoList.Delete(userId,id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,statusResponse{
		Status: "ok",
	})

}