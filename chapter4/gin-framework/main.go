package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/guozhe001/brwswg/chapter4/dbutils"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

var DB *sql.DB

// StationResource holds information about locations
type StationResource struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`
}

func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")
	err := DB.QueryRow("select ID, NAME, OPENING_TIME, CLOSING_TIME from station where ID = ?", id).
		Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": station})
	}
}

func CreateStation(c *gin.Context) {
	var station StationResource
	if err := c.BindJSON(&station); err == nil {
		statement, _ := DB.Prepare("insert into station(NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
		result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			id, _ := result.LastInsertId()
			station.ID = int(id)
			c.JSON(200, gin.H{"result": station})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func RemoveStation(c *gin.Context) {
	id := c.Param("station_id")
	statement, _ := DB.Prepare("delete from station where id = ?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.String(http.StatusOK, "OK")
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "/Users/guozhe/code/my-project/brwswg/chapter4/railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)
	r := gin.Default()
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)
	r.Run(":8088")
}

// basicgin.go
//func main() {
//	r := gin.Default()
//	r.GET("/pingTime", func(c *gin.Context) {
//		c.JSON(200, gin.H{"serverTime": time.Now().UTC()})
//	})
//	r.Run(":8088")
//}
