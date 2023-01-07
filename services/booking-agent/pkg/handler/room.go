package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *Handler) updateRoom(c *gin.Context) {
	res, err := http.Get("http://localhost:8080/rooms")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var rooms []models.Room
	err = json.Unmarshal(resBody, &rooms)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.services.RoomList.UpdateRooms(rooms)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

func (h *Handler) getAllRoom(c *gin.Context) {

	roomList, err := h.services.RoomList.GetAllRooms()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "find",
		"rooms":  roomList,
	})
}

func (h *Handler) getRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	roomItem, err := h.services.RoomList.GetRoomById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "find",
		"room":   roomItem,
	})
}

func (h *Handler) getAllRoomBySearch(c *gin.Context) {
	var params models.SearchParams
	if err := c.Bind(&params); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid search params")
		return
	}
	if params.ArrivalDate.Sub(params.DepartureDate) > 0 {
		newErrorResponse(c, http.StatusBadRequest, "arrival date after departure date")
		return
	}
	roomList, err := h.services.RoomList.GetAllRoomsBySearch(params)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "find",
		"rooms":  roomList,
	})
}
