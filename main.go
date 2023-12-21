package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func main() {
	router := gin.Default()
	router.POST("/offset", calculateCarbon)

	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
