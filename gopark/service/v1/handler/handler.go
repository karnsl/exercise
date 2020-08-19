package handler

import (
	"database/sql"

	"github.com/karnsl/exercise/gopark/service/v1/admin"

	"github.com/gin-gonic/gin"
	"github.com/karnsl/exercise/gopark/service/v1/register"
	"github.com/karnsl/exercise/gopark/service/v1/reserve"
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

	lot := reserve.Input{
		Db: input.Db,
	}

	r := gin.Default()

	r.POST("/register/email", reg.Email)
	r.POST("/admin/place", admin.AddPlace)
	r.POST("/admin/lot", admin.AddParkingLot)
	r.POST("/lot/reserve", lot.ReserveLot)
	r.POST("/lot/unlock", lot.Unlock)
	r.GET("/lot/list", lot.ListAvailableLot)
	r.GET("/account/list", reg.ListAccounts)

	return r
}
