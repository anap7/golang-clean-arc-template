package fixture

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"github.com/maragudk/migrate"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func Up(migrationsDir fs.FS) *sql.DB {
	ConnectionString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		"golang",
		"Golang_17396",
		"orders",
	)

	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		panic(err) 
	}
	return db
}

func Down(db *sql.DB, migrationsDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		panic(err) 
	}

	db.Close()
}