package service

import (
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/pkg/repository"
)

type BankAccountListService struct {
	repo repository.BankAccountList
}

func NewBankAccountListService(repo repository.BankAccountList) *BankAccountListService {
	return &BankAccountListService{repo: repo}
}

func (s *BankAccountListService) CreateAccount(acc models.BankAccount) (int, error) {
	return s.repo.CreateAccount(acc)
}

func (s *BankAccountListService) GetAllAccount() ([]models.BankAccount, error) {
	return s.repo.GetAllAccount()
}
