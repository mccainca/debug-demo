package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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
	fmt.Printf("%+v", square)

	squared := square.XVal * square.YVal

	c.JSON(200, squared)

}
