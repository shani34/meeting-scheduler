-- Create events table
CREATE TABLE IF NOT EXISTS events (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration INTEGER NOT NULL, -- Duration in minutes
    created_by VARCHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create event_time_slots table
CREATE TABLE IF NOT EXISTS event_time_slots (
    id VARCHAR(36) PRIMARY KEY,
    event_id VARCHAR(36) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    time_zone VARCHAR(50) NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

-- Create participant_availabilities table
CREATE TABLE IF NOT EXISTS participant_availabilities (
    id VARCHAR(36) PRIMARY KEY,
    event_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

-- Create availability_time_slots table
CREATE TABLE IF NOT EXISTS availability_time_slots (
    id VARCHAR(36) PRIMARY KEY,
    availability_id VARCHAR(36) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    time_zone VARCHAR(50) NOT NULL,
    FOREIGN KEY (availability_id) REFERENCES participant_availabilities(id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_event_time_slots_event_id ON event_time_slots(event_id);
CREATE INDEX IF NOT EXISTS idx_participant_availabilities_event_id ON participant_availabilities(event_id);
CREATE INDEX IF NOT EXISTS idx_availability_time_slots_availability_id ON availability_time_slots(availability_id); 