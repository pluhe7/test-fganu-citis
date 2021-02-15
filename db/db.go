package db

import (
	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "root",
		Database: "records-test",
	})
	return db
}

func Disconnect(db *pg.DB) {
	db.Close()
}
