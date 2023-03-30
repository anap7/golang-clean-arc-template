package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) OpenConnection() (*sql.DB, error) {
	ConnectionString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		"golang",
		"Golang_17396",
		"orders",
	)
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal(err)
		return db, err
	}

	return db, nil
}

func (t *TransactionRepositoryDb) Insert(id string, accountId string, amount float64, status string, errorMessage string) error {
	stmt, err := t.db.Prepare(`
		insert into transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		values (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		id,
		accountId,
		amount,
		status, 
		errorMessage,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}