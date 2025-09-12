package models

import (
	"time"

	"github.com/google/uuid"
)

// EventType defines the type of event
type EventType string

const (
	EventTypePageView EventType = "page_view"
	EventTypeClick    EventType = "click"
	EventTypePurchase EventType = "purchase"
	EventTypeSignup   EventType = "signup"
	EventTypeLogin    EventType = "login"
	EventTypeSearch   EventType = "search"
	EventTypeCustom   EventType = "custom"
)

// Event represents an analytics event
type Event struct {
	ID        string                 `json:"id"`
	Type      EventType              `json:"type"`
	Timestamp time.Time              `json:"timestamp"`
	UserID    string                 `json:"user_id,omitempty"`
	SessionID string                 `json:"session_id,omitempty"`
	Source    string                 `json:"source,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// NewEvent creates a new event
func NewEvent(eventType EventType, userID, sessionID, source string, data map[string]interface{}) *Event {
	return &Event{
		ID:        generateID(),
		Type:      eventType,
		Timestamp: time.Now(),
		UserID:    userID,
		SessionID: sessionID,
		Source:    source,
		Data:      data,
	}
}

// Validate validates the event
func (e *Event) Validate() error {
	if e.ID == "" {
		return ErrEventIDRequired
	}
	if e.Type == "" {
		return ErrEventTypeRequired
	}
	if e.Timestamp.IsZero() {
		return ErrEventTimestampRequired
	}
	return nil
}

// GetAggregateID returns the aggregate ID for the event
func (e *Event) GetAggregateID() string {
	return e.ID
}

// GetTimestamp returns the timestamp of the event
func (e *Event) GetTimestamp() time.Time {
	return e.Timestamp
}

// Event errors
var (
	ErrEventIDRequired        = &EventError{msg: "event ID is required"}
	ErrEventTypeRequired      = &EventError{msg: "event type is required"}
	ErrEventTimestampRequired = &EventError{msg: "event timestamp is required"}
)

// EventError represents an event-related error
type EventError struct {
	msg string
}

func (e *EventError) Error() string {
	return e.msg
}

// generateID generates a unique ID
func generateID() string {
	return uuid.New().String()
}