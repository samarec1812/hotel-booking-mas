package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	err := c.ShouldBindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "created",
		"user_id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie(authorizationHeader, token, 3600, "/booking", "localhost", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "authorization",
	})

}
