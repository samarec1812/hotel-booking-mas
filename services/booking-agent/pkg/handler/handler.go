package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		//auth.GET("/sign-in", h.getSignIn)
		//auth.GET("/sign-up", h.getSignUp)
	}
	api := router.Group("/booking", h.middlewareLogger)
	{
		api.GET("/hotels/update", h.updateRoom)
		api.GET("/hotels/all", h.getAllRoom)
		api.GET("/hotels/:id", h.getRoom)
		api.GET("/hotels", h.getAllRoomBySearch)

		booking := api.Group("/hotels/:id")
		{
			booking.POST("/create", h.createBooking)
			booking.POST("/payment", h.createPayment)
		}
	}

	return router
}
