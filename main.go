package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type offset struct {
	ID     string  `json:"id"`
	Weight float64 `json:"weight"`
	Price  float64 `json:"price"`
	Type   bool    `json:"type"`
}

type result struct {
	OffsetPrice float64
}

func calculateCarbon(c *gin.Context) {
	var newOffset offset

	if err := c.BindJSON(&newOffset); err != nil {
		return
	}

	if newOffset.Type == false {
		newPrice := newOffset.Price * 0.1
		c.IndentedJSON(http.StatusCreated, newPrice)
	} else {
		newPrice := newOffset.Weight * 0.1
		c.IndentedJSON(http.StatusCreated, newPrice)
	}
}

func helloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello world!")
}

func helloMob(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Ten alle tijde MOB niet vergeten stevie")
}

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	router.GET("/check", helloWorld)
	router.GET("/mob", helloMob)
	router.POST("/offset", calculateCarbon)

	log.Fatal(router.Run("0.0.0.0:" + port))
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
