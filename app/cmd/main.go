package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/anap7/golang-clean-arc-template/app/adapter/repository"
	"github.com/anap7/golang-clean-arc-template/app/core/usecase/process_transaction"
)

func main() {
	db, err := repository.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput {
		ID: "306660",
		AccountID: "66666",
		Amount: -1,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	outputJson, _ := json.Marshal(output)
	fmt.Println(string(outputJson))
}