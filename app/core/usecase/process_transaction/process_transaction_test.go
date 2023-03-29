package process_transaction

import (
	"testing"

	"github.com/anap7/golang-clean-arc-template/app/core/repository"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionValidWhenItsValid(t *testing.T) {
	input := TransactionDtoInput {
		ID: "1",
		AccountID: "1",
		Amount: 200,
	}

	expectedOutput := TransactionDtoOutput {
		ID: "1",
		Status: "Approved",
		ErrorMessage: "",
	}

	//repository := repository.TransactionRepository()

	usecase := NewProcessTransaction()
	output, err := usecase.Execute(input)

	//Espera que o erro seja em branco
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}