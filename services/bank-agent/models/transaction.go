package models

import "time"

type Transaction struct {
	ID            int       `json:"-" db:"id"`
	BankAccountID int       `json:"-" db:"bank_acc_id"`
	StartBalance  int       `json:"start_balance" db:"start_balance"`
	EndBalance    int       `json:"end_balance" db:"end_balance"`
	Amount        int       `json:"amount" db:"amount"`
	Description   string    `json:"description" db:"description"`
	Date          time.Time `json:"date" db:"date"`
}

type Payment struct {
	BankAccountName string `json:"bank_account_name"`
	Password        string `json:"password"`
	SecureCode      int    `json:"secure_code"`
	Amount          int    `json:"amount"`
}
