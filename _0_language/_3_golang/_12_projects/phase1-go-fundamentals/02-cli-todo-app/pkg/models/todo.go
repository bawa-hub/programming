package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Priority represents the priority level of a todo
type Priority int

// Priority constants using iota
const (
	Low Priority = iota
	Medium
	High
	Urgent
)

// Status represents the status of a todo
type Status int

// Status constants using iota
const (
	Pending Status = iota
	InProgress
	Completed
	Cancelled
)

// Category represents a todo category
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// Todo represents a single todo item
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Priority    Priority  `json:"priority"`
	Status      Status    `json:"status"`
	Category    *Category `json:"category,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// TodoList represents a collection of todos
type TodoList struct {
	Todos     []*Todo     `json:"todos"`
	Categories []*Category `json:"categories"`
	NextID    int         `json:"next_id"`
}

// Stringer interface implementations
func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	case Urgent:
		return "Urgent"
	default:
		return "Unknown"
	}
}

func (s Status) String() string {
	switch s {
	case Pending:
		return "Pending"
	case InProgress:
		return "In Progress"
	case Completed:
		return "Completed"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

func (c *Category) String() string {
	return fmt.Sprintf("Category{ID: %d, Name: %s, Color: %s}", c.ID, c.Name, c.Color)
}

func (t *Todo) String() string {
	status := t.Status.String()
	priority := t.Priority.String()
	category := "None"
	if t.Category != nil {
		category = t.Category.Name
	}
	
	dueDate := "No due date"
	if t.DueDate != nil {
		dueDate = t.DueDate.Format("2006-01-02 15:04")
	}
	
	return fmt.Sprintf("Todo{ID: %d, Title: %s, Status: %s, Priority: %s, Category: %s, Due: %s}", 
		t.ID, t.Title, status, priority, category, dueDate)
}

// Methods for Todo

// IsOverdue checks if the todo is overdue
func (t *Todo) IsOverdue() bool {
	if t.DueDate == nil || t.Status == Completed || t.Status == Cancelled {
		return false
	}
	return time.Now().After(*t.DueDate)
}

// DaysUntilDue returns the number of days until due date
func (t *Todo) DaysUntilDue() int {
	if t.DueDate == nil {
		return -1 // No due date
	}
	
	now := time.Now()
	diff := t.DueDate.Sub(now)
	return int(diff.Hours() / 24)
}

// MarkCompleted marks the todo as completed
func (t *Todo) MarkCompleted() {
	t.Status = Completed
	now := time.Now()
	t.CompletedAt = &now
	t.UpdatedAt = now
}

// MarkInProgress marks the todo as in progress
func (t *Todo) MarkInProgress() {
	t.Status = InProgress
	t.UpdatedAt = time.Now()
}

// UpdatePriority updates the priority of the todo
func (t *Todo) UpdatePriority(priority Priority) {
	t.Priority = priority
	t.UpdatedAt = time.Now()
}

// AddTag adds a tag to the todo
func (t *Todo) AddTag(tag string) {
	// Check if tag already exists
	for _, existingTag := range t.Tags {
		if existingTag == tag {
			return // Tag already exists
		}
	}
	t.Tags = append(t.Tags, tag)
	t.UpdatedAt = time.Now()
}

// RemoveTag removes a tag from the todo
func (t *Todo) RemoveTag(tag string) {
	for i, existingTag := range t.Tags {
		if existingTag == tag {
			t.Tags = append(t.Tags[:i], t.Tags[i+1:]...)
			t.UpdatedAt = time.Now()
			return
		}
	}
}

// SetDueDate sets the due date for the todo
func (t *Todo) SetDueDate(dueDate time.Time) {
	t.DueDate = &dueDate
	t.UpdatedAt = time.Now()
}

// SetCategory sets the category for the todo
func (t *Todo) SetCategory(category *Category) {
	t.Category = category
	t.UpdatedAt = time.Now()
}

// GetFormattedDueDate returns a formatted due date string
func (t *Todo) GetFormattedDueDate() string {
	if t.DueDate == nil {
		return "No due date"
	}
	return t.DueDate.Format("2006-01-02 15:04")
}

// GetFormattedCreatedAt returns a formatted creation date string
func (t *Todo) GetFormattedCreatedAt() string {
	return t.CreatedAt.Format("2006-01-02 15:04")
}

// GetFormattedUpdatedAt returns a formatted update date string
func (t *Todo) GetFormattedUpdatedAt() string {
	return t.UpdatedAt.Format("2006-01-02 15:04")
}

// GetFormattedCompletedAt returns a formatted completion date string
func (t *Todo) GetFormattedCompletedAt() string {
	if t.CompletedAt == nil {
		return "Not completed"
	}
	return t.CompletedAt.Format("2006-01-02 15:04")
}

// GetPriorityColor returns a color code for the priority
func (p Priority) GetPriorityColor() string {
	switch p {
	case Low:
		return "🟢" // Green
	case Medium:
		return "🟡" // Yellow
	case High:
		return "🟠" // Orange
	case Urgent:
		return "🔴" // Red
	default:
		return "⚪" // White
	}
}

// GetStatusIcon returns an icon for the status
func (s Status) GetStatusIcon() string {
	switch s {
	case Pending:
		return "⏳" // Hourglass
	case InProgress:
		return "🔄" // Refresh
	case Completed:
		return "✅" // Check mark
	case Cancelled:
		return "❌" // X mark
	default:
		return "❓" // Question mark
	}
}

// Methods for TodoList

// AddTodo adds a new todo to the list
func (tl *TodoList) AddTodo(todo *Todo) {
	todo.ID = tl.NextID
	tl.NextID++
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	tl.Todos = append(tl.Todos, todo)
}

// GetTodoByID returns a todo by its ID
func (tl *TodoList) GetTodoByID(id int) *Todo {
	for _, todo := range tl.Todos {
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

// RemoveTodo removes a todo by its ID
func (tl *TodoList) RemoveTodo(id int) bool {
	for i, todo := range tl.Todos {
		if todo.ID == id {
			tl.Todos = append(tl.Todos[:i], tl.Todos[i+1:]...)
			return true
		}
	}
	return false
}

// GetTodosByStatus returns todos filtered by status
func (tl *TodoList) GetTodosByStatus(status Status) []*Todo {
	var result []*Todo
	for _, todo := range tl.Todos {
		if todo.Status == status {
			result = append(result, todo)
		}
	}
	return result
}

// GetTodosByPriority returns todos filtered by priority
func (tl *TodoList) GetTodosByPriority(priority Priority) []*Todo {
	var result []*Todo
	for _, todo := range tl.Todos {
		if todo.Priority == priority {
			result = append(result, todo)
		}
	}
	return result
}

// GetTodosByCategory returns todos filtered by category
func (tl *TodoList) GetTodosByCategory(categoryID int) []*Todo {
	var result []*Todo
	for _, todo := range tl.Todos {
		if todo.Category != nil && todo.Category.ID == categoryID {
			result = append(result, todo)
		}
	}
	return result
}

// GetOverdueTodos returns todos that are overdue
func (tl *TodoList) GetOverdueTodos() []*Todo {
	var result []*Todo
	for _, todo := range tl.Todos {
		if todo.IsOverdue() {
			result = append(result, todo)
		}
	}
	return result
}

// GetTodosDueToday returns todos due today
func (tl *TodoList) GetTodosDueToday() []*Todo {
	var result []*Todo
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)
	
	for _, todo := range tl.Todos {
		if todo.DueDate != nil && 
		   todo.DueDate.After(today) && 
		   todo.DueDate.Before(tomorrow) &&
		   todo.Status != Completed {
			result = append(result, todo)
		}
	}
	return result
}

// GetCompletedTodos returns completed todos
func (tl *TodoList) GetCompletedTodos() []*Todo {
	return tl.GetTodosByStatus(Completed)
}

// GetPendingTodos returns pending todos
func (tl *TodoList) GetPendingTodos() []*Todo {
	return tl.GetTodosByStatus(Pending)
}

// GetInProgressTodos returns in-progress todos
func (tl *TodoList) GetInProgressTodos() []*Todo {
	return tl.GetTodosByStatus(InProgress)
}

// GetStats returns statistics about the todo list
func (tl *TodoList) GetStats() map[string]int {
	stats := make(map[string]int)
	stats["total"] = len(tl.Todos)
	stats["pending"] = len(tl.GetPendingTodos())
	stats["in_progress"] = len(tl.GetInProgressTodos())
	stats["completed"] = len(tl.GetCompletedTodos())
	stats["overdue"] = len(tl.GetOverdueTodos())
	stats["due_today"] = len(tl.GetTodosDueToday())
	
	return stats
}

// AddCategory adds a new category
func (tl *TodoList) AddCategory(category *Category) {
	category.ID = len(tl.Categories) + 1
	category.CreatedAt = time.Now()
	tl.Categories = append(tl.Categories, category)
}

// GetCategoryByID returns a category by its ID
func (tl *TodoList) GetCategoryByID(id int) *Category {
	for _, category := range tl.Categories {
		if category.ID == id {
			return category
		}
	}
	return nil
}

// GetCategoryByName returns a category by its name
func (tl *TodoList) GetCategoryByName(name string) *Category {
	for _, category := range tl.Categories {
		if category.Name == name {
			return category
		}
	}
	return nil
}

// RemoveCategory removes a category by its ID
func (tl *TodoList) RemoveCategory(id int) bool {
	for i, category := range tl.Categories {
		if category.ID == id {
			tl.Categories = append(tl.Categories[:i], tl.Categories[i+1:]...)
			return true
		}
	}
	return false
}

// ToJSON converts the todo list to JSON
func (tl *TodoList) ToJSON() ([]byte, error) {
	return json.MarshalIndent(tl, "", "  ")
}

// FromJSON creates a todo list from JSON
func (tl *TodoList) FromJSON(data []byte) error {
	return json.Unmarshal(data, tl)
}

// NewTodoList creates a new empty todo list
func NewTodoList() *TodoList {
	return &TodoList{
		Todos:      make([]*Todo, 0),
		Categories: make([]*Category, 0),
		NextID:     1,
	}
}

// NewTodo creates a new todo with default values
func NewTodo(title, description string) *Todo {
	now := time.Now()
	return &Todo{
		Title:     title,
		Description: description,
		Priority:  Medium,
		Status:    Pending,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// NewCategory creates a new category
func NewCategory(name, description, color string) *Category {
	return &Category{
		Name:        name,
		Description: description,
		Color:       color,
		CreatedAt:   time.Now(),
	}
}
