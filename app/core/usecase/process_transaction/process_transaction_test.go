package process_transaction

import (
	"fmt"
	"log"
	"testing"

	"github.com/anap7/golang-clean-arc-template/app/adapter/repository"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionValidWhenItsValid(t *testing.T) {
	input := TransactionDtoInput {
		ID: "777",
		AccountID: "777",
		Amount: 15,
	}

	expectedOutput := TransactionDtoOutput {
		ID: "777",
		Status: "approved",
		ErrorMessage: "",
	}

	db, err := repository.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)

	usecase := NewProcessTransaction(repo)
	output, err := usecase.Execute(input)
	fmt.Println(output)
	fmt.Println(expectedOutput)

	//Espera que o erro seja em branco
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}