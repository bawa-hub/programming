package main

import (
	"fmt"
	"strings"
	"time"
)

// =============================================================================
// BASIC BUILDER PATTERN
// =============================================================================

// Product - Complex object to be built
type Computer struct {
	CPU        string
	RAM        int
	Storage    int
	GPU        string
	Motherboard string
	PowerSupply string
	Case       string
	Price      float64
}

func (c *Computer) String() string {
	return fmt.Sprintf("Computer: CPU=%s, RAM=%dGB, Storage=%dGB, GPU=%s, Price=$%.2f",
		c.CPU, c.RAM, c.Storage, c.GPU, c.Price)
}

// Builder interface
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetRAM(ram int) ComputerBuilder
	SetStorage(storage int) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder
	SetMotherboard(motherboard string) ComputerBuilder
	SetPowerSupply(powerSupply string) ComputerBuilder
	SetCase(caseType string) ComputerBuilder
	CalculatePrice() ComputerBuilder
	Build() *Computer
}

// Concrete Builder
type GamingComputerBuilder struct {
	computer *Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{
		computer: &Computer{},
	}
}

func (gcb *GamingComputerBuilder) SetCPU(cpu string) ComputerBuilder {
	gcb.computer.CPU = cpu
	return gcb
}

func (gcb *GamingComputerBuilder) SetRAM(ram int) ComputerBuilder {
	gcb.computer.RAM = ram
	return gcb
}

func (gcb *GamingComputerBuilder) SetStorage(storage int) ComputerBuilder {
	gcb.computer.Storage = storage
	return gcb
}

func (gcb *GamingComputerBuilder) SetGPU(gpu string) ComputerBuilder {
	gcb.computer.GPU = gpu
	return gcb
}

func (gcb *GamingComputerBuilder) SetMotherboard(motherboard string) ComputerBuilder {
	gcb.computer.Motherboard = motherboard
	return gcb
}

func (gcb *GamingComputerBuilder) SetPowerSupply(powerSupply string) ComputerBuilder {
	gcb.computer.PowerSupply = powerSupply
	return gcb
}

func (gcb *GamingComputerBuilder) SetCase(caseType string) ComputerBuilder {
	gcb.computer.Case = caseType
	return gcb
}

func (gcb *GamingComputerBuilder) CalculatePrice() ComputerBuilder {
	// Simple price calculation based on components
	price := 0.0
	
	switch gcb.computer.CPU {
	case "Intel i9":
		price += 500
	case "Intel i7":
		price += 350
	case "AMD Ryzen 9":
		price += 450
	case "AMD Ryzen 7":
		price += 300
	}
	
	price += float64(gcb.computer.RAM) * 10
	price += float64(gcb.computer.Storage) * 0.1
	
	switch gcb.computer.GPU {
	case "RTX 4090":
		price += 1600
	case "RTX 4080":
		price += 1200
	case "RTX 4070":
		price += 600
	case "RTX 4060":
		price += 300
	}
	
	gcb.computer.Price = price
	return gcb
}

func (gcb *GamingComputerBuilder) Build() *Computer {
	return gcb.computer
}

// =============================================================================
// BUILDER WITH DIRECTOR PATTERN
// =============================================================================

// Director - Orchestrates the construction process
type ComputerDirector struct {
	builder ComputerBuilder
}

func NewComputerDirector(builder ComputerBuilder) *ComputerDirector {
	return &ComputerDirector{builder: builder}
}

func (cd *ComputerDirector) SetBuilder(builder ComputerBuilder) {
	cd.builder = builder
}

func (cd *ComputerDirector) BuildGamingPC() *Computer {
	return cd.builder.
		SetCPU("Intel i9").
		SetRAM(32).
		SetStorage(1000).
		SetGPU("RTX 4090").
		SetMotherboard("Z790").
		SetPowerSupply("1000W").
		SetCase("Full Tower").
		CalculatePrice().
		Build()
}

func (cd *ComputerDirector) BuildOfficePC() *Computer {
	return cd.builder.
		SetCPU("Intel i5").
		SetRAM(16).
		SetStorage(500).
		SetGPU("Integrated").
		SetMotherboard("B660").
		SetPowerSupply("500W").
		SetCase("Mid Tower").
		CalculatePrice().
		Build()
}

func (cd *ComputerDirector) BuildBudgetPC() *Computer {
	return cd.builder.
		SetCPU("AMD Ryzen 5").
		SetRAM(8).
		SetStorage(250).
		SetGPU("RTX 4060").
		SetMotherboard("B550").
		SetPowerSupply("600W").
		SetCase("Mini Tower").
		CalculatePrice().
		Build()
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. SQL QUERY BUILDER
type SQLQuery struct {
	selectClause string
	fromClause   string
	whereClause  string
	orderByClause string
	limitClause  string
}

func (sq *SQLQuery) String() string {
	var parts []string
	if sq.selectClause != "" {
		parts = append(parts, sq.selectClause)
	}
	if sq.fromClause != "" {
		parts = append(parts, sq.fromClause)
	}
	if sq.whereClause != "" {
		parts = append(parts, sq.whereClause)
	}
	if sq.orderByClause != "" {
		parts = append(parts, sq.orderByClause)
	}
	if sq.limitClause != "" {
		parts = append(parts, sq.limitClause)
	}
	return strings.Join(parts, " ")
}

type SQLQueryBuilder struct {
	query *SQLQuery
}

func NewSQLQueryBuilder() *SQLQueryBuilder {
	return &SQLQueryBuilder{
		query: &SQLQuery{},
	}
}

func (sqb *SQLQueryBuilder) Select(columns ...string) *SQLQueryBuilder {
	sqb.query.selectClause = "SELECT " + strings.Join(columns, ", ")
	return sqb
}

func (sqb *SQLQueryBuilder) From(table string) *SQLQueryBuilder {
	sqb.query.fromClause = "FROM " + table
	return sqb
}

func (sqb *SQLQueryBuilder) Where(condition string) *SQLQueryBuilder {
	sqb.query.whereClause = "WHERE " + condition
	return sqb
}

func (sqb *SQLQueryBuilder) OrderBy(column string, direction string) *SQLQueryBuilder {
	sqb.query.orderByClause = "ORDER BY " + column + " " + direction
	return sqb
}

func (sqb *SQLQueryBuilder) Limit(count int) *SQLQueryBuilder {
	sqb.query.limitClause = fmt.Sprintf("LIMIT %d", count)
	return sqb
}

func (sqb *SQLQueryBuilder) Build() *SQLQuery {
	return sqb.query
}

// 2. HTTP REQUEST BUILDER
type HTTPRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
	Timeout time.Duration
}

func (hr *HTTPRequest) String() string {
	return fmt.Sprintf("HTTP %s %s\nHeaders: %v\nBody: %s\nTimeout: %v",
		hr.Method, hr.URL, hr.Headers, hr.Body, hr.Timeout)
}

type HTTPRequestBuilder struct {
	request *HTTPRequest
}

func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		request: &HTTPRequest{
			Headers: make(map[string]string),
		},
	}
}

func (hrb *HTTPRequestBuilder) SetMethod(method string) *HTTPRequestBuilder {
	hrb.request.Method = method
	return hrb
}

func (hrb *HTTPRequestBuilder) SetURL(url string) *HTTPRequestBuilder {
	hrb.request.URL = url
	return hrb
}

func (hrb *HTTPRequestBuilder) AddHeader(key, value string) *HTTPRequestBuilder {
	hrb.request.Headers[key] = value
	return hrb
}

func (hrb *HTTPRequestBuilder) SetBody(body string) *HTTPRequestBuilder {
	hrb.request.Body = body
	return hrb
}

func (hrb *HTTPRequestBuilder) SetTimeout(timeout time.Duration) *HTTPRequestBuilder {
	hrb.request.Timeout = timeout
	return hrb
}

func (hrb *HTTPRequestBuilder) Build() *HTTPRequest {
	return hrb.request
}

// 3. CONFIGURATION BUILDER
type AppConfig struct {
	AppName        string
	Version        string
	Port           int
	DatabaseURL    string
	RedisURL       string
	LogLevel       string
	MaxConnections int
	Timeout        time.Duration
	Features       map[string]bool
}

func (ac *AppConfig) String() string {
	return fmt.Sprintf("App: %s v%s, Port: %d, DB: %s, Redis: %s, LogLevel: %s, MaxConn: %d, Timeout: %v, Features: %v",
		ac.AppName, ac.Version, ac.Port, ac.DatabaseURL, ac.RedisURL, ac.LogLevel, ac.MaxConnections, ac.Timeout, ac.Features)
}

type ConfigBuilder struct {
	config *AppConfig
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: &AppConfig{
			Features: make(map[string]bool),
		},
	}
}

func (cb *ConfigBuilder) SetAppName(name string) *ConfigBuilder {
	cb.config.AppName = name
	return cb
}

func (cb *ConfigBuilder) SetVersion(version string) *ConfigBuilder {
	cb.config.Version = version
	return cb
}

func (cb *ConfigBuilder) SetPort(port int) *ConfigBuilder {
	cb.config.Port = port
	return cb
}

func (cb *ConfigBuilder) SetDatabaseURL(url string) *ConfigBuilder {
	cb.config.DatabaseURL = url
	return cb
}

func (cb *ConfigBuilder) SetRedisURL(url string) *ConfigBuilder {
	cb.config.RedisURL = url
	return cb
}

func (cb *ConfigBuilder) SetLogLevel(level string) *ConfigBuilder {
	cb.config.LogLevel = level
	return cb
}

func (cb *ConfigBuilder) SetMaxConnections(max int) *ConfigBuilder {
	cb.config.MaxConnections = max
	return cb
}

func (cb *ConfigBuilder) SetTimeout(timeout time.Duration) *ConfigBuilder {
	cb.config.Timeout = timeout
	return cb
}

func (cb *ConfigBuilder) EnableFeature(feature string) *ConfigBuilder {
	cb.config.Features[feature] = true
	return cb
}

func (cb *ConfigBuilder) DisableFeature(feature string) *ConfigBuilder {
	cb.config.Features[feature] = false
	return cb
}

func (cb *ConfigBuilder) Build() *AppConfig {
	return cb.config
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== BUILDER PATTERN DEMONSTRATION ===\n")

	// 1. BASIC BUILDER
	fmt.Println("1. BASIC BUILDER:")
	gamingBuilder := NewGamingComputerBuilder()
	gamingPC := gamingBuilder.
		SetCPU("Intel i9").
		SetRAM(32).
		SetStorage(1000).
		SetGPU("RTX 4090").
		SetMotherboard("Z790").
		SetPowerSupply("1000W").
		SetCase("Full Tower").
		CalculatePrice().
		Build()
	
	fmt.Printf("Gaming PC: %s\n", gamingPC)
	fmt.Println()

	// 2. BUILDER WITH DIRECTOR
	fmt.Println("2. BUILDER WITH DIRECTOR:")
	director := NewComputerDirector(NewGamingComputerBuilder())
	
	gamingPC2 := director.BuildGamingPC()
	fmt.Printf("Director-built Gaming PC: %s\n", gamingPC2)
	
	officePC := director.BuildOfficePC()
	fmt.Printf("Director-built Office PC: %s\n", officePC)
	
	budgetPC := director.BuildBudgetPC()
	fmt.Printf("Director-built Budget PC: %s\n", budgetPC)
	fmt.Println()

	// 3. REAL-WORLD EXAMPLES
	fmt.Println("3. REAL-WORLD EXAMPLES:")

	// SQL Query Builder
	fmt.Println("SQL Query Builder:")
	query := NewSQLQueryBuilder().
		Select("id", "name", "email").
		From("users").
		Where("age > 18").
		OrderBy("name", "ASC").
		Limit(10).
		Build()
	
	fmt.Printf("SQL Query: %s\n", query)
	fmt.Println()

	// HTTP Request Builder
	fmt.Println("HTTP Request Builder:")
	httpRequest := NewHTTPRequestBuilder().
		SetMethod("POST").
		SetURL("https://api.example.com/users").
		AddHeader("Content-Type", "application/json").
		AddHeader("Authorization", "Bearer token123").
		SetBody(`{"name": "John", "email": "john@example.com"}`).
		SetTimeout(30 * time.Second).
		Build()
	
	fmt.Printf("HTTP Request:\n%s\n", httpRequest)
	fmt.Println()

	// Configuration Builder
	fmt.Println("Configuration Builder:")
	config := NewConfigBuilder().
		SetAppName("MyApp").
		SetVersion("1.0.0").
		SetPort(8080).
		SetDatabaseURL("postgresql://localhost:5432/mydb").
		SetRedisURL("redis://localhost:6379").
		SetLogLevel("INFO").
		SetMaxConnections(100).
		SetTimeout(30 * time.Second).
		EnableFeature("authentication").
		EnableFeature("rate_limiting").
		DisableFeature("debug_mode").
		Build()
	
	fmt.Printf("App Config: %s\n", config)
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
