package main

import (
	"fmt"

	"github.com/alawii17/goEvent_be/config"
	"github.com/alawii17/goEvent_be/models"
	"github.com/alawii17/goEvent_be/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
    config.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Registration{})

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message" : "Welcome to GO EVENT API"})
    })

    fmt.Println("Server running on port 8080")
    routes.UserRoutes(r)
    r.Run(":8080")

    routes.UserRoutes(r)
    routes.EventRoutes(r)
    routes.RegistrationRoutes(r)
}