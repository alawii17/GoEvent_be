package routes

import (
	"github.com/alawii17/goEvent_be/controllers"
	"github.com/alawii17/goEvent_be/middleware"
	"github.com/gin-gonic/gin"
)

func EventRoutes(r *gin.Engine) {
	eventGroup := r.Group("/events")
	{
		eventGroup.GET("/", controllers.GetEvents)
		eventGroup.GET("/:id", controllers.GetEventByID)
		eventGroup.POST("/",  middleware.AuthMiddleware(), controllers.CreateEvent)
		eventGroup.PUT("/:id", middleware.AuthMiddleware(), controllers.UpdateEvent)
		eventGroup.DELETE("/:id",  middleware.AuthMiddleware(), controllers.DeleteEvent)
	}
}