package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
)

type BookingListPostgres struct {
	db *sqlx.DB
}

func NewBookingListPostgres(db *sqlx.DB) *BookingListPostgres {
	return &BookingListPostgres{db: db}
}

func (r *BookingListPostgres) CreateBooking(booking models.Booking) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, room_id, arrival_date, departure_date, status) values (:user_id, :room_id, :arrival_date, :departure_date, :status)", bookingsTable)
	_, err := r.db.NamedExec(query, booking)

	return 0, err
}
