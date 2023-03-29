package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "07324"
	transaction.AccountID = "1"
	transaction.Amount = 2000

	err := transaction.IsValid()

	//Espera receber um erro
	assert.Error(t, err)
	//Espera que a mensagem seja a mensagem x
	assert.Equal(t, "You don't have limit for this transaction", err.Error())
}

func TestTransactionWithAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "07324"
	transaction.AccountID = "1"
	transaction.Amount = 0

	err := transaction.IsValid()

	//Espera receber um erro
	assert.Error(t, err)
	//Espera que a mensagem seja a mensagem x
	assert.Equal(t, "The amount must be greater than 1", err.Error())
}
