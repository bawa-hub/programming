package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// =============================================================================
// BASIC ADAPTER PATTERN
// =============================================================================

// Target interface - what the client expects
type MediaPlayer interface {
	Play(audioType string, fileName string)
}

// Adaptee - existing class with incompatible interface
type AdvancedMediaPlayer interface {
	PlayVLC(fileName string)
	PlayMP4(fileName string)
	PlayAVI(fileName string)
}

// Concrete Adaptee implementations
type VLCPlayer struct{}

func (v *VLCPlayer) PlayVLC(fileName string) {
	fmt.Printf("Playing VLC file: %s\n", fileName)
}

func (v *VLCPlayer) PlayMP4(fileName string) {
	// VLC can't play MP4
}

func (v *VLCPlayer) PlayAVI(fileName string) {
	// VLC can't play AVI
}

type MP4Player struct{}

func (m *MP4Player) PlayVLC(fileName string) {
	// MP4 player can't play VLC
}

func (m *MP4Player) PlayMP4(fileName string) {
	fmt.Printf("Playing MP4 file: %s\n", fileName)
}

func (m *MP4Player) PlayAVI(fileName string) {
	// MP4 player can't play AVI
}

type AVIPlayer struct{}

func (a *AVIPlayer) PlayVLC(fileName string) {
	// AVI player can't play VLC
}

func (a *AVIPlayer) PlayMP4(fileName string) {
	// AVI player can't play MP4
}

func (a *AVIPlayer) PlayAVI(fileName string) {
	fmt.Printf("Playing AVI file: %s\n", fileName)
}

// Adapter - adapts the adaptee to the target interface
type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	switch audioType {
	case "vlc":
		return &MediaAdapter{advancedPlayer: &VLCPlayer{}}
	case "mp4":
		return &MediaAdapter{advancedPlayer: &MP4Player{}}
	case "avi":
		return &MediaAdapter{advancedPlayer: &AVIPlayer{}}
	default:
		return nil
	}
}

func (ma *MediaAdapter) Play(audioType string, fileName string) {
	switch audioType {
	case "vlc":
		ma.advancedPlayer.PlayVLC(fileName)
	case "mp4":
		ma.advancedPlayer.PlayMP4(fileName)
	case "avi":
		ma.advancedPlayer.PlayAVI(fileName)
	}
}

// Concrete Target implementation
type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

func (ap *AudioPlayer) Play(audioType string, fileName string) {
	// Built-in support for mp3
	if audioType == "mp3" {
		fmt.Printf("Playing MP3 file: %s\n", fileName)
	} else if audioType == "vlc" || audioType == "mp4" || audioType == "avi" {
		// Use adapter for other formats
		ap.mediaAdapter = NewMediaAdapter(audioType)
		ap.mediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Printf("Invalid media type: %s\n", audioType)
	}
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. PAYMENT GATEWAY ADAPTER
type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string) (string, error)
	RefundPayment(transactionID string) error
}

// Legacy payment system
type LegacyPaymentSystem struct {
	apiKey string
}

func (lps *LegacyPaymentSystem) Charge(amount int, currencyCode string) (string, error) {
	fmt.Printf("Legacy system charging %d %s\n", amount, currencyCode)
	return fmt.Sprintf("legacy_txn_%d", time.Now().Unix()), nil
}

func (lps *LegacyPaymentSystem) Reverse(transactionID string) error {
	fmt.Printf("Legacy system reversing transaction: %s\n", transactionID)
	return nil
}

// Adapter for legacy payment system
type LegacyPaymentAdapter struct {
	legacySystem *LegacyPaymentSystem
}

func NewLegacyPaymentAdapter(apiKey string) *LegacyPaymentAdapter {
	return &LegacyPaymentAdapter{
		legacySystem: &LegacyPaymentSystem{apiKey: apiKey},
	}
}

func (lpa *LegacyPaymentAdapter) ProcessPayment(amount float64, currency string) (string, error) {
	// Convert float to int (legacy system uses cents)
	amountInCents := int(amount * 100)
	
	// Convert currency to currency code
	currencyCode := strings.ToUpper(currency[:3])
	
	return lpa.legacySystem.Charge(amountInCents, currencyCode)
}

func (lpa *LegacyPaymentAdapter) RefundPayment(transactionID string) error {
	return lpa.legacySystem.Reverse(transactionID)
}

// Modern payment system
type ModernPaymentSystem struct {
	apiKey string
}

func (mps *ModernPaymentSystem) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("Modern system processing $%.2f %s\n", amount, currency)
	return fmt.Sprintf("modern_txn_%d", time.Now().Unix()), nil
}

func (mps *ModernPaymentSystem) RefundPayment(transactionID string) error {
	fmt.Printf("Modern system refunding transaction: %s\n", transactionID)
	return nil
}

// 2. DATABASE ADAPTER
type Database interface {
	Query(sql string) ([]map[string]interface{}, error)
	Execute(sql string) error
	Close() error
}

// Legacy database system
type LegacyDatabase struct {
	connectionString string
}

func (ldb *LegacyDatabase) ExecuteQuery(query string) ([]map[string]string, error) {
	fmt.Printf("Legacy DB executing: %s\n", query)
	return []map[string]string{
		{"id": "1", "name": "John"},
		{"id": "2", "name": "Jane"},
	}, nil
}

func (ldb *LegacyDatabase) ExecuteCommand(command string) error {
	fmt.Printf("Legacy DB executing command: %s\n", command)
	return nil
}

func (ldb *LegacyDatabase) Disconnect() error {
	fmt.Println("Legacy DB disconnecting")
	return nil
}

// Adapter for legacy database
type LegacyDatabaseAdapter struct {
	legacyDB *LegacyDatabase
}

func NewLegacyDatabaseAdapter(connectionString string) *LegacyDatabaseAdapter {
	return &LegacyDatabaseAdapter{
		legacyDB: &LegacyDatabase{connectionString: connectionString},
	}
}

func (ldba *LegacyDatabaseAdapter) Query(sql string) ([]map[string]interface{}, error) {
	// Convert legacy result to expected format
	legacyResult, err := ldba.legacyDB.ExecuteQuery(sql)
	if err != nil {
		return nil, err
	}
	
	// Convert []map[string]string to []map[string]interface{}
	result := make([]map[string]interface{}, len(legacyResult))
	for i, row := range legacyResult {
		result[i] = make(map[string]interface{})
		for k, v := range row {
			result[i][k] = v
		}
	}
	
	return result, nil
}

func (ldba *LegacyDatabaseAdapter) Execute(sql string) error {
	return ldba.legacyDB.ExecuteCommand(sql)
}

func (ldba *LegacyDatabaseAdapter) Close() error {
	return ldba.legacyDB.Disconnect()
}

// 3. API VERSION ADAPTER
type APIClient interface {
	GetUser(id string) (map[string]interface{}, error)
	CreateUser(user map[string]interface{}) (string, error)
	UpdateUser(id string, user map[string]interface{}) error
}

// Legacy API (v1)
type LegacyAPI struct {
	baseURL string
	apiKey  string
}

func (la *LegacyAPI) FetchUser(userID string) (map[string]string, error) {
	fmt.Printf("Legacy API fetching user: %s\n", userID)
	return map[string]string{
		"user_id":    userID,
		"user_name":  "John Doe",
		"user_email": "john@example.com",
		"created_at": "2023-01-01",
	}, nil
}

func (la *LegacyAPI) AddUser(userData map[string]string) (string, error) {
	fmt.Printf("Legacy API creating user: %v\n", userData)
	return "legacy_user_123", nil
}

func (la *LegacyAPI) ModifyUser(userID string, userData map[string]string) error {
	fmt.Printf("Legacy API updating user %s: %v\n", userID, userData)
	return nil
}

// Adapter for legacy API
type LegacyAPIAdapter struct {
	legacyAPI *LegacyAPI
}

func NewLegacyAPIAdapter(baseURL, apiKey string) *LegacyAPIAdapter {
	return &LegacyAPIAdapter{
		legacyAPI: &LegacyAPI{baseURL: baseURL, apiKey: apiKey},
	}
}

func (laa *LegacyAPIAdapter) GetUser(id string) (map[string]interface{}, error) {
	legacyUser, err := laa.legacyAPI.FetchUser(id)
	if err != nil {
		return nil, err
	}
	
	// Convert legacy format to modern format
	user := make(map[string]interface{})
	user["id"] = legacyUser["user_id"]
	user["name"] = legacyUser["user_name"]
	user["email"] = legacyUser["user_email"]
	user["createdAt"] = legacyUser["created_at"]
	
	return user, nil
}

func (laa *LegacyAPIAdapter) CreateUser(user map[string]interface{}) (string, error) {
	// Convert modern format to legacy format
	legacyUser := make(map[string]string)
	legacyUser["user_name"] = user["name"].(string)
	legacyUser["user_email"] = user["email"].(string)
	
	return laa.legacyAPI.AddUser(legacyUser)
}

func (laa *LegacyAPIAdapter) UpdateUser(id string, user map[string]interface{}) error {
	// Convert modern format to legacy format
	legacyUser := make(map[string]string)
	if name, ok := user["name"].(string); ok {
		legacyUser["user_name"] = name
	}
	if email, ok := user["email"].(string); ok {
		legacyUser["user_email"] = email
	}
	
	return laa.legacyAPI.ModifyUser(id, legacyUser)
}

// 4. FILE FORMAT ADAPTER
type FileReader interface {
	Read(filePath string) ([]map[string]interface{}, error)
	GetFormat() string
}

// CSV file reader
type CSVReader struct{}

func (cr *CSVReader) ReadCSV(filePath string) ([]map[string]string, error) {
	fmt.Printf("Reading CSV file: %s\n", filePath)
	// Simulate CSV reading
	return []map[string]string{
		{"name": "John", "age": "30", "city": "New York"},
		{"name": "Jane", "age": "25", "city": "Los Angeles"},
	}, nil
}

func (cr *CSVReader) GetFormat() string {
	return "CSV"
}

// Adapter for CSV reader
type CSVReaderAdapter struct {
	csvReader *CSVReader
}

func NewCSVReaderAdapter() *CSVReaderAdapter {
	return &CSVReaderAdapter{
		csvReader: &CSVReader{},
	}
}

func (cra *CSVReaderAdapter) Read(filePath string) ([]map[string]interface{}, error) {
	csvData, err := cra.csvReader.ReadCSV(filePath)
	if err != nil {
		return nil, err
	}
	
	// Convert CSV data to expected format
	result := make([]map[string]interface{}, len(csvData))
	for i, row := range csvData {
		result[i] = make(map[string]interface{})
		for k, v := range row {
			// Try to convert to appropriate type
			if intVal, err := strconv.Atoi(v); err == nil {
				result[i][k] = intVal
			} else {
				result[i][k] = v
			}
		}
	}
	
	return result, nil
}

func (cra *CSVReaderAdapter) GetFormat() string {
	return cra.csvReader.GetFormat()
}

// =============================================================================
// TWO-WAY ADAPTER EXAMPLE
// =============================================================================

// Two-way adapter for bidirectional communication
type TwoWayAdapter struct {
	legacySystem *LegacyPaymentSystem
	modernSystem *ModernPaymentSystem
}

func NewTwoWayAdapter(legacyAPIKey, modernAPIKey string) *TwoWayAdapter {
	return &TwoWayAdapter{
		legacySystem: &LegacyPaymentSystem{apiKey: legacyAPIKey},
		modernSystem: &ModernPaymentSystem{apiKey: modernAPIKey},
	}
}

// Legacy interface methods
func (twa *TwoWayAdapter) Charge(amount int, currencyCode string) (string, error) {
	// Convert to modern format and use modern system
	amountFloat := float64(amount) / 100.0
	currency := currencyCode
	return twa.modernSystem.ProcessPayment(amountFloat, currency)
}

func (twa *TwoWayAdapter) Reverse(transactionID string) error {
	return twa.modernSystem.RefundPayment(transactionID)
}

// Modern interface methods
func (twa *TwoWayAdapter) ProcessPayment(amount float64, currency string) (string, error) {
	// Convert to legacy format and use legacy system
	amountInCents := int(amount * 100)
	currencyCode := strings.ToUpper(currency[:3])
	return twa.legacySystem.Charge(amountInCents, currencyCode)
}

func (twa *TwoWayAdapter) RefundPayment(transactionID string) error {
	return twa.legacySystem.Reverse(transactionID)
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== ADAPTER PATTERN DEMONSTRATION ===\n")

	// 1. BASIC ADAPTER
	fmt.Println("1. BASIC ADAPTER:")
	audioPlayer := &AudioPlayer{}
	
	audioPlayer.Play("mp3", "song.mp3")
	audioPlayer.Play("vlc", "movie.vlc")
	audioPlayer.Play("mp4", "video.mp4")
	audioPlayer.Play("avi", "movie.avi")
	audioPlayer.Play("wav", "sound.wav")
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Payment Gateway Adapter
	fmt.Println("Payment Gateway Adapter:")
	legacyPaymentAdapter := NewLegacyPaymentAdapter("legacy_api_key")
	modernPaymentSystem := &ModernPaymentSystem{apiKey: "modern_api_key"}
	
	// Use legacy system through adapter
	transactionID1, _ := legacyPaymentAdapter.ProcessPayment(100.50, "USD")
	fmt.Printf("Legacy transaction ID: %s\n", transactionID1)
	legacyPaymentAdapter.RefundPayment(transactionID1)
	
	// Use modern system directly
	transactionID2, _ := modernPaymentSystem.ProcessPayment(200.75, "EUR")
	fmt.Printf("Modern transaction ID: %s\n", transactionID2)
	modernPaymentSystem.RefundPayment(transactionID2)
	fmt.Println()

	// Database Adapter
	fmt.Println("Database Adapter:")
	legacyDBAdapter := NewLegacyDatabaseAdapter("legacy_db_connection")
	
	// Use legacy database through adapter
	users, _ := legacyDBAdapter.Query("SELECT * FROM users")
	fmt.Printf("Users from legacy DB: %v\n", users)
	
	legacyDBAdapter.Execute("INSERT INTO users (name) VALUES ('New User')")
	legacyDBAdapter.Close()
	fmt.Println()

	// API Version Adapter
	fmt.Println("API Version Adapter:")
	legacyAPIAdapter := NewLegacyAPIAdapter("https://api.legacy.com", "legacy_key")
	
	// Use legacy API through adapter
	user, _ := legacyAPIAdapter.GetUser("123")
	fmt.Printf("User from legacy API: %v\n", user)
	
	newUserID, _ := legacyAPIAdapter.CreateUser(map[string]interface{}{
		"name":  "Alice",
		"email": "alice@example.com",
	})
	fmt.Printf("Created user with ID: %s\n", newUserID)
	
	legacyAPIAdapter.UpdateUser("123", map[string]interface{}{
		"name": "John Updated",
	})
	fmt.Println()

	// File Format Adapter
	fmt.Println("File Format Adapter:")
	csvAdapter := NewCSVReaderAdapter()
	
	data, _ := csvAdapter.Read("users.csv")
	fmt.Printf("CSV data: %v\n", data)
	fmt.Printf("File format: %s\n", csvAdapter.GetFormat())
	fmt.Println()

	// 3. TWO-WAY ADAPTER
	fmt.Println("3. TWO-WAY ADAPTER:")
	twoWayAdapter := NewTwoWayAdapter("legacy_key", "modern_key")
	
	// Use as legacy system
	legacyTxnID, _ := twoWayAdapter.Charge(15000, "USD") // $150.00
	fmt.Printf("Legacy transaction: %s\n", legacyTxnID)
	twoWayAdapter.Reverse(legacyTxnID)
	
	// Use as modern system
	modernTxnID, _ := twoWayAdapter.ProcessPayment(250.00, "EUR")
	fmt.Printf("Modern transaction: %s\n", modernTxnID)
	twoWayAdapter.RefundPayment(modernTxnID)
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
