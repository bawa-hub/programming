package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileType represents the type of file system entry
type FileType int

const (
	FileTypeUnknown FileType = iota
	FileTypeRegular
	FileTypeDirectory
	FileTypeSymlink
	FileTypeSocket
	FileTypeNamedPipe
	FileTypeDevice
	FileTypeCharDevice
)

// FileInfo represents detailed information about a file system entry
type FileInfo struct {
	Path         string    `json:"path"`
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	Mode         os.FileMode `json:"mode"`
	ModTime      time.Time `json:"mod_time"`
	IsDir        bool      `json:"is_dir"`
	IsSymlink    bool      `json:"is_symlink"`
	FileType     FileType  `json:"file_type"`
	Extension    string    `json:"extension,omitempty"`
	Permissions  string    `json:"permissions"`
	Owner        string    `json:"owner,omitempty"`
	Group        string    `json:"group,omitempty"`
	Inode        uint64    `json:"inode,omitempty"`
	HardLinks    uint64    `json:"hard_links,omitempty"`
	Device       uint64    `json:"device,omitempty"`
	Blocks       int64     `json:"blocks,omitempty"`
	BlockSize    int64     `json:"block_size,omitempty"`
	AccessTime   time.Time `json:"access_time,omitempty"`
	ChangeTime   time.Time `json:"change_time,omitempty"`
	BirthTime    time.Time `json:"birth_time,omitempty"`
	ScanTime     time.Time `json:"scan_time"`
	Depth        int       `json:"depth"`
	ParentPath   string    `json:"parent_path,omitempty"`
	RelativePath string    `json:"relative_path,omitempty"`
}

// ScanResult represents the result of a file system scan
type ScanResult struct {
	RootPath     string     `json:"root_path"`
	TotalFiles   int64      `json:"total_files"`
	TotalDirs    int64      `json:"total_dirs"`
	TotalSize    int64      `json:"total_size"`
	ScanDuration time.Duration `json:"scan_duration"`
	StartTime    time.Time  `json:"start_time"`
	EndTime      time.Time  `json:"end_time"`
	Files        []*FileInfo `json:"files,omitempty"`
	Errors       []ScanError `json:"errors,omitempty"`
	Statistics   *Statistics `json:"statistics,omitempty"`
}

// ScanError represents an error encountered during scanning
type ScanError struct {
	Path      string    `json:"path"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

// Statistics represents file system statistics
type Statistics struct {
	FileCountByType    map[FileType]int64            `json:"file_count_by_type"`
	FileCountByExt     map[string]int64              `json:"file_count_by_extension"`
	SizeByType         map[FileType]int64            `json:"size_by_type"`
	SizeByExt          map[string]int64              `json:"size_by_extension"`
	LargestFiles       []*FileInfo                   `json:"largest_files"`
	OldestFiles        []*FileInfo                   `json:"oldest_files"`
	NewestFiles        []*FileInfo                   `json:"newest_files"`
	DuplicateFiles     []*DuplicateGroup             `json:"duplicate_files,omitempty"`
	EmptyDirs          []*FileInfo                   `json:"empty_directories,omitempty"`
	SymlinkTargets     map[string][]string           `json:"symlink_targets,omitempty"`
	PermissionStats    map[string]int64              `json:"permission_stats,omitempty"`
	DepthDistribution  map[int]int64                 `json:"depth_distribution,omitempty"`
	SizeDistribution   map[string]int64              `json:"size_distribution,omitempty"`
	TimeDistribution   map[string]int64              `json:"time_distribution,omitempty"`
}

// DuplicateGroup represents a group of duplicate files
type DuplicateGroup struct {
	Size     int64       `json:"size"`
	Hash     string      `json:"hash"`
	Files    []*FileInfo `json:"files"`
	Count    int         `json:"count"`
}

// Filter represents a file filter
type Filter struct {
	Name        string      `json:"name"`
	Pattern     string      `json:"pattern"`
	FileTypes   []FileType  `json:"file_types,omitempty"`
	Extensions  []string    `json:"extensions,omitempty"`
	MinSize     int64       `json:"min_size,omitempty"`
	MaxSize     int64       `json:"max_size,omitempty"`
	MinAge      time.Duration `json:"min_age,omitempty"`
	MaxAge      time.Duration `json:"max_age,omitempty"`
	Permissions os.FileMode `json:"permissions,omitempty"`
	Depth       int         `json:"depth,omitempty"`
	Regex       string      `json:"regex,omitempty"`
	CaseSensitive bool      `json:"case_sensitive"`
	Negate      bool        `json:"negate"`
}

// ScanOptions represents options for file system scanning
type ScanOptions struct {
	MaxDepth        int           `json:"max_depth"`
	FollowSymlinks  bool          `json:"follow_symlinks"`
	IncludeHidden   bool          `json:"include_hidden"`
	IncludeSystem   bool          `json:"include_system"`
	CalculateHashes bool          `json:"calculate_hashes"`
	FindDuplicates  bool          `json:"find_duplicates"`
	Concurrency     int           `json:"concurrency"`
	BufferSize      int           `json:"buffer_size"`
	Timeout         time.Duration `json:"timeout"`
	Filters         []*Filter     `json:"filters,omitempty"`
	ExcludePaths    []string      `json:"exclude_paths,omitempty"`
	IncludePaths    []string      `json:"include_paths,omitempty"`
	SortBy          string        `json:"sort_by"`
	SortOrder       string        `json:"sort_order"`
	Limit           int64         `json:"limit"`
	Progress        bool          `json:"progress"`
	Verbose         bool          `json:"verbose"`
}

// String methods for FileType
func (ft FileType) String() string {
	switch ft {
	case FileTypeRegular:
		return "regular"
	case FileTypeDirectory:
		return "directory"
	case FileTypeSymlink:
		return "symlink"
	case FileTypeSocket:
		return "socket"
	case FileTypeNamedPipe:
		return "named_pipe"
	case FileTypeDevice:
		return "device"
	case FileTypeCharDevice:
		return "char_device"
	default:
		return "unknown"
	}
}

// Icon returns an icon for the file type
func (ft FileType) Icon() string {
	switch ft {
	case FileTypeRegular:
		return "📄"
	case FileTypeDirectory:
		return "📁"
	case FileTypeSymlink:
		return "🔗"
	case FileTypeSocket:
		return "🔌"
	case FileTypeNamedPipe:
		return "🔧"
	case FileTypeDevice:
		return "💾"
	case FileTypeCharDevice:
		return "⌨️"
	default:
		return "❓"
	}
}

// String returns a string representation of FileInfo
func (fi *FileInfo) String() string {
	return fmt.Sprintf("%s %s %s (%d bytes)", 
		fi.FileType.Icon(), 
		fi.Name, 
		fi.Permissions, 
		fi.Size)
}

// FullString returns a detailed string representation
func (fi *FileInfo) FullString() string {
	return fmt.Sprintf("%s %s %s %s %d bytes %s %s", 
		fi.FileType.Icon(),
		fi.Permissions,
		fi.Owner,
		fi.Group,
		fi.Size,
		fi.ModTime.Format("2006-01-02 15:04:05"),
		fi.Path)
}

// IsEmpty returns true if the file is empty
func (fi *FileInfo) IsEmpty() bool {
	return fi.Size == 0
}

// IsExecutable returns true if the file is executable
func (fi *FileInfo) IsExecutable() bool {
	return fi.Mode&0111 != 0
}

// IsReadable returns true if the file is readable
func (fi *FileInfo) IsReadable() bool {
	return fi.Mode&0444 != 0
}

// IsWritable returns true if the file is writable
func (fi *FileInfo) IsWritable() bool {
	return fi.Mode&0222 != 0
}

// GetRelativePath returns the relative path from a base path
func (fi *FileInfo) GetRelativePath(basePath string) string {
	rel, err := filepath.Rel(basePath, fi.Path)
	if err != nil {
		return fi.Path
	}
	return rel
}

// GetExtension returns the file extension
func (fi *FileInfo) GetExtension() string {
	if fi.Extension != "" {
		return fi.Extension
	}
	ext := filepath.Ext(fi.Name)
	if ext != "" {
		return strings.ToLower(ext[1:]) // Remove the dot
	}
	return ""
}

// GetMimeType returns a basic MIME type based on extension
func (fi *FileInfo) GetMimeType() string {
	ext := fi.GetExtension()
	
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
		"ppt":  "application/vnd.ms-powerpoint",
		"pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
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

// GetSizeString returns a human-readable size string
func (fi *FileInfo) GetSizeString() string {
	return formatBytes(fi.Size)
}

// GetAgeString returns a human-readable age string
func (fi *FileInfo) GetAgeString() string {
	return formatDuration(time.Since(fi.ModTime))
}

// MatchesFilter checks if the file matches a filter
func (fi *FileInfo) MatchesFilter(filter *Filter) bool {
	// Check file type
	if len(filter.FileTypes) > 0 {
		found := false
		for _, ft := range filter.FileTypes {
			if fi.FileType == ft {
				found = true
				break
			}
		}
		if !found {
			return filter.Negate
		}
	}
	
	// Check extensions
	if len(filter.Extensions) > 0 {
		found := false
		fileExt := fi.GetExtension()
		for _, ext := range filter.Extensions {
			if strings.EqualFold(fileExt, ext) {
				found = true
				break
			}
		}
		if !found {
			return filter.Negate
		}
	}
	
	// Check size
	if filter.MinSize > 0 && fi.Size < filter.MinSize {
		return filter.Negate
	}
	if filter.MaxSize > 0 && fi.Size > filter.MaxSize {
		return filter.Negate
	}
	
	// Check age
	if filter.MinAge > 0 {
		age := time.Since(fi.ModTime)
		if age < filter.MinAge {
			return filter.Negate
		}
	}
	if filter.MaxAge > 0 {
		age := time.Since(fi.ModTime)
		if age > filter.MaxAge {
			return filter.Negate
		}
	}
	
	// Check permissions
	if filter.Permissions != 0 {
		if fi.Mode&filter.Permissions != filter.Permissions {
			return filter.Negate
		}
	}
	
	// Check depth
	if filter.Depth > 0 && fi.Depth > filter.Depth {
		return filter.Negate
	}
	
	// Check name pattern
	if filter.Pattern != "" {
		matched := false
		if filter.CaseSensitive {
			matched = strings.Contains(fi.Name, filter.Pattern)
		} else {
			matched = strings.Contains(strings.ToLower(fi.Name), strings.ToLower(filter.Pattern))
		}
		if !matched {
			return filter.Negate
		}
	}
	
	return !filter.Negate
}

// NewFileInfo creates a new FileInfo from os.FileInfo
func NewFileInfo(path string, info os.FileInfo, basePath string) *FileInfo {
	fi := &FileInfo{
		Path:        path,
		Name:        info.Name(),
		Size:        info.Size(),
		Mode:        info.Mode(),
		ModTime:     info.ModTime(),
		IsDir:       info.IsDir(),
		IsSymlink:   info.Mode()&os.ModeSymlink != 0,
		ScanTime:    time.Now(),
		ParentPath:  filepath.Dir(path),
	}
	
	// Set file type
	if fi.IsDir {
		fi.FileType = FileTypeDirectory
	} else if fi.IsSymlink {
		fi.FileType = FileTypeSymlink
	} else if info.Mode()&os.ModeSocket != 0 {
		fi.FileType = FileTypeSocket
	} else if info.Mode()&os.ModeNamedPipe != 0 {
		fi.FileType = FileTypeNamedPipe
	} else if info.Mode()&os.ModeDevice != 0 {
		if info.Mode()&os.ModeCharDevice != 0 {
			fi.FileType = FileTypeCharDevice
		} else {
			fi.FileType = FileTypeDevice
		}
	} else {
		fi.FileType = FileTypeRegular
	}
	
	// Set extension
	fi.Extension = fi.GetExtension()
	
	// Set permissions string
	fi.Permissions = info.Mode().String()
	
	// Set relative path
	if basePath != "" {
		fi.RelativePath = fi.GetRelativePath(basePath)
		fi.Depth = strings.Count(fi.RelativePath, string(filepath.Separator))
	}
	
	return fi
}

// NewScanResult creates a new ScanResult
func NewScanResult(rootPath string) *ScanResult {
	return &ScanResult{
		RootPath:  rootPath,
		StartTime: time.Now(),
		Files:     make([]*FileInfo, 0),
		Errors:    make([]ScanError, 0),
		Statistics: &Statistics{
			FileCountByType:   make(map[FileType]int64),
			FileCountByExt:    make(map[string]int64),
			SizeByType:        make(map[FileType]int64),
			SizeByExt:         make(map[string]int64),
			LargestFiles:      make([]*FileInfo, 0),
			OldestFiles:       make([]*FileInfo, 0),
			NewestFiles:       make([]*FileInfo, 0),
			DuplicateFiles:    make([]*DuplicateGroup, 0),
			EmptyDirs:         make([]*FileInfo, 0),
			SymlinkTargets:    make(map[string][]string),
			PermissionStats:   make(map[string]int64),
			DepthDistribution: make(map[int]int64),
			SizeDistribution:  make(map[string]int64),
			TimeDistribution:  make(map[string]int64),
		},
	}
}

// AddFile adds a file to the scan result
func (sr *ScanResult) AddFile(file *FileInfo) {
	sr.Files = append(sr.Files, file)
	
	if file.IsDir {
		sr.TotalDirs++
	} else {
		sr.TotalFiles++
	}
	
	sr.TotalSize += file.Size
}

// AddError adds an error to the scan result
func (sr *ScanResult) AddError(path string, err error, errorType string) {
	sr.Errors = append(sr.Errors, ScanError{
		Path:      path,
		Error:     err.Error(),
		Timestamp: time.Now(),
		Type:      errorType,
	})
}

// Finish marks the scan as complete
func (sr *ScanResult) Finish() {
	sr.EndTime = time.Now()
	sr.ScanDuration = sr.EndTime.Sub(sr.StartTime)
}

// GetFileCount returns the total number of files
func (sr *ScanResult) GetFileCount() int64 {
	return sr.TotalFiles + sr.TotalDirs
}

// GetAverageFileSize returns the average file size
func (sr *ScanResult) GetAverageFileSize() int64 {
	if sr.TotalFiles == 0 {
		return 0
	}
	return sr.TotalSize / sr.TotalFiles
}

// GetScanRate returns the scan rate (files per second)
func (sr *ScanResult) GetScanRate() float64 {
	if sr.ScanDuration == 0 {
		return 0
	}
	return float64(sr.GetFileCount()) / sr.ScanDuration.Seconds()
}

// Utility functions

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

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%.0fs", d.Seconds())
	} else if d < time.Hour {
		return fmt.Sprintf("%.0fm", d.Minutes())
	} else if d < 24*time.Hour {
		return fmt.Sprintf("%.1fh", d.Hours())
	} else {
		days := d.Hours() / 24
		return fmt.Sprintf("%.1fd", days)
	}
}
