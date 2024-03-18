package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = nil

func initDb() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	return db
}
