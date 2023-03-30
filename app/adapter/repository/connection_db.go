package repository

import (
	"database/sql"
	"fmt"
	"log"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
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
