package reserve

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Reservation a message request to reserve parking lot
type Reservation struct {
	Username string `json:"username"`
	ID       int16  `json:"id"`
	PlaceID  int16  `json:"place_id"`
}

// Input API
type Input struct {
	Db *sql.DB
}

// ReserveLot to reserve parking lot
func (input Input) ReserveLot(c *gin.Context) {
	sqlStr := "UPDATE lot SET username = $1 WHERE id = $2 AND place_id = $3"

	var reservation Reservation
	err := c.BindJSON(&reservation)
	if err != nil {
		log.Println("Bind JSON request Reservation failed...", err)
		c.JSON(http.StatusBadRequest, "Bind JSON request Reservation failed...")
	} else {
		// log.Println(account)
		_, err = input.Db.Exec(sqlStr, reservation.Username, reservation.ID, reservation.PlaceID)
		if err != nil {
			log.Println("Failed to reserve parking lot...", err)
			c.JSON(http.StatusInternalServerError, "Failed to reserve parking lot...")
		} else {

			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	}
}

// ListAvailableLot list all available parking lots
func (input Input) ListAvailableLot(c *gin.Context) {
	sqlStr := "SELECT * FROM lot WHERE username IS NULL"
	rows, err := input.Db.Query(sqlStr)
	if err != nil {
		log.Println("Query failed")
	}
	for rows.Next() {
		var id int16
		var placeID int16
		var building string
		var floor string
		var zone string
		var number int8
		var username string
		if err := rows.Scan(&id, &placeID, &building, &floor, &zone, &number, &username); err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"place_id": placeID,
			"building": building,
			"floor":    floor,
			"zone":     zone,
			"number":   number,
			"username": username,
		})
	}
	if !rows.NextResultSet() {
		log.Println(rows.Err())
	}
}
