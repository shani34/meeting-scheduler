package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shani34/meeting-scheduler/api/models"
	"github.com/shani34/meeting-scheduler/api/services"
)

// EventHandler handles HTTP requests related to events
type EventHandler struct {
	schedulerService *services.SchedulerService
	// Add other dependencies (e.g., database client)
}

// NewEventHandler creates a new instance of EventHandler
func NewEventHandler(schedulerService *services.SchedulerService) *EventHandler {
	return &EventHandler{
		schedulerService: schedulerService,
	}
}

// CreateEvent handles the creation of a new event
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req models.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := &models.Event{
		ID:        uuid.New().String(),
		Title:     req.Title,
		Duration:  req.Duration,
		TimeSlots: req.TimeSlots,
		CreatedBy: c.GetString("user_id"), // Assuming user_id is set by auth middleware
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// TODO: Save event to database
	c.JSON(http.StatusCreated, event)
}

// GetEvent retrieves an event by ID
func (h *EventHandler) GetEvent(c *gin.Context) {
	eventID := c.Param("id")
	// TODO: Fetch event from database
	c.JSON(http.StatusOK, gin.H{"message": "Event retrieved", "id": eventID})
}

// UpdateEvent updates an existing event
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	eventID := c.Param("id")
	var req models.UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update event in database
	c.JSON(http.StatusOK, gin.H{"message": "Event updated", "id": eventID})
}

// DeleteEvent deletes an event
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	eventID := c.Param("id")
	// TODO: Delete event from database
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted", "id": eventID})
}

// GetRecommendedTimeSlots retrieves recommended time slots for an event
func (h *EventHandler) GetRecommendedTimeSlots(c *gin.Context) {
	eventID := c.Param("id")
	
	// TODO: Fetch event and participant availabilities from database
	event := &models.Event{
		ID: eventID,
		// Add other fields
	}
	
	participantAvailabilities := []models.ParticipantAvailability{
		// Add participant availabilities
	}

	recommendations := h.schedulerService.FindOptimalTimeSlots(event, participantAvailabilities)
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