package todo

import (
	"cli-todo-app/internal/storage"
	"cli-todo-app/pkg/models"
	"cli-todo-app/pkg/utils"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Service represents the todo service
type Service struct {
	storage    storage.Storage
	todoList   *models.TodoList
	autoSave   *storage.AutoSave
}

// NewService creates a new todo service
func NewService(storageBackend storage.Storage) *Service {
	service := &Service{
		storage: storageBackend,
	}
	
	// Load existing data
	if err := service.Load(); err != nil {
		fmt.Printf("Warning: failed to load existing data: %v\n", err)
		service.todoList = models.NewTodoList()
	}
	
	// Initialize auto-save
	service.autoSave = storage.NewAutoSave(storageBackend, 30*time.Second)
	service.autoSave.Start(service.todoList)
	
	return service
}

// Close closes the service and stops auto-save
func (s *Service) Close() {
	if s.autoSave != nil {
		s.autoSave.Stop()
	}
}

// Load loads data from storage
func (s *Service) Load() error {
	todoList, err := s.storage.Load()
	if err != nil {
		return err
	}
	s.todoList = todoList
	return nil
}

// Save saves data to storage
func (s *Service) Save() error {
	if err := s.storage.Save(s.todoList); err != nil {
		return err
	}
	s.autoSave.UpdateTodoList(s.todoList)
	return nil
}

// Todo operations

// AddTodo adds a new todo
func (s *Service) AddTodo(title, description string) (*models.Todo, error) {
	if strings.TrimSpace(title) == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}
	
	todo := models.NewTodo(title, description)
	s.todoList.AddTodo(todo)
	
	if err := s.Save(); err != nil {
		return nil, fmt.Errorf("failed to save todo: %w", err)
	}
	
	return todo, nil
}

// GetTodo retrieves a todo by ID
func (s *Service) GetTodo(id int) *models.Todo {
	return s.todoList.GetTodoByID(id)
}

// UpdateTodo updates a todo
func (s *Service) UpdateTodo(id int, title, description string) error {
	todo := s.todoList.GetTodoByID(id)
	if todo == nil {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	if strings.TrimSpace(title) != "" {
		todo.Title = title
	}
	if description != "" {
		todo.Description = description
	}
	todo.UpdatedAt = time.Now()
	
	return s.Save()
}

// DeleteTodo deletes a todo
func (s *Service) DeleteTodo(id int) error {
	if !s.todoList.RemoveTodo(id) {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	return s.Save()
}

// CompleteTodo marks a todo as completed
func (s *Service) CompleteTodo(id int) error {
	todo := s.todoList.GetTodoByID(id)
	if todo == nil {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	todo.MarkCompleted()
	return s.Save()
}

// StartTodo marks a todo as in progress
func (s *Service) StartTodo(id int) error {
	todo := s.todoList.GetTodoByID(id)
	if todo == nil {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	todo.MarkInProgress()
	return s.Save()
}

// SetPriority sets the priority of a todo
func (s *Service) SetPriority(id int, priority models.Priority) error {
	todo := s.todoList.GetTodoByID(id)
	if todo == nil {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	todo.UpdatePriority(priority)
	return s.Save()
}

// SetDueDate sets the due date of a todo
func (s *Service) SetDueDate(id int, dueDate time.Time) error {
	todo := s.todoList.GetTodoByID(id)
	if todo == nil {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	todo.SetDueDate(dueDate)
	return s.Save()
}

// AddTag adds a tag to a todo
func (s *Service) AddTag(id int, tag string) error {
	todo := s.todoList.GetTodoByID(id)
	if todo == nil {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	if strings.TrimSpace(tag) == "" {
		return fmt.Errorf("tag cannot be empty")
	}
	
	todo.AddTag(strings.TrimSpace(tag))
	return s.Save()
}

// RemoveTag removes a tag from a todo
func (s *Service) RemoveTag(id int, tag string) error {
	todo := s.todoList.GetTodoByID(id)
	if todo == nil {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	todo.RemoveTag(tag)
	return s.Save()
}

// SetCategory sets the category of a todo
func (s *Service) SetCategory(id int, categoryID int) error {
	todo := s.todoList.GetTodoByID(id)
	if todo == nil {
		return fmt.Errorf("todo with ID %d not found", id)
	}
	
	category := s.todoList.GetCategoryByID(categoryID)
	if category == nil {
		return fmt.Errorf("category with ID %d not found", categoryID)
	}
	
	todo.SetCategory(category)
	return s.Save()
}

// List operations

// GetAllTodos returns all todos
func (s *Service) GetAllTodos() []*models.Todo {
	return s.todoList.Todos
}

// GetTodosByStatus returns todos filtered by status
func (s *Service) GetTodosByStatus(status models.Status) []*models.Todo {
	return s.todoList.GetTodosByStatus(status)
}

// GetTodosByPriority returns todos filtered by priority
func (s *Service) GetTodosByPriority(priority models.Priority) []*models.Todo {
	return s.todoList.GetTodosByPriority(priority)
}

// GetTodosByCategory returns todos filtered by category
func (s *Service) GetTodosByCategory(categoryID int) []*models.Todo {
	return s.todoList.GetTodosByCategory(categoryID)
}

// GetOverdueTodos returns overdue todos
func (s *Service) GetOverdueTodos() []*models.Todo {
	return s.todoList.GetOverdueTodos()
}

// GetTodosDueToday returns todos due today
func (s *Service) GetTodosDueToday() []*models.Todo {
	return s.todoList.GetTodosDueToday()
}

// SearchTodos searches todos by title or description
func (s *Service) SearchTodos(query string) []*models.Todo {
	if strings.TrimSpace(query) == "" {
		return s.todoList.Todos
	}
	
	query = strings.ToLower(strings.TrimSpace(query))
	var results []*models.Todo
	
	for _, todo := range s.todoList.Todos {
		title := strings.ToLower(todo.Title)
		description := strings.ToLower(todo.Description)
		
		if strings.Contains(title, query) || strings.Contains(description, query) {
			results = append(results, todo)
		}
	}
	
	return results
}

// SearchTodosByTag searches todos by tag
func (s *Service) SearchTodosByTag(tag string) []*models.Todo {
	if strings.TrimSpace(tag) == "" {
		return s.todoList.Todos
	}
	
	tag = strings.ToLower(strings.TrimSpace(tag))
	var results []*models.Todo
	
	for _, todo := range s.todoList.Todos {
		for _, todoTag := range todo.Tags {
			if strings.Contains(strings.ToLower(todoTag), tag) {
				results = append(results, todo)
				break
			}
		}
	}
	
	return results
}

// SortTodos sorts todos by the specified criteria
func (s *Service) SortTodos(todos []*models.Todo, sortBy string, ascending bool) []*models.Todo {
	result := make([]*models.Todo, len(todos))
	copy(result, todos)
	
	switch strings.ToLower(sortBy) {
	case "id":
		sort.Slice(result, func(i, j int) bool {
			if ascending {
				return result[i].ID < result[j].ID
			}
			return result[i].ID > result[j].ID
		})
	case "title":
		sort.Slice(result, func(i, j int) bool {
			if ascending {
				return result[i].Title < result[j].Title
			}
			return result[i].Title > result[j].Title
		})
	case "priority":
		sort.Slice(result, func(i, j int) bool {
			if ascending {
				return result[i].Priority < result[j].Priority
			}
			return result[i].Priority > result[j].Priority
		})
	case "status":
		sort.Slice(result, func(i, j int) bool {
			if ascending {
				return result[i].Status < result[j].Status
			}
			return result[i].Status > result[j].Status
		})
	case "created":
		sort.Slice(result, func(i, j int) bool {
			if ascending {
				return result[i].CreatedAt.Before(result[j].CreatedAt)
			}
			return result[i].CreatedAt.After(result[j].CreatedAt)
		})
	case "updated":
		sort.Slice(result, func(i, j int) bool {
			if ascending {
				return result[i].UpdatedAt.Before(result[j].UpdatedAt)
			}
			return result[i].UpdatedAt.After(result[j].UpdatedAt)
		})
	case "due":
		sort.Slice(result, func(i, j int) bool {
			// Handle nil due dates
			if result[i].DueDate == nil && result[j].DueDate == nil {
				return false
			}
			if result[i].DueDate == nil {
				return !ascending
			}
			if result[j].DueDate == nil {
				return ascending
			}
			
			if ascending {
				return result[i].DueDate.Before(*result[j].DueDate)
			}
			return result[i].DueDate.After(*result[j].DueDate)
		})
	}
	
	return result
}

// GetStats returns statistics about todos
func (s *Service) GetStats() map[string]int {
	return s.todoList.GetStats()
}

// Category operations

// AddCategory adds a new category
func (s *Service) AddCategory(name, description, color string) (*models.Category, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("category name cannot be empty")
	}
	
	// Check if category already exists
	if s.todoList.GetCategoryByName(name) != nil {
		return nil, fmt.Errorf("category with name '%s' already exists", name)
	}
	
	category := models.NewCategory(name, description, color)
	s.todoList.AddCategory(category)
	
	if err := s.Save(); err != nil {
		return nil, fmt.Errorf("failed to save category: %w", err)
	}
	
	return category, nil
}

// GetCategory retrieves a category by ID
func (s *Service) GetCategory(id int) *models.Category {
	return s.todoList.GetCategoryByID(id)
}

// GetCategoryByName retrieves a category by name
func (s *Service) GetCategoryByName(name string) *models.Category {
	return s.todoList.GetCategoryByName(name)
}

// GetAllCategories returns all categories
func (s *Service) GetAllCategories() []*models.Category {
	return s.todoList.Categories
}

// UpdateCategory updates a category
func (s *Service) UpdateCategory(id int, name, description, color string) error {
	category := s.todoList.GetCategoryByID(id)
	if category == nil {
		return fmt.Errorf("category with ID %d not found", id)
	}
	
	if strings.TrimSpace(name) != "" {
		// Check if new name conflicts with existing category
		if existing := s.todoList.GetCategoryByName(name); existing != nil && existing.ID != id {
			return fmt.Errorf("category with name '%s' already exists", name)
		}
		category.Name = name
	}
	if description != "" {
		category.Description = description
	}
	if color != "" {
		category.Color = color
	}
	
	return s.Save()
}

// DeleteCategory deletes a category
func (s *Service) DeleteCategory(id int) error {
	if !s.todoList.RemoveCategory(id) {
		return fmt.Errorf("category with ID %d not found", id)
	}
	
	// Remove category from all todos
	for _, todo := range s.todoList.Todos {
		if todo.Category != nil && todo.Category.ID == id {
			todo.Category = nil
		}
	}
	
	return s.Save()
}

// Export operations

// ExportToJSON exports todos to JSON
func (s *Service) ExportToJSON(filePath string) error {
	return storage.ExportToJSON(s.todoList, filePath)
}

// ExportToCSV exports todos to CSV
func (s *Service) ExportToCSV(filePath string) error {
	return storage.ExportToCSV(s.todoList, filePath)
}

// ImportFromJSON imports todos from JSON
func (s *Service) ImportFromJSON(filePath string) error {
	todoList, err := storage.ImportFromJSON(filePath)
	if err != nil {
		return err
	}
	
	s.todoList = todoList
	return s.Save()
}

// Utility operations

// ClearCompleted removes all completed todos
func (s *Service) ClearCompleted() error {
	completedTodos := s.todoList.GetCompletedTodos()
	
	for _, todo := range completedTodos {
		s.todoList.RemoveTodo(todo.ID)
	}
	
	return s.Save()
}

// ClearOverdue removes all overdue todos
func (s *Service) ClearOverdue() error {
	overdueTodos := s.todoList.GetOverdueTodos()
	
	for _, todo := range overdueTodos {
		s.todoList.RemoveTodo(todo.ID)
	}
	
	return s.Save()
}

// ArchiveCompleted marks all completed todos as archived (by setting a special tag)
func (s *Service) ArchiveCompleted() error {
	completedTodos := s.todoList.GetCompletedTodos()
	
	for _, todo := range completedTodos {
		todo.AddTag("archived")
	}
	
	return s.Save()
}

// GetTodosByDateRange returns todos created within a date range
func (s *Service) GetTodosByDateRange(start, end time.Time) []*models.Todo {
	var results []*models.Todo
	
	for _, todo := range s.todoList.Todos {
		if todo.CreatedAt.After(start) && todo.CreatedAt.Before(end) {
			results = append(results, todo)
		}
	}
	
	return results
}

// GetTodosByTag returns todos that have any of the specified tags
func (s *Service) GetTodosByTag(tags []string) []*models.Todo {
	var results []*models.Todo
	
	for _, todo := range s.todoList.Todos {
		for _, todoTag := range todo.Tags {
			for _, searchTag := range tags {
				if strings.EqualFold(todoTag, searchTag) {
					results = append(results, todo)
					goto nextTodo
				}
			}
		}
	nextTodo:
	}
	
	return results
}

// GetRecentTodos returns recently updated todos
func (s *Service) GetRecentTodos(limit int) []*models.Todo {
	todos := make([]*models.Todo, len(s.todoList.Todos))
	copy(todos, s.todoList.Todos)
	
	// Sort by updated time (most recent first)
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].UpdatedAt.After(todos[j].UpdatedAt)
	})
	
	if limit > 0 && limit < len(todos) {
		return todos[:limit]
	}
	
	return todos
}

// GetTodoCountByStatus returns count of todos by status
func (s *Service) GetTodoCountByStatus() map[models.Status]int {
	counts := make(map[models.Status]int)
	
	for _, todo := range s.todoList.Todos {
		counts[todo.Status]++
	}
	
	return counts
}

// GetTodoCountByPriority returns count of todos by priority
func (s *Service) GetTodoCountByPriority() map[models.Priority]int {
	counts := make(map[models.Priority]int)
	
	for _, todo := range s.todoList.Todos {
		counts[todo.Priority]++
	}
	
	return counts
}

// GetMostUsedTags returns the most frequently used tags
func (s *Service) GetMostUsedTags(limit int) []utils.TagCount {
	tagCounts := make(map[string]int)
	
	for _, todo := range s.todoList.Todos {
		for _, tag := range todo.Tags {
			tagCounts[tag]++
		}
	}
	
	var tagCountsSlice []utils.TagCount
	for tag, count := range tagCounts {
		tagCountsSlice = append(tagCountsSlice, utils.TagCount{Tag: tag, Count: count})
	}
	
	// Sort by count (descending)
	sort.Slice(tagCountsSlice, func(i, j int) bool {
		return tagCountsSlice[i].Count > tagCountsSlice[j].Count
	})
	
	if limit > 0 && limit < len(tagCountsSlice) {
		return tagCountsSlice[:limit]
	}
	
	return tagCountsSlice
}
