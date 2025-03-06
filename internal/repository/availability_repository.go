package repository

import (
	"database/sql"
	"time"

	"github.com/shani34/meeting-scheduler/api/models"
)

// AvailabilityRepository handles database operations for participant availabilities
type AvailabilityRepository struct {
	db *sql.DB
}

// NewAvailabilityRepository creates a new instance of AvailabilityRepository
func NewAvailabilityRepository(db *sql.DB) *AvailabilityRepository {
	return &AvailabilityRepository{db: db}
}

// CreateAvailability creates a new participant availability in the database
func (r *AvailabilityRepository) CreateAvailability(availability *models.ParticipantAvailability) error {
	query := `
		INSERT INTO participant_availabilities (id, event_id, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(query,
		availability.ID,
		availability.EventID,
		availability.UserID,
		availability.CreatedAt,
		availability.UpdatedAt,
	)
	if err != nil {
		return err
	}

	// Insert time slots
	for _, slot := range availability.TimeSlots {
		slotQuery := `
			INSERT INTO availability_time_slots (availability_id, start_time, end_time, time_zone)
			VALUES ($1, $2, $3, $4)
		`
		_, err = r.db.Exec(slotQuery,
			availability.ID,
			slot.StartTime,
			slot.EndTime,
			slot.TimeZone,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetAvailability retrieves a participant availability by ID
func (r *AvailabilityRepository) GetAvailability(id string) (*models.ParticipantAvailability, error) {
	availability := &models.ParticipantAvailability{}
	query := `
		SELECT id, event_id, user_id, created_at, updated_at
		FROM participant_availabilities
		WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&availability.ID,
		&availability.EventID,
		&availability.UserID,
		&availability.CreatedAt,
		&availability.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Get time slots
	slotsQuery := `
		SELECT start_time, end_time, time_zone
		FROM availability_time_slots
		WHERE availability_id = $1
	`
	rows, err := r.db.Query(slotsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var slot models.TimeSlot
		err := rows.Scan(&slot.StartTime, &slot.EndTime, &slot.TimeZone)
		if err != nil {
			return nil, err
		}
		availability.TimeSlots = append(availability.TimeSlots, slot)
	}

	return availability, nil
}

// UpdateAvailability updates an existing participant availability
func (r *AvailabilityRepository) UpdateAvailability(availability *models.ParticipantAvailability) error {
	query := `
		UPDATE participant_availabilities
		SET updated_at = $1
		WHERE id = $2
	`
	_, err := r.db.Exec(query,
		time.Now(),
		availability.ID,
	)
	if err != nil {
		return err
	}

	// Delete existing time slots
	_, err = r.db.Exec("DELETE FROM availability_time_slots WHERE availability_id = $1", availability.ID)
	if err != nil {
		return err
	}

	// Insert new time slots
	for _, slot := range availability.TimeSlots {
		slotQuery := `
			INSERT INTO availability_time_slots (availability_id, start_time, end_time, time_zone)
			VALUES ($1, $2, $3, $4)
		`
		_, err = r.db.Exec(slotQuery,
			availability.ID,
			slot.StartTime,
			slot.EndTime,
			slot.TimeZone,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteAvailability deletes a participant availability
func (r *AvailabilityRepository) DeleteAvailability(id string) error {
	// Delete time slots first
	_, err := r.db.Exec("DELETE FROM availability_time_slots WHERE availability_id = $1", id)
	if err != nil {
		return err
	}

	// Delete the availability
	query := "DELETE FROM participant_availabilities WHERE id = $1"
	_, err = r.db.Exec(query, id)
	return err
} 