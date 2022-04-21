package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "traflow_development"
	password = "traflow_development"
	dbname   = "practice6"
)

func openDBConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbClient, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbClient, nil
}
