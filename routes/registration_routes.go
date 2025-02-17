package routes

import (
	"github.com/alawii17/goEvent_be/controllers"
	"github.com/alawii17/goEvent_be/middleware"
	"github.com/gin-gonic/gin"
)

func RegistrationRoutes(r *gin.Engine) {
	registrationGroup := r.Group("/registrations")
	{
		registrationGroup.POST("/:event_id", middleware.AuthMiddleware(), controllers.RegisterForEvent)
		registrationGroup.GET("/", middleware.AuthMiddleware(), controllers.GetUserRegistrations)
	}
}