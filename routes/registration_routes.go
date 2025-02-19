package routes

import (
	"github.com/alawii17/goEvent_be/controllers"
	"github.com/alawii17/goEvent_be/middleware"
	"github.com/gin-gonic/gin"
)

func RegistrationRoutes(r *gin.Engine) {
	registrationGroup := r.Group("/registrations", middleware.AuthMiddleware()) // Middleware diterapkan ke semua route dalam group
	{
		registrationGroup.POST("/:event_id", controllers.RegisterForEvent) // Mendaftar event
		registrationGroup.GET("/", controllers.GetUserRegistrations) // Mendapatkan daftar event yang didaftarkan user
	}
}
