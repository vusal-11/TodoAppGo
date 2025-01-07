package handler

import (
	"net/http"
	todoapp "todo-app"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todoapp.User
	
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id , err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"` // Исправлено закрытие кавычки
    Password string `json:"password" binding:"required"` // Исправлено тег json
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	token , err := h.services.Authorization.GenerateToken(input.Username,input.Password)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":token,
	})

}