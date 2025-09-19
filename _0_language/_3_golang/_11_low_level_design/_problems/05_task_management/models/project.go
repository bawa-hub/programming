package models

import (
	"fmt"
	"sync"
	"time"
)

// Project Status
type ProjectStatus int

const (
	ProjectPlanning ProjectStatus = iota
	ProjectActive
	ProjectOnHold
	ProjectCompleted
	ProjectCancelled
)

func (ps ProjectStatus) String() string {
	switch ps {
	case ProjectPlanning:
		return "Planning"
	case ProjectActive:
		return "Active"
	case ProjectOnHold:
		return "On Hold"
	case ProjectCompleted:
		return "Completed"
	case ProjectCancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

// Project Settings
type ProjectSettings struct {
	AllowSelfAssignment bool
	RequireTimeTracking bool
	DefaultTaskPriority Priority
	WorkflowEnabled     bool
	NotificationSettings map[string]bool
}

// Project
type Project struct {
	ID          string
	Name        string
	Description string
	Status      ProjectStatus
	Owner       *User
	Members     []*User
	Settings    ProjectSettings
	CreatedAt   time.Time
	UpdatedAt   time.Time
	mu          sync.RWMutex
}

func NewProject(name, description string, owner *User) *Project {
	return &Project{
		ID:        fmt.Sprintf("P%d", time.Now().UnixNano()),
		Name:      name,
		Description: description,
		Status:    ProjectPlanning,
		Owner:     owner,
		Members:   make([]*User, 0),
		Settings: ProjectSettings{
			AllowSelfAssignment: true,
			RequireTimeTracking: false,
			DefaultTaskPriority: Medium,
			WorkflowEnabled:     false,
			NotificationSettings: map[string]bool{
				"task_assigned":   true,
				"task_completed":  true,
				"deadline_alert":  true,
				"status_changed":  true,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *Project) AddMember(user *User) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	// Check if user is already a member
	for _, member := range p.Members {
		if member.ID == user.ID {
			return fmt.Errorf("user is already a member of this project")
		}
	}
	
	p.Members = append(p.Members, user)
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Project) RemoveMember(userID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	for i, member := range p.Members {
		if member.ID == userID {
			p.Members = append(p.Members[:i], p.Members[i+1:]...)
			p.UpdatedAt = time.Now()
			return nil
		}
	}
	
	return fmt.Errorf("user not found in project")
}

func (p *Project) GetMembers() []*User {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	members := make([]*User, len(p.Members))
	copy(members, p.Members)
	return members
}

func (p *Project) IsMember(userID string) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	// Owner is always a member
	if p.Owner.ID == userID {
		return true
	}
	
	for _, member := range p.Members {
		if member.ID == userID {
			return true
		}
	}
	return false
}

func (p *Project) UpdateStatus(status ProjectStatus) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Status = status
	p.UpdatedAt = time.Now()
}

func (p *Project) UpdateSettings(settings ProjectSettings) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Settings = settings
	p.UpdatedAt = time.Now()
}

func (p *Project) UpdateDetails(name, description string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Name = name
	p.Description = description
	p.UpdatedAt = time.Now()
}

func (p *Project) GetMemberCount() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.Members) + 1 // +1 for owner
}

func (p *Project) CanUserAccess(userID string) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	// Owner can always access
	if p.Owner.ID == userID {
		return true
	}
	
	// Check if user is a member
	for _, member := range p.Members {
		if member.ID == userID {
			return true
		}
	}
	return false
}
