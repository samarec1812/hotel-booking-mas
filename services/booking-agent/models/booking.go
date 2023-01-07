package models

import (
	"time"
)

type Booking struct {
	Id            int       `json:"-" db:"id"`
	UserId        int       `json:"user_id" db:"user_id"`
	RoomId        int       `json:"room_id" db:"room_id"`
	ArrivalDate   time.Time `json:"arrival_date" db:"arrival_date"`
	DepartureDate time.Time `json:"departure_date" db:"departure_date"`
	Status        string    `json:"status" db:"status"`
}

type Payment struct {
	BankAccountName string `json:"bank_account_name"`
	Password        string `json:"password"`
	SecureCode      int    `json:"secure_code"`
	Amount          int    `json:"amount"`
}

//
//type JSTime time.Time
//
//func (c *JSTime) UnmarshalJSON(b []byte) error {
//	value := strings.Trim(string(b), `"`) //get rid of "
//	if value == "" || value == "null" {
//		return nil
//	}
//
//	t, err := time.Parse(, value) //parse time
//	if err != nil {
//		return err
//	}
//	*c = JSTime(t) //set result using the pointer
//	return nil
//}
//
//func (c JSTime) MarshalJSON() ([]byte, error) {
//	return []byte(`"` + time.Time(c).Format("2022-12-14") + `"`), nil
//}
