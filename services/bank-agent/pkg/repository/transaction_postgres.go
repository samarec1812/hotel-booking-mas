package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
	"time"
)

type TransactionsListPostgres struct {
	db *sqlx.DB
}

func NewTransactionsListPostgres(db *sqlx.DB) *TransactionsListPostgres {
	return &TransactionsListPostgres{db: db}
}

func (r *TransactionsListPostgres) GetAllTransactions() ([]models.Transaction, error) {
	return nil, nil
}

func (r *TransactionsListPostgres) CreateTransaction(payment models.Payment) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	var acc models.BankAccount
	queryGetAccount := fmt.Sprintf("SELECT id, name, password, secure_code, balance from %s where name=?, password=?, secure_code=?", bankAccountTable)
	err = tx.Get(acc, queryGetAccount, payment.BankAccountName, payment.Password, payment.SecureCode)
	if err != nil {
		tx.Rollback()
		return err
	}
	if acc.Balance < payment.Amount {
		tx.Rollback()
		return errors.New("not enough money in this account")
	}

	queryChangeBalance := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2", bankAccountTable)
	endBalance := acc.Balance - payment.Amount
	_, err = tx.Exec(queryChangeBalance, endBalance, acc.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	queryInsertTransaction := fmt.Sprintf("INSERT INTO %s (bank_acc_id, start_balance, end_balance, amount, description, date) values ($1, $2, $3, $4, $5, $6)", transactionsTable)
	_, err = tx.Exec(queryInsertTransaction, acc.Id, acc.Balance, endBalance, payment.Amount, "бронирование номера", time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
