package main

import (
	"fmt"
	"strings"
	"time"
)

// =============================================================================
// BASIC DECORATOR PATTERN
// =============================================================================

// Component interface - defines the interface for objects that can have responsibilities added dynamically
type Coffee interface {
	GetDescription() string
	GetCost() float64
}

// Concrete Component - Basic coffee
type BasicCoffee struct{}

func (bc *BasicCoffee) GetDescription() string {
	return "Basic Coffee"
}

func (bc *BasicCoffee) GetCost() float64 {
	return 2.0
}

// Decorator - Base decorator that wraps a coffee
type CoffeeDecorator struct {
	coffee Coffee
}

func (cd *CoffeeDecorator) GetDescription() string {
	return cd.coffee.GetDescription()
}

func (cd *CoffeeDecorator) GetCost() float64 {
	return cd.coffee.GetCost()
}

// Concrete Decorators
type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (md *MilkDecorator) GetDescription() string {
	return md.coffee.GetDescription() + ", Milk"
}

func (md *MilkDecorator) GetCost() float64 {
	return md.coffee.GetCost() + 0.5
}

type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (sd *SugarDecorator) GetDescription() string {
	return sd.coffee.GetDescription() + ", Sugar"
}

func (sd *SugarDecorator) GetCost() float64 {
	return sd.coffee.GetCost() + 0.2
}

type WhipDecorator struct {
	CoffeeDecorator
}

func NewWhipDecorator(coffee Coffee) *WhipDecorator {
	return &WhipDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (wd *WhipDecorator) GetDescription() string {
	return wd.coffee.GetDescription() + ", Whip"
}

func (wd *WhipDecorator) GetCost() float64 {
	return wd.coffee.GetCost() + 0.8
}

type VanillaDecorator struct {
	CoffeeDecorator
}

func NewVanillaDecorator(coffee Coffee) *VanillaDecorator {
	return &VanillaDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (vd *VanillaDecorator) GetDescription() string {
	return vd.coffee.GetDescription() + ", Vanilla"
}

func (vd *VanillaDecorator) GetCost() float64 {
	return vd.coffee.GetCost() + 0.6
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. TEXT FORMATTING DECORATOR
type TextFormatter interface {
	Format(text string) string
	GetDescription() string
}

type PlainText struct{}

func (pt *PlainText) Format(text string) string {
	return text
}

func (pt *PlainText) GetDescription() string {
	return "Plain Text"
}

type TextDecorator struct {
	formatter TextFormatter
}

func (td *TextDecorator) Format(text string) string {
	return td.formatter.Format(text)
}

func (td *TextDecorator) GetDescription() string {
	return td.formatter.GetDescription()
}

type BoldDecorator struct {
	TextDecorator
}

func NewBoldDecorator(formatter TextFormatter) *BoldDecorator {
	return &BoldDecorator{
		TextDecorator: TextDecorator{formatter: formatter},
	}
}

func (bd *BoldDecorator) Format(text string) string {
	return "<b>" + bd.formatter.Format(text) + "</b>"
}

func (bd *BoldDecorator) GetDescription() string {
	return bd.formatter.GetDescription() + " + Bold"
}

type ItalicDecorator struct {
	TextDecorator
}

func NewItalicDecorator(formatter TextFormatter) *ItalicDecorator {
	return &ItalicDecorator{
		TextDecorator: TextDecorator{formatter: formatter},
	}
}

func (id *ItalicDecorator) Format(text string) string {
	return "<i>" + id.formatter.Format(text) + "</i>"
}

func (id *ItalicDecorator) GetDescription() string {
	return id.formatter.GetDescription() + " + Italic"
}

type UnderlineDecorator struct {
	TextDecorator
}

func NewUnderlineDecorator(formatter TextFormatter) *UnderlineDecorator {
	return &UnderlineDecorator{
		TextDecorator: TextDecorator{formatter: formatter},
	}
}

func (ud *UnderlineDecorator) Format(text string) string {
	return "<u>" + ud.formatter.Format(text) + "</u>"
}

func (ud *UnderlineDecorator) GetDescription() string {
	return ud.formatter.GetDescription() + " + Underline"
}

type ColorDecorator struct {
	TextDecorator
	color string
}

func NewColorDecorator(formatter TextFormatter, color string) *ColorDecorator {
	return &ColorDecorator{
		TextDecorator: TextDecorator{formatter: formatter},
		color:         color,
	}
}

func (cd *ColorDecorator) Format(text string) string {
	return fmt.Sprintf("<span style='color: %s'>%s</span>", cd.color, cd.formatter.Format(text))
}

func (cd *ColorDecorator) GetDescription() string {
	return cd.formatter.GetDescription() + " + Color(" + cd.color + ")"
}

// 2. HTTP REQUEST DECORATOR
type HTTPRequest interface {
	Send() (string, error)
	GetURL() string
	GetMethod() string
	GetHeaders() map[string]string
}

type BasicHTTPRequest struct {
	url     string
	method  string
	headers map[string]string
	body    string
}

func NewBasicHTTPRequest(url, method string) *BasicHTTPRequest {
	return &BasicHTTPRequest{
		url:     url,
		method:  method,
		headers: make(map[string]string),
		body:    "",
	}
}

func (bhr *BasicHTTPRequest) Send() (string, error) {
	fmt.Printf("Sending %s request to %s\n", bhr.method, bhr.url)
	fmt.Printf("Headers: %v\n", bhr.headers)
	if bhr.body != "" {
		fmt.Printf("Body: %s\n", bhr.body)
	}
	return "Response from " + bhr.url, nil
}

func (bhr *BasicHTTPRequest) GetURL() string {
	return bhr.url
}

func (bhr *BasicHTTPRequest) GetMethod() string {
	return bhr.method
}

func (bhr *BasicHTTPRequest) GetHeaders() map[string]string {
	return bhr.headers
}

func (bhr *BasicHTTPRequest) SetBody(body string) {
	bhr.body = body
}

type HTTPRequestDecorator struct {
	request HTTPRequest
}

func (hrd *HTTPRequestDecorator) Send() (string, error) {
	return hrd.request.Send()
}

func (hrd *HTTPRequestDecorator) GetURL() string {
	return hrd.request.GetURL()
}

func (hrd *HTTPRequestDecorator) GetMethod() string {
	return hrd.request.GetMethod()
}

func (hrd *HTTPRequestDecorator) GetHeaders() map[string]string {
	return hrd.request.GetHeaders()
}

type AuthenticationDecorator struct {
	HTTPRequestDecorator
	token string
}

func NewAuthenticationDecorator(request HTTPRequest, token string) *AuthenticationDecorator {
	return &AuthenticationDecorator{
		HTTPRequestDecorator: HTTPRequestDecorator{request: request},
		token:                token,
	}
}

func (ad *AuthenticationDecorator) Send() (string, error) {
	ad.request.GetHeaders()["Authorization"] = "Bearer " + ad.token
	fmt.Printf("Added authentication token: %s\n", ad.token)
	return ad.request.Send()
}

type LoggingDecorator struct {
	HTTPRequestDecorator
}

func NewLoggingDecorator(request HTTPRequest) *LoggingDecorator {
	return &LoggingDecorator{
		HTTPRequestDecorator: HTTPRequestDecorator{request: request},
	}
}

func (ld *LoggingDecorator) Send() (string, error) {
	fmt.Printf("Logging request: %s %s at %s\n", ld.request.GetMethod(), ld.request.GetURL(), time.Now().Format("2006-01-02 15:04:05"))
	response, err := ld.request.Send()
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	} else {
		fmt.Printf("Request successful: %s\n", response)
	}
	return response, err
}

type RetryDecorator struct {
	HTTPRequestDecorator
	maxRetries int
}

func NewRetryDecorator(request HTTPRequest, maxRetries int) *RetryDecorator {
	return &RetryDecorator{
		HTTPRequestDecorator: HTTPRequestDecorator{request: request},
		maxRetries:           maxRetries,
	}
}

func (rd *RetryDecorator) Send() (string, error) {
	var lastErr error
	for i := 0; i <= rd.maxRetries; i++ {
		if i > 0 {
			fmt.Printf("Retry attempt %d/%d\n", i, rd.maxRetries)
		}
		response, err := rd.request.Send()
		if err == nil {
			return response, nil
		}
		lastErr = err
	}
	return "", fmt.Errorf("request failed after %d retries: %v", rd.maxRetries, lastErr)
}

// 3. STREAM DECORATOR
type Stream interface {
	Read() ([]byte, error)
	Write(data []byte) error
	Close() error
	GetType() string
}

type BasicStream struct {
	data []byte
	pos  int
}

func NewBasicStream() *BasicStream {
	return &BasicStream{
		data: make([]byte, 0),
		pos:  0,
	}
}

func (bs *BasicStream) Read() ([]byte, error) {
	if bs.pos >= len(bs.data) {
		return nil, fmt.Errorf("end of stream")
	}
	result := bs.data[bs.pos:]
	bs.pos = len(bs.data)
	return result, nil
}

func (bs *BasicStream) Write(data []byte) error {
	bs.data = append(bs.data, data...)
	return nil
}

func (bs *BasicStream) Close() error {
	bs.data = nil
	bs.pos = 0
	return nil
}

func (bs *BasicStream) GetType() string {
	return "Basic Stream"
}

type StreamDecorator struct {
	stream Stream
}

func (sd *StreamDecorator) Read() ([]byte, error) {
	return sd.stream.Read()
}

func (sd *StreamDecorator) Write(data []byte) error {
	return sd.stream.Write(data)
}

func (sd *StreamDecorator) Close() error {
	return sd.stream.Close()
}

func (sd *StreamDecorator) GetType() string {
	return sd.stream.GetType()
}

type CompressionDecorator struct {
	StreamDecorator
}

func NewCompressionDecorator(stream Stream) *CompressionDecorator {
	return &CompressionDecorator{
		StreamDecorator: StreamDecorator{stream: stream},
	}
}

func (cd *CompressionDecorator) Write(data []byte) error {
	// Simulate compression
	compressed := []byte("COMPRESSED:" + string(data))
	fmt.Printf("Compressing data: %s -> %s\n", string(data), string(compressed))
	return cd.stream.Write(compressed)
}

func (cd *CompressionDecorator) Read() ([]byte, error) {
	data, err := cd.stream.Read()
	if err != nil {
		return nil, err
	}
	// Simulate decompression
	if strings.HasPrefix(string(data), "COMPRESSED:") {
		decompressed := []byte(strings.TrimPrefix(string(data), "COMPRESSED:"))
		fmt.Printf("Decompressing data: %s -> %s\n", string(data), string(decompressed))
		return decompressed, nil
	}
	return data, nil
}

func (cd *CompressionDecorator) GetType() string {
	return cd.stream.GetType() + " + Compression"
}

type EncryptionDecorator struct {
	StreamDecorator
	key string
}

func NewEncryptionDecorator(stream Stream, key string) *EncryptionDecorator {
	return &EncryptionDecorator{
		StreamDecorator: StreamDecorator{stream: stream},
		key:             key,
	}
}

func (ed *EncryptionDecorator) Write(data []byte) error {
	// Simulate encryption
	encrypted := []byte("ENCRYPTED:" + string(data))
	fmt.Printf("Encrypting data with key %s: %s -> %s\n", ed.key, string(data), string(encrypted))
	return ed.stream.Write(encrypted)
}

func (ed *EncryptionDecorator) Read() ([]byte, error) {
	data, err := ed.stream.Read()
	if err != nil {
		return nil, err
	}
	// Simulate decryption
	if strings.HasPrefix(string(data), "ENCRYPTED:") {
		decrypted := []byte(strings.TrimPrefix(string(data), "ENCRYPTED:"))
		fmt.Printf("Decrypting data with key %s: %s -> %s\n", ed.key, string(data), string(decrypted))
		return decrypted, nil
	}
	return data, nil
}

func (ed *EncryptionDecorator) GetType() string {
	return ed.stream.GetType() + " + Encryption"
}

// =============================================================================
// DECORATOR WITH STATE
// =============================================================================

type StatefulDecorator struct {
	component Coffee
	state     map[string]interface{}
}

func NewStatefulDecorator(component Coffee) *StatefulDecorator {
	return &StatefulDecorator{
		component: component,
		state:     make(map[string]interface{}),
	}
}

func (sd *StatefulDecorator) GetDescription() string {
	return sd.component.GetDescription()
}

func (sd *StatefulDecorator) GetCost() float64 {
	return sd.component.GetCost()
}

func (sd *StatefulDecorator) SetState(key string, value interface{}) {
	sd.state[key] = value
}

func (sd *StatefulDecorator) GetState(key string) interface{} {
	return sd.state[key]
}

func (sd *StatefulDecorator) GetAllState() map[string]interface{} {
	return sd.state
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== DECORATOR PATTERN DEMONSTRATION ===\n")

	// 1. BASIC DECORATOR
	fmt.Println("1. BASIC DECORATOR:")
	
	// Create basic coffee
	coffee := &BasicCoffee{}
	fmt.Printf("Basic coffee: %s, Cost: $%.2f\n", coffee.GetDescription(), coffee.GetCost())
	
	// Add decorators
	coffeeWithMilk := NewMilkDecorator(coffee)
	fmt.Printf("Coffee with milk: %s, Cost: $%.2f\n", coffeeWithMilk.GetDescription(), coffeeWithMilk.GetCost())
	
	coffeeWithMilkAndSugar := NewSugarDecorator(coffeeWithMilk)
	fmt.Printf("Coffee with milk and sugar: %s, Cost: $%.2f\n", coffeeWithMilkAndSugar.GetDescription(), coffeeWithMilkAndSugar.GetCost())
	
	coffeeWithEverything := NewWhipDecorator(NewVanillaDecorator(coffeeWithMilkAndSugar))
	fmt.Printf("Coffee with everything: %s, Cost: $%.2f\n", coffeeWithEverything.GetDescription(), coffeeWithEverything.GetCost())
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Text Formatting Decorator
	fmt.Println("Text Formatting Decorator:")
	plainText := &PlainText{}
	formattedText := NewBoldDecorator(NewItalicDecorator(NewColorDecorator(plainText, "red")))
	
	text := "Hello, World!"
	fmt.Printf("Original text: %s\n", text)
	fmt.Printf("Formatted text: %s\n", formattedText.Format(text))
	fmt.Printf("Formatter description: %s\n", formattedText.GetDescription())
	fmt.Println()

	// HTTP Request Decorator
	fmt.Println("HTTP Request Decorator:")
	basicRequest := NewBasicHTTPRequest("https://api.example.com/users", "GET")
	basicRequest.SetBody("{\"name\": \"John\"}")
	
	// Add decorators
	authenticatedRequest := NewAuthenticationDecorator(basicRequest, "abc123")
	loggedRequest := NewLoggingDecorator(authenticatedRequest)
	retryRequest := NewRetryDecorator(loggedRequest, 3)
	
	response, err := retryRequest.Send()
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	} else {
		fmt.Printf("Request successful: %s\n", response)
	}
	fmt.Println()

	// Stream Decorator
	fmt.Println("Stream Decorator:")
	basicStream := NewBasicStream()
	compressedStream := NewCompressionDecorator(basicStream)
	encryptedStream := NewEncryptionDecorator(compressedStream, "secretkey")
	
	// Write data
	data := []byte("Hello, World!")
	encryptedStream.Write(data)
	
	// Read data
	readData, err := encryptedStream.Read()
	if err != nil {
		fmt.Printf("Read failed: %v\n", err)
	} else {
		fmt.Printf("Read data: %s\n", string(readData))
	}
	
	encryptedStream.Close()
	fmt.Printf("Stream type: %s\n", encryptedStream.GetType())
	fmt.Println()

	// 3. DECORATOR WITH STATE
	fmt.Println("3. DECORATOR WITH STATE:")
	statefulDecorator := NewStatefulDecorator(coffee)
	statefulDecorator.SetState("temperature", "hot")
	statefulDecorator.SetState("size", "large")
	statefulDecorator.SetState("customization", "extra strong")
	
	fmt.Printf("Coffee: %s, Cost: $%.2f\n", statefulDecorator.GetDescription(), statefulDecorator.GetCost())
	fmt.Printf("State: %v\n", statefulDecorator.GetAllState())
	fmt.Printf("Temperature: %v\n", statefulDecorator.GetState("temperature"))
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
