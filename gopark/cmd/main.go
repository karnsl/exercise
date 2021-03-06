package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/karnsl/exercise/gopark/service/v1/handler"
	_ "github.com/lib/pq"
)

func main() {

	const (
		dbUser     = "postgres"
		dbPassword = "P@ssw0rd"
		dbHost     = "172.17.106.61"
		dbPort     = 5432
		dbName     = "Gopark"
	)

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Failed to open connection to db")
	}
	log.Println("Database Connected.")
	defer db.Close()

	h := handler.Input{
		Db: db,
	}

	r := h.Router()
	r.Run(":8080")
}
