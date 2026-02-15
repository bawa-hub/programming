# CLI Todo App - Go Fundamentals Mastery Project

A comprehensive command-line todo management application that demonstrates mastery of all Go fundamentals including primitive data types, arrays/slices, structs, interfaces, pointers, and more.

## 🎯 Project Overview

This CLI Todo App is designed to help you master Go fundamentals through hands-on practice. It implements a complete todo management system using all the core Go concepts you've learned.

## ✨ Features

### Core Todo Management
- **Create, edit, and delete todos** with titles and descriptions
- **Set priorities** (Low, Medium, High, Urgent) with visual indicators
- **Set due dates** with flexible date parsing
- **Track status** (Pending, In Progress, Completed, Cancelled)
- **Add tags** for organization and filtering
- **Categorize todos** with custom categories

### Advanced Features
- **Search and filter** todos by various criteria
- **Sort todos** by different fields
- **Statistics and reporting** with detailed analytics
- **Export/Import** data in JSON and CSV formats
- **Auto-save** with backup functionality
- **Overdue detection** and due date warnings

### Data Management
- **Persistent storage** with JSON files
- **Backup system** for data safety
- **Memory mode** for testing without persistence
- **Data validation** and error handling

## 🏗️ Architecture

The application demonstrates clean architecture principles:

```
cli-todo-app/
├── cmd/                    # Application entry point
│   └── main.go
├── internal/               # Internal packages
│   ├── cli/               # Command-line interface
│   ├── storage/           # Data persistence layer
│   └── todo/              # Business logic
├── pkg/                   # Public packages
│   ├── models/            # Data models and structs
│   └── utils/             # Utility functions
├── data/                  # Data storage directory
└── build/                 # Build output
```

## 🚀 Getting Started

### Prerequisites
- Go 1.19 or later
- Make (optional, for build automation)

### Installation

1. **Clone or download the project**
2. **Build the application**
   ```bash
   make build
   # or
   go build -o build/cli-todo-app ./cmd
   ```

3. **Run the application**
   ```bash
   make run
   # or
   ./build/cli-todo-app
   ```

### Quick Start

1. **Add your first todo**
   ```
   todo> add "Learn Go programming"
   ✓ Created todo #1: Learn Go programming
   ```

2. **List all todos**
   ```
   todo> list
   ID   Status    Priority  Title                    Due Date        Category
   ---- --------  --------  -----------------------  --------------  ----------
   1    ⏳ Pending  🟡 Medium  Learn Go programming     No due date     None
   ```

3. **Set priority and due date**
   ```
   todo> priority 1 high
   ✓ Set priority of todo #1 to High
   
   todo> due 1 2024-12-31
   ✓ Set due date of todo #1 to 2024-12-31 12:00
   ```

4. **Complete the todo**
   ```
   todo> complete 1
   ✓ Completed todo #1
   ```

## 📚 Go Concepts Demonstrated

### 1. Primitive Data Types
- **Integers**: `int`, `int64` for IDs and counts
- **Strings**: `string` for titles, descriptions, tags
- **Booleans**: `bool` for status flags
- **Time**: `time.Time` for dates and timestamps
- **Constants**: `iota` for priority and status enums

### 2. Arrays and Slices
- **Dynamic collections**: `[]*Todo` for todo lists
- **Slice operations**: `append()`, `copy()`, filtering
- **Range loops**: Iterating over collections
- **Memory management**: Understanding slice capacity

### 3. Structs
- **Data modeling**: `Todo`, `Category`, `TodoList` structs
- **Methods**: Value and pointer receivers
- **Embedding**: Struct composition
- **JSON tags**: Serialization metadata
- **String methods**: Custom string representation

### 4. Interfaces
- **Storage abstraction**: `Storage` interface
- **Polymorphism**: Different storage implementations
- **Type assertions**: Runtime type checking
- **Interface composition**: Combining interfaces

### 5. Pointers
- **Memory management**: Pointer receivers for methods
- **Reference passing**: Avoiding data copying
- **Nil safety**: Checking for nil pointers
- **Pointer arithmetic**: Safe pointer operations

### 6. Advanced Concepts
- **Error handling**: Custom error types and wrapping
- **Concurrency**: Goroutines for auto-save
- **File I/O**: JSON serialization/deserialization
- **Command parsing**: CLI argument processing
- **Data validation**: Input validation and sanitization

## 🎮 Usage Examples

### Basic Todo Operations
```bash
# Add a todo
add "Complete Go tutorial" "Finish the comprehensive Go tutorial"

# List todos
list
list pending
list completed
list overdue

# Show todo details
show 1

# Edit a todo
edit 1 "Complete Go tutorial and exercises"

# Set priority
priority 1 high

# Set due date
due 1 2024-12-31 18:00

# Add tags
tag 1 add programming
tag 1 add tutorial

# Complete a todo
complete 1
```

### Category Management
```bash
# List categories
category list

# Add a category
category add "Work" "Work-related tasks" "red"

# Set todo category
category set 1 1

# Delete a category
category delete 1
```

### Search and Filter
```bash
# Search todos
search "Go programming"

# Filter by status
filter status pending

# Filter by priority
filter priority high

# Filter by category
filter category 1

# Filter by tag
filter tag programming
```

### Sorting
```bash
# Sort by title
sort title

# Sort by priority (descending)
sort priority desc

# Sort by due date
sort due asc
```

### Statistics
```bash
# Show statistics
stats
```

### Export/Import
```bash
# Export to JSON
export json todos_backup.json

# Export to CSV
export csv todos.csv

# Import from JSON
import todos_backup.json
```

## 🧪 Testing

The project includes comprehensive tests demonstrating Go testing patterns:

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific test packages
make test-models
make test-storage
make test-todo
```

## 🔧 Development

### Build Commands
```bash
make build        # Build the application
make run          # Run the application
make run-memory   # Run in memory mode
make test         # Run tests
make fmt          # Format code
make lint         # Run linter
make clean        # Clean build artifacts
```

### Project Structure
- **Models**: Data structures and business logic
- **Storage**: Persistence layer with multiple backends
- **CLI**: User interface and command parsing
- **Utils**: Utility functions and helpers
- **Tests**: Comprehensive test coverage

## 📊 Data Models

### Todo
```go
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
```

### Category
```go
type Category struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    Color       string `json:"color,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
}
```

## 🎓 Learning Outcomes

After working with this project, you will have mastered:

1. **Go Syntax**: All fundamental language constructs
2. **Data Types**: Primitive types, structs, and interfaces
3. **Memory Management**: Pointers, slices, and garbage collection
4. **Error Handling**: Custom errors and error propagation
5. **File I/O**: JSON serialization and file operations
6. **Concurrency**: Goroutines and channels
7. **Testing**: Unit tests and test coverage
8. **CLI Development**: Command parsing and user interaction
9. **Project Structure**: Clean architecture and package organization
10. **Real-world Applications**: Building complete applications

## 🚀 Next Steps

This project provides a solid foundation for:
- **System Programming**: File system operations, process management
- **Web Development**: HTTP servers, REST APIs, microservices
- **Concurrency**: Advanced goroutine patterns, worker pools
- **Database Integration**: SQL and NoSQL database operations
- **Cloud Development**: Containerization and deployment

## 🤝 Contributing

This is a learning project, but feel free to:
- Add new features
- Improve error handling
- Add more test cases
- Optimize performance
- Enhance the CLI interface

## 📝 License

This project is for educational purposes and demonstrates Go fundamentals mastery.

---

**Happy Coding! 🎉**

This CLI Todo App demonstrates that you've mastered Go fundamentals and are ready to build more complex applications. The concepts you've learned here will serve as the foundation for all your future Go development.
