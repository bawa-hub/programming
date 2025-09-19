package models

import (
	"fmt"
	"sync"
	"time"
)

// Time Entry
type TimeEntry struct {
	ID          string
	Task        *Task
	User        *User
	Duration    time.Duration
	Date        time.Time
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	mu          sync.RWMutex
}

func NewTimeEntry(task *Task, user *User, duration time.Duration, date time.Time, description string) *TimeEntry {
	return &TimeEntry{
		ID:          fmt.Sprintf("TE%d", time.Now().UnixNano()),
		Task:        task,
		User:        user,
		Duration:    duration,
		Date:        date,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (te *TimeEntry) UpdateDuration(duration time.Duration) {
	te.mu.Lock()
	defer te.mu.Unlock()
	te.Duration = duration
	te.UpdatedAt = time.Now()
}

func (te *TimeEntry) UpdateDescription(description string) {
	te.mu.Lock()
	defer te.mu.Unlock()
	te.Description = description
	te.UpdatedAt = time.Now()
}

func (te *TimeEntry) GetDuration() time.Duration {
	te.mu.RLock()
	defer te.mu.RUnlock()
	return te.Duration
}

func (te *TimeEntry) GetDurationInHours() float64 {
	te.mu.RLock()
	defer te.mu.RUnlock()
	return te.Duration.Hours()
}

func (te *TimeEntry) GetDurationInMinutes() float64 {
	te.mu.RLock()
	defer te.mu.RUnlock()
	return te.Duration.Minutes()
}

func (te *TimeEntry) GetUser() *User {
	te.mu.RLock()
	defer te.mu.RUnlock()
	return te.User
}

func (te *TimeEntry) GetTask() *Task {
	te.mu.RLock()
	defer te.mu.RUnlock()
	return te.Task
}

func (te *TimeEntry) GetDate() time.Time {
	te.mu.RLock()
	defer te.mu.RUnlock()
	return te.Date
}

func (te *TimeEntry) GetDescription() string {
	te.mu.RLock()
	defer te.mu.RUnlock()
	return te.Description
}

func (te *TimeEntry) CanBeEditedBy(user *User) bool {
	te.mu.RLock()
	defer te.mu.RUnlock()
	
	// User can edit their own time entries
	// Project managers and admins can edit any time entry
	return te.User.ID == user.ID || 
		   user.Role == Admin || 
		   user.Role == ProjectManager
}

func (te *TimeEntry) CanBeDeletedBy(user *User) bool {
	te.mu.RLock()
	defer te.mu.RUnlock()
	
	// User can delete their own time entries
	// Project managers and admins can delete any time entry
	return te.User.ID == user.ID || 
		   user.Role == Admin || 
		   user.Role == ProjectManager
}

// Time Report
type TimeReport struct {
	UserID      string
	ProjectID   string
	TotalHours  float64
	Entries     []*TimeEntry
	DateRange   DateRange
	GeneratedAt time.Time
}

type DateRange struct {
	StartDate time.Time
	EndDate   time.Time
}

func NewTimeReport(userID, projectID string, entries []*TimeEntry, dateRange DateRange) *TimeReport {
	totalHours := 0.0
	for _, entry := range entries {
		totalHours += entry.GetDurationInHours()
	}
	
	return &TimeReport{
		UserID:      userID,
		ProjectID:   projectID,
		TotalHours:  totalHours,
		Entries:     entries,
		DateRange:   dateRange,
		GeneratedAt: time.Now(),
	}
}

func (tr *TimeReport) GetTotalHours() float64 {
	return tr.TotalHours
}

func (tr *TimeReport) GetEntryCount() int {
	return len(tr.Entries)
}

func (tr *TimeReport) GetAverageHoursPerDay() float64 {
	days := tr.DateRange.EndDate.Sub(tr.DateRange.StartDate).Hours() / 24
	if days == 0 {
		return 0
	}
	return tr.TotalHours / days
}
