package models

import "time"

// TimeSlot represents a time slot with start and end times
type TimeSlot struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	TimeZone  string    `json:"time_zone"`
}

// Event represents a meeting event
type Event struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Duration  int        `json:"duration"` // Duration in minutes
	TimeSlots []TimeSlot `json:"time_slots"`
	CreatedBy string     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// ParticipantAvailability represents a participant's available time slots
type ParticipantAvailability struct {
	ID        string     `json:"id"`
	EventID   string     `json:"event_id"`
	UserID    string     `json:"user_id"`
	TimeSlots []TimeSlot `json:"time_slots"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// RecommendedTimeSlot represents a recommended meeting time slot
type RecommendedTimeSlot struct {
	TimeSlot     TimeSlot `json:"time_slot"`
	Participants []string `json:"participants"`
	MissingUsers []string `json:"missing_users"`
	Score        int      `json:"score"`
}

// CreateEventRequest represents the request body for creating an event
type CreateEventRequest struct {
	Title     string     `json:"title" binding:"required"`
	Duration  int        `json:"duration" binding:"required"`
	TimeSlots []TimeSlot `json:"time_slots" binding:"required"`
}

// UpdateEventRequest represents the request body for updating an event
type UpdateEventRequest struct {
	Title     string     `json:"title"`
	Duration  int        `json:"duration"`
	TimeSlots []TimeSlot `json:"time_slots"`
}

// CreateAvailabilityRequest represents the request body for creating participant availability
type CreateAvailabilityRequest struct {
	EventID   string     `json:"event_id" binding:"required"`
	TimeSlots []TimeSlot `json:"time_slots" binding:"required"`
}

// UpdateAvailabilityRequest represents the request body for updating participant availability
type UpdateAvailabilityRequest struct {
	TimeSlots []TimeSlot `json:"time_slots" binding:"required"`
} 