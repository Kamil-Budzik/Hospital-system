package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "appointment-service",
		})
	})

	// Start server
	port := os.Getenv("SERVER_PORT")
	log.Printf("Appointment service starting on port %s", port)
	r.Run(":" + port)
}
