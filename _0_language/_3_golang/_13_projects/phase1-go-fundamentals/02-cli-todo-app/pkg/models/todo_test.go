package models

import (
	"strings"
	"testing"
	"time"
)

func TestTodoCreation(t *testing.T) {
	todo := NewTodo("Test Todo", "This is a test todo")
	
	if todo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got '%s'", todo.Title)
	}
	
	if todo.Description != "This is a test todo" {
		t.Errorf("Expected description 'This is a test todo', got '%s'", todo.Description)
	}
	
	if todo.Priority != Medium {
		t.Errorf("Expected priority Medium, got %v", todo.Priority)
	}
	
	if todo.Status != Pending {
		t.Errorf("Expected status Pending, got %v", todo.Status)
	}
	
	if todo.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
	
	if todo.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should not be zero")
	}
}

func TestTodoMethods(t *testing.T) {
	todo := NewTodo("Test Todo", "Test description")
	
	// Test MarkCompleted
	todo.MarkCompleted()
	if todo.Status != Completed {
		t.Errorf("Expected status Completed, got %v", todo.Status)
	}
	
	if todo.CompletedAt == nil {
		t.Error("CompletedAt should not be nil after marking completed")
	}
	
	// Test MarkInProgress
	todo.MarkInProgress()
	if todo.Status != InProgress {
		t.Errorf("Expected status InProgress, got %v", todo.Status)
	}
	
	// Test UpdatePriority
	todo.UpdatePriority(High)
	if todo.Priority != High {
		t.Errorf("Expected priority High, got %v", todo.Priority)
	}
	
	// Test AddTag
	todo.AddTag("test")
	if len(todo.Tags) != 1 || todo.Tags[0] != "test" {
		t.Errorf("Expected tag 'test', got %v", todo.Tags)
	}
	
	// Test RemoveTag
	todo.RemoveTag("test")
	if len(todo.Tags) != 0 {
		t.Errorf("Expected no tags, got %v", todo.Tags)
	}
	
	// Test SetDueDate
	dueDate := time.Now().Add(24 * time.Hour)
	todo.SetDueDate(dueDate)
	if todo.DueDate == nil || !todo.DueDate.Equal(dueDate) {
		t.Error("Due date not set correctly")
	}
}

func TestTodoIsOverdue(t *testing.T) {
	todo := NewTodo("Test Todo", "Test description")
	
	// Test with no due date
	if todo.IsOverdue() {
		t.Error("Todo with no due date should not be overdue")
	}
	
	// Test with future due date
	futureDate := time.Now().Add(24 * time.Hour)
	todo.SetDueDate(futureDate)
	if todo.IsOverdue() {
		t.Error("Todo with future due date should not be overdue")
	}
	
	// Test with past due date
	pastDate := time.Now().Add(-24 * time.Hour)
	todo.SetDueDate(pastDate)
	if !todo.IsOverdue() {
		t.Error("Todo with past due date should be overdue")
	}
	
	// Test completed todo (should not be overdue)
	todo.MarkCompleted()
	if todo.IsOverdue() {
		t.Error("Completed todo should not be overdue")
	}
}

func TestTodoDaysUntilDue(t *testing.T) {
	todo := NewTodo("Test Todo", "Test description")
	
	// Test with no due date
	if todo.DaysUntilDue() != -1 {
		t.Errorf("Expected -1 for no due date, got %d", todo.DaysUntilDue())
	}
	
	// Test with due date
	dueDate := time.Now().Add(48 * time.Hour)
	todo.SetDueDate(dueDate)
	days := todo.DaysUntilDue()
	if days < 1 || days > 3 {
		t.Errorf("Expected 1-3 days, got %d", days)
	}
}

func TestPriorityString(t *testing.T) {
	tests := []struct {
		priority Priority
		expected string
	}{
		{Low, "Low"},
		{Medium, "Medium"},
		{High, "High"},
		{Urgent, "Urgent"},
		{Priority(999), "Unknown"},
	}
	
	for _, test := range tests {
		if test.priority.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.priority.String())
		}
	}
}

func TestStatusString(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{Pending, "Pending"},
		{InProgress, "In Progress"},
		{Completed, "Completed"},
		{Cancelled, "Cancelled"},
		{Status(999), "Unknown"},
	}
	
	for _, test := range tests {
		if test.status.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.status.String())
		}
	}
}

func TestTodoListOperations(t *testing.T) {
	todoList := NewTodoList()
	
	// Test AddTodo
	todo1 := NewTodo("Todo 1", "Description 1")
	todoList.AddTodo(todo1)
	
	if len(todoList.Todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todoList.Todos))
	}
	
	if todo1.ID != 1 {
		t.Errorf("Expected ID 1, got %d", todo1.ID)
	}
	
	// Test GetTodoByID
	retrievedTodo := todoList.GetTodoByID(1)
	if retrievedTodo != todo1 {
		t.Error("Retrieved todo should be the same as added todo")
	}
	
	// Test RemoveTodo
	if !todoList.RemoveTodo(1) {
		t.Error("RemoveTodo should return true for existing todo")
	}
	
	if len(todoList.Todos) != 0 {
		t.Errorf("Expected 0 todos after removal, got %d", len(todoList.Todos))
	}
	
	// Test RemoveTodo with non-existent ID
	if todoList.RemoveTodo(999) {
		t.Error("RemoveTodo should return false for non-existent todo")
	}
}

func TestTodoListFiltering(t *testing.T) {
	todoList := NewTodoList()
	
	// Add test todos
	todo1 := NewTodo("Todo 1", "Description 1")
	todo1.Priority = High
	todo1.Status = Pending
	todoList.AddTodo(todo1)
	
	todo2 := NewTodo("Todo 2", "Description 2")
	todo2.Priority = Low
	todo2.Status = Completed
	todoList.AddTodo(todo2)
	
	todo3 := NewTodo("Todo 3", "Description 3")
	todo3.Priority = Medium
	todo3.Status = InProgress
	todoList.AddTodo(todo3)
	
	// Test GetTodosByStatus
	pendingTodos := todoList.GetTodosByStatus(Pending)
	if len(pendingTodos) != 1 {
		t.Errorf("Expected 1 pending todo, got %d", len(pendingTodos))
	}
	
	completedTodos := todoList.GetTodosByStatus(Completed)
	if len(completedTodos) != 1 {
		t.Errorf("Expected 1 completed todo, got %d", len(completedTodos))
	}
	
	// Test GetTodosByPriority
	highPriorityTodos := todoList.GetTodosByPriority(High)
	if len(highPriorityTodos) != 1 {
		t.Errorf("Expected 1 high priority todo, got %d", len(highPriorityTodos))
	}
	
	// Test GetStats
	stats := todoList.GetStats()
	if stats["total"] != 3 {
		t.Errorf("Expected 3 total todos, got %d", stats["total"])
	}
	
	if stats["pending"] != 1 {
		t.Errorf("Expected 1 pending todo, got %d", stats["pending"])
	}
	
	if stats["completed"] != 1 {
		t.Errorf("Expected 1 completed todo, got %d", stats["completed"])
	}
}

func TestCategoryOperations(t *testing.T) {
	todoList := NewTodoList()
	
	// Test AddCategory
	category := NewCategory("Work", "Work-related tasks", "blue")
	todoList.AddCategory(category)
	
	if len(todoList.Categories) != 1 {
		t.Errorf("Expected 1 category, got %d", len(todoList.Categories))
	}
	
	if category.ID != 1 {
		t.Errorf("Expected category ID 1, got %d", category.ID)
	}
	
	// Test GetCategoryByID
	retrievedCategory := todoList.GetCategoryByID(1)
	if retrievedCategory != category {
		t.Error("Retrieved category should be the same as added category")
	}
	
	// Test GetCategoryByName
	categoryByName := todoList.GetCategoryByName("Work")
	if categoryByName != category {
		t.Error("Category retrieved by name should be the same as added category")
	}
	
	// Test RemoveCategory
	if !todoList.RemoveCategory(1) {
		t.Error("RemoveCategory should return true for existing category")
	}
	
	if len(todoList.Categories) != 0 {
		t.Errorf("Expected 0 categories after removal, got %d", len(todoList.Categories))
	}
}

func TestTodoListJSON(t *testing.T) {
	todoList := NewTodoList()
	
	// Add test data
	todo := NewTodo("Test Todo", "Test Description")
	todoList.AddTodo(todo)
	
	category := NewCategory("Test", "Test Category", "red")
	todoList.AddCategory(category)
	
	// Test ToJSON
	jsonData, err := todoList.ToJSON()
	if err != nil {
		t.Errorf("ToJSON failed: %v", err)
	}
	
	if len(jsonData) == 0 {
		t.Error("JSON data should not be empty")
	}
	
	// Test FromJSON
	newTodoList := NewTodoList()
	err = newTodoList.FromJSON(jsonData)
	if err != nil {
		t.Errorf("FromJSON failed: %v", err)
	}
	
	if len(newTodoList.Todos) != 1 {
		t.Errorf("Expected 1 todo after JSON round-trip, got %d", len(newTodoList.Todos))
	}
	
	if len(newTodoList.Categories) != 1 {
		t.Errorf("Expected 1 category after JSON round-trip, got %d", len(newTodoList.Categories))
	}
}

func TestPriorityGetPriorityColor(t *testing.T) {
	tests := []struct {
		priority Priority
		expected string
	}{
		{Low, "🟢"},
		{Medium, "🟡"},
		{High, "🟠"},
		{Urgent, "🔴"},
		{Priority(999), "⚪"},
	}
	
	for _, test := range tests {
		if test.priority.GetPriorityColor() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.priority.GetPriorityColor())
		}
	}
}

func TestStatusGetStatusIcon(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{Pending, "⏳"},
		{InProgress, "🔄"},
		{Completed, "✅"},
		{Cancelled, "❌"},
		{Status(999), "❓"},
	}
	
	for _, test := range tests {
		if test.status.GetStatusIcon() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.status.GetStatusIcon())
		}
	}
}

func TestTodoString(t *testing.T) {
	todo := NewTodo("Test Todo", "Test Description")
	todo.Priority = High
	todo.Status = InProgress
	
	str := todo.String()
	if !strings.Contains(str, "Test Todo") {
		t.Error("String representation should contain title")
	}
	
	if !strings.Contains(str, "High") {
		t.Error("String representation should contain priority")
	}
	
	if !strings.Contains(str, "In Progress") {
		t.Error("String representation should contain status")
	}
}

func TestCategoryString(t *testing.T) {
	category := NewCategory("Test Category", "Test Description", "blue")
	
	str := category.String()
	if !strings.Contains(str, "Test Category") {
		t.Error("String representation should contain name")
	}
	
	if !strings.Contains(str, "blue") {
		t.Error("String representation should contain color")
	}
}
