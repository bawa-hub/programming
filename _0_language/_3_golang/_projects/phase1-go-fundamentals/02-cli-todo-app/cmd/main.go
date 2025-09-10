package main

import (
	"cli-todo-app/internal/cli"
	"cli-todo-app/internal/storage"
	"cli-todo-app/internal/todo"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define command-line flags
	var (
		dataFile = flag.String("data", "data/todos.json", "Path to data file")
		backupFile = flag.String("backup", "data/backup.json", "Path to backup file")
		memoryMode = flag.Bool("memory", false, "Use in-memory storage (no persistence)")
		help = flag.Bool("help", false, "Show help message")
	)
	
	flag.Parse()

	// Show help if requested
	if *help {
		showHelp()
		return
	}

	// Create data directory if it doesn't exist
	if !*memoryMode {
		dataDir := filepath.Dir(*dataFile)
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			fmt.Printf("Error creating data directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Initialize storage
	var storageBackend storage.Storage
	if *memoryMode {
		storageBackend = storage.NewMemoryStorage()
		fmt.Println("Using in-memory storage (data will not persist)")
	} else {
		storageBackend = storage.NewFileStorage(*dataFile)
	}

	// Create storage manager with backup
	var storageManager *storage.StorageManager
	if *memoryMode {
		// For memory mode, use memory storage for both primary and backup
		storageManager = storage.NewStorageManager(storageBackend, storage.NewMemoryStorage())
	} else {
		// For file mode, use file storage for primary and backup
		backupStorage := storage.NewFileStorage(*backupFile)
		storageManager = storage.NewStorageManager(storageBackend, backupStorage)
	}

	// Create todo service
	service := todo.NewService(storageManager)
	defer service.Close()

	// Create and run CLI
	cliApp := cli.New(service)
	
	fmt.Println("🚀 Starting CLI Todo App...")
	fmt.Printf("📁 Data file: %s\n", *dataFile)
	if !*memoryMode {
		fmt.Printf("💾 Backup file: %s\n", *backupFile)
	}
	fmt.Println()

	// Run the CLI
	cliApp.Run()
}

func showHelp() {
	fmt.Println("CLI Todo App - A comprehensive todo management system")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  cli-todo-app [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -data string")
	fmt.Println("        Path to data file (default \"data/todos.json\")")
	fmt.Println("  -backup string")
	fmt.Println("        Path to backup file (default \"data/backup.json\")")
	fmt.Println("  -memory")
	fmt.Println("        Use in-memory storage (no persistence)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  cli-todo-app")
	fmt.Println("  cli-todo-app -data /path/to/todos.json")
	fmt.Println("  cli-todo-app -memory")
	fmt.Println()
	fmt.Println("Features:")
	fmt.Println("  • Create, edit, and delete todos")
	fmt.Println("  • Set priorities and due dates")
	fmt.Println("  • Organize with categories and tags")
	fmt.Println("  • Search and filter todos")
	fmt.Println("  • Export/import data")
	fmt.Println("  • Statistics and reporting")
	fmt.Println("  • Auto-save and backup")
	fmt.Println()
	fmt.Println("This application demonstrates all Go fundamentals:")
	fmt.Println("  • Primitive data types (int, string, bool, time.Time)")
	fmt.Println("  • Arrays and slices for collections")
	fmt.Println("  • Structs for data modeling")
	fmt.Println("  • Interfaces for polymorphism")
	fmt.Println("  • Pointers for memory management")
	fmt.Println("  • Error handling and validation")
	fmt.Println("  • File I/O and JSON processing")
	fmt.Println("  • Concurrency with goroutines")
	fmt.Println("  • Command-line argument parsing")
}
