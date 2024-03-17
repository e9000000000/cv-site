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

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, title TEXT, text TEXT)")
    if err != nil {
        panic(err)
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT UNIQUE, password_hash TEXT, token TEXT, is_admin BOOLEAN)")
    if err != nil {
        panic(err)
    }

    return db
}
