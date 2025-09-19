package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC TEMPLATE METHOD PATTERN
// =============================================================================

// Abstract class
type AbstractClass struct{}

func (ac *AbstractClass) TemplateMethod() {
	fmt.Println("Template Method: Starting algorithm")
	ac.Step1()
	ac.Step2()
	ac.Step3()
	ac.Step4()
	fmt.Println("Template Method: Algorithm completed")
}

func (ac *AbstractClass) Step1() {
	fmt.Println("AbstractClass: Step1 - Default implementation")
}

func (ac *AbstractClass) Step2() {
	// Abstract method - must be implemented by subclasses
	panic("Step2 must be implemented by subclass")
}

func (ac *AbstractClass) Step3() {
	fmt.Println("AbstractClass: Step3 - Default implementation")
}

func (ac *AbstractClass) Step4() {
	// Hook method - can be overridden by subclasses
	fmt.Println("AbstractClass: Step4 - Default hook implementation")
}

// Concrete class A
type ConcreteClassA struct {
	AbstractClass
}

func (cca *ConcreteClassA) Step2() {
	fmt.Println("ConcreteClassA: Step2 - Custom implementation A")
}

func (cca *ConcreteClassA) Step4() {
	fmt.Println("ConcreteClassA: Step4 - Overridden hook implementation A")
}

// Concrete class B
type ConcreteClassB struct {
	AbstractClass
}

func (ccb *ConcreteClassB) Step2() {
	fmt.Println("ConcreteClassB: Step2 - Custom implementation B")
}

func (ccb *ConcreteClassB) Step4() {
	fmt.Println("ConcreteClassB: Step4 - Overridden hook implementation B")
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. DATA PROCESSING PIPELINE
type DataProcessor interface {
	LoadData() []string
	ProcessData(data []string) []string
	SaveData(data []string) error
	GetName() string
}

type AbstractDataProcessor struct{}

func (adp *AbstractDataProcessor) ProcessPipeline() error {
	fmt.Println("Data Processing Pipeline: Starting")
	
	// Load data
	data := adp.LoadData()
	fmt.Printf("Loaded %d records\n", len(data))
	
	// Process data
	processedData := adp.ProcessData(data)
	fmt.Printf("Processed %d records\n", len(processedData))
	
	// Save data
	err := adp.SaveData(processedData)
	if err != nil {
		return err
	}
	
	fmt.Println("Data Processing Pipeline: Completed")
	return nil
}

func (adp *AbstractDataProcessor) LoadData() []string {
	// Default implementation - must be overridden
	panic("LoadData must be implemented by subclass")
}

func (adp *AbstractDataProcessor) ProcessData(data []string) []string {
	// Default implementation - can be overridden
	fmt.Println("AbstractDataProcessor: Processing data with default logic")
	return data
}

func (adp *AbstractDataProcessor) SaveData(data []string) error {
	// Default implementation - must be overridden
	panic("SaveData must be implemented by subclass")
}

func (adp *AbstractDataProcessor) GetName() string {
	return "AbstractDataProcessor"
}

// CSV Data Processor
type CSVDataProcessor struct {
	AbstractDataProcessor
	filename string
}

func NewCSVDataProcessor(filename string) *CSVDataProcessor {
	return &CSVDataProcessor{filename: filename}
}

func (cdp *CSVDataProcessor) LoadData() []string {
	fmt.Printf("CSVDataProcessor: Loading data from %s\n", cdp.filename)
	// Simulate loading CSV data
	return []string{"record1", "record2", "record3", "record4", "record5"}
}

func (cdp *CSVDataProcessor) ProcessData(data []string) []string {
	fmt.Println("CSVDataProcessor: Processing CSV data")
	// Simulate CSV processing
	processed := make([]string, len(data))
	for i, record := range data {
		processed[i] = "processed_" + record
	}
	return processed
}

func (cdp *CSVDataProcessor) SaveData(data []string) error {
	fmt.Printf("CSVDataProcessor: Saving %d records to %s\n", len(data), cdp.filename)
	// Simulate saving
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (cdp *CSVDataProcessor) GetName() string {
	return "CSVDataProcessor"
}

// JSON Data Processor
type JSONDataProcessor struct {
	AbstractDataProcessor
	filename string
}

func NewJSONDataProcessor(filename string) *JSONDataProcessor {
	return &JSONDataProcessor{filename: filename}
}

func (jdp *JSONDataProcessor) LoadData() []string {
	fmt.Printf("JSONDataProcessor: Loading data from %s\n", jdp.filename)
	// Simulate loading JSON data
	return []string{"json_record1", "json_record2", "json_record3"}
}

func (jdp *JSONDataProcessor) ProcessData(data []string) []string {
	fmt.Println("JSONDataProcessor: Processing JSON data")
	// Simulate JSON processing
	processed := make([]string, len(data))
	for i, record := range data {
		processed[i] = "json_processed_" + record
	}
	return processed
}

func (jdp *JSONDataProcessor) SaveData(data []string) error {
	fmt.Printf("JSONDataProcessor: Saving %d records to %s\n", len(data), jdp.filename)
	// Simulate saving
	time.Sleep(150 * time.Millisecond)
	return nil
}

func (jdp *JSONDataProcessor) GetName() string {
	return "JSONDataProcessor"
}

// 2. DOCUMENT GENERATION
type DocumentGenerator interface {
	CreateHeader() string
	CreateBody() string
	CreateFooter() string
	GenerateDocument() string
	GetName() string
}

type AbstractDocumentGenerator struct{}

func (adg *AbstractDocumentGenerator) GenerateDocument() string {
	fmt.Println("Document Generation: Starting")
	
	header := adg.CreateHeader()
	body := adg.CreateBody()
	footer := adg.CreateFooter()
	
	document := header + "\n" + body + "\n" + footer
	
	fmt.Println("Document Generation: Completed")
	return document
}

func (adg *AbstractDocumentGenerator) CreateHeader() string {
	// Default implementation - must be overridden
	panic("CreateHeader must be implemented by subclass")
}

func (adg *AbstractDocumentGenerator) CreateBody() string {
	// Default implementation - must be overridden
	panic("CreateBody must be implemented by subclass")
}

func (adg *AbstractDocumentGenerator) CreateFooter() string {
	// Default implementation - can be overridden
	return "Generated by AbstractDocumentGenerator"
}

func (adg *AbstractDocumentGenerator) GetName() string {
	return "AbstractDocumentGenerator"
}

// HTML Document Generator
type HTMLDocumentGenerator struct {
	AbstractDocumentGenerator
	title string
}

func NewHTMLDocumentGenerator(title string) *HTMLDocumentGenerator {
	return &HTMLDocumentGenerator{title: title}
}

func (hdg *HTMLDocumentGenerator) CreateHeader() string {
	return fmt.Sprintf("<html><head><title>%s</title></head><body>", hdg.title)
}

func (hdg *HTMLDocumentGenerator) CreateBody() string {
	return "<h1>Welcome to HTML Document</h1><p>This is the body content.</p>"
}

func (hdg *HTMLDocumentGenerator) CreateFooter() string {
	return "</body></html>"
}

func (hdg *HTMLDocumentGenerator) GetName() string {
	return "HTMLDocumentGenerator"
}

// PDF Document Generator
type PDFDocumentGenerator struct {
	AbstractDocumentGenerator
	title string
}

func NewPDFDocumentGenerator(title string) *PDFDocumentGenerator {
	return &PDFDocumentGenerator{title: title}
}

func (pdg *PDFDocumentGenerator) CreateHeader() string {
	return fmt.Sprintf("PDF Document: %s\n", pdg.title)
}

func (pdg *PDFDocumentGenerator) CreateBody() string {
	return "This is the PDF document body content.\n"
}

func (pdg *PDFDocumentGenerator) CreateFooter() string {
	return "End of PDF Document"
}

func (pdg *PDFDocumentGenerator) GetName() string {
	return "PDFDocumentGenerator"
}

// 3. DATABASE OPERATIONS
type DatabaseOperation interface {
	Connect() error
	ExecuteQuery() error
	Close() error
	GetName() string
}

type AbstractDatabaseOperation struct{}

func (ado *AbstractDatabaseOperation) Execute() error {
	fmt.Println("Database Operation: Starting")
	
	err := ado.Connect()
	if err != nil {
		return err
	}
	
	err = ado.ExecuteQuery()
	if err != nil {
		ado.Close()
		return err
	}
	
	err = ado.Close()
	if err != nil {
		return err
	}
	
	fmt.Println("Database Operation: Completed")
	return nil
}

func (ado *AbstractDatabaseOperation) Connect() error {
	// Default implementation - must be overridden
	panic("Connect must be implemented by subclass")
}

func (ado *AbstractDatabaseOperation) ExecuteQuery() error {
	// Default implementation - must be overridden
	panic("ExecuteQuery must be implemented by subclass")
}

func (ado *AbstractDatabaseOperation) Close() error {
	// Default implementation - must be overridden
	panic("Close must be implemented by subclass")
}

func (ado *AbstractDatabaseOperation) GetName() string {
	return "AbstractDatabaseOperation"
}

// MySQL Database Operation
type MySQLDatabaseOperation struct {
	AbstractDatabaseOperation
	host     string
	port     int
	username string
	password string
	database string
}

func NewMySQLDatabaseOperation(host string, port int, username, password, database string) *MySQLDatabaseOperation {
	return &MySQLDatabaseOperation{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}
}

func (mdo *MySQLDatabaseOperation) Connect() error {
	fmt.Printf("MySQL: Connecting to %s:%d/%s as %s\n", mdo.host, mdo.port, mdo.database, mdo.username)
	// Simulate connection
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (mdo *MySQLDatabaseOperation) ExecuteQuery() error {
	fmt.Println("MySQL: Executing SELECT * FROM users")
	// Simulate query execution
	time.Sleep(200 * time.Millisecond)
	return nil
}

func (mdo *MySQLDatabaseOperation) Close() error {
	fmt.Println("MySQL: Closing connection")
	// Simulate connection close
	time.Sleep(50 * time.Millisecond)
	return nil
}

func (mdo *MySQLDatabaseOperation) GetName() string {
	return "MySQLDatabaseOperation"
}

// PostgreSQL Database Operation
type PostgreSQLDatabaseOperation struct {
	AbstractDatabaseOperation
	host     string
	port     int
	username string
	password string
	database string
}

func NewPostgreSQLDatabaseOperation(host string, port int, username, password, database string) *PostgreSQLDatabaseOperation {
	return &PostgreSQLDatabaseOperation{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}
}

func (pdo *PostgreSQLDatabaseOperation) Connect() error {
	fmt.Printf("PostgreSQL: Connecting to %s:%d/%s as %s\n", pdo.host, pdo.port, pdo.database, pdo.username)
	// Simulate connection
	time.Sleep(120 * time.Millisecond)
	return nil
}

func (pdo *PostgreSQLDatabaseOperation) ExecuteQuery() error {
	fmt.Println("PostgreSQL: Executing SELECT * FROM products")
	// Simulate query execution
	time.Sleep(180 * time.Millisecond)
	return nil
}

func (pdo *PostgreSQLDatabaseOperation) Close() error {
	fmt.Println("PostgreSQL: Closing connection")
	// Simulate connection close
	time.Sleep(60 * time.Millisecond)
	return nil
}

func (pdo *PostgreSQLDatabaseOperation) GetName() string {
	return "PostgreSQLDatabaseOperation"
}

// 4. BUILD PROCESS
type BuildProcess interface {
	Checkout() error
	InstallDependencies() error
	Build() error
	Test() error
	Package() error
	Deploy() error
	GetName() string
}

type AbstractBuildProcess struct{}

func (abp *AbstractBuildProcess) Execute() error {
	fmt.Println("Build Process: Starting")
	
	steps := []func() error{
		abp.Checkout,
		abp.InstallDependencies,
		abp.Build,
		abp.Test,
		abp.Package,
		abp.Deploy,
	}
	
	for i, step := range steps {
		fmt.Printf("Build Process: Step %d\n", i+1)
		err := step()
		if err != nil {
			fmt.Printf("Build Process: Failed at step %d\n", i+1)
			return err
		}
	}
	
	fmt.Println("Build Process: Completed")
	return nil
}

func (abp *AbstractBuildProcess) Checkout() error {
	// Default implementation - must be overridden
	panic("Checkout must be implemented by subclass")
}

func (abp *AbstractBuildProcess) InstallDependencies() error {
	// Default implementation - must be overridden
	panic("InstallDependencies must be implemented by subclass")
}

func (abp *AbstractBuildProcess) Build() error {
	// Default implementation - must be overridden
	panic("Build must be implemented by subclass")
}

func (abp *AbstractBuildProcess) Test() error {
	// Default implementation - can be overridden
	fmt.Println("AbstractBuildProcess: Running default tests")
	return nil
}

func (abp *AbstractBuildProcess) Package() error {
	// Default implementation - must be overridden
	panic("Package must be implemented by subclass")
}

func (abp *AbstractBuildProcess) Deploy() error {
	// Default implementation - can be overridden
	fmt.Println("AbstractBuildProcess: Deploying to default environment")
	return nil
}

func (abp *AbstractBuildProcess) GetName() string {
	return "AbstractBuildProcess"
}

// Node.js Build Process
type NodeJSBuildProcess struct {
	AbstractBuildProcess
	projectPath string
}

func NewNodeJSBuildProcess(projectPath string) *NodeJSBuildProcess {
	return &NodeJSBuildProcess{projectPath: projectPath}
}

func (njbp *NodeJSBuildProcess) Checkout() error {
	fmt.Printf("Node.js: Checking out code from %s\n", njbp.projectPath)
	// Simulate checkout
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (njbp *NodeJSBuildProcess) InstallDependencies() error {
	fmt.Println("Node.js: Running npm install")
	// Simulate npm install
	time.Sleep(200 * time.Millisecond)
	return nil
}

func (njbp *NodeJSBuildProcess) Build() error {
	fmt.Println("Node.js: Running npm run build")
	// Simulate build
	time.Sleep(300 * time.Millisecond)
	return nil
}

func (njbp *NodeJSBuildProcess) Test() error {
	fmt.Println("Node.js: Running npm test")
	// Simulate test
	time.Sleep(150 * time.Millisecond)
	return nil
}

func (njbp *NodeJSBuildProcess) Package() error {
	fmt.Println("Node.js: Creating package")
	// Simulate packaging
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (njbp *NodeJSBuildProcess) Deploy() error {
	fmt.Println("Node.js: Deploying to production")
	// Simulate deployment
	time.Sleep(250 * time.Millisecond)
	return nil
}

func (njbp *NodeJSBuildProcess) GetName() string {
	return "NodeJSBuildProcess"
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== TEMPLATE METHOD PATTERN DEMONSTRATION ===\n")

	// 1. BASIC TEMPLATE METHOD
	fmt.Println("1. BASIC TEMPLATE METHOD:")
	concreteA := &ConcreteClassA{}
	concreteA.TemplateMethod()
	fmt.Println()
	
	concreteB := &ConcreteClassB{}
	concreteB.TemplateMethod()
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Data Processing Pipeline
	fmt.Println("Data Processing Pipeline:")
	csvProcessor := NewCSVDataProcessor("data.csv")
	csvProcessor.ProcessPipeline()
	fmt.Println()
	
	jsonProcessor := NewJSONDataProcessor("data.json")
	jsonProcessor.ProcessPipeline()
	fmt.Println()

	// Document Generation
	fmt.Println("Document Generation:")
	htmlGenerator := NewHTMLDocumentGenerator("My HTML Document")
	htmlDoc := htmlGenerator.GenerateDocument()
	fmt.Printf("Generated HTML Document:\n%s\n\n", htmlDoc)
	
	pdfGenerator := NewPDFDocumentGenerator("My PDF Document")
	pdfDoc := pdfGenerator.GenerateDocument()
	fmt.Printf("Generated PDF Document:\n%s\n\n", pdfDoc)

	// Database Operations
	fmt.Println("Database Operations:")
	mysqlOp := NewMySQLDatabaseOperation("localhost", 3306, "user", "pass", "mydb")
	mysqlOp.Execute()
	fmt.Println()
	
	postgresOp := NewPostgreSQLDatabaseOperation("localhost", 5432, "user", "pass", "mydb")
	postgresOp.Execute()
	fmt.Println()

	// Build Process
	fmt.Println("Build Process:")
	nodejsBuild := NewNodeJSBuildProcess("/path/to/project")
	nodejsBuild.Execute()
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
