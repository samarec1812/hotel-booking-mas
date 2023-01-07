package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
	"log"
	"net/http"
)

func (h *Handler) getAll(c *gin.Context) {
	acc, err := h.services.BankAccountList.GetAllAccount()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "find",
		"accounts": acc,
	})
}

func (h *Handler) create(c *gin.Context) {
	var acc models.BankAccount

	if err := c.ShouldBindJSON(&acc); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.services.BankAccountList.CreateAccount(acc)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("Bank account created")
	c.JSON(http.StatusCreated, gin.H{
		"status": "created",
	})
}
