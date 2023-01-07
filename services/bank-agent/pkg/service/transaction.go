package service

import (
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/pkg/repository"
)

type TransactionsListService struct {
	repo repository.TransactionsList
}

func NewTransactionsListService(repo repository.TransactionsList) *TransactionsListService {
	return &TransactionsListService{repo: repo}
}

func (s *TransactionsListService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.GetAllTransactions()
}

func (s *TransactionsListService) CreateTransaction(payment models.Payment) error {
	return s.repo.CreateTransaction(payment)
}
