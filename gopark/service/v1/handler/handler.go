package handler

import (
	"github.com/karnsl/exercise/gopark/service/v1/admin"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/karnsl/exercise/gopark/service/v1/register"
)

// Input a router input
type Input struct {
	Db *sql.DB
}

// Router app router
func (input Input) Router() *gin.Engine {

	reg := register.Input{
		Db: input.Db,
	}

	admin := admin.Input{
		Db: input.Db,
	}

	r := gin.Default()

	r.POST("/register/email", reg.Email)
	r.POST("/admin/place", admin.AddPlace)
	r.POST("/admin/lot", admin.AddParkingLot)
	r.GET("/accounts", reg.ListAccounts)
	
	return r
}
