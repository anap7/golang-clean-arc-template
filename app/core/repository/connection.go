package repository

import "database/sql"

type ConnectionRepository interface {
	OpenConnection() (*sql.DB, error)
}