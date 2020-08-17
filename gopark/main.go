package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

const (
	dbUser     = "postgres"
	dbPassword = "P@ssw0rd"
	dbHost     = "192.168.1.51"
	dbPort     = 5432
	dbName     = "Gopark"
)

// Accounts is a representation of table accounts.
type Accounts struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Channel     string `json:"channel"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
}

func main() {

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Println("Failed to open connection to db")
		panic(err)
	}
	fmt.Println("Database Connected.")
	defer db.Close()

	r := gin.Default()

	r.POST("/register/email", registerWithEmail)
	r.GET("/accounts", listAccounts)
	r.Run()
}

func registerWithEmail(c *gin.Context) {
	sqlStr := "INSERT INTO accounts VALUES($1, $2, $3, $4, $5)"

	var accounts Accounts
	err := c.BindJSON(&accounts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(accounts)
	_, err = db.Exec(sqlStr, accounts.Username, accounts.Password, "email", accounts.DisplayName, true)
	if err != nil {
		panic(err)
	}
}

func listAccounts(c *gin.Context) {
	sqlStr := "SELECT username, display_name FROM accounts"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("Query failed")
		panic(err)
	}
	for rows.Next() {
		var username string
		var displayName string
		if err := rows.Scan(&username, &displayName); err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"username":    username,
			"displayName": displayName,
		})
	}
	if !rows.NextResultSet() {
		fmt.Println(rows.Err())
	}
}
