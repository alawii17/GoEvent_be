package main

import (
	"fmt"

	"github.com/alawii17/goEvent_be/config"
	"github.com/alawii17/goEvent_be/models"
	"github.com/alawii17/goEvent_be/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Registration{})

	r := gin.Default()

	// Konfigurasi CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Izinkan frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Izinkan header Authorization
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to GO EVENT API"})
	})

	// Tambahkan routes
	routes.UserRoutes(r)
	routes.EventRoutes(r)
	routes.RegistrationRoutes(r)

	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}