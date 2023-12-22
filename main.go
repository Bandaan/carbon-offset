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

func calculateCarbon(c *gin.Context) {
	// Your existing CORS headers
	c.Header("Access-Control-Allow-Origin", "*") // Adjust as needed
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	// Check if it's a preflight request
	if c.Request.Method == http.MethodOptions {
		c.Header("Access-Control-Allow-Methods", "POST")
		c.Header("Access-Control-Max-Age", "86400") // 24 hours
		c.JSON(http.StatusOK, nil)
		return
	}

	var newOffset offset

	if err := c.ShouldBindJSON(&newOffset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	var newPrice float64

	if newOffset.Type == false {
		newPrice = newOffset.Price * 0.1
	} else {
		newPrice = newOffset.Weight * 0.1
	}

	c.JSON(http.StatusOK, gin.H{"OffsetPrice": newPrice})
}

func helloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello world!")
}

func helloMob(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Ten alle tijde MOB niet vergeten stevie")
}

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // Adjust as needed
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Next()
	})

	// Handle OPTIONS request for all routes
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Max-Age", "86400") // 24 hours
		c.JSON(http.StatusOK, nil)
	})

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
