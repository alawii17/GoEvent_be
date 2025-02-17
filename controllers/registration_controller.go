package controllers

import (
	"net/http"

	"github.com/alawii17/goEvent_be/config"
	"github.com/alawii17/goEvent_be/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(c *gin.Context) {
	userID,_ := c.Get("user_id")
	eventID := c.Param("event_id")

	registration := models.Registration{
		UserID: userID.(uint),
		EventID: uint(eventID[0] - '0'),
	}

	config.DB.Create(&registration)
	c.JSON(http.StatusOK, gin.H{"message": "Registered for event successfully!"})
}

func GetUserRegistrations(c *gin.Context) {
	userID,_ := c.Get("user_id")
	var registrations []models.Registration
	config.DB.Where("user_id = ?", userID).Find(&registrations)

	c.JSON(http.StatusOK, registrations)
}