package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/pkg/jetstream"
	"net/http"
	"strconv"
)

func (h *Handler) createBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "need authorization")
		return
	}
	var input models.Booking
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.UserId = userId
	input.RoomId = id
	input.Status = "requires payment"
	fmt.Println(input)
	bookingId, err := h.services.BookingList.CreateBooking(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "created",
		"booking_id": bookingId,
	})
}

func (h *Handler) createPayment(c *gin.Context) {
	var input models.Payment
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "need authorization")
		return
	}

	err = jetstream.CreateOrder(input, userId, jetstream.JS)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	resp := jetstream.ResponseOrder(jetstream.JS)
	if resp.Status != "ok" {
		newErrorResponse(c, http.StatusBadRequest, resp.Status)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
