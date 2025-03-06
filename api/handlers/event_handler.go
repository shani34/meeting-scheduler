package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shani34/meeting-scheduler/api/models"
	"github.com/shani34/meeting-scheduler/api/services"
	"github.com/shani34/meeting-scheduler/internal/repository"
)

// EventHandler handles HTTP requests for event-related operations
type EventHandler struct {
	eventRepo      *repository.EventRepository
	availabilityRepo *repository.AvailabilityRepository
	scheduler      *services.SchedulerService
}

// NewEventHandler creates a new instance of EventHandler
func NewEventHandler(eventRepo *repository.EventRepository, availabilityRepo *repository.AvailabilityRepository, scheduler *services.SchedulerService) *EventHandler {
	return &EventHandler{
		eventRepo:      eventRepo,
		availabilityRepo: availabilityRepo,
		scheduler:      scheduler,
	}
}

// CreateEvent handles the creation of a new event
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set event ID and timestamps
	event.ID = uuid.New().String()
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	// Create event in database
	if err := h.eventRepo.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

// GetEvent handles retrieving an event by ID
func (h *EventHandler) GetEvent(c *gin.Context) {
	eventID := c.Query("id")
	if eventID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event ID is required"})
		return
	}

	event, err := h.eventRepo.GetEvent(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// UpdateEvent handles updating an existing event
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	eventID := c.Query("id")
	if eventID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event ID is required"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure event ID matches
	event.ID = eventID
	event.UpdatedAt = time.Now()

	if err := h.eventRepo.UpdateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// DeleteEvent handles deleting an event
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	eventID := c.Query("id")
	if eventID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event ID is required"})
		return
	}

	if err := h.eventRepo.DeleteEvent(eventID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.Status(http.StatusNoContent)
}

// SubmitAvailability handles submitting participant availability
func (h *EventHandler) SubmitAvailability(c *gin.Context) {
	eventID := c.Query("event_id")
	if eventID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event ID is required"})
		return
	}

	var availability models.ParticipantAvailability
	if err := c.ShouldBindJSON(&availability); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Set availability ID and timestamps
	availability.ID = uuid.New().String()
	availability.EventID = eventID
	availability.CreatedAt = time.Now()
	availability.UpdatedAt = time.Now()

	if err := h.availabilityRepo.CreateAvailability(&availability); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit availability"})
		return
	}

	c.JSON(http.StatusCreated, availability)
}

// GetOptimalTimeSlots handles finding optimal time slots for an event
func (h *EventHandler) GetOptimalTimeSlots(c *gin.Context) {
	eventID := c.Query("event_id")
	if eventID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event ID is required"})
		return
	}

	// Get event details
	event, err := h.eventRepo.GetEvent(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Get all participant availabilities
	availabilities, err := h.eventRepo.GetParticipantAvailabilities(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get participant availabilities"})
		return
	}

	// Find optimal time slots
	recommendations := h.scheduler.FindOptimalTimeSlots(event, availabilities)

	c.JSON(http.StatusOK, recommendations)
}

// CreateAvailability handles the creation of participant availability
func (h *EventHandler) CreateAvailability(c *gin.Context) {
	var req models.CreateAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	availability := &models.ParticipantAvailability{
		ID:        uuid.New().String(),
		EventID:   req.EventID,
		UserID:    c.GetString("user_id"), // Assuming user_id is set by auth middleware
		TimeSlots: req.TimeSlots,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// TODO: Save availability to database
	c.JSON(http.StatusCreated, availability)
}

// UpdateAvailability updates participant availability
func (h *EventHandler) UpdateAvailability(c *gin.Context) {
	availabilityID := c.Param("id")
	var req models.UpdateAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update availability in database
	c.JSON(http.StatusOK, gin.H{"message": "Availability updated", "id": availabilityID})
}

// DeleteAvailability deletes participant availability
func (h *EventHandler) DeleteAvailability(c *gin.Context) {
	availabilityID := c.Param("id")
	// TODO: Delete availability from database
	c.JSON(http.StatusOK, gin.H{"message": "Availability deleted", "id": availabilityID})
} 