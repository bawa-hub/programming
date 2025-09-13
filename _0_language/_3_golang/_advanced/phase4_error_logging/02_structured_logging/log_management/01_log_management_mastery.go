// 📋 LOG MANAGEMENT MASTERY
// Advanced log management including rotation, retention, and lifecycle management
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

// ============================================================================
// LOG ROTATION TYPES
// ============================================================================

type RotationStrategy int

const (
	RotationBySize RotationStrategy = iota
	RotationByTime
	RotationBySizeAndTime
)

type LogFile struct {
	Path      string    `json:"path"`
	Size      int64     `json:"size"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
	Compressed bool     `json:"compressed"`
}

type RotationConfig struct {
	Strategy     RotationStrategy `json:"strategy"`
	MaxSize      int64           `json:"max_size"`      // bytes
	MaxAge       time.Duration   `json:"max_age"`       // duration
	MaxFiles     int             `json:"max_files"`     // number of files
	Compress     bool            `json:"compress"`      // compress old files
	CompressAge  time.Duration   `json:"compress_age"`  // age to compress
	BackupDir    string          `json:"backup_dir"`    // backup directory
}

// ============================================================================
// LOG ROTATOR
// ============================================================================

type LogRotator struct {
	config     RotationConfig
	currentFile *os.File
	currentPath string
	mu          sync.RWMutex
	files       []LogFile
}

func NewLogRotator(config RotationConfig) *LogRotator {
	// Ensure backup directory exists
	if config.BackupDir != "" {
		os.MkdirAll(config.BackupDir, 0755)
	}
	
	return &LogRotator{
		config: config,
		files:  make([]LogFile, 0),
	}
}

func (lr *LogRotator) OpenLogFile(basePath string) error {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	
	// Close current file if open
	if lr.currentFile != nil {
		lr.currentFile.Close()
	}
	
	// Open new log file
	file, err := os.OpenFile(basePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	
	lr.currentFile = file
	lr.currentPath = basePath
	
	// Get file info
	info, err := file.Stat()
	if err != nil {
		return err
	}
	
	// Add to files list
	logFile := LogFile{
		Path:     basePath,
		Size:     info.Size(),
		Created:  info.ModTime(),
		Modified: info.ModTime(),
	}
	lr.files = append(lr.files, logFile)
	
	return nil
}

func (lr *LogRotator) Write(data []byte) (int, error) {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	
	if lr.currentFile == nil {
		return 0, fmt.Errorf("no log file open")
	}
	
	// Check if rotation is needed
	if lr.shouldRotate() {
		if err := lr.rotate(); err != nil {
			return 0, err
		}
	}
	
	return lr.currentFile.Write(data)
}

func (lr *LogRotator) shouldRotate() bool {
	if lr.currentFile == nil {
		return false
	}
	
	// Get current file size
	info, err := lr.currentFile.Stat()
	if err != nil {
		return false
	}
	
	switch lr.config.Strategy {
	case RotationBySize:
		return info.Size() >= lr.config.MaxSize
	case RotationByTime:
		return time.Since(info.ModTime()) >= lr.config.MaxAge
	case RotationBySizeAndTime:
		return info.Size() >= lr.config.MaxSize || time.Since(info.ModTime()) >= lr.config.MaxAge
	default:
		return false
	}
}

func (lr *LogRotator) rotate() error {
	if lr.currentFile == nil {
		return fmt.Errorf("no current file to rotate")
	}
	
	// Close current file
	lr.currentFile.Close()
	
	// Generate rotated filename
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	rotatedPath := fmt.Sprintf("%s.%s", lr.currentPath, timestamp)
	
	// Move current file to rotated name
	if err := os.Rename(lr.currentPath, rotatedPath); err != nil {
		return err
	}
	
	// Update file info
	for i, file := range lr.files {
		if file.Path == lr.currentPath {
			lr.files[i].Path = rotatedPath
			lr.files[i].Modified = time.Now()
			break
		}
	}
	
	// Open new log file
	file, err := os.OpenFile(lr.currentPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	
	lr.currentFile = file
	
	// Cleanup old files
	lr.cleanup()
	
	// Compress old files if needed
	if lr.config.Compress {
		go lr.compressOldFiles()
	}
	
	return nil
}

func (lr *LogRotator) cleanup() {
	// Sort files by modification time (oldest first)
	sort.Slice(lr.files, func(i, j int) bool {
		return lr.files[i].Modified.Before(lr.files[j].Modified)
	})
	
	// Remove excess files
	if lr.config.MaxFiles > 0 && len(lr.files) > lr.config.MaxFiles {
		filesToRemove := lr.files[:len(lr.files)-lr.config.MaxFiles]
		for _, file := range filesToRemove {
			os.Remove(file.Path)
		}
		lr.files = lr.files[lr.config.MaxFiles:]
	}
	
	// Remove files older than max age
	if lr.config.MaxAge > 0 {
		cutoff := time.Now().Add(-lr.config.MaxAge)
		var remainingFiles []LogFile
		for _, file := range lr.files {
			if file.Modified.After(cutoff) {
				remainingFiles = append(remainingFiles, file)
			} else {
				os.Remove(file.Path)
			}
		}
		lr.files = remainingFiles
	}
}

func (lr *LogRotator) compressOldFiles() {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	
	cutoff := time.Now().Add(-lr.config.CompressAge)
	for i, file := range lr.files {
		if !file.Compressed && file.Modified.Before(cutoff) {
			if err := lr.compressFile(file.Path); err == nil {
				lr.files[i].Compressed = true
				lr.files[i].Path = file.Path + ".gz"
			}
		}
	}
}

func (lr *LogRotator) compressFile(filePath string) error {
	// Open source file
	srcFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	
	// Create compressed file
	compressedFile, err := os.Create(filePath + ".gz")
	if err != nil {
		return err
	}
	defer compressedFile.Close()
	
	// Create gzip writer
	gzWriter := gzip.NewWriter(compressedFile)
	defer gzWriter.Close()
	
	// Copy and compress
	_, err = io.Copy(gzWriter, srcFile)
	if err != nil {
		return err
	}
	
	// Remove original file
	return os.Remove(filePath)
}

func (lr *LogRotator) Close() error {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	
	if lr.currentFile != nil {
		return lr.currentFile.Close()
	}
	return nil
}

func (lr *LogRotator) GetFiles() []LogFile {
	lr.mu.RLock()
	defer lr.mu.RUnlock()
	
	files := make([]LogFile, len(lr.files))
	copy(files, lr.files)
	return files
}

// ============================================================================
// LOG RETENTION MANAGER
// ============================================================================

type RetentionPolicy struct {
	MaxAge       time.Duration `json:"max_age"`
	MaxSize      int64         `json:"max_size"`
	MaxFiles     int           `json:"max_files"`
	Compress     bool          `json:"compress"`
	Archive      bool          `json:"archive"`
	ArchivePath  string        `json:"archive_path"`
}

type RetentionManager struct {
	policies map[string]RetentionPolicy
	mu       sync.RWMutex
}

func NewRetentionManager() *RetentionManager {
	return &RetentionManager{
		policies: make(map[string]RetentionPolicy),
	}
}

func (rm *RetentionManager) AddPolicy(name string, policy RetentionPolicy) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.policies[name] = policy
}

func (rm *RetentionManager) ApplyPolicy(name string, logDir string) error {
	rm.mu.RLock()
	policy, exists := rm.policies[name]
	rm.mu.RUnlock()
	
	if !exists {
		return fmt.Errorf("policy %s not found", name)
	}
	
	return rm.applyRetentionPolicy(policy, logDir)
}

func (rm *RetentionManager) applyRetentionPolicy(policy RetentionPolicy, logDir string) error {
	// Get all log files
	files, err := rm.getLogFiles(logDir)
	if err != nil {
		return err
	}
	
	// Sort by modification time (oldest first)
	sort.Slice(files, func(i, j int) bool {
		return files[i].Modified.Before(files[j].Modified)
	})
	
	// Apply age-based retention
	if policy.MaxAge > 0 {
		cutoff := time.Now().Add(-policy.MaxAge)
		for _, file := range files {
			if file.Modified.Before(cutoff) {
				if policy.Archive {
					rm.archiveFile(file, policy.ArchivePath)
				}
				os.Remove(file.Path)
			}
		}
	}
	
	// Apply size-based retention
	if policy.MaxSize > 0 {
		var totalSize int64
		for i := len(files) - 1; i >= 0; i-- {
			totalSize += files[i].Size
			if totalSize > policy.MaxSize {
				if policy.Archive {
					rm.archiveFile(files[i], policy.ArchivePath)
				}
				os.Remove(files[i].Path)
			}
		}
	}
	
	// Apply file count retention
	if policy.MaxFiles > 0 && len(files) > policy.MaxFiles {
		filesToRemove := files[:len(files)-policy.MaxFiles]
		for _, file := range filesToRemove {
			if policy.Archive {
				rm.archiveFile(file, policy.ArchivePath)
			}
			os.Remove(file.Path)
		}
	}
	
	return nil
}

func (rm *RetentionManager) getLogFiles(logDir string) ([]LogFile, error) {
	var files []LogFile
	
	err := filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() && (filepath.Ext(path) == ".log" || filepath.Ext(path) == ".gz") {
			files = append(files, LogFile{
				Path:      path,
				Size:      info.Size(),
				Created:   info.ModTime(),
				Modified:  info.ModTime(),
				Compressed: filepath.Ext(path) == ".gz",
			})
		}
		
		return nil
	})
	
	return files, err
}

func (rm *RetentionManager) archiveFile(file LogFile, archivePath string) error {
	// Create archive directory if it doesn't exist
	if err := os.MkdirAll(archivePath, 0755); err != nil {
		return err
	}
	
	// Generate archive filename
	archiveFile := filepath.Join(archivePath, filepath.Base(file.Path))
	
	// Copy file to archive
	src, err := os.Open(file.Path)
	if err != nil {
		return err
	}
	defer src.Close()
	
	dst, err := os.Create(archiveFile)
	if err != nil {
		return err
	}
	defer dst.Close()
	
	_, err = io.Copy(dst, src)
	return err
}

// ============================================================================
// LOG LIFECYCLE MANAGER
// ============================================================================

type LifecycleStage int

const (
	StageIngestion LifecycleStage = iota
	StageProcessing
	StageStorage
	StageIndexing
	StageRetention
	StageArchival
	StageDeletion
)

type LogLifecycleManager struct {
	stages map[LifecycleStage]func(LogFile) error
	mu     sync.RWMutex
}

func NewLogLifecycleManager() *LogLifecycleManager {
	llm := &LogLifecycleManager{
		stages: make(map[LifecycleStage]func(LogFile) error),
	}
	
	// Set default stages
	llm.SetStage(StageIngestion, llm.ingestLog)
	llm.SetStage(StageProcessing, llm.processLog)
	llm.SetStage(StageStorage, llm.storeLog)
	llm.SetStage(StageIndexing, llm.indexLog)
	llm.SetStage(StageRetention, llm.applyRetention)
	llm.SetStage(StageArchival, llm.archiveLog)
	llm.SetStage(StageDeletion, llm.deleteLog)
	
	return llm
}

func (llm *LogLifecycleManager) SetStage(stage LifecycleStage, handler func(LogFile) error) {
	llm.mu.Lock()
	defer llm.mu.Unlock()
	llm.stages[stage] = handler
}

func (llm *LogLifecycleManager) ProcessLog(file LogFile) error {
	llm.mu.RLock()
	defer llm.mu.RUnlock()
	
	// Process through all stages
	for stage := StageIngestion; stage <= StageDeletion; stage++ {
		if handler, exists := llm.stages[stage]; exists {
			if err := handler(file); err != nil {
				return fmt.Errorf("stage %d failed: %v", stage, err)
			}
		}
	}
	
	return nil
}

func (llm *LogLifecycleManager) ingestLog(file LogFile) error {
	fmt.Printf("   📥 Ingesting log: %s\n", file.Path)
	return nil
}

func (llm *LogLifecycleManager) processLog(file LogFile) error {
	fmt.Printf("   🔄 Processing log: %s\n", file.Path)
	return nil
}

func (llm *LogLifecycleManager) storeLog(file LogFile) error {
	fmt.Printf("   💾 Storing log: %s\n", file.Path)
	return nil
}

func (llm *LogLifecycleManager) indexLog(file LogFile) error {
	fmt.Printf("   🔍 Indexing log: %s\n", file.Path)
	return nil
}

func (llm *LogLifecycleManager) applyRetention(file LogFile) error {
	fmt.Printf("   ⏰ Applying retention to: %s\n", file.Path)
	return nil
}

func (llm *LogLifecycleManager) archiveLog(file LogFile) error {
	fmt.Printf("   📦 Archiving log: %s\n", file.Path)
	return nil
}

func (llm *LogLifecycleManager) deleteLog(file LogFile) error {
	fmt.Printf("   🗑️  Deleting log: %s\n", file.Path)
	return nil
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateLogRotation() {
	fmt.Println("\n=== Log Rotation ===")
	
	// Create rotation config
	config := RotationConfig{
		Strategy:    RotationBySize,
		MaxSize:     1024, // 1KB for demo
		MaxFiles:    3,
		Compress:    true,
		CompressAge: 1 * time.Minute,
		BackupDir:   "logs/backup",
	}
	
	// Create rotator
	rotator := NewLogRotator(config)
	defer rotator.Close()
	
	// Open log file
	rotator.OpenLogFile("logs/app.log")
	
	// Write some data to trigger rotation
	for i := 0; i < 10; i++ {
		line := fmt.Sprintf("Log line %d: %s\n", i, time.Now().Format(time.RFC3339))
		rotator.Write([]byte(line))
	}
	
	// List rotated files
	files := rotator.GetFiles()
	fmt.Printf("   📊 Created %d log files\n", len(files))
	for _, file := range files {
		fmt.Printf("   📄 %s (%.2f KB)\n", file.Path, float64(file.Size)/1024)
	}
}

func demonstrateRetentionPolicies() {
	fmt.Println("\n=== Retention Policies ===")
	
	// Create retention manager
	manager := NewRetentionManager()
	
	// Add retention policy
	policy := RetentionPolicy{
		MaxAge:      24 * time.Hour,  // 1 day
		MaxSize:     10 * 1024 * 1024, // 10MB
		MaxFiles:    5,
		Compress:    true,
		Archive:     true,
		ArchivePath: "logs/archive",
	}
	
	manager.AddPolicy("app_logs", policy)
	
	// Create some test log files
	os.MkdirAll("logs", 0755)
	for i := 0; i < 3; i++ {
		filename := fmt.Sprintf("logs/test_%d.log", i)
		file, _ := os.Create(filename)
		file.WriteString(fmt.Sprintf("Test log content %d\n", i))
		file.Close()
	}
	
	// Apply retention policy
	err := manager.ApplyPolicy("app_logs", "logs")
	if err != nil {
		fmt.Printf("   ❌ Error applying policy: %v\n", err)
	} else {
		fmt.Println("   📊 Retention policy applied successfully")
	}
}

func demonstrateLogLifecycle() {
	fmt.Println("\n=== Log Lifecycle Management ===")
	
	// Create lifecycle manager
	lifecycleManager := NewLogLifecycleManager()
	
	// Create test log file
	testFile := LogFile{
		Path:     "logs/lifecycle_test.log",
		Size:     1024,
		Created:  time.Now(),
		Modified: time.Now(),
	}
	
	// Process through lifecycle
	err := lifecycleManager.ProcessLog(testFile)
	if err != nil {
		fmt.Printf("   ❌ Lifecycle processing failed: %v\n", err)
	} else {
		fmt.Println("   📊 Log processed through complete lifecycle")
	}
}

func demonstrateCompression() {
	fmt.Println("\n=== Log Compression ===")
	
	// Create a test log file
	testFile := "logs/compression_test.log"
	os.MkdirAll("logs", 0755)
	
	file, err := os.Create(testFile)
	if err != nil {
		fmt.Printf("   ❌ Error creating test file: %v\n", err)
		return
	}
	
	// Write some data
	for i := 0; i < 100; i++ {
		file.WriteString(fmt.Sprintf("Log line %d: This is some test data for compression\n", i))
	}
	file.Close()
	
	// Get original size
	info, _ := os.Stat(testFile)
	originalSize := info.Size()
	
	// Compress the file
	rotator := NewLogRotator(RotationConfig{})
	err = rotator.compressFile(testFile)
	if err != nil {
		fmt.Printf("   ❌ Compression failed: %v\n", err)
		return
	}
	
	// Get compressed size
	compressedFile := testFile + ".gz"
	info, _ = os.Stat(compressedFile)
	compressedSize := info.Size()
	
	compressionRatio := float64(compressedSize) / float64(originalSize) * 100
	
	fmt.Printf("   📊 Original size: %d bytes\n", originalSize)
	fmt.Printf("   📊 Compressed size: %d bytes\n", compressedSize)
	fmt.Printf("   📊 Compression ratio: %.1f%%\n", compressionRatio)
	
	// Cleanup
	os.Remove(testFile)
	os.Remove(compressedFile)
}

func demonstrateLogManagement() {
	fmt.Println("\n=== Complete Log Management ===")
	
	// Create comprehensive log management system
	config := RotationConfig{
		Strategy:    RotationBySizeAndTime,
		MaxSize:     2048, // 2KB
		MaxAge:      1 * time.Minute,
		MaxFiles:    3,
		Compress:    true,
		CompressAge: 30 * time.Second,
		BackupDir:   "logs/backup",
	}
	
	rotator := NewLogRotator(config)
	defer rotator.Close()
	
	// Open log file
	rotator.OpenLogFile("logs/management_test.log")
	
	// Write logs over time to trigger various management actions
	for i := 0; i < 20; i++ {
		line := fmt.Sprintf("[%s] Log entry %d: %s\n", 
			time.Now().Format("15:04:05"), i, "This is a test log entry for management")
		rotator.Write([]byte(line))
		
		// Small delay to simulate real logging
		time.Sleep(50 * time.Millisecond)
	}
	
	// Show final state
	files := rotator.GetFiles()
	fmt.Printf("   📊 Total log files: %d\n", len(files))
	for _, file := range files {
		status := "active"
		if file.Compressed {
			status = "compressed"
		}
		fmt.Printf("   📄 %s (%.2f KB) - %s\n", file.Path, float64(file.Size)/1024, status)
	}
}

func main() {
	fmt.Println("📋 LOG MANAGEMENT MASTERY")
	fmt.Println("=========================")
	
	demonstrateLogRotation()
	demonstrateRetentionPolicies()
	demonstrateLogLifecycle()
	demonstrateCompression()
	demonstrateLogManagement()
	
	fmt.Println("\n🎉 LOG MANAGEMENT MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Log rotation strategies")
	fmt.Println("✅ Retention policies and cleanup")
	fmt.Println("✅ Log lifecycle management")
	fmt.Println("✅ Compression and archival")
	fmt.Println("✅ Complete log management system")
	
	fmt.Println("\n🚀 You are now ready for Performance Optimization Mastery!")
}
