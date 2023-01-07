package service

import (
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type BookingList interface {
	CreateBooking(booking models.Booking) (int, error)
}

type RoomList interface {
	UpdateRooms(rooms []models.Room) (int, error)
	GetAllRooms() ([]models.Room, error)
	GetAllRoomsBySearch(params models.SearchParams) ([]models.Room, error)
	GetRoomById(roomId int) (models.Room, error)
	//GetRoomList()
}

type Service struct {
	Authorization
	BookingList
	RoomList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		RoomList:      NewRoomListService(repos),
		BookingList:   NewBookingListService(repos),
	}
}
