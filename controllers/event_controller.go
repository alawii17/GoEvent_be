package controllers

import (
	"net/http"

	"github.com/alawii17/goEvent_be/config"
	"github.com/alawii17/goEvent_be/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	var events []models.Event
	config.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

func GetEventByID(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	if err := config.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	userID,_ := c.Get("user_id")

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.CreatedBy = userID.(uint)
	config.DB.Create(&event)

	c.JSON(http.StatusCreated, gin.H{"message": "Event created succesfully!", "event": event})
}

func UpdateEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	if err := config.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event not found"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	if event.CreatedBy != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this event"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&event).Where("id = ?", id).Updates(event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!", "event": event})
}


func DeleteEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	if err := config.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not Found"})
		return
	}

	if event.CreatedBy != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this event"})
		return
	}

	config.DB.Delete(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}