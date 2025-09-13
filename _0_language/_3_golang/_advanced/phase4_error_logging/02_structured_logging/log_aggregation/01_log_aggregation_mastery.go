// 📊 LOG AGGREGATION MASTERY
// Advanced log collection, processing, and aggregation systems
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================================
// LOG ENTRY STRUCTURES
// ============================================================================

type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Source    string                 `json:"source"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	ID        string                 `json:"id"`
}

type LogBatch struct {
	Entries []LogEntry `json:"entries"`
	Size    int        `json:"size"`
	Source  string     `json:"source"`
}

type LogQuery struct {
	Levels     []string               `json:"levels,omitempty"`
	Sources    []string               `json:"sources,omitempty"`
	StartTime  *time.Time             `json:"start_time,omitempty"`
	EndTime    *time.Time             `json:"end_time,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Fields     map[string]interface{} `json:"fields,omitempty"`
	Limit      int                    `json:"limit,omitempty"`
	Offset     int                    `json:"offset,omitempty"`
}

type LogStats struct {
	TotalEntries    int64            `json:"total_entries"`
	EntriesByLevel  map[string]int64 `json:"entries_by_level"`
	EntriesBySource map[string]int64 `json:"entries_by_source"`
	TimeRange       struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"time_range"`
}

// ============================================================================
// LOG COLLECTOR
// ============================================================================

type LogCollector struct {
	sources    map[string]chan LogEntry
	processors []LogProcessor
	mu         sync.RWMutex
	running    bool
	wg         sync.WaitGroup
}

type LogProcessor interface {
	Process(entry LogEntry) error
	GetName() string
}

func NewLogCollector() *LogCollector {
	return &LogCollector{
		sources:    make(map[string]chan LogEntry),
		processors: make([]LogProcessor, 0),
	}
}

func (lc *LogCollector) AddSource(name string, bufferSize int) chan LogEntry {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	
	ch := make(chan LogEntry, bufferSize)
	lc.sources[name] = ch
	return ch
}

func (lc *LogCollector) AddProcessor(processor LogProcessor) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.processors = append(lc.processors, processor)
}

func (lc *LogCollector) Start() {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	
	if lc.running {
		return
	}
	
	lc.running = true
	
	// Start processing for each source
	for sourceName, sourceChan := range lc.sources {
		lc.wg.Add(1)
		go lc.processSource(sourceName, sourceChan)
	}
}

func (lc *LogCollector) Stop() {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	
	if !lc.running {
		return
	}
	
	lc.running = false
	
	// Close all source channels
	for _, ch := range lc.sources {
		close(ch)
	}
	
	lc.wg.Wait()
}

func (lc *LogCollector) processSource(sourceName string, sourceChan chan LogEntry) {
	defer lc.wg.Done()
	
	for entry := range sourceChan {
		entry.Source = sourceName
		entry.ID = generateLogID()
		
		// Process through all processors
		for _, processor := range lc.processors {
			if err := processor.Process(entry); err != nil {
				log.Printf("Error processing log entry: %v", err)
			}
		}
	}
}

func generateLogID() string {
	return fmt.Sprintf("log-%d", time.Now().UnixNano())
}

// ============================================================================
// LOG PROCESSORS
// ============================================================================

type ConsoleProcessor struct {
	name string
}

func NewConsoleProcessor() *ConsoleProcessor {
	return &ConsoleProcessor{name: "console"}
}

func (cp *ConsoleProcessor) Process(entry LogEntry) error {
	jsonData, _ := json.Marshal(entry)
	fmt.Printf("[%s] %s\n", cp.name, string(jsonData))
	return nil
}

func (cp *ConsoleProcessor) GetName() string {
	return cp.name
}

type FileProcessor struct {
	name     string
	filename string
	file     *os.File
	mu       sync.Mutex
}

func NewFileProcessor(filename string) *FileProcessor {
	return &FileProcessor{
		name:     "file",
		filename: filename,
	}
}

func (fp *FileProcessor) Process(entry LogEntry) error {
	fp.mu.Lock()
	defer fp.mu.Unlock()
	
	if fp.file == nil {
		file, err := os.OpenFile(fp.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		fp.file = file
	}
	
	jsonData, _ := json.Marshal(entry)
	_, err := fp.file.Write(append(jsonData, '\n'))
	return err
}

func (fp *FileProcessor) GetName() string {
	return fp.name
}

func (fp *FileProcessor) Close() error {
	fp.mu.Lock()
	defer fp.mu.Unlock()
	
	if fp.file != nil {
		return fp.file.Close()
	}
	return nil
}

type FilterProcessor struct {
	name      string
	predicate func(LogEntry) bool
	next      LogProcessor
}

func NewFilterProcessor(predicate func(LogEntry) bool, next LogProcessor) *FilterProcessor {
	return &FilterProcessor{
		name:      "filter",
		predicate: predicate,
		next:      next,
	}
}

func (fp *FilterProcessor) Process(entry LogEntry) error {
	if fp.predicate(entry) {
		return fp.next.Process(entry)
	}
	return nil
}

func (fp *FilterProcessor) GetName() string {
	return fp.name
}

// ============================================================================
// LOG STORAGE
// ============================================================================

type InMemoryLogStorage struct {
	entries []LogEntry
	mu      sync.RWMutex
	stats   LogStats
}

func NewInMemoryLogStorage() *InMemoryLogStorage {
	return &InMemoryLogStorage{
		entries: make([]LogEntry, 0),
		stats: LogStats{
			EntriesByLevel:  make(map[string]int64),
			EntriesBySource: make(map[string]int64),
		},
	}
}

func (ls *InMemoryLogStorage) Store(entry LogEntry) error {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	
	ls.entries = append(ls.entries, entry)
	
	// Update stats
	atomic.AddInt64(&ls.stats.TotalEntries, 1)
	atomic.AddInt64(&ls.stats.EntriesByLevel[entry.Level], 1)
	atomic.AddInt64(&ls.stats.EntriesBySource[entry.Source], 1)
	
	// Update time range
	if ls.stats.TimeRange.Start.IsZero() || entry.Timestamp.Before(ls.stats.TimeRange.Start) {
		ls.stats.TimeRange.Start = entry.Timestamp
	}
	if ls.stats.TimeRange.End.IsZero() || entry.Timestamp.After(ls.stats.TimeRange.End) {
		ls.stats.TimeRange.End = entry.Timestamp
	}
	
	return nil
}

func (ls *InMemoryLogStorage) Query(query LogQuery) ([]LogEntry, error) {
	ls.mu.RLock()
	defer ls.mu.RUnlock()
	
	var results []LogEntry
	
	for _, entry := range ls.entries {
		if ls.matchesQuery(entry, query) {
			results = append(results, entry)
		}
	}
	
	// Sort by timestamp (newest first)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Timestamp.After(results[j].Timestamp)
	})
	
	// Apply limit and offset
	start := query.Offset
	if start >= len(results) {
		return []LogEntry{}, nil
	}
	
	end := start + query.Limit
	if end > len(results) {
		end = len(results)
	}
	if query.Limit == 0 {
		end = len(results)
	}
	
	return results[start:end], nil
}

func (ls *InMemoryLogStorage) matchesQuery(entry LogEntry, query LogQuery) bool {
	// Check level filter
	if len(query.Levels) > 0 {
		found := false
		for _, level := range query.Levels {
			if entry.Level == level {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	
	// Check source filter
	if len(query.Sources) > 0 {
		found := false
		for _, source := range query.Sources {
			if entry.Source == source {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	
	// Check time range
	if query.StartTime != nil && entry.Timestamp.Before(*query.StartTime) {
		return false
	}
	if query.EndTime != nil && entry.Timestamp.After(*query.EndTime) {
		return false
	}
	
	// Check message filter
	if query.Message != "" && !strings.Contains(entry.Message, query.Message) {
		return false
	}
	
	// Check field filters
	for key, value := range query.Fields {
		if entry.Fields[key] != value {
			return false
		}
	}
	
	return true
}

func (ls *InMemoryLogStorage) GetStats() LogStats {
	ls.mu.RLock()
	defer ls.mu.RUnlock()
	return ls.stats
}

// ============================================================================
// LOG ANALYZER
// ============================================================================

type LogAnalyzer struct {
	storage *InMemoryLogStorage
}

func NewLogAnalyzer(storage *InMemoryLogStorage) *LogAnalyzer {
	return &LogAnalyzer{storage: storage}
}

func (la *LogAnalyzer) AnalyzeLevels() map[string]int64 {
	stats := la.storage.GetStats()
	return stats.EntriesByLevel
}

func (la *LogAnalyzer) AnalyzeSources() map[string]int64 {
	stats := la.storage.GetStats()
	return stats.EntriesBySource
}

func (la *LogAnalyzer) AnalyzeTimeRange() (time.Time, time.Time) {
	stats := la.storage.GetStats()
	return stats.TimeRange.Start, stats.TimeRange.End
}

func (la *LogAnalyzer) FindPatterns(pattern string) ([]LogEntry, error) {
	query := LogQuery{
		Message: pattern,
	}
	return la.storage.Query(query)
}

func (la *LogAnalyzer) GetTopErrors(limit int) ([]LogEntry, error) {
	query := LogQuery{
		Levels: []string{"ERROR", "FATAL"},
		Limit:  limit,
	}
	return la.storage.Query(query)
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateLogCollection() {
	fmt.Println("\n=== Log Collection ===")
	
	collector := NewLogCollector()
	
	// Add console processor
	collector.AddProcessor(NewConsoleProcessor())
	
	// Add file processor
	fileProcessor := NewFileProcessor("logs.json")
	collector.AddProcessor(fileProcessor)
	defer fileProcessor.Close()
	
	// Start collector
	collector.Start()
	defer collector.Stop()
	
	// Add sources
	appSource := collector.AddSource("app", 100)
	webSource := collector.AddSource("web", 100)
	dbSource := collector.AddSource("database", 100)
	
	// Send some log entries
	appSource <- LogEntry{
		Timestamp: time.Now(),
		Level:     "INFO",
		Message:   "Application started",
		Fields:    map[string]interface{}{"port": 8080},
	}
	
	webSource <- LogEntry{
		Timestamp: time.Now(),
		Level:     "WARN",
		Message:   "High memory usage",
		Fields:    map[string]interface{}{"memory_usage": "85%"},
	}
	
	dbSource <- LogEntry{
		Timestamp: time.Now(),
		Level:     "ERROR",
		Message:   "Database connection failed",
		Fields:    map[string]interface{}{"error": "timeout"},
	}
	
	// Wait a bit for processing
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("   📊 Logs collected from multiple sources")
}

func demonstrateLogProcessing() {
	fmt.Println("\n=== Log Processing ===")
	
	collector := NewLogCollector()
	
	// Create storage
	storage := NewInMemoryLogStorage()
	
	// Add storage processor
	collector.AddProcessor(&StorageProcessor{storage: storage})
	
	// Add filter processor (only ERROR and WARN levels)
	errorFilter := NewFilterProcessor(
		func(entry LogEntry) bool {
			return entry.Level == "ERROR" || entry.Level == "WARN"
		},
		NewConsoleProcessor(),
	)
	collector.AddProcessor(errorFilter)
	
	// Start collector
	collector.Start()
	defer collector.Stop()
	
	// Add source
	source := collector.AddSource("filtered", 100)
	
	// Send various log levels
	levels := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	for _, level := range levels {
		source <- LogEntry{
			Timestamp: time.Now(),
			Level:     level,
			Message:   fmt.Sprintf("Test %s message", level),
		}
	}
	
	// Wait for processing
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("   📊 Logs processed and filtered")
}

type StorageProcessor struct {
	storage *InMemoryLogStorage
}

func (sp *StorageProcessor) Process(entry LogEntry) error {
	return sp.storage.Store(entry)
}

func (sp *StorageProcessor) GetName() string {
	return "storage"
}

func demonstrateLogQuerying() {
	fmt.Println("\n=== Log Querying ===")
	
	// Create storage and add some test data
	storage := NewInMemoryLogStorage()
	
	// Add test entries
	testEntries := []LogEntry{
		{
			Timestamp: time.Now().Add(-1 * time.Hour),
			Level:     "INFO",
			Message:   "User login",
			Source:    "auth",
			Fields:    map[string]interface{}{"user_id": "123"},
		},
		{
			Timestamp: time.Now().Add(-30 * time.Minute),
			Level:     "ERROR",
			Message:   "Database timeout",
			Source:    "database",
			Fields:    map[string]interface{}{"query": "SELECT * FROM users"},
		},
		{
			Timestamp: time.Now().Add(-15 * time.Minute),
			Level:     "WARN",
			Message:   "High CPU usage",
			Source:    "monitor",
			Fields:    map[string]interface{}{"cpu_usage": "90%"},
		},
	}
	
	for _, entry := range testEntries {
		storage.Store(entry)
	}
	
	// Query by level
	errorQuery := LogQuery{Levels: []string{"ERROR"}}
	errors, _ := storage.Query(errorQuery)
	fmt.Printf("   📊 Found %d ERROR logs\n", len(errors))
	
	// Query by source
	authQuery := LogQuery{Sources: []string{"auth"}}
	authLogs, _ := storage.Query(authQuery)
	fmt.Printf("   📊 Found %d auth logs\n", len(authLogs))
	
	// Query by time range
	now := time.Now()
	recentQuery := LogQuery{
		StartTime: &now.Add(-1 * time.Hour),
		EndTime:   &now,
	}
	recentLogs, _ := storage.Query(recentQuery)
	fmt.Printf("   📊 Found %d recent logs\n", len(recentLogs))
}

func demonstrateLogAnalytics() {
	fmt.Println("\n=== Log Analytics ===")
	
	// Create storage with test data
	storage := NewInMemoryLogStorage()
	analyzer := NewLogAnalyzer(storage)
	
	// Add test data
	levels := []string{"INFO", "WARN", "ERROR", "INFO", "DEBUG", "ERROR"}
	sources := []string{"app", "web", "database", "app", "web", "database"}
	
	for i, level := range levels {
		storage.Store(LogEntry{
			Timestamp: time.Now().Add(time.Duration(i) * time.Minute),
			Level:     level,
			Message:   fmt.Sprintf("Test %s message", level),
			Source:    sources[i],
		})
	}
	
	// Analyze levels
	levelStats := analyzer.AnalyzeLevels()
	fmt.Printf("   📊 Level distribution: %+v\n", levelStats)
	
	// Analyze sources
	sourceStats := analyzer.AnalyzeSources()
	fmt.Printf("   📊 Source distribution: %+v\n", sourceStats)
	
	// Get top errors
	topErrors, _ := analyzer.GetTopErrors(3)
	fmt.Printf("   📊 Top %d errors found\n", len(topErrors))
}

func demonstrateRealTimeMonitoring() {
	fmt.Println("\n=== Real-time Monitoring ===")
	
	collector := NewLogCollector()
	storage := NewInMemoryLogStorage()
	
	// Add storage processor
	collector.AddProcessor(&StorageProcessor{storage: storage})
	
	// Add alert processor
	alertProcessor := NewAlertProcessor()
	collector.AddProcessor(alertProcessor)
	
	// Start collector
	collector.Start()
	defer collector.Stop()
	
	// Add source
	source := collector.AddSource("monitor", 100)
	
	// Simulate real-time logs
	go func() {
		for i := 0; i < 10; i++ {
			level := "INFO"
			if i%3 == 0 {
				level = "ERROR"
			} else if i%5 == 0 {
				level = "WARN"
			}
			
			source <- LogEntry{
				Timestamp: time.Now(),
				Level:     level,
				Message:   fmt.Sprintf("Real-time log %d", i),
				Fields:    map[string]interface{}{"iteration": i},
			}
			
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// Wait for processing
	time.Sleep(1 * time.Second)
	
	fmt.Println("   📊 Real-time monitoring completed")
}

type AlertProcessor struct {
	alertCount int
}

func NewAlertProcessor() *AlertProcessor {
	return &AlertProcessor{}
}

func (ap *AlertProcessor) Process(entry LogEntry) error {
	if entry.Level == "ERROR" || entry.Level == "FATAL" {
		ap.alertCount++
		fmt.Printf("   🚨 ALERT #%d: %s - %s\n", ap.alertCount, entry.Level, entry.Message)
	}
	return nil
}

func (ap *AlertProcessor) GetName() string {
	return "alert"
}

func main() {
	fmt.Println("📊 LOG AGGREGATION MASTERY")
	fmt.Println("==========================")
	
	demonstrateLogCollection()
	demonstrateLogProcessing()
	demonstrateLogQuerying()
	demonstrateLogAnalytics()
	demonstrateRealTimeMonitoring()
	
	fmt.Println("\n🎉 LOG AGGREGATION MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Multi-source log collection")
	fmt.Println("✅ Log processing and filtering")
	fmt.Println("✅ Log storage and querying")
	fmt.Println("✅ Log analytics and statistics")
	fmt.Println("✅ Real-time monitoring and alerting")
	
	fmt.Println("\n🚀 You are now ready for Performance Impact Mastery!")
}
