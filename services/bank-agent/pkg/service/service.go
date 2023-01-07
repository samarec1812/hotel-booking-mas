package service

import (
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/pkg/repository"
)

type BankAccountList interface {
	CreateAccount(acc models.BankAccount) (int, error)
	GetAllAccount() ([]models.BankAccount, error)
}

type TransactionsList interface {
	GetAllTransactions() ([]models.Transaction, error)
	CreateTransaction(payment models.Payment) error
}
type Service struct {
	BankAccountList
	TransactionsList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		BankAccountList:  NewBankAccountListService(repos),
		TransactionsList: NewTransactionsListService(repos),
	}
}
