package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karnsl/exercise/gopark/service/v1/handler"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	const (
		dbUser     = "postgres"
		dbPassword = "P@ssw0rd"
		dbHost     = "localhost"
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

	router := h.Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/lot/list", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
