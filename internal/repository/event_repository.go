package repository

import (
	"database/sql"
	"time"

	"github.com/shani34/meeting-scheduler/api/models"
)

// EventRepository handles database operations for events
type EventRepository struct {
	db *sql.DB
}

// NewEventRepository creates a new instance of EventRepository
func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

// CreateEvent creates a new event in the database
func (r *EventRepository) CreateEvent(event *models.Event) error {
	query := `
		INSERT INTO events (id, title, duration, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query,
		event.ID,
		event.Title,
		event.Duration,
		event.CreatedBy,
		event.CreatedAt,
		event.UpdatedAt,
	)
	if err != nil {
		return err
	}

	// Insert time slots
	for _, slot := range event.TimeSlots {
		slotQuery := `
			INSERT INTO event_time_slots (event_id, start_time, end_time, time_zone)
			VALUES ($1, $2, $3, $4)
		`
		_, err = r.db.Exec(slotQuery,
			event.ID,
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

// GetEvent retrieves an event by ID
func (r *EventRepository) GetEvent(id string) (*models.Event, error) {
	event := &models.Event{}
	query := `
		SELECT id, title, duration, created_by, created_at, updated_at
		FROM events
		WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&event.ID,
		&event.Title,
		&event.Duration,
		&event.CreatedBy,
		&event.CreatedAt,
		&event.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Get time slots
	slotsQuery := `
		SELECT start_time, end_time, time_zone
		FROM event_time_slots
		WHERE event_id = $1
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
		event.TimeSlots = append(event.TimeSlots, slot)
	}

	return event, nil
}

// UpdateEvent updates an existing event
func (r *EventRepository) UpdateEvent(event *models.Event) error {
	query := `
		UPDATE events
		SET title = $1, duration = $2, updated_at = $3
		WHERE id = $4
	`
	_, err := r.db.Exec(query,
		event.Title,
		event.Duration,
		time.Now(),
		event.ID,
	)
	if err != nil {
		return err
	}

	// Delete existing time slots
	_, err = r.db.Exec("DELETE FROM event_time_slots WHERE event_id = $1", event.ID)
	if err != nil {
		return err
	}

	// Insert new time slots
	for _, slot := range event.TimeSlots {
		slotQuery := `
			INSERT INTO event_time_slots (event_id, start_time, end_time, time_zone)
			VALUES ($1, $2, $3, $4)
		`
		_, err = r.db.Exec(slotQuery,
			event.ID,
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

// DeleteEvent deletes an event
func (r *EventRepository) DeleteEvent(id string) error {
	// Delete time slots first
	_, err := r.db.Exec("DELETE FROM event_time_slots WHERE event_id = $1", id)
	if err != nil {
		return err
	}

	// Delete the event
	query := "DELETE FROM events WHERE id = $1"
	_, err = r.db.Exec(query, id)
	return err
}

// GetParticipantAvailabilities retrieves all participant availabilities for an event
func (r *EventRepository) GetParticipantAvailabilities(eventID string) ([]models.ParticipantAvailability, error) {
	query := `
		SELECT id, event_id, user_id, created_at, updated_at
		FROM participant_availabilities
		WHERE event_id = $1
	`
	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var availabilities []models.ParticipantAvailability
	for rows.Next() {
		var availability models.ParticipantAvailability
		err := rows.Scan(
			&availability.ID,
			&availability.EventID,
			&availability.UserID,
			&availability.CreatedAt,
			&availability.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Get time slots for this availability
		slotsQuery := `
			SELECT start_time, end_time, time_zone
			FROM availability_time_slots
			WHERE availability_id = $1
		`
		slotRows, err := r.db.Query(slotsQuery, availability.ID)
		if err != nil {
			return nil, err
		}

		for slotRows.Next() {
			var slot models.TimeSlot
			err := slotRows.Scan(&slot.StartTime, &slot.EndTime, &slot.TimeZone)
			if err != nil {
				slotRows.Close()
				return nil, err
			}
			availability.TimeSlots = append(availability.TimeSlots, slot)
		}
		slotRows.Close()

		availabilities = append(availabilities, availability)
	}

	return availabilities, nil
} 