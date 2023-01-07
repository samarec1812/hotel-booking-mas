package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type RoomList interface {
	UpdateRooms(rooms []models.Room) (int, error)
	GetAllRooms() ([]models.Room, error)
	GetRoomById(roomId int) (models.Room, error)
	GetAllRoomsBySearch(params models.SearchParams) ([]models.Room, error)
}

type BookingList interface {
	CreateBooking(booking models.Booking) (int, error)
}

type Repository struct {
	Authorization
	RoomList
	BookingList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		RoomList:      NewRoomListPostgres(db),
		BookingList:   NewBookingListPostgres(db),
	}
}
