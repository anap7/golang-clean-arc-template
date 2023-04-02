package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/anap7/golang-clean-arc-template/app/adapter/repository"
	"github.com/anap7/golang-clean-arc-template/app/core/usecase/process_transaction"
	"github.com/anap7/golang-clean-arc-template/app/handler"
)

func Create(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := ioutil.ReadAll(r.Body)

	if error != nil {
		// A helper that handle with errors response
		handler.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	transaction := process_transaction.TransactionDtoInput{}

	if error = json.Unmarshal(bodyRequest, &transaction); error != nil {
		// A helper that handle with errors response
		handler.Error(w, http.StatusBadRequest, error)
		return
	}

	db, err := repository.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)

	usecase := process_transaction.NewProcessTransaction(repo)
	output, err := usecase.Execute(transaction)
	if err != nil {
		log.Fatal(err)
	}

	handler.JSON(w, http.StatusCreated, output)
}