package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const DB_NAME = "UMS"
const DB_SERVER = "tcp(127.0.0.1)"
const DB_PORT = ""
const DB_USER = "root"
const DB_PASSWORD = "password"

func connect() *sql.DB {
	connString := getConnectionString()

	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func getConnectionString() string {
	connString := DB_USER

	if DB_PASSWORD != "" {
		connString += ":" + DB_PASSWORD
	}

	connString += "@" + DB_SERVER

	if DB_PORT != "" {
		connString += ":" + DB_PORT
	}

	connString += "/" + DB_NAME

	return connString
}
