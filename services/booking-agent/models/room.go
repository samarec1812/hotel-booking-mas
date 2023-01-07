package models

import "time"

type Room struct {
	Id           int    `json:"-,omitempty" db:"id"`
	Name         string `json:"name" db:"room_name"`
	Description  string `json:"description" db:"description"`
	Hotel        string `json:"hotel" db:"hotel_name"`
	Price        string `json:"price" db:"price"`
	Accommodates int    `json:"accommodates" db:"accommodates"`
}

type SearchParams struct {
	ArrivalDate   time.Time `form:"checkin"`
	DepartureDate time.Time `form:"checkout"`
}
