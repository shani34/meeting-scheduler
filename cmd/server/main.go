package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shani34/meeting-scheduler/api/handlers"
	"github.com/shani34/meeting-scheduler/api/services"
	"github.com/shani34/meeting-scheduler/internal/config"
	"github.com/shani34/meeting-scheduler/internal/database"
	"github.com/shani34/meeting-scheduler/internal/repository"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := database.NewConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	eventRepo := repository.NewEventRepository(db)
	availabilityRepo := repository.NewAvailabilityRepository(db)

	// Initialize services
	scheduler := services.NewSchedulerService()

	// Initialize handlers
	eventHandler := handlers.NewEventHandler(eventRepo, availabilityRepo, scheduler)

	// Initialize router
	router := gin.Default()

	// Event routes
	router.POST("/events", eventHandler.CreateEvent)
	router.GET("/events", eventHandler.GetEvent)
	router.PUT("/events", eventHandler.UpdateEvent)
	router.DELETE("/events", eventHandler.DeleteEvent)

	// Availability routes
	router.POST("/availabilities", eventHandler.SubmitAvailability)
	router.GET("/events/optimal-slots", eventHandler.GetOptimalTimeSlots)

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 