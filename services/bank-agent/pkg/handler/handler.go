package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	bank := router.Group("/bank")
	{
		bank.GET("/account", h.getAll)
		bank.POST("/account/create", h.create)
		bank.POST("/transaction/create", h.createTransaction)

		//auth.POST("/sign-in", h.signIn

	}

	return router
}
