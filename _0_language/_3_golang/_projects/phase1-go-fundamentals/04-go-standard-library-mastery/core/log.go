package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Custom log levels
const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Custom logger with levels
type LevelLogger struct {
	logger *log.Logger
	level  int
}

func NewLevelLogger(w io.Writer, prefix string, flag int, level int) *LevelLogger {
	return &LevelLogger{
		logger: log.New(w, prefix, flag),
		level:  level,
	}
}

func (l *LevelLogger) Debug(v ...interface{}) {
	if l.level <= DEBUG {
		l.logger.Printf("[DEBUG] %s", fmt.Sprint(v...))
	}
}

func (l *LevelLogger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.logger.Printf("[INFO] %s", fmt.Sprint(v...))
	}
}

func (l *LevelLogger) Warning(v ...interface{}) {
	if l.level <= WARNING {
		l.logger.Printf("[WARNING] %s", fmt.Sprint(v...))
	}
}

func (l *LevelLogger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.logger.Printf("[ERROR] %s", fmt.Sprint(v...))
	}
}

func (l *LevelLogger) Fatal(v ...interface{}) {
	if l.level <= FATAL {
		l.logger.Fatalf("[FATAL] %s", fmt.Sprint(v...))
	}
}

func main() {
	fmt.Println("🚀 Go log Package Mastery Examples")
	fmt.Println("===================================")

	// 1. Basic Logging
	fmt.Println("\n1. Basic Logging:")
	
	// Simple logging
	log.Print("Hello, World!")
	log.Printf("Current time: %s", time.Now().Format("2006-01-02 15:04:05"))
	log.Println("This is a log message with newline")
	
	// Log with different levels
	log.Print("INFO: Application started")
	log.Print("WARNING: Low memory warning")
	log.Print("ERROR: Database connection failed")

	// 2. Log Flags
	fmt.Println("\n2. Log Flags:")
	
	// Set different flags
	flags := []int{
		log.Ldate,
		log.Ltime,
		log.Lmicroseconds,
		log.Llongfile,
		log.Lshortfile,
		log.LUTC,
		log.LstdFlags,
	}
	
	flagNames := []string{
		"Ldate",
		"Ltime", 
		"Lmicroseconds",
		"Llongfile",
		"Lshortfile",
		"LUTC",
		"LstdFlags",
	}
	
	for i, flag := range flags {
		log.SetFlags(flag)
		log.Printf("Flag %s: %s", flagNames[i], "This is a test message")
	}

	// 3. Log Prefix
	fmt.Println("\n3. Log Prefix:")
	
	// Set different prefixes
	prefixes := []string{"", "APP: ", "ERROR: ", "DEBUG: ", "USER: "}
	
	for _, prefix := range prefixes {
		log.SetPrefix(prefix)
		log.Print("This is a test message")
	}

	// 4. Log Output
	fmt.Println("\n4. Log Output:")
	
	// Log to different outputs
	log.SetOutput(os.Stdout)
	log.Print("This goes to stdout")
	
	// Log to stderr
	log.SetOutput(os.Stderr)
	log.Print("This goes to stderr")
	
	// Reset to stdout
	log.SetOutput(os.Stdout)

	// 5. Fatal and Panic
	fmt.Println("\n5. Fatal and Panic:")
	
	// Note: We won't actually call Fatal or Panic to avoid terminating the program
	fmt.Println("Fatal and Panic functions would terminate the program")
	fmt.Println("log.Fatal() - logs and calls os.Exit(1)")
	fmt.Println("log.Panic() - logs and calls panic()")
	
	// Example of what they would do:
	// log.Fatal("This would terminate the program")
	// log.Panic("This would panic the program")

	// 6. Custom Logger
	fmt.Println("\n6. Custom Logger:")
	
	// Create custom logger
	customLogger := log.New(os.Stdout, "CUSTOM: ", log.LstdFlags)
	customLogger.Print("This is from custom logger")
	customLogger.Printf("User %s performed action %s", "alice", "login")
	customLogger.Println("Custom logger with newline")

	// 7. Multiple Loggers
	fmt.Println("\n7. Multiple Loggers:")
	
	// Create different loggers for different purposes
	infoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	errorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	debugLogger := log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
	
	infoLogger.Print("Application started successfully")
	errorLogger.Print("Failed to connect to database")
	debugLogger.Print("Processing user request")

	// 8. Log to File
	fmt.Println("\n8. Log to File:")
	
	// Create log file
	logFile, err := os.CreateTemp("", "example_*.log")
	if err != nil {
		log.Printf("Error creating log file: %v", err)
	} else {
		defer os.Remove(logFile.Name())
		
		// Log to file
		fileLogger := log.New(logFile, "FILE: ", log.LstdFlags)
		fileLogger.Print("This message goes to the log file")
		fileLogger.Printf("Timestamp: %s", time.Now().Format(time.RFC3339))
		
		// Read and display file contents
		logFile.Seek(0, 0)
		content, _ := io.ReadAll(logFile)
		fmt.Printf("Log file contents:\n%s", string(content))
	}

	// 9. Structured Logging
	fmt.Println("\n9. Structured Logging:")
	
	// Create structured logger
	structuredLogger := log.New(os.Stdout, "", 0)
	
	// Log structured data
	userID := "12345"
	action := "login"
	timestamp := time.Now().Format(time.RFC3339)
	
	structuredLogger.Printf("timestamp=%s user_id=%s action=%s status=success", 
		timestamp, userID, action)
	structuredLogger.Printf("timestamp=%s user_id=%s action=%s status=failed error=invalid_credentials", 
		timestamp, userID, "login")

	// 10. Log Levels with Custom Logger
	fmt.Println("\n10. Log Levels with Custom Logger:")
	
	// Create level logger
	levelLogger := NewLevelLogger(os.Stdout, "", log.LstdFlags, INFO)
	
	levelLogger.Debug("This debug message won't be shown (level too low)")
	levelLogger.Info("This info message will be shown")
	levelLogger.Warning("This warning message will be shown")
	levelLogger.Error("This error message will be shown")
	// levelLogger.Fatal("This fatal message would terminate the program")

	// 11. Log Rotation Simulation
	fmt.Println("\n11. Log Rotation Simulation:")
	
	// Simulate log rotation
	for i := 1; i <= 3; i++ {
		log.Printf("Log entry %d - simulating log rotation", i)
		time.Sleep(100 * time.Millisecond)
	}

	// 12. Log Filtering
	fmt.Println("\n12. Log Filtering:")
	
	// Create filtered logger
	filteredLogger := NewFilteredLogger(os.Stdout, "FILTERED: ", log.LstdFlags, "ERROR")
	
	filteredLogger.Log("INFO", "This info message will be filtered out")
	filteredLogger.Log("ERROR", "This error message will be shown")
	filteredLogger.Log("WARNING", "This warning message will be filtered out")

	// 13. Log Metrics
	fmt.Println("\n13. Log Metrics:")
	
	// Count log messages by level
	logMetrics := make(map[string]int)
	
	messages := []struct {
		level   string
		message string
	}{
		{"INFO", "User logged in"},
		{"ERROR", "Database connection failed"},
		{"WARNING", "Memory usage high"},
		{"INFO", "User logged out"},
		{"ERROR", "File not found"},
	}
	
	for _, msg := range messages {
		logMetrics[msg.level]++
		log.Printf("[%s] %s", msg.level, msg.message)
	}
	
	fmt.Println("Log metrics:")
	for level, count := range logMetrics {
		fmt.Printf("  %s: %d messages\n", level, count)
	}

	// 14. Log Context
	fmt.Println("\n14. Log Context:")
	
	// Add context to logs
	contextLogger := NewContextLogger(os.Stdout, "CONTEXT: ", log.LstdFlags)
	
	contextLogger.SetContext("user_id", "12345")
	contextLogger.SetContext("session_id", "abc123")
	contextLogger.Log("User performed action", "login")
	contextLogger.Log("User performed action", "logout")
	
	contextLogger.SetContext("user_id", "67890")
	contextLogger.Log("User performed action", "login")

	// 15. Log Performance
	fmt.Println("\n15. Log Performance:")
	
	// Benchmark logging performance
	iterations := 10000
	
	start := time.Now()
	for i := 0; i < iterations; i++ {
		log.Printf("Performance test message %d", i)
	}
	logTime := time.Since(start)
	
	fmt.Printf("Logged %d messages in %v\n", iterations, logTime)
	fmt.Printf("Average time per message: %v\n", logTime/time.Duration(iterations))

	// 16. Log Formatting
	fmt.Println("\n16. Log Formatting:")
	
	// Different log formats
	formats := []struct {
		name   string
		format string
	}{
		{"Simple", "%s"},
		{"With timestamp", "[%s] %s"},
		{"With level", "[%s] %s: %s"},
		{"JSON", `{"timestamp":"%s","level":"%s","message":"%s"}`},
	}
	
	for _, fmt := range formats {
		log.Printf("Format: %s", fmt.name)
		log.Printf(fmt.format, time.Now().Format("15:04:05"), "INFO", "Test message")
	}

	// 17. Log Buffering
	fmt.Println("\n17. Log Buffering:")
	
	// Create buffered logger
	bufferedLogger := NewBufferedLogger(os.Stdout, "BUFFERED: ", log.LstdFlags, 5)
	
	// Add messages to buffer
	for i := 1; i <= 7; i++ {
		bufferedLogger.Log(fmt.Sprintf("Buffered message %d", i))
	}
	
	// Flush buffer
	bufferedLogger.Flush()

	// 18. Log Validation
	fmt.Println("\n18. Log Validation:")
	
	// Validate log messages
	validator := NewLogValidator()
	
	testMessages := []string{
		"Valid log message",
		"", // Empty message
		strings.Repeat("x", 1000), // Very long message
		"Message with\nnewline",
		"Message with\ttab",
	}
	
	for i, msg := range testMessages {
		if validator.Validate(msg) {
			log.Printf("Message %d is valid: %s", i+1, msg)
		} else {
			log.Printf("Message %d is invalid: %s", i+1, msg)
		}
	}

	fmt.Println("\n🎉 log Package Mastery Complete!")
}

// Helper types and functions

// FilteredLogger filters log messages by level
type FilteredLogger struct {
	logger *log.Logger
	level  string
}

func NewFilteredLogger(w io.Writer, prefix string, flag int, level string) *FilteredLogger {
	return &FilteredLogger{
		logger: log.New(w, prefix, flag),
		level:  level,
	}
}

func (l *FilteredLogger) Log(level, message string) {
	if level == l.level {
		l.logger.Printf("[%s] %s", level, message)
	}
}

// ContextLogger adds context to log messages
type ContextLogger struct {
	logger  *log.Logger
	context map[string]string
}

func NewContextLogger(w io.Writer, prefix string, flag int) *ContextLogger {
	return &ContextLogger{
		logger:  log.New(w, prefix, flag),
		context: make(map[string]string),
	}
}

func (l *ContextLogger) SetContext(key, value string) {
	l.context[key] = value
}

func (l *ContextLogger) Log(message string, details ...string) {
	var contextStr strings.Builder
	for key, value := range l.context {
		contextStr.WriteString(fmt.Sprintf("%s=%s ", key, value))
	}
	
	if len(details) > 0 {
		contextStr.WriteString(strings.Join(details, " "))
	}
	
	l.logger.Printf("%s%s", contextStr.String(), message)
}

// BufferedLogger buffers log messages
type BufferedLogger struct {
	logger *log.Logger
	buffer []string
	maxSize int
}

func NewBufferedLogger(w io.Writer, prefix string, flag int, maxSize int) *BufferedLogger {
	return &BufferedLogger{
		logger:  log.New(w, prefix, flag),
		buffer:  make([]string, 0, maxSize),
		maxSize: maxSize,
	}
}

func (l *BufferedLogger) Log(message string) {
	l.buffer = append(l.buffer, message)
	if len(l.buffer) >= l.maxSize {
		l.Flush()
	}
}

func (l *BufferedLogger) Flush() {
	for _, msg := range l.buffer {
		l.logger.Print(msg)
	}
	l.buffer = l.buffer[:0]
}

// LogValidator validates log messages
type LogValidator struct {
	maxLength int
}

func NewLogValidator() *LogValidator {
	return &LogValidator{
		maxLength: 1000,
	}
}

func (v *LogValidator) Validate(message string) bool {
	if message == "" {
		return false
	}
	if len(message) > v.maxLength {
		return false
	}
	if strings.Contains(message, "\n") || strings.Contains(message, "\t") {
		return false
	}
	return true
}
