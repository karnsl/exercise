package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/karnsl/exercise/gopark/service/v1/model"

	"github.com/gin-gonic/gin"
	"github.com/karnsl/exercise/gopark/service/v1/model/model"
)

// Input API
type Input struct {
	Db *sql.DB
}

// AddPlace a method to add new place
func (input Input) AddPlace(c *gin.Context) {
	sqlStr := "INSERT INTO places VALUES $1"

	var place model.Place

	err := c.BindJSON(&place)
	if err != nil {
		log.Println("Bind JSON place failed...", err)
		c.JSON(http.StatusBadRequest, "Bind JSON place failed")
	} else {
		_, err = input.Db.Exec(sqlStr, place.Name)
		if err != nil {
			log.Println("Failed to add a new place...", err)
			c.JSON(http.StatusInternalServerError, "Failed to add a new place...")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	}
}

// AddParkingLot a method to add new parking lot
func (input Input) AddParkingLot(c *gin.Context) {
	sqlStr := "INSERT INTO lot VALUES ($1 $2 $3 $4 $5 $6)"

	var lot model.Lot

	err := c.BindJSON(&lot)
	if err != nil {
		log.Println("Bind JSON place failed...", err)
		c.JSON(http.StatusBadRequest, "Bind JSON place failed")
	} else {
		_, err = input.Db.Exec(sqlStr, lot.PlaceID, lot.Building, lot.Floor, lot.Zone, lot.Number, nil)
		if err != nil {
			log.Println("Failed to add a new place...", err)
			c.JSON(http.StatusInternalServerError, "Failed to add a new place...")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	}
}
