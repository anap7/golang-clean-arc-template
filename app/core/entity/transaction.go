package entity

import (
	"errors"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	ErrorMessage string
}

// Construtor
func NewTransaction() *Transaction {
	return &Transaction{}
}

// Metodo
func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("You don't have limit for this transaction")
	} 
	
	if t.Amount < 1 {
		return errors.New("The amount must be greater than 1")
	}
	return nil
}