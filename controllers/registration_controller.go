package controllers

import (
	"net/http"
	"strconv"

	"github.com/alawii17/goEvent_be/config"
	"github.com/alawii17/goEvent_be/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	eventIDParam := c.Param("event_id")
	eventID, err := strconv.ParseUint(eventIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event models.Event
	if err := config.DB.First(&event, eventID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var existingRegistration models.Registration
	if err := config.DB.Where("user_id = ? AND event_id = ?", userID, eventID).First(&existingRegistration).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already registered for this event"})
		return
	}

	registration := models.Registration{
		UserID:  userID.(uint),
		EventID: uint(eventID),
	}
	if err := config.DB.Create(&registration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register for event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered for event successfully!"})
}

func GetUserRegistrations(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var registrations []models.Registration
	if err := config.DB.Where("user_id = ?", userID).Find(&registrations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve registrations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"registrations": registrations})
}
