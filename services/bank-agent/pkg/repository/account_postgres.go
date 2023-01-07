package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
)

type BankAccountListPostgres struct {
	db *sqlx.DB
}

func NewBankAccountListPostgres(db *sqlx.DB) *BankAccountListPostgres {
	return &BankAccountListPostgres{db: db}
}

func (r *BankAccountListPostgres) CreateAccount(acc models.BankAccount) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, password, secure_code, balance) values (:name, :password, :secure_code, :balance)", bankAccountTable)
	_, err := r.db.NamedExec(query, acc)

	return 0, err
}

func (r *BankAccountListPostgres) GetAllAccount() ([]models.BankAccount, error) {
	var acc []models.BankAccount
	query := fmt.Sprintf("SELECT id, name, password, secure_code, balance from %s", bankAccountTable)
	nstmt, err := r.db.Preparex(query)
	if err != nil {
		return []models.BankAccount{}, err
	}
	err = nstmt.Select(&acc)

	return acc, err
}

//func (r *BankAccountListPostgres) GetAllRooms() ([]models.Room, error) {
//	var rooms []models.Room
//
//	query := fmt.Sprintf("SELECT id, room_name, description, hotel_name, price, accommodates from %s", roomsTable)
//	nstmt, err := r.db.Preparex(query)
//	if err != nil {
//		return []models.Room{}, err
//	}
//	err = nstmt.Select(&rooms)
//
//	return rooms, err
//}
