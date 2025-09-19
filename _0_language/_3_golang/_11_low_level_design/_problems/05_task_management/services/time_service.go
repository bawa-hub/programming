package services

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"task_management/models"
)

// Time Service
type TimeService struct {
	timeEntries map[string]*models.TimeEntry
	mu          sync.RWMutex
}

func NewTimeService() *TimeService {
	return &TimeService{
		timeEntries: make(map[string]*models.TimeEntry),
	}
}

// Time Entry Management
func (tms *TimeService) LogTime(taskID string, user *models.User, duration time.Duration, date time.Time, description string) (*models.TimeEntry, error) {
	tms.mu.Lock()
	defer tms.mu.Unlock()
	
	// In a real system, we would validate that the task exists and user has access
	// For demo purposes, we'll create a mock task
	task := &models.Task{ID: taskID}
	
	timeEntry := models.NewTimeEntry(task, user, duration, date, description)
	tms.timeEntries[timeEntry.ID] = timeEntry
	
	return timeEntry, nil
}

func (tms *TimeService) GetTimeEntries(taskID string) []*models.TimeEntry {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var entries []*models.TimeEntry
	for _, entry := range tms.timeEntries {
		if entry.GetTask().ID == taskID {
			entries = append(entries, entry)
		}
	}
	
	// Sort by date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].GetDate().After(entries[j].GetDate())
	})
	
	return entries
}

func (tms *TimeService) GetUserTimeEntries(userID string, startDate, endDate time.Time) []*models.TimeEntry {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var entries []*models.TimeEntry
	for _, entry := range tms.timeEntries {
		if entry.GetUser().ID == userID {
			entryDate := entry.GetDate()
			if entryDate.After(startDate) && entryDate.Before(endDate) {
				entries = append(entries, entry)
			}
		}
	}
	
	// Sort by date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].GetDate().After(entries[j].GetDate())
	})
	
	return entries
}

func (tms *TimeService) GetProjectTimeEntries(projectID string, startDate, endDate time.Time) []*models.TimeEntry {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var entries []*models.TimeEntry
	for _, entry := range tms.timeEntries {
		// In a real system, we would check the task's project
		// For demo purposes, we'll use a simple filter
		entryDate := entry.GetDate()
		if entryDate.After(startDate) && entryDate.Before(endDate) {
			entries = append(entries, entry)
		}
	}
	
	// Sort by date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].GetDate().After(entries[j].GetDate())
	})
	
	return entries
}

func (tms *TimeService) UpdateTimeEntry(entryID string, duration time.Duration, description string, user *models.User) error {
	tms.mu.Lock()
	defer tms.mu.Unlock()
	
	entry := tms.timeEntries[entryID]
	if entry == nil {
		return fmt.Errorf("time entry not found")
	}
	
	// Check if user can edit this time entry
	if !entry.CanBeEditedBy(user) {
		return fmt.Errorf("user does not have permission to edit this time entry")
	}
	
	entry.UpdateDuration(duration)
	entry.UpdateDescription(description)
	
	return nil
}

func (tms *TimeService) DeleteTimeEntry(entryID string, user *models.User) error {
	tms.mu.Lock()
	defer tms.mu.Unlock()
	
	entry := tms.timeEntries[entryID]
	if entry == nil {
		return fmt.Errorf("time entry not found")
	}
	
	// Check if user can delete this time entry
	if !entry.CanBeDeletedBy(user) {
		return fmt.Errorf("user does not have permission to delete this time entry")
	}
	
	delete(tms.timeEntries, entryID)
	return nil
}

// Reporting
func (tms *TimeService) GenerateUserTimeReport(userID string, startDate, endDate time.Time) *models.TimeReport {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	entries := tms.GetUserTimeEntries(userID, startDate, endDate)
	dateRange := models.DateRange{StartDate: startDate, EndDate: endDate}
	
	return models.NewTimeReport(userID, "", entries, dateRange)
}

func (tms *TimeService) GenerateProjectTimeReport(projectID string, startDate, endDate time.Time) *models.TimeReport {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	entries := tms.GetProjectTimeEntries(projectID, startDate, endDate)
	dateRange := models.DateRange{StartDate: startDate, EndDate: endDate}
	
	return models.NewTimeReport("", projectID, entries, dateRange)
}

func (tms *TimeService) GetTotalTimeForTask(taskID string) time.Duration {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var totalDuration time.Duration
	for _, entry := range tms.timeEntries {
		if entry.GetTask().ID == taskID {
			totalDuration += entry.GetDuration()
		}
	}
	
	return totalDuration
}

func (tms *TimeService) GetTotalTimeForUser(userID string, startDate, endDate time.Time) time.Duration {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var totalDuration time.Duration
	for _, entry := range tms.timeEntries {
		if entry.GetUser().ID == userID {
			entryDate := entry.GetDate()
			if entryDate.After(startDate) && entryDate.Before(endDate) {
				totalDuration += entry.GetDuration()
			}
		}
	}
	
	return totalDuration
}

func (tms *TimeService) GetTotalTimeForProject(projectID string, startDate, endDate time.Time) time.Duration {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var totalDuration time.Duration
	for _, entry := range tms.timeEntries {
		// In a real system, we would check the task's project
		// For demo purposes, we'll use a simple filter
		entryDate := entry.GetDate()
		if entryDate.After(startDate) && entryDate.Before(endDate) {
			totalDuration += entry.GetDuration()
		}
	}
	
	return totalDuration
}

// Analytics
func (tms *TimeService) GetTimeEntriesByDateRange(startDate, endDate time.Time) []*models.TimeEntry {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var entries []*models.TimeEntry
	for _, entry := range tms.timeEntries {
		entryDate := entry.GetDate()
		if entryDate.After(startDate) && entryDate.Before(endDate) {
			entries = append(entries, entry)
		}
	}
	
	// Sort by date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].GetDate().After(entries[j].GetDate())
	})
	
	return entries
}

func (tms *TimeService) GetTimeEntriesByUser(userID string) []*models.TimeEntry {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var entries []*models.TimeEntry
	for _, entry := range tms.timeEntries {
		if entry.GetUser().ID == userID {
			entries = append(entries, entry)
		}
	}
	
	// Sort by date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].GetDate().After(entries[j].GetDate())
	})
	
	return entries
}

func (tms *TimeService) GetTimeEntriesByTask(taskID string) []*models.TimeEntry {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var entries []*models.TimeEntry
	for _, entry := range tms.timeEntries {
		if entry.GetTask().ID == taskID {
			entries = append(entries, entry)
		}
	}
	
	// Sort by date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].GetDate().After(entries[j].GetDate())
	})
	
	return entries
}

// Helper Methods
func (tms *TimeService) GetTimeEntry(entryID string) *models.TimeEntry {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	return tms.timeEntries[entryID]
}

func (tms *TimeService) GetAllTimeEntries() []*models.TimeEntry {
	tms.mu.RLock()
	defer tms.mu.RUnlock()
	
	var entries []*models.TimeEntry
	for _, entry := range tms.timeEntries {
		entries = append(entries, entry)
	}
	
	// Sort by date
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].GetDate().After(entries[j].GetDate())
	})
	
	return entries
}
