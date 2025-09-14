package main

import (
	"fmt"
	"log"
)

// 🚀 GOLANG DATA TYPES MASTERY - CRUD APPLICATION
// This application demonstrates ALL Go data types through practical CRUD operations

func main() {
	fmt.Println("🎯 GOLANG DATA TYPES MASTERY - CRUD APPLICATION")
	fmt.Println("================================================")
	
	// Initialize the application
	app := NewApp()
	
	// Run the interactive menu
	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
