package models

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFileTypeString(t *testing.T) {
	tests := []struct {
		fileType FileType
		expected string
	}{
		{FileTypeRegular, "regular"},
		{FileTypeDirectory, "directory"},
		{FileTypeSymlink, "symlink"},
		{FileTypeSocket, "socket"},
		{FileTypeNamedPipe, "named_pipe"},
		{FileTypeDevice, "device"},
		{FileTypeCharDevice, "char_device"},
		{FileTypeUnknown, "unknown"},
	}
	
	for _, test := range tests {
		if test.fileType.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.fileType.String())
		}
	}
}

func TestFileTypeIcon(t *testing.T) {
	tests := []struct {
		fileType FileType
		expected string
	}{
		{FileTypeRegular, "📄"},
		{FileTypeDirectory, "📁"},
		{FileTypeSymlink, "🔗"},
		{FileTypeSocket, "🔌"},
		{FileTypeNamedPipe, "🔧"},
		{FileTypeDevice, "💾"},
		{FileTypeCharDevice, "⌨️"},
		{FileTypeUnknown, "❓"},
	}
	
	for _, test := range tests {
		if test.fileType.Icon() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.fileType.Icon())
		}
	}
}

func TestNewFileInfo(t *testing.T) {
	// Create a temporary file for testing
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.txt")
	
	err := os.WriteFile(tempFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// Get file info
	info, err := os.Stat(tempFile)
	if err != nil {
		t.Fatalf("Failed to stat test file: %v", err)
	}
	
	// Create FileInfo
	fileInfo := NewFileInfo(tempFile, info, tempDir)
	
	// Test basic properties
	if fileInfo.Path != tempFile {
		t.Errorf("Expected path %s, got %s", tempFile, fileInfo.Path)
	}
	
	if fileInfo.Name != "test.txt" {
		t.Errorf("Expected name test.txt, got %s", fileInfo.Name)
	}
	
	if fileInfo.Size != 12 {
		t.Errorf("Expected size 12, got %d", fileInfo.Size)
	}
	
	if fileInfo.FileType != FileTypeRegular {
		t.Errorf("Expected FileTypeRegular, got %v", fileInfo.FileType)
	}
	
	if fileInfo.IsDir {
		t.Error("Expected IsDir to be false")
	}
	
	if fileInfo.IsSymlink {
		t.Error("Expected IsSymlink to be false")
	}
	
	if fileInfo.Extension != "txt" {
		t.Errorf("Expected extension txt, got %s", fileInfo.Extension)
	}
	
	if fileInfo.RelativePath != "test.txt" {
		t.Errorf("Expected relative path test.txt, got %s", fileInfo.RelativePath)
	}
	
	if fileInfo.Depth != 0 {
		t.Errorf("Expected depth 0, got %d", fileInfo.Depth)
	}
}

func TestFileInfoMethods(t *testing.T) {
	// Create a temporary file
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.txt")
	
	err := os.WriteFile(tempFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	info, err := os.Stat(tempFile)
	if err != nil {
		t.Fatalf("Failed to stat test file: %v", err)
	}
	
	fileInfo := NewFileInfo(tempFile, info, tempDir)
	
	// Test IsEmpty
	if fileInfo.IsEmpty() {
		t.Error("Expected file to not be empty")
	}
	
	// Test IsExecutable
	if fileInfo.IsExecutable() {
		t.Error("Expected file to not be executable")
	}
	
	// Test IsReadable
	if !fileInfo.IsReadable() {
		t.Error("Expected file to be readable")
	}
	
	// Test IsWritable
	if !fileInfo.IsWritable() {
		t.Error("Expected file to be writable")
	}
	
	// Test GetExtension
	if fileInfo.GetExtension() != "txt" {
		t.Errorf("Expected extension txt, got %s", fileInfo.GetExtension())
	}
	
	// Test GetMimeType
	expectedMimeType := "text/plain"
	if fileInfo.GetMimeType() != expectedMimeType {
		t.Errorf("Expected MIME type %s, got %s", expectedMimeType, fileInfo.GetMimeType())
	}
	
	// Test GetSizeString
	sizeStr := fileInfo.GetSizeString()
	if sizeStr == "" {
		t.Error("Expected non-empty size string")
	}
	
	// Test GetAgeString
	ageStr := fileInfo.GetAgeString()
	if ageStr == "" {
		t.Error("Expected non-empty age string")
	}
}

func TestFileInfoMatchesFilter(t *testing.T) {
	// Create a temporary file
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.txt")
	
	err := os.WriteFile(tempFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	info, err := os.Stat(tempFile)
	if err != nil {
		t.Fatalf("Failed to stat test file: %v", err)
	}
	
	fileInfo := NewFileInfo(tempFile, info, tempDir)
	
	// Test file type filter
	filter := &Filter{
		FileTypes: []FileType{FileTypeRegular},
	}
	
	if !fileInfo.MatchesFilter(filter) {
		t.Error("Expected file to match file type filter")
	}
	
	// Test extension filter
	filter = &Filter{
		Extensions: []string{"txt"},
	}
	
	if !fileInfo.MatchesFilter(filter) {
		t.Error("Expected file to match extension filter")
	}
	
	// Test size filter
	filter = &Filter{
		MinSize: 1,
		MaxSize: 100,
	}
	
	if !fileInfo.MatchesFilter(filter) {
		t.Error("Expected file to match size filter")
	}
	
	// Test negated filter
	filter = &Filter{
		FileTypes: []FileType{FileTypeDirectory},
		Negate:    true,
	}
	
	if !fileInfo.MatchesFilter(filter) {
		t.Error("Expected file to match negated filter")
	}
}

func TestNewScanResult(t *testing.T) {
	rootPath := "/test/path"
	result := NewScanResult(rootPath)
	
	if result.RootPath != rootPath {
		t.Errorf("Expected root path %s, got %s", rootPath, result.RootPath)
	}
	
	if result.TotalFiles != 0 {
		t.Errorf("Expected total files 0, got %d", result.TotalFiles)
	}
	
	if result.TotalDirs != 0 {
		t.Errorf("Expected total dirs 0, got %d", result.TotalDirs)
	}
	
	if result.TotalSize != 0 {
		t.Errorf("Expected total size 0, got %d", result.TotalSize)
	}
	
	if result.Files == nil {
		t.Error("Expected Files to be initialized")
	}
	
	if result.Errors == nil {
		t.Error("Expected Errors to be initialized")
	}
	
	if result.Statistics == nil {
		t.Error("Expected Statistics to be initialized")
	}
}

func TestScanResultAddFile(t *testing.T) {
	result := NewScanResult("/test")
	
	// Create a mock file
	file := &FileInfo{
		Path:  "/test/file.txt",
		Name:  "file.txt",
		Size:  100,
		IsDir: false,
	}
	
	result.AddFile(file)
	
	if result.TotalFiles != 1 {
		t.Errorf("Expected total files 1, got %d", result.TotalFiles)
	}
	
	if result.TotalDirs != 0 {
		t.Errorf("Expected total dirs 0, got %d", result.TotalDirs)
	}
	
	if result.TotalSize != 100 {
		t.Errorf("Expected total size 100, got %d", result.TotalSize)
	}
	
	if len(result.Files) != 1 {
		t.Errorf("Expected 1 file in Files slice, got %d", len(result.Files))
	}
}

func TestScanResultAddError(t *testing.T) {
	result := NewScanResult("/test")
	
	result.AddError("/test/file.txt", os.ErrNotExist, "stat")
	
	if len(result.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(result.Errors))
	}
	
	error := result.Errors[0]
	if error.Path != "/test/file.txt" {
		t.Errorf("Expected error path /test/file.txt, got %s", error.Path)
	}
	
	if error.Type != "stat" {
		t.Errorf("Expected error type stat, got %s", error.Type)
	}
}

func TestScanResultFinish(t *testing.T) {
	result := NewScanResult("/test")
	
	// Set start time to 1 second ago
	result.StartTime = time.Now().Add(-time.Second)
	
	result.Finish()
	
	if result.EndTime.IsZero() {
		t.Error("Expected EndTime to be set")
	}
	
	if result.ScanDuration == 0 {
		t.Error("Expected ScanDuration to be set")
	}
	
	// Check that duration is approximately 1 second
	if result.ScanDuration < time.Second || result.ScanDuration > 2*time.Second {
		t.Errorf("Expected scan duration around 1 second, got %v", result.ScanDuration)
	}
}

func TestScanResultGetFileCount(t *testing.T) {
	result := NewScanResult("/test")
	
	// Add some files and directories
	result.AddFile(&FileInfo{IsDir: false})
	result.AddFile(&FileInfo{IsDir: true})
	result.AddFile(&FileInfo{IsDir: false})
	
	expectedCount := int64(3)
	if result.GetFileCount() != expectedCount {
		t.Errorf("Expected file count %d, got %d", expectedCount, result.GetFileCount())
	}
}

func TestScanResultGetAverageFileSize(t *testing.T) {
	result := NewScanResult("/test")
	
	// Add some files
	result.AddFile(&FileInfo{Size: 100, IsDir: false})
	result.AddFile(&FileInfo{Size: 200, IsDir: false})
	result.AddFile(&FileInfo{Size: 300, IsDir: false})
	
	expectedAverage := int64(200)
	if result.GetAverageFileSize() != expectedAverage {
		t.Errorf("Expected average file size %d, got %d", expectedAverage, result.GetAverageFileSize())
	}
}

func TestScanResultGetScanRate(t *testing.T) {
	result := NewScanResult("/test")
	
	// Set up scan duration
	result.StartTime = time.Now().Add(-time.Second)
	result.EndTime = time.Now()
	result.ScanDuration = time.Second
	
	// Add some files
	result.AddFile(&FileInfo{IsDir: false})
	result.AddFile(&FileInfo{IsDir: true})
	
	expectedRate := 2.0 // 2 files per second
	if result.GetScanRate() != expectedRate {
		t.Errorf("Expected scan rate %f, got %f", expectedRate, result.GetScanRate())
	}
}

func TestFilter(t *testing.T) {
	// Create a test file
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.txt")
	
	err := os.WriteFile(tempFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	info, err := os.Stat(tempFile)
	if err != nil {
		t.Fatalf("Failed to stat test file: %v", err)
	}
	
	fileInfo := NewFileInfo(tempFile, info, tempDir)
	
	// Test various filters
	tests := []struct {
		name     string
		filter   *Filter
		expected bool
	}{
		{
			name: "file type filter - match",
			filter: &Filter{
				FileTypes: []FileType{FileTypeRegular},
			},
			expected: true,
		},
		{
			name: "file type filter - no match",
			filter: &Filter{
				FileTypes: []FileType{FileTypeDirectory},
			},
			expected: false,
		},
		{
			name: "extension filter - match",
			filter: &Filter{
				Extensions: []string{"txt"},
			},
			expected: true,
		},
		{
			name: "extension filter - no match",
			filter: &Filter{
				Extensions: []string{"pdf"},
			},
			expected: false,
		},
		{
			name: "size filter - match",
			filter: &Filter{
				MinSize: 1,
				MaxSize: 1000,
			},
			expected: true,
		},
		{
			name: "size filter - too small",
			filter: &Filter{
				MinSize: 1000,
			},
			expected: false,
		},
		{
			name: "size filter - too large",
			filter: &Filter{
				MaxSize: 1,
			},
			expected: false,
		},
		{
			name: "pattern filter - match",
			filter: &Filter{
				Pattern: "test",
			},
			expected: true,
		},
		{
			name: "pattern filter - no match",
			filter: &Filter{
				Pattern: "notfound",
			},
			expected: false,
		},
		{
			name: "negated filter - match",
			filter: &Filter{
				FileTypes: []FileType{FileTypeDirectory},
				Negate:    true,
			},
			expected: true,
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := fileInfo.MatchesFilter(test.filter)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
