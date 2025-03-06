package services

import (
	"sort"
	"time"

	"github.com/shani34/meeting-scheduler/api/models"
)

// SchedulerService handles the business logic for finding optimal meeting times
type SchedulerService struct {
	// Add any dependencies here (e.g., database client)
}

// NewSchedulerService creates a new instance of SchedulerService
func NewSchedulerService() *SchedulerService {
	return &SchedulerService{}
}

// FindOptimalTimeSlots finds the best meeting time slots based on all participants' availability
func (s *SchedulerService) FindOptimalTimeSlots(
	event *models.Event,
	participantAvailabilities []models.ParticipantAvailability,
) []models.RecommendedTimeSlot {
	if len(participantAvailabilities) == 0 {
		return nil
	}

	// Convert all time slots to UTC for comparison
	eventSlots := convertToUTC(event.TimeSlots)
	participantSlots := make([][]models.TimeSlot, len(participantAvailabilities))
	for i, pa := range participantAvailabilities {
		participantSlots[i] = convertToUTC(pa.TimeSlots)
	}

	// Find overlapping time slots
	overlappingSlots := findOverlappingSlots(eventSlots, participantSlots)

	// Calculate scores for each overlapping slot
	recommendations := make([]models.RecommendedTimeSlot, 0)
	for _, slot := range overlappingSlots {
		participants, missingUsers := s.calculateParticipantAvailability(slot, participantAvailabilities)
		recommendations = append(recommendations, models.RecommendedTimeSlot{
			TimeSlot:     slot,
			Participants: participants,
			MissingUsers: missingUsers,
			Score:        len(participants),
		})
	}

	// Sort recommendations by score (highest first)
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})

	return recommendations
}

// calculateParticipantAvailability determines which participants can attend a given time slot
func (s *SchedulerService) calculateParticipantAvailability(
	slot models.TimeSlot,
	participantAvailabilities []models.ParticipantAvailability,
) ([]string, []string) {
	participants := make([]string, 0)
	missingUsers := make([]string, 0)

	for _, pa := range participantAvailabilities {
		if isSlotAvailable(slot, pa.TimeSlots) {
			participants = append(participants, pa.UserID)
		} else {
			missingUsers = append(missingUsers, pa.UserID)
		}
	}

	return participants, missingUsers
}

// Helper functions

func convertToUTC(slots []models.TimeSlot) []models.TimeSlot {
	utcSlots := make([]models.TimeSlot, len(slots))
	for i, slot := range slots {
		loc, _ := time.LoadLocation(slot.TimeZone)
		utcSlots[i] = models.TimeSlot{
			StartTime: slot.StartTime.In(loc).UTC(),
			EndTime:   slot.EndTime.In(loc).UTC(),
			TimeZone:  "UTC",
		}
	}
	return utcSlots
}

func findOverlappingSlots(eventSlots []models.TimeSlot, participantSlots [][]models.TimeSlot) []models.TimeSlot {
	overlapping := make([]models.TimeSlot, 0)

	for _, eventSlot := range eventSlots {
		// Find all slots that overlap with the event slot
		for _, participantSlotList := range participantSlots {
			for _, participantSlot := range participantSlotList {
				if isOverlapping(eventSlot, participantSlot) {
					overlapping = append(overlapping, eventSlot)
					break
				}
			}
		}
	}

	return overlapping
}

func isOverlapping(slot1, slot2 models.TimeSlot) bool {
	return slot1.StartTime.Before(slot2.EndTime) && slot2.StartTime.Before(slot1.EndTime)
}

func isSlotAvailable(slot models.TimeSlot, availableSlots []models.TimeSlot) bool {
	for _, availableSlot := range availableSlots {
		if isOverlapping(slot, availableSlot) {
			return true
		}
	}
	return false
} 