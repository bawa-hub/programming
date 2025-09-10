package scanner

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"file-scanner/pkg/models"
	"file-scanner/pkg/patterns"
)

// Scanner represents a file system scanner
type Scanner struct {
	options    *models.ScanOptions
	workerPool *patterns.WorkerPool
	rateLimiter *patterns.RateLimiter
	semaphore  *patterns.Semaphore
	ctx        context.Context
	cancel     context.CancelFunc
	mu         sync.RWMutex
	stats      *ScanStats
}

// ScanStats represents scanning statistics
type ScanStats struct {
	FilesScanned    int64
	DirectoriesScanned int64
	ErrorsEncountered  int64
	BytesProcessed     int64
	StartTime          time.Time
	EndTime            time.Time
	mu                 sync.RWMutex
}

// NewScanner creates a new file system scanner
func NewScanner(options *models.ScanOptions) *Scanner {
	if options == nil {
		options = &models.ScanOptions{
			MaxDepth:       10,
			FollowSymlinks: false,
			IncludeHidden:  false,
			IncludeSystem:  false,
			Concurrency:    runtime.NumCPU(),
			BufferSize:     1024 * 1024, // 1MB
			Timeout:        30 * time.Minute,
			Progress:       true,
			Verbose:        false,
		}
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
	
	// Create worker pool
	workerPool := patterns.NewWorkerPool(options.Concurrency, options.Concurrency*2)
	
	// Create rate limiter (1000 operations per second)
	rateLimiter := patterns.NewRateLimiter(time.Millisecond, 1000)
	
	// Create semaphore for file operations
	semaphore := patterns.NewSemaphore(options.Concurrency * 2)
	
	return &Scanner{
		options:     options,
		workerPool:  workerPool,
		rateLimiter: rateLimiter,
		semaphore:   semaphore,
		ctx:         ctx,
		cancel:      cancel,
		stats: &ScanStats{
			StartTime: time.Now(),
		},
	}
}

// Scan scans a directory and returns the results
func (s *Scanner) Scan(rootPath string) (*models.ScanResult, error) {
	// Validate root path
	if rootPath == "" {
		return nil, fmt.Errorf("root path cannot be empty")
	}
	
	// Check if path exists
	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, fmt.Errorf("failed to stat root path: %w", err)
	}
	
	if !info.IsDir() {
		return nil, fmt.Errorf("root path must be a directory")
	}
	
	// Create scan result
	result := models.NewScanResult(rootPath)
	
	// Start progress reporting
	var progressWg sync.WaitGroup
	if s.options.Progress {
		progressWg.Add(1)
		go s.reportProgress(&progressWg, result)
	}
	
	// Start scanning (simplified without worker pool for now)
	err = s.scanDirectory(rootPath, result, 0)
	
	// Stop progress reporting
	if s.options.Progress {
		progressWg.Done()
		progressWg.Wait()
	}
	
	// Finish scan
	result.Finish()
	
	// Calculate statistics
	s.calculateStatistics(result)
	
	return result, err
}

// scanDirectory recursively scans a directory
func (s *Scanner) scanDirectory(path string, result *models.ScanResult, depth int) error {
	// Check depth limit
	if s.options.MaxDepth > 0 && depth > s.options.MaxDepth {
		return nil
	}
	
	// Check context cancellation
	select {
	case <-s.ctx.Done():
		return s.ctx.Err()
	default:
	}
	
	// Rate limiting
	s.rateLimiter.Wait()
	
	// Note: Semaphore removed to prevent deadlock
	// s.semaphore.Acquire()
	// defer s.semaphore.Release()
	
	// Read directory
	entries, err := os.ReadDir(path)
	if err != nil {
		result.AddError(path, err, "readdir")
		atomic.AddInt64(&s.stats.ErrorsEncountered, 1)
		return err
	}
	
	// Process entries
	for _, entry := range entries {
		// Check context cancellation
		select {
		case <-s.ctx.Done():
			return s.ctx.Err()
		default:
		}
		
		entryPath := filepath.Join(path, entry.Name())
		
		// Apply filters
		if !s.shouldProcessEntry(entryPath, entry) {
			continue
		}
		
		// Get file info
		fileInfo, err := entry.Info()
		if err != nil {
			result.AddError(entryPath, err, "stat")
			atomic.AddInt64(&s.stats.ErrorsEncountered, 1)
			continue
		}
		
		// Create FileInfo
		fi := models.NewFileInfo(entryPath, fileInfo, result.RootPath)
		fi.Depth = depth
		
		// Apply filters
		if !s.matchesFilters(fi) {
			continue
		}
		
		// Add to result
		result.AddFile(fi)
		
		// Update statistics
		if fi.IsDir {
			atomic.AddInt64(&s.stats.DirectoriesScanned, 1)
		} else {
			atomic.AddInt64(&s.stats.FilesScanned, 1)
			atomic.AddInt64(&s.stats.BytesProcessed, fi.Size)
		}
		
		// Process file if needed
		if s.options.CalculateHashes && !fi.IsDir {
			job := &HashJob{
				FileInfo: fi,
				Options:  s.options,
			}
			s.workerPool.Submit(job)
		}
		
		// Recursively scan subdirectories
		if fi.IsDir && (s.options.FollowSymlinks || !fi.IsSymlink) {
			if err := s.scanDirectory(entryPath, result, depth+1); err != nil {
				if err == context.Canceled {
					return err
				}
				// Continue with other directories
			}
		}
	}
	
	return nil
}

// shouldProcessEntry checks if an entry should be processed
func (s *Scanner) shouldProcessEntry(path string, entry os.DirEntry) bool {
	name := entry.Name()
	
	// Check hidden files
	if !s.options.IncludeHidden && strings.HasPrefix(name, ".") {
		return false
	}
	
	// Check system files
	if !s.options.IncludeSystem {
		if name == "System Volume Information" || 
		   name == "$RECYCLE.BIN" || 
		   name == "Thumbs.db" {
			return false
		}
	}
	
	// Check exclude paths
	for _, excludePath := range s.options.ExcludePaths {
		if strings.Contains(path, excludePath) {
			return false
		}
	}
	
	// Check include paths
	if len(s.options.IncludePaths) > 0 {
		found := false
		for _, includePath := range s.options.IncludePaths {
			if strings.Contains(path, includePath) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	
	return true
}

// matchesFilters checks if a file matches the filters
func (s *Scanner) matchesFilters(fi *models.FileInfo) bool {
	if len(s.options.Filters) == 0 {
		return true
	}
	
	for _, filter := range s.options.Filters {
		if !fi.MatchesFilter(filter) {
			return false
		}
	}
	
	return true
}

// calculateStatistics calculates comprehensive statistics
func (s *Scanner) calculateStatistics(result *models.ScanResult) {
	stats := result.Statistics
	
	// Initialize maps
	stats.FileCountByType = make(map[models.FileType]int64)
	stats.FileCountByExt = make(map[string]int64)
	stats.SizeByType = make(map[models.FileType]int64)
	stats.SizeByExt = make(map[string]int64)
	stats.PermissionStats = make(map[string]int64)
	stats.DepthDistribution = make(map[int]int64)
	stats.SizeDistribution = make(map[string]int64)
	stats.TimeDistribution = make(map[string]int64)
	
	// Process files
	for _, file := range result.Files {
		// Count by type
		stats.FileCountByType[file.FileType]++
		stats.SizeByType[file.FileType] += file.Size
		
		// Count by extension
		ext := file.GetExtension()
		if ext != "" {
			stats.FileCountByExt[ext]++
			stats.SizeByExt[ext] += file.Size
		}
		
		// Permission statistics
		perm := file.Permissions
		stats.PermissionStats[perm]++
		
		// Depth distribution
		stats.DepthDistribution[file.Depth]++
		
		// Size distribution
		sizeCategory := s.getSizeCategory(file.Size)
		stats.SizeDistribution[sizeCategory]++
		
		// Time distribution
		timeCategory := s.getTimeCategory(file.ModTime)
		stats.TimeDistribution[timeCategory]++
	}
	
	// Find largest files
	s.findLargestFiles(result, stats)
	
	// Find oldest and newest files
	s.findOldestNewestFiles(result, stats)
	
	// Find empty directories
	s.findEmptyDirectories(result, stats)
	
	// Find duplicates if requested
	if s.options.FindDuplicates {
		s.findDuplicates(result, stats)
	}
}

// getSizeCategory returns a size category for a file
func (s *Scanner) getSizeCategory(size int64) string {
	switch {
	case size == 0:
		return "0 bytes"
	case size < 1024:
		return "1KB-"
	case size < 1024*1024:
		return "1KB-1MB"
	case size < 1024*1024*1024:
		return "1MB-1GB"
	case size < 1024*1024*1024*1024:
		return "1GB-1TB"
	default:
		return "1TB+"
	}
}

// getTimeCategory returns a time category for a file
func (s *Scanner) getTimeCategory(modTime time.Time) string {
	age := time.Since(modTime)
	switch {
	case age < 24*time.Hour:
		return "Today"
	case age < 7*24*time.Hour:
		return "This week"
	case age < 30*24*time.Hour:
		return "This month"
	case age < 365*24*time.Hour:
		return "This year"
	default:
		return "Older"
	}
}

// findLargestFiles finds the largest files
func (s *Scanner) findLargestFiles(result *models.ScanResult, stats *models.Statistics) {
	// Sort files by size (descending)
	files := make([]*models.FileInfo, len(result.Files))
	copy(files, result.Files)
	
	// Simple bubble sort for largest files
	for i := 0; i < len(files)-1; i++ {
		for j := 0; j < len(files)-i-1; j++ {
			if files[j].Size < files[j+1].Size {
				files[j], files[j+1] = files[j+1], files[j]
			}
		}
	}
	
	// Take top 10
	limit := 10
	if len(files) < limit {
		limit = len(files)
	}
	stats.LargestFiles = files[:limit]
}

// findOldestNewestFiles finds the oldest and newest files
func (s *Scanner) findOldestNewestFiles(result *models.ScanResult, stats *models.Statistics) {
	if len(result.Files) == 0 {
		return
	}
	
	oldest := result.Files[0]
	newest := result.Files[0]
	
	for _, file := range result.Files {
		if file.ModTime.Before(oldest.ModTime) {
			oldest = file
		}
		if file.ModTime.After(newest.ModTime) {
			newest = file
		}
	}
	
	stats.OldestFiles = []*models.FileInfo{oldest}
	stats.NewestFiles = []*models.FileInfo{newest}
}

// findEmptyDirectories finds empty directories
func (s *Scanner) findEmptyDirectories(result *models.ScanResult, stats *models.Statistics) {
	for _, file := range result.Files {
		if file.IsDir && file.Size == 0 {
			stats.EmptyDirs = append(stats.EmptyDirs, file)
		}
	}
}

// findDuplicates finds duplicate files
func (s *Scanner) findDuplicates(result *models.ScanResult, stats *models.Statistics) {
	// Group files by size first
	sizeGroups := make(map[int64][]*models.FileInfo)
	for _, file := range result.Files {
		if !file.IsDir {
			sizeGroups[file.Size] = append(sizeGroups[file.Size], file)
		}
	}
	
	// Check for duplicates within each size group
	for size, files := range sizeGroups {
		if len(files) < 2 {
			continue
		}
		
		// Calculate hashes for files of the same size
		hashGroups := make(map[string][]*models.FileInfo)
		for _, file := range files {
			hash := s.calculateFileHash(file.Path)
			if hash != "" {
				hashGroups[hash] = append(hashGroups[hash], file)
			}
		}
		
		// Find groups with multiple files
		for hash, hashFiles := range hashGroups {
			if len(hashFiles) > 1 {
				stats.DuplicateFiles = append(stats.DuplicateFiles, &models.DuplicateGroup{
					Size:  size,
					Hash:  hash,
					Files: hashFiles,
					Count: len(hashFiles),
				})
			}
		}
	}
}

// calculateFileHash calculates the hash of a file
func (s *Scanner) calculateFileHash(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()
	
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}
	
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// reportProgress reports scanning progress
func (s *Scanner) reportProgress(wg *sync.WaitGroup, result *models.ScanResult) {
	defer wg.Done()
	
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			filesScanned := atomic.LoadInt64(&s.stats.FilesScanned)
			dirsScanned := atomic.LoadInt64(&s.stats.DirectoriesScanned)
			errors := atomic.LoadInt64(&s.stats.ErrorsEncountered)
			bytesProcessed := atomic.LoadInt64(&s.stats.BytesProcessed)
			
			fmt.Printf("\rScanning... Files: %d, Dirs: %d, Errors: %d, Size: %s", 
				filesScanned, dirsScanned, errors, formatBytes(bytesProcessed))
			
		case <-s.ctx.Done():
			return
		}
	}
}

// HashJob represents a job for calculating file hashes
type HashJob struct {
	FileInfo *models.FileInfo
	Options  *models.ScanOptions
}

// ID returns the job ID
func (hj *HashJob) ID() string {
	return fmt.Sprintf("hash_%s", hj.FileInfo.Path)
}

// Process processes the hash job
func (hj *HashJob) Process() (patterns.Result, error) {
	start := time.Now()
	
	// Calculate hashes
	hashes := hj.calculateHashes()
	
	return &HashResult{
		JobIDValue:    hj.ID(),
		FileInfo:      hj.FileInfo,
		Hashes:        hashes,
		DurationValue: time.Since(start),
		SuccessValue:  true,
	}, nil
}

// Timeout returns the job timeout
func (hj *HashJob) Timeout() time.Duration {
	return 30 * time.Second
}

// calculateHashes calculates MD5, SHA1, and SHA256 hashes
func (hj *HashJob) calculateHashes() map[string]string {
	hashes := make(map[string]string)
	
	file, err := os.Open(hj.FileInfo.Path)
	if err != nil {
		return hashes
	}
	defer file.Close()
	
	// Read file in chunks
	buffer := make([]byte, hj.Options.BufferSize)
	md5Hash := md5.New()
	sha1Hash := sha1.New()
	sha256Hash := sha256.New()
	
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return hashes
		}
		
		md5Hash.Write(buffer[:n])
		sha1Hash.Write(buffer[:n])
		sha256Hash.Write(buffer[:n])
	}
	
	hashes["md5"] = fmt.Sprintf("%x", md5Hash.Sum(nil))
	hashes["sha1"] = fmt.Sprintf("%x", sha1Hash.Sum(nil))
	hashes["sha256"] = fmt.Sprintf("%x", sha256Hash.Sum(nil))
	
	return hashes
}

// HashResult represents the result of a hash calculation
type HashResult struct {
	JobIDValue    string
	FileInfo      *models.FileInfo
	Hashes        map[string]string
	DurationValue time.Duration
	SuccessValue  bool
}

// JobID returns the job ID
func (hr *HashResult) JobID() string {
	return hr.JobIDValue
}

// Data returns the result data
func (hr *HashResult) Data() interface{} {
	return hr.Hashes
}

// Duration returns the processing duration
func (hr *HashResult) Duration() time.Duration {
	return hr.DurationValue
}

// Success returns true if the job was successful
func (hr *HashResult) Success() bool {
	return hr.SuccessValue
}

// formatBytes formats bytes into human-readable format
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
