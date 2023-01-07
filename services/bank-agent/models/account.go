package models

type BankAccount struct {
	Id         int    `json:"-,omitempty" db:"id"`
	Name       string `json:"name" db:"name"`
	Password   string `json:"password" db:"password"`
	SecureCode int    `json:"secure_code" db:"secure_code"`
	Balance    int    `json:"balance" db:"balance"`
}
