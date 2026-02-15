package models

import (
	"fmt"
	"sync"
	"time"
)

// Task Priority
type Priority int

const (
	Critical Priority = iota
	High
	Medium
	Low
)

func (p Priority) String() string {
	switch p {
	case Critical:
		return "Critical"
	case High:
		return "High"
	case Medium:
		return "Medium"
	case Low:
		return "Low"
	default:
		return "Unknown"
	}
}

// Task Status
type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	InReview
	Done
	Cancelled
)

func (ts TaskStatus) String() string {
	switch ts {
	case Todo:
		return "To Do"
	case InProgress:
		return "In Progress"
	case InReview:
		return "In Review"
	case Done:
		return "Done"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

// Task
type Task struct {
	ID           string
	Title        string
	Description  string
	Status       TaskStatus
	Priority     Priority
	Assignee     *User
	Project      *Project
	DueDate      *time.Time
	Dependencies []*Task
	Subtasks     []*Task
	ParentTask   *Task
	CreatedBy    *User
	CreatedAt    time.Time
	UpdatedAt    time.Time
	mu           sync.RWMutex
}

func NewTask(title, description string, project *Project, createdBy *User) *Task {
	return &Task{
		ID:          fmt.Sprintf("TK%d", time.Now().UnixNano()),
		Title:       title,
		Description: description,
		Status:      Todo,
		Priority:    Medium,
		Project:     project,
		CreatedBy:   createdBy,
		Dependencies: make([]*Task, 0),
		Subtasks:    make([]*Task, 0),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) AssignTo(user *User) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	// Check if user has access to the project
	if !t.Project.CanUserAccess(user.ID) {
		return fmt.Errorf("user does not have access to this project")
	}
	
	t.Assignee = user
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) Unassign() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Assignee = nil
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateStatus(status TaskStatus) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	// Check if status transition is valid
	if !t.isValidStatusTransition(t.Status, status) {
		return fmt.Errorf("invalid status transition from %s to %s", t.Status.String(), status.String())
	}
	
	t.Status = status
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) isValidStatusTransition(from, to TaskStatus) bool {
	// Define valid transitions
	validTransitions := map[TaskStatus][]TaskStatus{
		Todo:       {InProgress, Cancelled},
		InProgress: {InReview, Done, Cancelled},
		InReview:   {InProgress, Done, Cancelled},
		Done:       {InProgress}, // Allow reopening
		Cancelled:  {Todo},       // Allow reopening
	}
	
	allowedStatuses, exists := validTransitions[from]
	if !exists {
		return false
	}
	
	for _, allowedStatus := range allowedStatuses {
		if allowedStatus == to {
			return true
		}
	}
	return false
}

func (t *Task) SetPriority(priority Priority) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Priority = priority
	t.UpdatedAt = time.Now()
}

func (t *Task) SetDueDate(dueDate *time.Time) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.DueDate = dueDate
	t.UpdatedAt = time.Now()
}

func (t *Task) AddDependency(dependency *Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	// Check for circular dependency
	if t.wouldCreateCircularDependency(dependency) {
		return fmt.Errorf("adding this dependency would create a circular dependency")
	}
	
	// Check if dependency is already added
	for _, dep := range t.Dependencies {
		if dep.ID == dependency.ID {
			return fmt.Errorf("dependency already exists")
		}
	}
	
	t.Dependencies = append(t.Dependencies, dependency)
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) wouldCreateCircularDependency(dependency *Task) bool {
	// Check if the dependency task depends on this task
	return t.isDependencyOf(dependency, t)
}

func (t *Task) isDependencyOf(task, target *Task) bool {
	for _, dep := range task.Dependencies {
		if dep.ID == target.ID {
			return true
		}
		if t.isDependencyOf(dep, target) {
			return true
		}
	}
	return false
}

func (t *Task) AddSubtask(subtask *Task) {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	subtask.ParentTask = t
	t.Subtasks = append(t.Subtasks, subtask)
	t.UpdatedAt = time.Now()
}

func (t *Task) GetSubtasks() []*Task {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	subtasks := make([]*Task, len(t.Subtasks))
	copy(subtasks, t.Subtasks)
	return subtasks
}

func (t *Task) GetDependencies() []*Task {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	dependencies := make([]*Task, len(t.Dependencies))
	copy(dependencies, t.Dependencies)
	return dependencies
}

func (t *Task) IsOverdue() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	if t.DueDate == nil {
		return false
	}
	
	return t.Status != Done && t.Status != Cancelled && time.Now().After(*t.DueDate)
}

func (t *Task) GetDaysUntilDue() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	if t.DueDate == nil {
		return -1 // No due date
	}
	
	days := int(t.DueDate.Sub(time.Now()).Hours() / 24)
	return days
}

func (t *Task) UpdateDetails(title, description string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Task) CanBeAssignedTo(user *User) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	// Check if user has access to the project
	return t.Project.CanUserAccess(user.ID)
}

func (t *Task) IsCompleted() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Status == Done
}

func (t *Task) IsInProgress() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Status == InProgress
}
