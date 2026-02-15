package analytics

import (
	"file-scanner/pkg/models"
	"fmt"
	"math"
	"sort"
	"time"
)

// Analyzer represents a file system analyzer
type Analyzer struct {
	result *models.ScanResult
}

// NewAnalyzer creates a new analyzer
func NewAnalyzer(result *models.ScanResult) *Analyzer {
	return &Analyzer{
		result: result,
	}
}

// GenerateReport generates a comprehensive analysis report
func (a *Analyzer) GenerateReport() *Report {
	report := &Report{
		Summary:     a.generateSummary(),
		FileTypes:   a.analyzeFileTypes(),
		Extensions:  a.analyzeExtensions(),
		SizeAnalysis: a.analyzeSizes(),
		TimeAnalysis: a.analyzeTimes(),
		Permissions: a.analyzePermissions(),
		Duplicates:  a.analyzeDuplicates(),
		EmptyDirs:   a.analyzeEmptyDirectories(),
		LargestFiles: a.analyzeLargestFiles(),
		OldestFiles: a.analyzeOldestFiles(),
		NewestFiles: a.analyzeNewestFiles(),
		DepthAnalysis: a.analyzeDepth(),
		Recommendations: a.generateRecommendations(),
	}
	
	return report
}

// Report represents a comprehensive analysis report
type Report struct {
	Summary         *SummaryReport     `json:"summary"`
	FileTypes       *FileTypeReport    `json:"file_types"`
	Extensions      *ExtensionReport   `json:"extensions"`
	SizeAnalysis    *SizeReport        `json:"size_analysis"`
	TimeAnalysis    *TimeReport        `json:"time_analysis"`
	Permissions     *PermissionReport  `json:"permissions"`
	Duplicates      *DuplicateReport   `json:"duplicates"`
	EmptyDirs       *EmptyDirReport    `json:"empty_directories"`
	LargestFiles    *FileListReport    `json:"largest_files"`
	OldestFiles     *FileListReport    `json:"oldest_files"`
	NewestFiles     *FileListReport    `json:"newest_files"`
	DepthAnalysis   *DepthReport       `json:"depth_analysis"`
	Recommendations []string           `json:"recommendations"`
}

// SummaryReport represents a summary of the scan
type SummaryReport struct {
	TotalFiles      int64         `json:"total_files"`
	TotalDirs       int64         `json:"total_directories"`
	TotalSize       int64         `json:"total_size"`
	AverageFileSize int64         `json:"average_file_size"`
	ScanDuration    time.Duration `json:"scan_duration"`
	ScanRate        float64       `json:"scan_rate"`
	ErrorCount      int64         `json:"error_count"`
	RootPath        string        `json:"root_path"`
}

// FileTypeReport represents file type analysis
type FileTypeReport struct {
	Counts map[string]int64 `json:"counts"`
	Sizes  map[string]int64 `json:"sizes"`
	Top    []TypeStat       `json:"top_types"`
}

// TypeStat represents a file type statistic
type TypeStat struct {
	Type  string `json:"type"`
	Count int64  `json:"count"`
	Size  int64  `json:"size"`
	Icon  string `json:"icon"`
}

// ExtensionReport represents file extension analysis
type ExtensionReport struct {
	Counts map[string]int64 `json:"counts"`
	Sizes  map[string]int64 `json:"sizes"`
	Top    []ExtStat        `json:"top_extensions"`
}

// ExtStat represents an extension statistic
type ExtStat struct {
	Extension string `json:"extension"`
	Count     int64  `json:"count"`
	Size      int64  `json:"size"`
	MimeType  string `json:"mime_type"`
}

// SizeReport represents size analysis
type SizeReport struct {
	TotalSize       int64            `json:"total_size"`
	AverageSize     int64            `json:"average_size"`
	MedianSize      int64            `json:"median_size"`
	LargestFile     int64            `json:"largest_file"`
	SmallestFile    int64            `json:"smallest_file"`
	Distribution    map[string]int64 `json:"distribution"`
	SizeCategories  []SizeCategory   `json:"size_categories"`
}

// SizeCategory represents a size category
type SizeCategory struct {
	Range string `json:"range"`
	Count int64  `json:"count"`
	Size  int64  `json:"size"`
}

// TimeReport represents time analysis
type TimeReport struct {
	OldestFile    time.Time         `json:"oldest_file"`
	NewestFile    time.Time         `json:"newest_file"`
	Distribution  map[string]int64  `json:"distribution"`
	TimeCategories []TimeCategory   `json:"time_categories"`
}

// TimeCategory represents a time category
type TimeCategory struct {
	Period string `json:"period"`
	Count  int64  `json:"count"`
}

// PermissionReport represents permission analysis
type PermissionReport struct {
	Counts        map[string]int64 `json:"counts"`
	TopPermissions []PermStat      `json:"top_permissions"`
}

// PermStat represents a permission statistic
type PermStat struct {
	Permission string `json:"permission"`
	Count      int64  `json:"count"`
}

// DuplicateReport represents duplicate file analysis
type DuplicateReport struct {
	TotalDuplicates int64                  `json:"total_duplicates"`
	WastedSpace     int64                  `json:"wasted_space"`
	Groups          []*models.DuplicateGroup `json:"groups"`
}

// EmptyDirReport represents empty directory analysis
type EmptyDirReport struct {
	Count int64              `json:"count"`
	Dirs  []*models.FileInfo `json:"directories"`
}

// FileListReport represents a list of files
type FileListReport struct {
	Count int64              `json:"count"`
	Files []*models.FileInfo `json:"files"`
}

// DepthReport represents depth analysis
type DepthReport struct {
	MaxDepth     int            `json:"max_depth"`
	Distribution map[int]int64  `json:"distribution"`
	AverageDepth float64        `json:"average_depth"`
}

// generateSummary generates a summary report
func (a *Analyzer) generateSummary() *SummaryReport {
	return &SummaryReport{
		TotalFiles:      a.result.TotalFiles,
		TotalDirs:       a.result.TotalDirs,
		TotalSize:       a.result.TotalSize,
		AverageFileSize: a.result.GetAverageFileSize(),
		ScanDuration:    a.result.ScanDuration,
		ScanRate:        a.result.GetScanRate(),
		ErrorCount:      int64(len(a.result.Errors)),
		RootPath:        a.result.RootPath,
	}
}

// analyzeFileTypes analyzes file types
func (a *Analyzer) analyzeFileTypes() *FileTypeReport {
	report := &FileTypeReport{
		Counts: make(map[string]int64),
		Sizes:  make(map[string]int64),
		Top:    make([]TypeStat, 0),
	}
	
	// Count files by type
	for _, file := range a.result.Files {
		typeName := file.FileType.String()
		report.Counts[typeName]++
		report.Sizes[typeName] += file.Size
	}
	
	// Create top types
	for typeName, count := range report.Counts {
		report.Top = append(report.Top, TypeStat{
			Type:  typeName,
			Count: count,
			Size:  report.Sizes[typeName],
			Icon:  getFileTypeIcon(typeName),
		})
	}
	
	// Sort by count
	sort.Slice(report.Top, func(i, j int) bool {
		return report.Top[i].Count > report.Top[j].Count
	})
	
	return report
}

// analyzeExtensions analyzes file extensions
func (a *Analyzer) analyzeExtensions() *ExtensionReport {
	report := &ExtensionReport{
		Counts: make(map[string]int64),
		Sizes:  make(map[string]int64),
		Top:    make([]ExtStat, 0),
	}
	
	// Count files by extension
	for _, file := range a.result.Files {
		if file.IsDir {
			continue
		}
		
		ext := file.GetExtension()
		if ext == "" {
			ext = "no_extension"
		}
		
		report.Counts[ext]++
		report.Sizes[ext] += file.Size
	}
	
	// Create top extensions
	for ext, count := range report.Counts {
		report.Top = append(report.Top, ExtStat{
			Extension: ext,
			Count:     count,
			Size:      report.Sizes[ext],
			MimeType:  getMimeType(ext),
		})
	}
	
	// Sort by count
	sort.Slice(report.Top, func(i, j int) bool {
		return report.Top[i].Count > report.Top[j].Count
	})
	
	return report
}

// analyzeSizes analyzes file sizes
func (a *Analyzer) analyzeSizes() *SizeReport {
	report := &SizeReport{
		TotalSize:      a.result.TotalSize,
		AverageSize:    a.result.GetAverageFileSize(),
		Distribution:   make(map[string]int64),
		SizeCategories: make([]SizeCategory, 0),
	}
	
	if len(a.result.Files) == 0 {
		return report
	}
	
	// Collect file sizes
	sizes := make([]int64, 0, len(a.result.Files))
	for _, file := range a.result.Files {
		if !file.IsDir {
			sizes = append(sizes, file.Size)
		}
	}
	
	if len(sizes) == 0 {
		return report
	}
	
	// Sort sizes
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})
	
	// Calculate statistics
	report.LargestFile = sizes[len(sizes)-1]
	report.SmallestFile = sizes[0]
	report.MedianSize = sizes[len(sizes)/2]
	
	// Create size distribution
	categories := []struct {
		name  string
		min   int64
		max   int64
	}{
		{"0 bytes", 0, 0},
		{"1B-1KB", 1, 1024},
		{"1KB-1MB", 1024, 1024*1024},
		{"1MB-10MB", 1024*1024, 10*1024*1024},
		{"10MB-100MB", 10*1024*1024, 100*1024*1024},
		{"100MB-1GB", 100*1024*1024, 1024*1024*1024},
		{"1GB+", 1024*1024*1024, math.MaxInt64},
	}
	
	for _, category := range categories {
		count := int64(0)
		size := int64(0)
		
		for _, fileSize := range sizes {
			if fileSize >= category.min && fileSize <= category.max {
				count++
				size += fileSize
			}
		}
		
		if count > 0 {
			report.Distribution[category.name] = count
			report.SizeCategories = append(report.SizeCategories, SizeCategory{
				Range: category.name,
				Count: count,
				Size:  size,
			})
		}
	}
	
	return report
}

// analyzeTimes analyzes file times
func (a *Analyzer) analyzeTimes() *TimeReport {
	report := &TimeReport{
		Distribution:    make(map[string]int64),
		TimeCategories:  make([]TimeCategory, 0),
	}
	
	if len(a.result.Files) == 0 {
		return report
	}
	
	// Find oldest and newest files
	oldest := a.result.Files[0].ModTime
	newest := a.result.Files[0].ModTime
	
	for _, file := range a.result.Files {
		if file.ModTime.Before(oldest) {
			oldest = file.ModTime
		}
		if file.ModTime.After(newest) {
			newest = file.ModTime
		}
	}
	
	report.OldestFile = oldest
	report.NewestFile = newest
	
	// Create time distribution
	now := time.Now()
	categories := []struct {
		name  string
		check func(time.Time) bool
	}{
		{"Today", func(t time.Time) bool { return t.After(now.Truncate(24 * time.Hour)) }},
		{"This week", func(t time.Time) bool { return t.After(now.AddDate(0, 0, -7)) }},
		{"This month", func(t time.Time) bool { return t.After(now.AddDate(0, 0, -30)) }},
		{"This year", func(t time.Time) bool { return t.After(now.AddDate(-1, 0, 0)) }},
		{"Older", func(t time.Time) bool { return t.Before(now.AddDate(-1, 0, 0)) }},
	}
	
	for _, category := range categories {
		count := int64(0)
		
		for _, file := range a.result.Files {
			if category.check(file.ModTime) {
				count++
			}
		}
		
		if count > 0 {
			report.Distribution[category.name] = count
			report.TimeCategories = append(report.TimeCategories, TimeCategory{
				Period: category.name,
				Count:  count,
			})
		}
	}
	
	return report
}

// analyzePermissions analyzes file permissions
func (a *Analyzer) analyzePermissions() *PermissionReport {
	report := &PermissionReport{
		Counts:         make(map[string]int64),
		TopPermissions: make([]PermStat, 0),
	}
	
	// Count permissions
	for _, file := range a.result.Files {
		perm := file.Permissions
		report.Counts[perm]++
	}
	
	// Create top permissions
	for perm, count := range report.Counts {
		report.TopPermissions = append(report.TopPermissions, PermStat{
			Permission: perm,
			Count:      count,
		})
	}
	
	// Sort by count
	sort.Slice(report.TopPermissions, func(i, j int) bool {
		return report.TopPermissions[i].Count > report.TopPermissions[j].Count
	})
	
	return report
}

// analyzeDuplicates analyzes duplicate files
func (a *Analyzer) analyzeDuplicates() *DuplicateReport {
	report := &DuplicateReport{
		TotalDuplicates: 0,
		WastedSpace:     0,
		Groups:          a.result.Statistics.DuplicateFiles,
	}
	
	for _, group := range a.result.Statistics.DuplicateFiles {
		report.TotalDuplicates += int64(group.Count)
		report.WastedSpace += int64(group.Count-1) * group.Size
	}
	
	return report
}

// analyzeEmptyDirectories analyzes empty directories
func (a *Analyzer) analyzeEmptyDirectories() *EmptyDirReport {
	return &EmptyDirReport{
		Count: int64(len(a.result.Statistics.EmptyDirs)),
		Dirs:  a.result.Statistics.EmptyDirs,
	}
}

// analyzeLargestFiles analyzes largest files
func (a *Analyzer) analyzeLargestFiles() *FileListReport {
	return &FileListReport{
		Count: int64(len(a.result.Statistics.LargestFiles)),
		Files: a.result.Statistics.LargestFiles,
	}
}

// analyzeOldestFiles analyzes oldest files
func (a *Analyzer) analyzeOldestFiles() *FileListReport {
	return &FileListReport{
		Count: int64(len(a.result.Statistics.OldestFiles)),
		Files: a.result.Statistics.OldestFiles,
	}
}

// analyzeNewestFiles analyzes newest files
func (a *Analyzer) analyzeNewestFiles() *FileListReport {
	return &FileListReport{
		Count: int64(len(a.result.Statistics.NewestFiles)),
		Files: a.result.Statistics.NewestFiles,
	}
}

// analyzeDepth analyzes directory depth
func (a *Analyzer) analyzeDepth() *DepthReport {
	report := &DepthReport{
		MaxDepth:     0,
		Distribution: make(map[int]int64),
		AverageDepth: 0,
	}
	
	if len(a.result.Files) == 0 {
		return report
	}
	
	totalDepth := 0
	for _, file := range a.result.Files {
		depth := file.Depth
		report.Distribution[depth]++
		totalDepth += depth
		
		if depth > report.MaxDepth {
			report.MaxDepth = depth
		}
	}
	
	report.AverageDepth = float64(totalDepth) / float64(len(a.result.Files))
	
	return report
}

// generateRecommendations generates recommendations based on analysis
func (a *Analyzer) generateRecommendations() []string {
	recommendations := make([]string, 0)
	
	// Check for large files
	if len(a.result.Statistics.LargestFiles) > 0 {
		largest := a.result.Statistics.LargestFiles[0]
		if largest.Size > 1024*1024*1024 { // 1GB
			recommendations = append(recommendations, 
				fmt.Sprintf("Consider archiving large file: %s (%s)", 
					largest.Name, formatBytes(largest.Size)))
		}
	}
	
	// Check for duplicates
	if len(a.result.Statistics.DuplicateFiles) > 0 {
		totalWasted := int64(0)
		for _, group := range a.result.Statistics.DuplicateFiles {
			totalWasted += int64(group.Count-1) * group.Size
		}
		if totalWasted > 100*1024*1024 { // 100MB
			recommendations = append(recommendations, 
				fmt.Sprintf("Consider removing duplicate files to save %s", 
					formatBytes(totalWasted)))
		}
	}
	
	// Check for empty directories
	if len(a.result.Statistics.EmptyDirs) > 10 {
		recommendations = append(recommendations, 
			fmt.Sprintf("Consider cleaning up %d empty directories", 
				len(a.result.Statistics.EmptyDirs)))
	}
	
	// Check for old files
	oldFiles := 0
	cutoff := time.Now().AddDate(-2, 0, 0) // 2 years ago
	for _, file := range a.result.Files {
		if file.ModTime.Before(cutoff) {
			oldFiles++
		}
	}
	if oldFiles > 1000 {
		recommendations = append(recommendations, 
			fmt.Sprintf("Consider archiving %d old files (older than 2 years)", oldFiles))
	}
	
	return recommendations
}

// Helper functions

func getFileTypeIcon(fileType string) string {
	switch fileType {
	case "regular":
		return "📄"
	case "directory":
		return "📁"
	case "symlink":
		return "🔗"
	case "socket":
		return "🔌"
	case "named_pipe":
		return "🔧"
	case "device":
		return "💾"
	case "char_device":
		return "⌨️"
	default:
		return "❓"
	}
}

func getMimeType(ext string) string {
	mimeTypes := map[string]string{
		"txt":  "text/plain",
		"html": "text/html",
		"css":  "text/css",
		"js":   "application/javascript",
		"json": "application/json",
		"xml":  "application/xml",
		"pdf":  "application/pdf",
		"doc":  "application/msword",
		"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"xls":  "application/vnd.ms-excel",
		"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"gif":  "image/gif",
		"svg":  "image/svg+xml",
		"mp3":  "audio/mpeg",
		"mp4":  "video/mp4",
		"avi":  "video/x-msvideo",
		"zip":  "application/zip",
		"tar":  "application/x-tar",
		"gz":   "application/gzip",
		"bz2":  "application/x-bzip2",
		"7z":   "application/x-7z-compressed",
		"exe":  "application/x-msdownload",
		"dll":  "application/x-msdownload",
		"so":   "application/x-sharedlib",
		"dylib": "application/x-mach-binary",
	}
	
	if mimeType, exists := mimeTypes[ext]; exists {
		return mimeType
	}
	
	return "application/octet-stream"
}

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
