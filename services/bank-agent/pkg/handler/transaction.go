package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
	"log"
	"net/http"
)

func (h *Handler) createTransaction(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.TransactionsList.CreateTransaction(payment)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("Transaction created")
	c.JSON(http.StatusCreated, gin.H{
		"status": "created",
	})
}
