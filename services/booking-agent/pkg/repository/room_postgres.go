package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
)

type RoomListPostgres struct {
	db *sqlx.DB
}

func NewRoomListPostgres(db *sqlx.DB) *RoomListPostgres {
	return &RoomListPostgres{db: db}
}

func (r *RoomListPostgres) UpdateRooms(rooms []models.Room) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (room_name, description, hotel_name, price, accommodates) values (:room_name, :description, :hotel_name, :price, :accommodates)", roomsTable)
	_, err := r.db.NamedExec(query, rooms)
	return 0, err
}

func (r *RoomListPostgres) GetAllRooms() ([]models.Room, error) {
	var rooms []models.Room

	query := fmt.Sprintf("SELECT id, room_name, description, hotel_name, price, accommodates from %s", roomsTable)
	nstmt, err := r.db.Preparex(query)
	if err != nil {
		return []models.Room{}, err
	}
	err = nstmt.Select(&rooms)

	return rooms, err
}

func (r *RoomListPostgres) GetRoomById(roomId int) (models.Room, error) {
	query := fmt.Sprintf("SELECT room_name, description, hotel_name, price, accommodates FROM %s WHERE id=$1", roomsTable)
	var room models.Room
	err := r.db.Get(&room, query, roomId)

	return room, err
}

func (r *RoomListPostgres) GetAllRoomsBySearch(params models.SearchParams) ([]models.Room, error) {
	var rooms []models.Room

	query := fmt.Sprintf("select room_name, description, hotel_name, price, accommodates from %s where id not in (select room_id from %s where ((arrival_date <= $1 and $2 <= departure_date) or (arrival_date > $1 and arrival_date < $2) or (departure_date > $1 and departure_date < $2)))", roomsTable, bookingsTable)
	//nstmt, err := r.db.Preparex(query)
	//if err != nil {
	//	return []models.Room{}, err
	//}
	//err = nstmt.Select(&rooms)
	err := r.db.Select(&rooms, query, params.ArrivalDate, params.DepartureDate)

	return rooms, err
}
