package service

import (
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/pkg/repository"
)

type BookingListService struct {
	repo repository.BookingList
}

func NewBookingListService(repo repository.BookingList) *BookingListService {
	return &BookingListService{repo: repo}
}

func (s *BookingListService) CreateBooking(booking models.Booking) (int, error) {
	return s.repo.CreateBooking(booking)
}
