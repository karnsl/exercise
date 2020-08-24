package register

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Account a model of table accounts.
type Account struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Channel     string `json:"channel"`
	DisplayName string `json:"display_name"`
	Active      bool   `json:"active"`
}

// Input API
type Input struct {
	Db *sql.DB
}

// Email register via email
func (input Input) Email(c *gin.Context) {
	sqlStr := "INSERT INTO accounts VALUES($1, $2, $3, $4, $5)"

	var account Account
	err := c.BindJSON(&account)
	if err != nil {
		log.Println("Bind JSON failed...", err)
		c.JSON(http.StatusBadRequest, "Bind JSON failed...")
		return
	}

	cli := Client{}

	if !userExists(cli, account.DisplayName, "https://jsonplaceholder.typicode.com/todos/1") {
		_, err = input.Db.Exec(sqlStr, account.Username, account.Password, "email", account.DisplayName, true)
		if err != nil {
			log.Println("Failed to register via email...", err)
			c.JSON(http.StatusInternalServerError, "Failed to register via email...")
			return
		}
	} else {
		log.Println("User already exist")
		c.JSON(http.StatusNotAcceptable, "User already exist")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
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

// User Http Get Response
type User struct {
	UserID    int    `json:"user_id"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func userExists(c Clienter, title string, url string) bool {

	bodyBytes, err := c.httpGet(url)
	if err != nil {
		log.Print(err)
	}

	// Convert response body to Todo struct
	var user User
	json.Unmarshal(bodyBytes, &user)
	// fmt.Printf("API Response as struct %+v\n", user)

	return title == user.Title
}

// Clienter client interface
type Clienter interface {
	httpGet(url string) ([]byte, error)
}

//Client client
type Client struct{}

func (cli Client) httpGet(url string) ([]byte, error) {
	// fmt.Println("1. Performing Http Get...")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
