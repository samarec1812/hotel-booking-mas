package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
)

type BankAccountList interface {
	CreateAccount(acc models.BankAccount) (int, error)
	GetAllAccount() ([]models.BankAccount, error)
}

type TransactionsList interface {
	GetAllTransactions() ([]models.Transaction, error)
	CreateTransaction(payment models.Payment) error
}

type Repository struct {
	BankAccountList
	TransactionsList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		BankAccountList:  NewBankAccountListPostgres(db),
		TransactionsList: NewTransactionsListPostgres(db),
	}
}
