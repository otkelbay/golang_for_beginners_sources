package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "traflow_development"
	password = "traflow_development"
	dbname   = "restoran"
)

type Order struct {
	CustomerName string
	OrderPrice   int
	OrderItems   string
	IsPaid       bool
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbClient, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := dbClient.Query("SELECT * FROM orders")

	for result.Next() {
		order := Order{}

		result.Scan(&order.CustomerName, &order.OrderPrice, &order.OrderItems, &order.IsPaid)

		fmt.Println(order)
	}
}
