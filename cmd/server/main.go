package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shanikumar/meeting-scheduler/api/handlers"
	"github.com/shanikumar/meeting-scheduler/api/services"
)

func main() {
	// Initialize services
	schedulerService := services.NewSchedulerService()
	eventHandler := handlers.NewEventHandler(schedulerService)

	// Set up router
	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// API routes
	api := router.Group("/api/v1")
	{
		// Event routes
		events := api.Group("/events")
		{
			events.POST("", eventHandler.CreateEvent)
			events.GET("/:id", eventHandler.GetEvent)
			events.PUT("/:id", eventHandler.UpdateEvent)
			events.DELETE("/:id", eventHandler.DeleteEvent)
			events.GET("/:id/recommendations", eventHandler.GetRecommendedTimeSlots)
		}

		// Availability routes
		availabilities := api.Group("/availabilities")
		{
			availabilities.POST("", eventHandler.CreateAvailability)
			availabilities.PUT("/:id", eventHandler.UpdateAvailability)
			availabilities.DELETE("/:id", eventHandler.DeleteAvailability)
		}
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 