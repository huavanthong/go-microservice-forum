package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	conStr := "postgres://postgres:admin1234@localhost:5432/discount_service?sslmode=disable"

	fmt.Println("Postgres addr: " + conStr)

	_, err := sqlx.Connect("postgres", conStr)

	if err != nil {
		fmt.Println("Could not connect...")
	} else {
		fmt.Println("Connecting successful")
	}
}
