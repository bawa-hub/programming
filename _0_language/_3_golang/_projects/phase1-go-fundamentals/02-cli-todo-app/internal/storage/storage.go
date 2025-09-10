package storage

import (
	"cli-todo-app/pkg/models"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Storage interface for different storage backends
type Storage interface {
	Save(todoList *models.TodoList) error
	Load() (*models.TodoList, error)
	Exists() bool
	Delete() error
	GetInfo() (os.FileInfo, error)
}

// FileStorage implements storage using JSON files
type FileStorage struct {
	filePath string
	mutex    sync.RWMutex
}

// MemoryStorage implements storage using in-memory storage
type MemoryStorage struct {
	data  *models.TodoList
	mutex sync.RWMutex
}

// NewFileStorage creates a new file storage instance
func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{
		filePath: filePath,
	}
}

// NewMemoryStorage creates a new memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: models.NewTodoList(),
	}
}

// FileStorage methods

// Save saves the todo list to a JSON file
func (fs *FileStorage) Save(todoList *models.TodoList) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	
	// Create directory if it doesn't exist
	dir := filepath.Dir(fs.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	// Convert to JSON
	data, err := todoList.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal todo list: %w", err)
	}
	
	// Write to file
	if err := ioutil.WriteFile(fs.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	return nil
}

// Load loads the todo list from a JSON file
func (fs *FileStorage) Load() (*models.TodoList, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()
	
	// Check if file exists
	if !fs.Exists() {
		return models.NewTodoList(), nil
	}
	
	// Read file
	data, err := ioutil.ReadFile(fs.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	
	// Parse JSON
	todoList := models.NewTodoList()
	if err := todoList.FromJSON(data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal todo list: %w", err)
	}
	
	return todoList, nil
}

// Exists checks if the storage file exists
func (fs *FileStorage) Exists() bool {
	_, err := os.Stat(fs.filePath)
	return !os.IsNotExist(err)
}

// Delete deletes the storage file
func (fs *FileStorage) Delete() error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	
	if !fs.Exists() {
		return nil
	}
	
	return os.Remove(fs.filePath)
}

// GetInfo returns file information
func (fs *FileStorage) GetInfo() (os.FileInfo, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()
	
	return os.Stat(fs.filePath)
}

// MemoryStorage methods

// Save saves the todo list to memory
func (ms *MemoryStorage) Save(todoList *models.TodoList) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	
	ms.data = todoList
	return nil
}

// Load loads the todo list from memory
func (ms *MemoryStorage) Load() (*models.TodoList, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	
	// Return a copy to prevent external modifications
	data, err := ms.data.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	
	todoList := models.NewTodoList()
	if err := todoList.FromJSON(data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}
	
	return todoList, nil
}

// Exists always returns true for memory storage
func (ms *MemoryStorage) Exists() bool {
	return true
}

// Delete clears the memory storage
func (ms *MemoryStorage) Delete() error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	
	ms.data = models.NewTodoList()
	return nil
}

// GetInfo returns nil for memory storage
func (ms *MemoryStorage) GetInfo() (os.FileInfo, error) {
	return nil, fmt.Errorf("memory storage does not have file info")
}

// StorageManager manages different storage backends
type StorageManager struct {
	primary   Storage
	backup    Storage
	lastSave  time.Time
	mutex     sync.RWMutex
}

// NewStorageManager creates a new storage manager
func NewStorageManager(primary, backup Storage) *StorageManager {
	return &StorageManager{
		primary: primary,
		backup:  backup,
	}
}

// Save saves to both primary and backup storage
func (sm *StorageManager) Save(todoList *models.TodoList) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	// Save to primary storage
	if err := sm.primary.Save(todoList); err != nil {
		// Try backup storage if primary fails
		if backupErr := sm.backup.Save(todoList); backupErr != nil {
			return fmt.Errorf("failed to save to both primary and backup: primary=%v, backup=%v", err, backupErr)
		}
		return fmt.Errorf("failed to save to primary storage, saved to backup: %w", err)
	}
	
	// Save to backup storage
	if err := sm.backup.Save(todoList); err != nil {
		// Log warning but don't fail the operation
		fmt.Printf("Warning: failed to save to backup storage: %v\n", err)
	}
	
	sm.lastSave = time.Now()
	return nil
}

// Load loads from primary storage, falls back to backup
func (sm *StorageManager) Load() (*models.TodoList, error) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	// Try primary storage first
	if sm.primary.Exists() {
		todoList, err := sm.primary.Load()
		if err == nil {
			return todoList, nil
		}
		fmt.Printf("Warning: failed to load from primary storage: %v\n", err)
	}
	
	// Fall back to backup storage
	if sm.backup.Exists() {
		todoList, err := sm.backup.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load from backup storage: %w", err)
		}
		return todoList, nil
	}
	
	// Return empty todo list if neither exists
	return models.NewTodoList(), nil
}

// GetLastSaveTime returns the last save time
func (sm *StorageManager) GetLastSaveTime() time.Time {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	return sm.lastSave
}

// Backup creates a backup of the current data
func (sm *StorageManager) Backup(backupPath string) error {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	// Load current data
	todoList, err := sm.Load()
	if err != nil {
		return fmt.Errorf("failed to load data for backup: %w", err)
	}
	
	// Create backup storage
	backupStorage := NewFileStorage(backupPath)
	
	// Save to backup
	if err := backupStorage.Save(todoList); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}
	
	return nil
}

// Restore restores data from a backup
func (sm *StorageManager) Restore(backupPath string) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	// Create backup storage
	backupStorage := NewFileStorage(backupPath)
	
	// Load from backup
	todoList, err := backupStorage.Load()
	if err != nil {
		return fmt.Errorf("failed to load backup: %w", err)
	}
	
	// Save to primary storage
	if err := sm.primary.Save(todoList); err != nil {
		return fmt.Errorf("failed to restore to primary storage: %w", err)
	}
	
	sm.lastSave = time.Now()
	return nil
}

// Delete deletes both primary and backup storage
func (sm *StorageManager) Delete() error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	// Delete primary storage
	if err := sm.primary.Delete(); err != nil {
		return fmt.Errorf("failed to delete primary storage: %w", err)
	}
	
	// Delete backup storage
	if err := sm.backup.Delete(); err != nil {
		return fmt.Errorf("failed to delete backup storage: %w", err)
	}
	
	return nil
}

// Exists checks if either primary or backup storage exists
func (sm *StorageManager) Exists() bool {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	return sm.primary.Exists() || sm.backup.Exists()
}

// GetInfo returns info from primary storage
func (sm *StorageManager) GetInfo() (os.FileInfo, error) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	return sm.primary.GetInfo()
}

// GetStorageInfo returns information about storage backends
func (sm *StorageManager) GetStorageInfo() map[string]interface{} {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	info := make(map[string]interface{})
	
	// Primary storage info
	info["primary_exists"] = sm.primary.Exists()
	if info["primary_exists"].(bool) {
		if fileInfo, err := sm.primary.GetInfo(); err == nil {
			info["primary_size"] = fileInfo.Size()
			info["primary_mod_time"] = fileInfo.ModTime()
		}
	}
	
	// Backup storage info
	info["backup_exists"] = sm.backup.Exists()
	if info["backup_exists"].(bool) {
		if fileInfo, err := sm.backup.GetInfo(); err == nil {
			info["backup_size"] = fileInfo.Size()
			info["backup_mod_time"] = fileInfo.ModTime()
		}
	}
	
	info["last_save"] = sm.lastSave
	
	return info
}

// AutoSave enables automatic saving at specified intervals
type AutoSave struct {
	storage    Storage
	interval   time.Duration
	todoList   *models.TodoList
	stopChan   chan bool
	mutex      sync.RWMutex
}

// NewAutoSave creates a new auto-save instance
func NewAutoSave(storage Storage, interval time.Duration) *AutoSave {
	return &AutoSave{
		storage:  storage,
		interval: interval,
		stopChan: make(chan bool),
	}
}

// Start starts the auto-save goroutine
func (as *AutoSave) Start(todoList *models.TodoList) {
	as.mutex.Lock()
	as.todoList = todoList
	as.mutex.Unlock()
	
	go as.autoSaveLoop()
}

// Stop stops the auto-save goroutine
func (as *AutoSave) Stop() {
	as.stopChan <- true
}

// UpdateTodoList updates the todo list reference
func (as *AutoSave) UpdateTodoList(todoList *models.TodoList) {
	as.mutex.Lock()
	as.todoList = todoList
	as.mutex.Unlock()
}

// autoSaveLoop runs the auto-save loop
func (as *AutoSave) autoSaveLoop() {
	ticker := time.NewTicker(as.interval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			as.mutex.RLock()
			todoList := as.todoList
			as.mutex.RUnlock()
			
			if todoList != nil {
				if err := as.storage.Save(todoList); err != nil {
					fmt.Printf("Auto-save failed: %v\n", err)
				}
			}
		case <-as.stopChan:
			return
		}
	}
}

// Export utilities for different formats

// ExportToJSON exports todo list to JSON
func ExportToJSON(todoList *models.TodoList, filePath string) error {
	data, err := todoList.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal todo list: %w", err)
	}
	
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}
	
	return nil
}

// ImportFromJSON imports todo list from JSON
func ImportFromJSON(filePath string) (*models.TodoList, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}
	
	todoList := models.NewTodoList()
	if err := todoList.FromJSON(data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	
	return todoList, nil
}

// ExportToCSV exports todo list to CSV format
func ExportToCSV(todoList *models.TodoList, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()
	
	// Write CSV header
	header := "ID,Title,Description,Priority,Status,Category,Tags,DueDate,CreatedAt,UpdatedAt,CompletedAt\n"
	if _, err := file.WriteString(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}
	
	// Write todo data
	for _, todo := range todoList.Todos {
		category := ""
		if todo.Category != nil {
			category = todo.Category.Name
		}
		
		tags := ""
		if len(todo.Tags) > 0 {
			tags = fmt.Sprintf("\"%s\"", strings.Join(todo.Tags, ","))
		}
		
		dueDate := ""
		if todo.DueDate != nil {
			dueDate = todo.DueDate.Format("2006-01-02 15:04:05")
		}
		
		completedAt := ""
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format("2006-01-02 15:04:05")
		}
		
		row := fmt.Sprintf("%d,\"%s\",\"%s\",%s,%s,\"%s\",%s,%s,%s,%s,%s\n",
			todo.ID,
			todo.Title,
			todo.Description,
			todo.Priority.String(),
			todo.Status.String(),
			category,
			tags,
			dueDate,
			todo.CreatedAt.Format("2006-01-02 15:04:05"),
			todo.UpdatedAt.Format("2006-01-02 15:04:05"),
			completedAt,
		)
		
		if _, err := file.WriteString(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %w", err)
		}
	}
	
	return nil
}
