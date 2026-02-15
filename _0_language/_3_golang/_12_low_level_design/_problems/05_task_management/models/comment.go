package models

import (
	"fmt"
	"sync"
	"time"
)

// Comment Type
type CommentType int

const (
	GeneralComment CommentType = iota
	StatusChangeComment
	AssignmentComment
	TimeLogComment
	SystemComment
)

func (ct CommentType) String() string {
	switch ct {
	case GeneralComment:
		return "General"
	case StatusChangeComment:
		return "Status Change"
	case AssignmentComment:
		return "Assignment"
	case TimeLogComment:
		return "Time Log"
	case SystemComment:
		return "System"
	default:
		return "Unknown"
	}
}

// Comment
type Comment struct {
	ID        string
	Content   string
	Author    *User
	Task      *Task
	Type      CommentType
	CreatedAt time.Time
	UpdatedAt time.Time
	mu        sync.RWMutex
}

func NewComment(content string, author *User, task *Task, commentType CommentType) *Comment {
	return &Comment{
		ID:        fmt.Sprintf("C%d", time.Now().UnixNano()),
		Content:   content,
		Author:    author,
		Task:      task,
		Type:      commentType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Comment) UpdateContent(content string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Content = content
	c.UpdatedAt = time.Now()
}

func (c *Comment) GetContent() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Content
}

func (c *Comment) GetAuthor() *User {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Author
}

func (c *Comment) GetTask() *Task {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Task
}

func (c *Comment) GetType() CommentType {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Type
}

func (c *Comment) GetCreatedAt() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.CreatedAt
}

func (c *Comment) IsSystemComment() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Type == SystemComment
}

func (c *Comment) CanBeEditedBy(user *User) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	// System comments cannot be edited
	if c.Type == SystemComment {
		return false
	}
	
	// Author can edit their own comments
	return c.Author.ID == user.ID
}

func (c *Comment) CanBeDeletedBy(user *User) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	// System comments cannot be deleted
	if c.Type == SystemComment {
		return false
	}
	
	// Author can delete their own comments
	// Project managers and admins can delete any comment
	return c.Author.ID == user.ID || 
		   user.Role == Admin || 
		   user.Role == ProjectManager
}
