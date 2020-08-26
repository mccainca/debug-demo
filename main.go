package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Square struct {
	XVal int `json:"x,string"`
	YVal int `json:"y,string"`
}

func main() {
	router := gin.Default()

	router.POST("/square", square)

	router.Run()
}

func square(c *gin.Context) {
	var square *Square
	err := c.ShouldBindJSON(&square)
	if err != nil {
		log.Println(err.Error())
	}
	go store(square)
	squared := square.XVal * square.YVal

	c.JSON(200, squared)

}

func store(square *Square) error {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in store", r)
		}
	}()

	database, _ := sql.Open("sqlite3", "./square.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS square (id INTEGER PRIMARY KEY, x INTEGER, y INTEGER, timestamp TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO square (x, y) VALUES (?, ?, datetime('now'))")
	statement.Exec(square.XVal, square.YVal)
	return nil
}
