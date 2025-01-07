package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kamil-budzik/hospital-system/auth-service/db"
	"github.com/kamil-budzik/hospital-system/auth-service/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	// Start server
	port := os.Getenv("SERVER_PORT")
	log.Printf("Auth service starting on port %s", port)
	server.Run(":" + port)
}
