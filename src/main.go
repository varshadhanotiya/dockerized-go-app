package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Define PostgreSQL connection parameters
const connStr = "user=postgres password=postgres dbname=nifty host=10.11.96.120 port=5432 sslmode=disable"

type Item struct {
	Date  time.Time
	Price float64
}

func handler(c *gin.Context) {
	fmt.Printf("entered handler\n")

	var inputData struct {
		StockName string `form:"stockName" binding:"required"`
		StartDate string `form:"startDate" binding:"required"`
		EndDate   string `form:"endDate" binding:"required"`
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Error in connecting to postgres\n", err)
		return
	}
	defer db.Close()

	fmt.Printf("connection request to db sent \n")

	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(inputData)

	// convert from string to time.Time object
	layout := "2006-01-02"

	parsedStartDate, errStartDate := time.Parse(layout, inputData.StartDate)
	if errStartDate != nil {
		// fmt.Println("Error:", errStartDate)
		log.Error("start date is in wrong format")
		return
	}

	parsedEndDate, errEndDate := time.Parse(layout, inputData.EndDate)
	if errEndDate != nil {
		log.Error("end date is in wrong format")
		// fmt.Println("Error:", errEndDate)
		return
	}

	rows, err := db.Query("SELECT nifty_date, \""+inputData.StockName+"\" FROM summary where nifty_date >= $1 and nifty_date <= $2 ",
		parsedStartDate, parsedEndDate)
	if err != nil {
		log.Error("Error in running sql query", err)
		// log.Fatal(err)
		return
	}
	defer rows.Close()

	var items [2]Item

	// Process the query results
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Date, &item.Price); err != nil {
			log.Fatal(err)
		}
		if items[0].Price < item.Price || items[0].Price == 0 {
			items[0].Price = item.Price
			items[0].Date = item.Date
		}
		if items[1].Price > item.Price || items[1].Price == 0 {
			items[1].Price = item.Price
			items[1].Date = item.Date
		}
	}

	fmt.Println(items)
	c.JSON(http.StatusOK, items)

}

func main() {
	// Gin connection
	r := gin.Default()
	r.GET("/getdata", handler)
	r.Run(":8080")
}
