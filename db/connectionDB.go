package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "host.docker.internal"
	port     = 5434
	user     = "default"
	password = "default"
	dbname   = "default"
)

func Connect() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	err = database.Ping()
	CheckError(err)

	fmt.Println("Successfull connect to db")
	return database
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
