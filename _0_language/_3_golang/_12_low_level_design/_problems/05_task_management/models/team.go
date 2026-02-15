package models

import (
	"fmt"
	"sync"
	"time"
)

// Team
type Team struct {
	ID          string
	Name        string
	Description string
	Members     []*User
	CreatedBy   *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
	mu          sync.RWMutex
}

func NewTeam(name, description string, createdBy *User) *Team {
	return &Team{
		ID:          fmt.Sprintf("T%d", time.Now().UnixNano()),
		Name:        name,
		Description: description,
		Members:     make([]*User, 0),
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Team) AddMember(user *User) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	// Check if user is already a member
	for _, member := range t.Members {
		if member.ID == user.ID {
			return fmt.Errorf("user is already a member of this team")
		}
	}
	
	t.Members = append(t.Members, user)
	user.SetTeam(t)
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Team) RemoveMember(userID string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	for i, member := range t.Members {
		if member.ID == userID {
			t.Members = append(t.Members[:i], t.Members[i+1:]...)
			t.UpdatedAt = time.Now()
			return nil
		}
	}
	
	return fmt.Errorf("user not found in team")
}

func (t *Team) GetMembers() []*User {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	members := make([]*User, len(t.Members))
	copy(members, t.Members)
	return members
}

func (t *Team) GetMemberCount() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.Members)
}

func (t *Team) IsMember(userID string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	for _, member := range t.Members {
		if member.ID == userID {
			return true
		}
	}
	return false
}

func (t *Team) UpdateDetails(name, description string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Name = name
	t.Description = description
	t.UpdatedAt = time.Now()
}
