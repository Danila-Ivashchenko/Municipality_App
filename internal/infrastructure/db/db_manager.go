package db

import "database/sql"

type DataBaseManager interface {
	GetDB() *sql.DB
}
