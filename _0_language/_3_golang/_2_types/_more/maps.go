package main

import "fmt"

// MapManager handles various map operations and patterns
type MapManager struct {
	// Basic maps with different key-value types
	StringIntMap    map[string]int            `json:"string_int_map"`
	IntStringMap    map[int]string             `json:"int_string_map"`
	StringBoolMap   map[string]bool            `json:"string_bool_map"`
	FloatStringMap  map[float64]string         `json:"float_string_map"`
	
	// Maps with complex value types
	StringSliceMap  map[string][]string        `json:"string_slice_map"`
	StringStructMap map[string]Person          `json:"string_struct_map"`
	IntInterfaceMap map[int]interface{}        `json:"int_interface_map"`
	
	// Nested maps
	NestedMap       map[string]map[string]int  `json:"nested_map"`
	DeepNestedMap   map[string]map[string]map[string]interface{} `json:"deep_nested_map"`
	
	// Maps with custom types
	CustomKeyMap    map[CustomKey]string       `json:"custom_key_map"`
	TimeMap         map[time.Time]string       `json:"time_map"`
	
	// Maps for different use cases
	CacheMap        map[string]CacheEntry      `json:"cache_map"`
	ConfigMap       map[string]ConfigValue     `json:"config_map"`
	StatsMap        map[string]Statistics      `json:"stats_map"`
	
	// Concurrent maps (using channels for synchronization)
	ConcurrentMap   map[string]int             `json:"concurrent_map"`
	MapMutex        chan bool                  `json:"-"` // For synchronization
}

// Custom types for map keys and values
type CustomKey struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CacheEntry struct {
	Value     interface{} `json:"value"`
	ExpiresAt time.Time  `json:"expires_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type ConfigValue struct {
	Value     interface{} `json:"value"`
	Type      string      `json:"type"`
	Required  bool        `json:"required"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Statistics struct {
	Count     int       `json:"count"`
	Sum       float64   `json:"sum"`
	Average   float64   `json:"average"`
	Min       float64   `json:"min"`
	Max       float64   `json:"max"`
	LastUpdate time.Time `json:"last_update"`
}

// NewMapManager creates a new map manager
func NewMapManager() *MapManager {
	return &MapManager{
		StringIntMap:    make(map[string]int),
		IntStringMap:    make(map[int]string),
		StringBoolMap:   make(map[string]bool),
		FloatStringMap:  make(map[float64]string),
		StringSliceMap:  make(map[string][]string),
		StringStructMap: make(map[string]Person),
		IntInterfaceMap: make(map[int]interface{}),
		NestedMap:       make(map[string]map[string]int),
		DeepNestedMap:   make(map[string]map[string]map[string]interface{}),
		CustomKeyMap:    make(map[CustomKey]string),
		TimeMap:         make(map[time.Time]string),
		CacheMap:        make(map[string]CacheEntry),
		ConfigMap:       make(map[string]ConfigValue),
		StatsMap:        make(map[string]Statistics),
		ConcurrentMap:   make(map[string]int),
		MapMutex:        make(chan bool, 1), // Buffered channel for mutex
	}
}

// CRUD Operations for Maps

// Create - Initialize and populate maps
func (mm *MapManager) Create() {
	fmt.Println("🔧 Creating and populating maps...")
	
	// Basic maps
	mm.StringIntMap["apple"] = 5
	mm.StringIntMap["banana"] = 3
	mm.StringIntMap["cherry"] = 8
	mm.StringIntMap["date"] = 2
	
	mm.IntStringMap[1] = "one"
	mm.IntStringMap[2] = "two"
	mm.IntStringMap[3] = "three"
	mm.IntStringMap[4] = "four"
	
	mm.StringBoolMap["is_active"] = true
	mm.StringBoolMap["is_verified"] = false
	mm.StringBoolMap["is_premium"] = true
	
	mm.FloatStringMap[3.14] = "pi"
	mm.FloatStringMap[2.71] = "e"
	mm.FloatStringMap[1.41] = "sqrt(2)"
	
	// Maps with complex values
	mm.StringSliceMap["fruits"] = []string{"apple", "banana", "cherry"}
	mm.StringSliceMap["colors"] = []string{"red", "green", "blue"}
	mm.StringSliceMap["numbers"] = []string{"one", "two", "three"}
	
	mm.StringStructMap["user1"] = Person{
		ID:    1,
		Name:  "Alice Johnson",
		Email: "alice@example.com",
		Age:   30,
	}
	mm.StringStructMap["user2"] = Person{
		ID:    2,
		Name:  "Bob Smith",
		Email: "bob@example.com",
		Age:   25,
	}
	
	mm.IntInterfaceMap[1] = "string value"
	mm.IntInterfaceMap[2] = 42
	mm.IntInterfaceMap[3] = true
	mm.IntInterfaceMap[4] = []int{1, 2, 3}
	
	// Nested maps
	mm.NestedMap["user1"] = map[string]int{"age": 30, "score": 95}
	mm.NestedMap["user2"] = map[string]int{"age": 25, "score": 87}
	mm.NestedMap["user3"] = map[string]int{"age": 35, "score": 92}
	
	mm.DeepNestedMap["config"] = map[string]map[string]interface{}{
		"database": {
			"host":     "localhost",
			"port":     5432,
			"username": "admin",
		},
		"cache": {
			"enabled": true,
			"ttl":     3600,
		},
	}
	
	// Maps with custom types
	mm.CustomKeyMap[CustomKey{ID: 1, Name: "key1"}] = "value1"
	mm.CustomKeyMap[CustomKey{ID: 2, Name: "key2"}] = "value2"
	
	mm.TimeMap[time.Now()] = "current_time"
	mm.TimeMap[time.Now().Add(-24*time.Hour)] = "yesterday"
	
	// Cache map
	mm.CacheMap["user:1"] = CacheEntry{
		Value:     "Alice Johnson",
		ExpiresAt: time.Now().Add(1 * time.Hour),
		CreatedAt: time.Now(),
	}
	mm.CacheMap["user:2"] = CacheEntry{
		Value:     "Bob Smith",
		ExpiresAt: time.Now().Add(2 * time.Hour),
		CreatedAt: time.Now(),
	}
	
	// Config map
	mm.ConfigMap["app_name"] = ConfigValue{
		Value:     "Golang CRUD Mastery",
		Type:      "string",
		Required:  true,
		UpdatedAt: time.Now(),
	}
	mm.ConfigMap["max_users"] = ConfigValue{
		Value:     1000,
		Type:      "int",
		Required:  true,
		UpdatedAt: time.Now(),
	}
	mm.ConfigMap["debug_mode"] = ConfigValue{
		Value:     false,
		Type:      "bool",
		Required:  false,
		UpdatedAt: time.Now(),
	}
	
	// Statistics map
	mm.StatsMap["requests"] = Statistics{
		Count:     1000,
		Sum:       50000.0,
		Average:   50.0,
		Min:       10.0,
		Max:       200.0,
		LastUpdate: time.Now(),
	}
	mm.StatsMap["errors"] = Statistics{
		Count:     25,
		Sum:       25.0,
		Average:   1.0,
		Min:       0.0,
		Max:       5.0,
		LastUpdate: time.Now(),
	}
	
	// Concurrent map
	mm.ConcurrentMap["counter"] = 0
	mm.ConcurrentMap["total"] = 0
	
	fmt.Println("✅ Maps created and populated successfully")
}

// Read - Display map contents
func (mm *MapManager) Read() {
	fmt.Println("\n📖 READING MAP CONTENTS:")
	fmt.Println("========================")
	
	// Read basic maps
	fmt.Printf("String-Int Map (%d items):\n", len(mm.StringIntMap))
	for key, value := range mm.StringIntMap {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	fmt.Printf("\nInt-String Map (%d items):\n", len(mm.IntStringMap))
	for key, value := range mm.IntStringMap {
		fmt.Printf("  %d: %s\n", key, value)
	}
	
	fmt.Printf("\nString-Bool Map (%d items):\n", len(mm.StringBoolMap))
	for key, value := range mm.StringBoolMap {
		fmt.Printf("  %s: %t\n", key, value)
	}
	
	// Read complex maps
	fmt.Printf("\nString-Slice Map (%d items):\n", len(mm.StringSliceMap))
	for key, value := range mm.StringSliceMap {
		fmt.Printf("  %s: %v\n", key, value)
	}
	
	fmt.Printf("\nString-Struct Map (%d items):\n", len(mm.StringStructMap))
	for key, value := range mm.StringStructMap {
		fmt.Printf("  %s: %+v\n", key, value)
	}
	
	fmt.Printf("\nInt-Interface Map (%d items):\n", len(mm.IntInterfaceMap))
	for key, value := range mm.IntInterfaceMap {
		fmt.Printf("  %d: %v (%T)\n", key, value, value)
	}
	
	// Read nested maps
	fmt.Printf("\nNested Map (%d items):\n", len(mm.NestedMap))
	for key, value := range mm.NestedMap {
		fmt.Printf("  %s: %v\n", key, value)
	}
	
	fmt.Printf("\nDeep Nested Map (%d items):\n", len(mm.DeepNestedMap))
	for key, value := range mm.DeepNestedMap {
		fmt.Printf("  %s: %v\n", key, value)
	}
	
	// Read custom maps
	fmt.Printf("\nCustom Key Map (%d items):\n", len(mm.CustomKeyMap))
	for key, value := range mm.CustomKeyMap {
		fmt.Printf("  %+v: %s\n", key, value)
	}
	
	fmt.Printf("\nTime Map (%d items):\n", len(mm.TimeMap))
	for key, value := range mm.TimeMap {
		fmt.Printf("  %s: %s\n", key.Format(time.RFC3339), value)
	}
	
	// Read specialized maps
	fmt.Printf("\nCache Map (%d items):\n", len(mm.CacheMap))
	for key, value := range mm.CacheMap {
		fmt.Printf("  %s: %v (expires: %s)\n", key, value.Value, value.ExpiresAt.Format(time.RFC3339))
	}
	
	fmt.Printf("\nConfig Map (%d items):\n", len(mm.ConfigMap))
	for key, value := range mm.ConfigMap {
		fmt.Printf("  %s: %v (%s, required: %t)\n", key, value.Value, value.Type, value.Required)
	}
	
	fmt.Printf("\nStats Map (%d items):\n", len(mm.StatsMap))
	for key, value := range mm.StatsMap {
		fmt.Printf("  %s: count=%d, avg=%.2f, min=%.2f, max=%.2f\n", 
			key, value.Count, value.Average, value.Min, value.Max)
	}
}

// Update - Modify map contents
func (mm *MapManager) Update() {
	fmt.Println("\n🔄 UPDATING MAP CONTENTS:")
	fmt.Println("=========================")
	
	// Update basic maps
	mm.StringIntMap["apple"] = 10
	mm.StringIntMap["grape"] = 6 // Add new key
	fmt.Println("Updated String-Int Map")
	
	mm.IntStringMap[5] = "five" // Add new key
	mm.IntStringMap[1] = "ONE"  // Update existing
	fmt.Println("Updated Int-String Map")
	
	// Update complex maps
	mm.StringSliceMap["fruits"] = append(mm.StringSliceMap["fruits"], "grape")
	mm.StringSliceMap["colors"] = append(mm.StringSliceMap["colors"], "yellow")
	fmt.Println("Updated String-Slice Map")
	
	// Update struct in map
	if person, exists := mm.StringStructMap["user1"]; exists {
		person.Age = 31
		person.UpdatedAt = time.Now()
		mm.StringStructMap["user1"] = person
	}
	fmt.Println("Updated String-Struct Map")
	
	// Update nested maps
	if user1, exists := mm.NestedMap["user1"]; exists {
		user1["score"] = 98
		user1["level"] = 5
		mm.NestedMap["user1"] = user1
	}
	fmt.Println("Updated Nested Map")
	
	// Update deep nested map
	if config, exists := mm.DeepNestedMap["config"]; exists {
		if database, exists := config["database"]; exists {
			database["password"] = "secret"
			config["database"] = database
			mm.DeepNestedMap["config"] = config
		}
	}
	fmt.Println("Updated Deep Nested Map")
	
	// Update cache
	if entry, exists := mm.CacheMap["user:1"]; exists {
		entry.Value = "Alice Johnson-Updated"
		entry.ExpiresAt = time.Now().Add(2 * time.Hour)
		mm.CacheMap["user:1"] = entry
	}
	fmt.Println("Updated Cache Map")
	
	// Update config
	if config, exists := mm.ConfigMap["app_name"]; exists {
		config.Value = "Golang CRUD Mastery - Updated"
		config.UpdatedAt = time.Now()
		mm.ConfigMap["app_name"] = config
	}
	fmt.Println("Updated Config Map")
	
	// Update statistics
	if stats, exists := mm.StatsMap["requests"]; exists {
		stats.Count += 100
		stats.Sum += 5000.0
		stats.Average = stats.Sum / float64(stats.Count)
		stats.LastUpdate = time.Now()
		mm.StatsMap["requests"] = stats
	}
	fmt.Println("Updated Stats Map")
	
	// Update concurrent map
	mm.MapMutex <- true // Lock
	mm.ConcurrentMap["counter"]++
	mm.ConcurrentMap["total"] += 10
	<-mm.MapMutex // Unlock
	fmt.Println("Updated Concurrent Map")
	
	fmt.Println("✅ Map contents updated successfully")
}

// Delete - Remove map contents
func (mm *MapManager) Delete() {
	fmt.Println("\n🗑️  DELETING MAP CONTENTS:")
	fmt.Println("==========================")
	
	// Delete from basic maps
	delete(mm.StringIntMap, "cherry")
	delete(mm.StringIntMap, "date")
	fmt.Println("Deleted items from String-Int Map")
	
	delete(mm.IntStringMap, 4)
	fmt.Println("Deleted item from Int-String Map")
	
	// Delete from complex maps
	delete(mm.StringSliceMap, "numbers")
	fmt.Println("Deleted item from String-Slice Map")
	
	delete(mm.StringStructMap, "user2")
	fmt.Println("Deleted item from String-Struct Map")
	
	// Delete from nested maps
	delete(mm.NestedMap, "user3")
	fmt.Println("Deleted item from Nested Map")
	
	// Delete from deep nested map
	if config, exists := mm.DeepNestedMap["config"]; exists {
		delete(config, "cache")
		mm.DeepNestedMap["config"] = config
	}
	fmt.Println("Deleted item from Deep Nested Map")
	
	// Delete from custom maps
	delete(mm.CustomKeyMap, CustomKey{ID: 2, Name: "key2"})
	fmt.Println("Deleted item from Custom Key Map")
	
	// Delete from cache
	delete(mm.CacheMap, "user:2")
	fmt.Println("Deleted item from Cache Map")
	
	// Delete from config
	delete(mm.ConfigMap, "debug_mode")
	fmt.Println("Deleted item from Config Map")
	
	// Delete from stats
	delete(mm.StatsMap, "errors")
	fmt.Println("Deleted item from Stats Map")
	
	// Clear entire maps
	mm.FloatStringMap = make(map[float64]string)
	mm.StringBoolMap = make(map[string]bool)
	fmt.Println("Cleared Float-String and String-Bool Maps")
	
	fmt.Println("✅ Map contents deleted successfully")
}

// Advanced Map Operations

// DemonstrateMapIteration shows different iteration patterns
func (mm *MapManager) DemonstrateMapIteration() {
	fmt.Println("\n🔄 MAP ITERATION DEMONSTRATION:")
	fmt.Println("===============================")
	
	// Basic iteration
	fmt.Println("Basic iteration (order is random):")
	for key, value := range mm.StringIntMap {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	// Iteration with index
	fmt.Println("\nIteration with index:")
	index := 0
	for key, value := range mm.IntStringMap {
		fmt.Printf("  %d. %d: %s\n", index+1, key, value)
		index++
	}
	
	// Iteration with condition
	fmt.Println("\nIteration with condition (values > 5):")
	for key, value := range mm.StringIntMap {
		if value > 5 {
			fmt.Printf("  %s: %d\n", key, value)
		}
	}
	
	// Iteration over keys only
	fmt.Println("\nKeys only iteration:")
	keys := make([]string, 0, len(mm.StringIntMap))
	for key := range mm.StringIntMap {
		keys = append(keys, key)
	}
	for i, key := range keys {
		fmt.Printf("  %d. %s\n", i+1, key)
	}
	
	// Iteration over values only
	fmt.Println("\nValues only iteration:")
	values := make([]int, 0, len(mm.StringIntMap))
	for _, value := range mm.StringIntMap {
		values = append(values, value)
	}
	for i, value := range values {
		fmt.Printf("  %d. %d\n", i+1, value)
	}
}

// DemonstrateMapSorting shows how to sort map contents
func (mm *MapManager) DemonstrateMapSorting() {
	fmt.Println("\n📊 MAP SORTING DEMONSTRATION:")
	fmt.Println("=============================")
	
	// Sort by key
	fmt.Println("Sorting String-Int Map by key:")
	keys := make([]string, 0, len(mm.StringIntMap))
	for key := range mm.StringIntMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	
	for _, key := range keys {
		fmt.Printf("  %s: %d\n", key, mm.StringIntMap[key])
	}
	
	// Sort by value
	fmt.Println("\nSorting String-Int Map by value:")
	type kv struct {
		Key   string
		Value int
	}
	
	var kvs []kv
	for key, value := range mm.StringIntMap {
		kvs = append(kvs, kv{key, value})
	}
	
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value < kvs[j].Value
	})
	
	for _, kv := range kvs {
		fmt.Printf("  %s: %d\n", kv.Key, kv.Value)
	}
	
	// Sort by value (descending)
	fmt.Println("\nSorting by value (descending):")
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})
	
	for _, kv := range kvs {
		fmt.Printf("  %s: %d\n", kv.Key, kv.Value)
	}
}

// DemonstrateMapFiltering shows how to filter map contents
func (mm *MapManager) DemonstrateMapFiltering() {
	fmt.Println("\n🔍 MAP FILTERING DEMONSTRATION:")
	fmt.Println("===============================")
	
	// Filter by value
	fmt.Println("Filtering String-Int Map (values > 5):")
	filtered := make(map[string]int)
	for key, value := range mm.StringIntMap {
		if value > 5 {
			filtered[key] = value
		}
	}
	
	for key, value := range filtered {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	// Filter by key pattern
	fmt.Println("\nFiltering by key pattern (contains 'a'):")
	patternFiltered := make(map[string]int)
	for key, value := range mm.StringIntMap {
		if strings.Contains(key, "a") {
			patternFiltered[key] = value
		}
	}
	
	for key, value := range patternFiltered {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	// Filter by value type (for interface maps)
	fmt.Println("\nFiltering Int-Interface Map (string values only):")
	stringValues := make(map[int]string)
	for key, value := range mm.IntInterfaceMap {
		if str, ok := value.(string); ok {
			stringValues[key] = str
		}
	}
	
	for key, value := range stringValues {
		fmt.Printf("  %d: %s\n", key, value)
	}
}

// DemonstrateMapTransformation shows how to transform map contents
func (mm *MapManager) DemonstrateMapTransformation() {
	fmt.Println("\n🔄 MAP TRANSFORMATION DEMONSTRATION:")
	fmt.Println("===================================")
	
	// Transform values
	fmt.Println("Transforming String-Int Map (multiply by 2):")
	transformed := make(map[string]int)
	for key, value := range mm.StringIntMap {
		transformed[key] = value * 2
	}
	
	for key, value := range transformed {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	// Transform keys
	fmt.Println("\nTransforming keys (to uppercase):")
	keyTransformed := make(map[string]int)
	for key, value := range mm.StringIntMap {
		keyTransformed[strings.ToUpper(key)] = value
	}
	
	for key, value := range keyTransformed {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	// Transform to different type
	fmt.Println("\nTransforming to different type (int to string):")
	typeTransformed := make(map[string]string)
	for key, value := range mm.StringIntMap {
		typeTransformed[key] = fmt.Sprintf("count_%d", value)
	}
	
	for key, value := range typeTransformed {
		fmt.Printf("  %s: %s\n", key, value)
	}
}

// DemonstrateMapAggregation shows how to aggregate map data
func (mm *MapManager) DemonstrateMapAggregation() {
	fmt.Println("\n📈 MAP AGGREGATION DEMONSTRATION:")
	fmt.Println("=================================")
	
	// Sum values
	fmt.Println("Sum of String-Int Map values:")
	sum := 0
	for _, value := range mm.StringIntMap {
		sum += value
	}
	fmt.Printf("  Total: %d\n", sum)
	
	// Count items
	fmt.Println("\nCount of items in each map:")
	fmt.Printf("  String-Int Map: %d\n", len(mm.StringIntMap))
	fmt.Printf("  Int-String Map: %d\n", len(mm.IntStringMap))
	fmt.Printf("  String-Slice Map: %d\n", len(mm.StringSliceMap))
	
	// Find min/max values
	fmt.Println("\nMin/Max values in String-Int Map:")
	if len(mm.StringIntMap) > 0 {
		minKey, maxKey := "", ""
		minValue, maxValue := 0, 0
		first := true
		
		for key, value := range mm.StringIntMap {
			if first {
				minKey, maxKey = key, key
				minValue, maxValue = value, value
				first = false
			} else {
				if value < minValue {
					minKey, minValue = key, value
				}
				if value > maxValue {
					maxKey, maxValue = key, value
				}
			}
		}
		
		fmt.Printf("  Min: %s = %d\n", minKey, minValue)
		fmt.Printf("  Max: %s = %d\n", maxKey, maxValue)
	}
	
	// Average value
	fmt.Println("\nAverage value in String-Int Map:")
	if len(mm.StringIntMap) > 0 {
		sum := 0
		for _, value := range mm.StringIntMap {
			sum += value
		}
		avg := float64(sum) / float64(len(mm.StringIntMap))
		fmt.Printf("  Average: %.2f\n", avg)
	}
}

// DemonstrateMapMerging shows how to merge maps
func (mm *MapManager) DemonstrateMapMerging() {
	fmt.Println("\n🔗 MAP MERGING DEMONSTRATION:")
	fmt.Println("=============================")
	
	// Create another map to merge
	otherMap := map[string]int{
		"grape": 4,
		"kiwi":  2,
		"apple": 15, // This will overwrite the existing value
	}
	
	fmt.Println("Original map:")
	for key, value := range mm.StringIntMap {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	fmt.Println("\nMap to merge:")
	for key, value := range otherMap {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	// Merge maps
	merged := make(map[string]int)
	
	// Copy original map
	for key, value := range mm.StringIntMap {
		merged[key] = value
	}
	
	// Merge with other map
	for key, value := range otherMap {
		merged[key] = value
	}
	
	fmt.Println("\nMerged map:")
	for key, value := range merged {
		fmt.Printf("  %s: %d\n", key, value)
	}
	
	// Merge with conflict resolution (sum values)
	fmt.Println("\nMerged map with conflict resolution (sum values):")
	mergedSum := make(map[string]int)
	
	// Copy original map
	for key, value := range mm.StringIntMap {
		mergedSum[key] = value
	}
	
	// Merge with sum for conflicts
	for key, value := range otherMap {
		if existing, exists := mergedSum[key]; exists {
			mergedSum[key] = existing + value
		} else {
			mergedSum[key] = value
		}
	}
	
	for key, value := range mergedSum {
		fmt.Printf("  %s: %d\n", key, value)
	}
}

// DemonstrateMapConcurrency shows concurrent map operations
func (mm *MapManager) DemonstrateMapConcurrency() {
	fmt.Println("\n⚡ MAP CONCURRENCY DEMONSTRATION:")
	fmt.Println("=================================")
	
	// Simulate concurrent writes
	fmt.Println("Simulating concurrent writes...")
	
	// Write to concurrent map
	mm.MapMutex <- true // Lock
	mm.ConcurrentMap["counter"] = 0
	mm.ConcurrentMap["total"] = 0
	<-mm.MapMutex // Unlock
	
	// Simulate multiple goroutines writing
	for i := 0; i < 5; i++ {
		go func(id int) {
			mm.MapMutex <- true // Lock
			mm.ConcurrentMap["counter"]++
			mm.ConcurrentMap["total"] += id
			<-mm.MapMutex // Unlock
		}(i)
	}
	
	// Wait a bit for goroutines to complete
	time.Sleep(100 * time.Millisecond)
	
	// Read final values
	mm.MapMutex <- true // Lock
	counter := mm.ConcurrentMap["counter"]
	total := mm.ConcurrentMap["total"]
	<-mm.MapMutex // Unlock
	
	fmt.Printf("Final counter: %d\n", counter)
	fmt.Printf("Final total: %d\n", total)
}

// DemonstrateMapValidation shows map validation patterns
func (mm *MapManager) DemonstrateMapValidation() {
	fmt.Println("\n✅ MAP VALIDATION DEMONSTRATION:")
	fmt.Println("===============================")
	
	// Validate map contents
	fmt.Println("Validating String-Int Map:")
	for key, value := range mm.StringIntMap {
		valid := true
		issues := []string{}
		
		if len(key) == 0 {
			valid = false
			issues = append(issues, "empty key")
		}
		
		if value < 0 {
			valid = false
			issues = append(issues, "negative value")
		}
		
		if value > 100 {
			valid = false
			issues = append(issues, "value too large")
		}
		
		status := "✅ Valid"
		if !valid {
			status = "❌ Invalid: " + strings.Join(issues, ", ")
		}
		
		fmt.Printf("  %s: %d - %s\n", key, value, status)
	}
	
	// Validate cache entries
	fmt.Println("\nValidating Cache Map:")
	for key, entry := range mm.CacheMap {
		valid := true
		issues := []string{}
		
		if entry.Value == nil {
			valid = false
			issues = append(issues, "nil value")
		}
		
		if time.Now().After(entry.ExpiresAt) {
			valid = false
			issues = append(issues, "expired")
		}
		
		if entry.CreatedAt.After(time.Now()) {
			valid = false
			issues = append(issues, "invalid creation time")
		}
		
		status := "✅ Valid"
		if !valid {
			status = "❌ Invalid: " + strings.Join(issues, ", ")
		}
		
		fmt.Printf("  %s: %s\n", key, status)
	}
}
