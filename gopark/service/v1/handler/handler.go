package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/karnsl/exercise/gopark/service/v1/register"
)

// Router app router
func Router(db *sql.DB) {

	reg := register.Input{
		Db: db,
	}

	r := gin.Default()

	r.POST("/register/email", reg.Email)
	r.GET("/accounts", reg.ListAccounts)
	r.Run()
}
