package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = nil

func runSql(q string) error {
    statement, err := db.Prepare(q)
    if err != nil {
        return err
    }
    _, err = statement.Exec()
    if err != nil {
        return err
    }

    return nil
}

func initDb() *sql.DB {
    var err error
    db, err = sql.Open("sqlite3", "./db.sqlite3")
    if err != nil {
        panic(err)
    }

    err = runSql("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, title TEXT, text TEXT)")
    if err != nil {
        panic(err)
    }

    err = runSql("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password_hash TEXT, is_admin BOOLEAN)")
    if err != nil {
        panic(err)
    }

    return db
}
