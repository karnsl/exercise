package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Place a model of table places
type Place struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
}

// Lot a model of table lot
type Lot struct {
	ID int16 `json:"id"`
	PlaceID int16 `json:"place_id"`
	Building string `json:"building"`
	Floor string `json:"floor"`
	Zone string `json:"zone"`
	Number int8 `json:"number"`
	Username string `json:"username"`
}

// Input API
type Input struct {
	Db *sql.DB
}

// AddPlace a method to add new place
func (input Input) AddPlace(c *gin.Context) {
	sqlStr := "INSERT INTO places VALUES $1"

	var place Place

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

	var lot Lot

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