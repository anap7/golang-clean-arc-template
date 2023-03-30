package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	//mysql driver
	"github.com/anap7/golang-clean-arc-template/app/adapter/repository"
	"github.com/anap7/golang-clean-arc-template/app/core/usecase/process_transaction"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ConnectionString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		"golang",
		"Golang_17396",
		"orders",
	)
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput {
		ID: "3",
		AccountID: "3",
		Amount: 200,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Println(err.Error())
	}


	outputJson, _ := json.Marshal(output)
	fmt.Println(string(outputJson))
}