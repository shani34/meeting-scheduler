package tests

import (
	"testing"
	"time"

	"github.com/shanikumar/meeting-scheduler/api/models"
	"github.com/shanikumar/meeting-scheduler/api/services"
	"github.com/stretchr/testify/assert"
)

func TestFindOptimalTimeSlots(t *testing.T) {
	scheduler := services.NewSchedulerService()

	// Create test event
	event := &models.Event{
		ID:       "test-event",
		Title:    "Test Meeting",
		Duration: 60,
		TimeSlots: []models.TimeSlot{
			{
				StartTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
				TimeZone:  "UTC",
			},
		},
	}

	// Create test participant availabilities
	participantAvailabilities := []models.ParticipantAvailability{
		{
			ID:      "availability-1",
			EventID: "test-event",
			UserID:  "user-1",
			TimeSlots: []models.TimeSlot{
				{
					StartTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC),
					TimeZone:  "UTC",
				},
			},
		},
		{
			ID:      "availability-2",
			EventID: "test-event",
			UserID:  "user-2",
			TimeSlots: []models.TimeSlot{
				{
					StartTime: time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC),
					TimeZone:  "UTC",
				},
			},
		},
	}

	// Test finding optimal time slots
	recommendations := scheduler.FindOptimalTimeSlots(event, participantAvailabilities)

	// Assertions
	assert.NotNil(t, recommendations)
	assert.Greater(t, len(recommendations), 0)

	// Check if the recommended slot has the correct participants
	for _, rec := range recommendations {
		assert.Contains(t, rec.Participants, "user-1")
		assert.Contains(t, rec.Participants, "user-2")
	}
}

func TestTimeSlotOverlap(t *testing.T) {
	scheduler := services.NewSchedulerService()

	slot1 := models.TimeSlot{
		StartTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		TimeZone:  "UTC",
	}

	slot2 := models.TimeSlot{
		StartTime: time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC),
		TimeZone:  "UTC",
	}

	// Test overlapping slots
	assert.True(t, scheduler.IsOverlapping(slot1, slot2))

	// Test non-overlapping slots
	slot3 := models.TimeSlot{
		StartTime: time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 1, 15, 0, 0, 0, time.UTC),
		TimeZone:  "UTC",
	}
	assert.False(t, scheduler.IsOverlapping(slot1, slot3))
} 