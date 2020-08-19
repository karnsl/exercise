package register

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karnsl/exercise/gopark/service/v1/model/model"
)

// Input API
type Input struct {
	Db *sql.DB
}

// Email register via email
func (input Input) Email(c *gin.Context) {
	sqlStr := "INSERT INTO accounts VALUES($1, $2, $3, $4, $5)"

	var account model.Account
	err := c.BindJSON(&account)
	if err != nil {
		log.Println("Bind JSON failed...", err)
		c.JSON(http.StatusBadRequest, "Bind JSON failed...")
	} else {
		// log.Println(account)
		_, err = input.Db.Exec(sqlStr, account.Username, account.Password, "email", account.DisplayName, true)
		if err != nil {
			log.Println("Failed to register via email...", err)
			c.JSON(http.StatusInternalServerError, "Failed to register via email...")
		} else {

			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	}
}

// ListAccounts list all accounts
func (input Input) ListAccounts(c *gin.Context) {
	sqlStr := "SELECT username, display_name FROM accounts"
	rows, err := input.Db.Query(sqlStr)
	if err != nil {
		log.Println("Query failed")
	}
	for rows.Next() {
		var username string
		var displayName string
		if err := rows.Scan(&username, &displayName); err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":       "ok",
			"username":     username,
			"display_name": displayName,
		})
	}
	if !rows.NextResultSet() {
		log.Println(rows.Err())
	}
}