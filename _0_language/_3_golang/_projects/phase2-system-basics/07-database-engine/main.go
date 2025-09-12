package main

import (
	"database-engine/concurrency"
	"database-engine/index"
	"database-engine/query"
	"database-engine/schema"
	"database-engine/storage"
	"database-engine/transaction"
	"database-engine/types"
	"flag"
	"fmt"
	"os"
	"time"
)

// DatabaseEngine represents the main database engine
type DatabaseEngine struct {
	schemaManager      *schema.SchemaManager
	storageEngine      storage.StorageEngine
	queryProcessor     *query.QueryProcessor
	transactionManager *transaction.TransactionManager
	indexManager       *index.IndexManager
	concurrencyManager *concurrency.ConcurrencyManager
}

// NewDatabaseEngine creates a new database engine
func NewDatabaseEngine() *DatabaseEngine {
	return &DatabaseEngine{
		schemaManager:      schema.NewSchemaManager(),
		storageEngine:      storage.NewBTreeStorageEngine(),
		queryProcessor:     query.NewQueryProcessor(),
		transactionManager: transaction.NewTransactionManager(),
		indexManager:       index.NewIndexManager(),
		concurrencyManager: concurrency.NewConcurrencyManager(),
	}
}

// Initialize initializes the database engine
func (db *DatabaseEngine) Initialize() error {
	fmt.Println("🗄️ Initializing Database Engine...")
	
	// Initialize components
	// Storage engine doesn't need explicit initialization
	
	if err := db.queryProcessor.Initialize(db.storageEngine, db.schemaManager); err != nil {
		return fmt.Errorf("failed to initialize query processor: %w", err)
	}
	
	if err := db.transactionManager.Initialize(db.storageEngine); err != nil {
		return fmt.Errorf("failed to initialize transaction manager: %w", err)
	}
	
	if err := db.indexManager.Initialize(db.storageEngine, db.schemaManager); err != nil {
		return fmt.Errorf("failed to initialize index manager: %w", err)
	}
	
	if err := db.concurrencyManager.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize concurrency manager: %w", err)
	}
	
	fmt.Println("✅ Database Engine initialized successfully")
	return nil
}

// ExecuteSQL executes a SQL statement
func (db *DatabaseEngine) ExecuteSQL(sql string) (*query.Result, error) {
	fmt.Printf("🔍 Executing SQL: %s\n", sql)
	
	// Parse SQL
	parser := query.NewParser()
	statement, err := parser.Parse(sql)
	if err != nil {
		return nil, fmt.Errorf("SQL parsing failed: %w", err)
	}
	
	// Execute statement
	result, err := db.queryProcessor.Execute(statement)
	if err != nil {
		return nil, fmt.Errorf("SQL execution failed: %w", err)
	}
	
	return result, nil
}

// CreateTable creates a new table
func (db *DatabaseEngine) CreateTable(tableName string, columns []*schema.Column) error {
	fmt.Printf("📋 Creating table: %s\n", tableName)
	
	// Create table schema
	table := &schema.Table{
		Name:    tableName,
		Columns: columns,
	}
	
	// Create table in schema manager
	if err := db.schemaManager.CreateTable(table); err != nil {
		return fmt.Errorf("failed to create table schema: %w", err)
	}
	
	// Create table in storage engine
	if err := db.storageEngine.CreateTable(table); err != nil {
		return fmt.Errorf("failed to create table storage: %w", err)
	}
	
	fmt.Printf("✅ Table %s created successfully\n", tableName)
	return nil
}

// InsertRow inserts a row into a table
func (db *DatabaseEngine) InsertRow(tableName string, values map[string]types.Value) error {
	fmt.Printf("➕ Inserting row into table: %s\n", tableName)
	
	// Create row
	row := &storage.Row{
		Key:    db.generateRowKey(tableName, values),
		Values: values,
	}
	
	// Insert row
	if err := db.storageEngine.InsertRow(tableName, row); err != nil {
		return fmt.Errorf("failed to insert row: %w", err)
	}
	
	fmt.Printf("✅ Row inserted successfully\n")
	return nil
}

// SelectRows selects rows from a table
func (db *DatabaseEngine) SelectRows(tableName string, filter storage.Filter) ([]*storage.Row, error) {
	fmt.Printf("🔍 Selecting rows from table: %s\n", tableName)
	
	// Scan table
	iterator, err := db.storageEngine.ScanTable(tableName, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}
	defer iterator.Close()
	
	// Collect rows
	var rows []*storage.Row
	for {
		row, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next row: %w", err)
		}
		if row == nil {
			break
		}
		rows = append(rows, row)
	}
	
	fmt.Printf("✅ Found %d rows\n", len(rows))
	return rows, nil
}

// UpdateRow updates a row in a table
func (db *DatabaseEngine) UpdateRow(tableName string, key storage.Key, values map[string]types.Value) error {
	fmt.Printf("✏️ Updating row in table: %s\n", tableName)
	
	// Create updated row
	row := &storage.Row{
		Key:    key,
		Values: values,
	}
	
	// Update row
	if err := db.storageEngine.UpdateRow(tableName, key, row); err != nil {
		return fmt.Errorf("failed to update row: %w", err)
	}
	
	fmt.Printf("✅ Row updated successfully\n")
	return nil
}

// DeleteRow deletes a row from a table
func (db *DatabaseEngine) DeleteRow(tableName string, key storage.Key) error {
	fmt.Printf("🗑️ Deleting row from table: %s\n", tableName)
	
	// Delete row
	if err := db.storageEngine.DeleteRow(tableName, key); err != nil {
		return fmt.Errorf("failed to delete row: %w", err)
	}
	
	fmt.Printf("✅ Row deleted successfully\n")
	return nil
}

// CreateIndex creates an index on a table
func (db *DatabaseEngine) CreateIndex(tableName, indexName string, columns []string) error {
	fmt.Printf("📊 Creating index: %s on table: %s\n", indexName, tableName)
	
	// Create index
	if err := db.indexManager.CreateIndex(tableName, indexName, columns); err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}
	
	fmt.Printf("✅ Index %s created successfully\n", indexName)
	return nil
}

// GetTableInfo gets information about a table
func (db *DatabaseEngine) GetTableInfo(tableName string) (*schema.TableInfo, error) {
	return db.schemaManager.GetTableInfo(tableName)
}

// ListTables lists all tables
func (db *DatabaseEngine) ListTables() []*schema.Table {
	return db.schemaManager.ListTables()
}

// BeginTransaction begins a new transaction
func (db *DatabaseEngine) BeginTransaction() *transaction.Transaction {
	return db.transactionManager.BeginTransaction()
}

// CommitTransaction commits a transaction
func (db *DatabaseEngine) CommitTransaction(tx *transaction.Transaction) error {
	return db.transactionManager.CommitTransaction(tx)
}

// RollbackTransaction rolls back a transaction
func (db *DatabaseEngine) RollbackTransaction(tx *transaction.Transaction) error {
	return db.transactionManager.RollbackTransaction(tx)
}

// Vacuum performs vacuum operation
func (db *DatabaseEngine) Vacuum() error {
	fmt.Println("🧹 Performing vacuum operation...")
	
	if err := db.storageEngine.Vacuum(); err != nil {
		return fmt.Errorf("vacuum failed: %w", err)
	}
	
	fmt.Println("✅ Vacuum completed successfully")
	return nil
}

// Analyze performs analysis operation
func (db *DatabaseEngine) Analyze() error {
	fmt.Println("📊 Performing analysis operation...")
	
	if err := db.storageEngine.Analyze(); err != nil {
		return fmt.Errorf("analysis failed: %w", err)
	}
	
	fmt.Println("✅ Analysis completed successfully")
	return nil
}

// CheckIntegrity checks database integrity
func (db *DatabaseEngine) CheckIntegrity() error {
	fmt.Println("🔍 Checking database integrity...")
	
	if err := db.storageEngine.CheckIntegrity(); err != nil {
		return fmt.Errorf("integrity check failed: %w", err)
	}
	
	fmt.Println("✅ Database integrity check passed")
	return nil
}

// Helper methods

// generateRowKey generates a row key
func (db *DatabaseEngine) generateRowKey(tableName string, values map[string]types.Value) storage.Key {
	// Simple key generation - in real implementation, use primary key
	return &storage.SimpleKey{
		Value: values["id"], // Assume 'id' column exists
	}
}

// Demo functions

// runDemo runs a demonstration of the database engine
func (db *DatabaseEngine) runDemo() error {
	fmt.Println("\n🎯 Running Database Engine Demo")
	fmt.Println("================================")
	
	// Create users table
	usersTable := []*schema.Column{
		{Name: "id", Type: types.Int64Type, Nullable: false, Unique: true},
		{Name: "name", Type: types.VarcharType, Nullable: false},
		{Name: "email", Type: types.VarcharType, Nullable: false, Unique: true},
		{Name: "age", Type: types.Int32Type, Nullable: true},
		{Name: "created_at", Type: types.TimestampTypeVar, Nullable: false},
	}
	
	if err := db.CreateTable("users", usersTable); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	
	// Insert sample data
	sampleUsers := []map[string]types.Value{
		{
			"id":         types.IntValue(1),
			"name":       types.StringValue("Alice Johnson"),
			"email":      types.StringValue("alice@example.com"),
			"age":        types.IntValue(25),
			"created_at": types.DateTimeValue(time.Now()),
		},
		{
			"id":         types.IntValue(2),
			"name":       types.StringValue("Bob Smith"),
			"email":      types.StringValue("bob@example.com"),
			"age":        types.IntValue(30),
			"created_at": types.DateTimeValue(time.Now()),
		},
		{
			"id":         types.IntValue(3),
			"name":       types.StringValue("Charlie Brown"),
			"email":      types.StringValue("charlie@example.com"),
			"age":        types.IntValue(35),
			"created_at": types.DateTimeValue(time.Now()),
		},
	}
	
	for _, user := range sampleUsers {
		if err := db.InsertRow("users", user); err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
		}
	}
	
	// Create index on email
	if err := db.CreateIndex("users", "idx_users_email", []string{"email"}); err != nil {
		return fmt.Errorf("failed to create email index: %w", err)
	}
	
	// Select all users
	rows, err := db.SelectRows("users", nil)
	if err != nil {
		return fmt.Errorf("failed to select users: %w", err)
	}
	
	fmt.Println("\n📋 Users Table:")
	fmt.Println("ID | Name           | Email              | Age | Created At")
	fmt.Println("---|----------------|--------------------|-----|-------------------")
	for _, row := range rows {
		fmt.Printf("%-2d | %-14s | %-18s | %-3d | %s\n",
			row.Values["id"].(types.IntValue),
			row.Values["name"].(types.StringValue),
			row.Values["email"].(types.StringValue),
			row.Values["age"].(types.IntValue),
			row.Values["created_at"].(types.DateTimeValue).String(),
		)
	}
	
	// Update a user
	updateValues := map[string]types.Value{
		"id":         types.IntValue(1),
		"name":       types.StringValue("Alice Johnson Updated"),
		"email":      types.StringValue("alice.updated@example.com"),
		"age":        types.IntValue(26),
		"created_at": types.DateTimeValue(time.Now()),
	}
	
	if err := db.UpdateRow("users", &storage.SimpleKey{Value: types.IntValue(1)}, updateValues); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	
	// Select updated user
	filter := &storage.SimpleFilter{
		Column: "id",
		Value:  types.IntValue(1),
	}
	
	updatedRows, err := db.SelectRows("users", filter)
	if err != nil {
		return fmt.Errorf("failed to select updated user: %w", err)
	}
	
	fmt.Println("\n✏️ Updated User:")
	for _, row := range updatedRows {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n",
			row.Values["id"].(types.IntValue),
			row.Values["name"].(types.StringValue),
			row.Values["email"].(types.StringValue),
			row.Values["age"].(types.IntValue),
		)
	}
	
	// Delete a user
	if err := db.DeleteRow("users", &storage.SimpleKey{Value: types.IntValue(3)}); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	
	// Select remaining users
	remainingRows, err := db.SelectRows("users", nil)
	if err != nil {
		return fmt.Errorf("failed to select remaining users: %w", err)
	}
	
	fmt.Printf("\n🗑️ Remaining Users (%d):\n", len(remainingRows))
	for _, row := range remainingRows {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n",
			row.Values["id"].(types.IntValue),
			row.Values["name"].(types.StringValue),
			row.Values["email"].(types.StringValue),
		)
	}
	
	// Get table info
	tableInfo, err := db.GetTableInfo("users")
	if err != nil {
		return fmt.Errorf("failed to get table info: %w", err)
	}
	
	fmt.Println("\n📊 Table Information:")
	fmt.Printf("Name: %s\n", tableInfo.Name)
	fmt.Printf("Columns: %d\n", len(tableInfo.Columns))
	fmt.Printf("Row Count: %d\n", tableInfo.RowCount)
	
	// Run maintenance operations
	if err := db.Vacuum(); err != nil {
		return fmt.Errorf("vacuum failed: %w", err)
	}
	
	if err := db.Analyze(); err != nil {
		return fmt.Errorf("analysis failed: %w", err)
	}
	
	if err := db.CheckIntegrity(); err != nil {
		return fmt.Errorf("integrity check failed: %w", err)
	}
	
	fmt.Println("\n✅ Demo completed successfully!")
	return nil
}

func main() {
	// Parse command line arguments
	var (
		demo     = flag.Bool("demo", false, "Run demonstration")
		sql      = flag.String("sql", "", "Execute SQL statement")
		interactive = flag.Bool("interactive", false, "Start interactive mode")
	)
	flag.Parse()
	
	// Create database engine
	db := NewDatabaseEngine()
	
	// Initialize database engine
	if err := db.Initialize(); err != nil {
		fmt.Printf("❌ Failed to initialize database engine: %v\n", err)
		os.Exit(1)
	}
	
	// Handle different modes
	if *demo {
		// Run demonstration
		if err := db.runDemo(); err != nil {
			fmt.Printf("❌ Demo failed: %v\n", err)
			os.Exit(1)
		}
	} else if *sql != "" {
		// Execute SQL statement
		result, err := db.ExecuteSQL(*sql)
		if err != nil {
			fmt.Printf("❌ SQL execution failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ SQL executed successfully: %v\n", result)
	} else if *interactive {
		// Start interactive mode
		fmt.Println("🗄️ Database Engine Interactive Mode")
		fmt.Println("Type 'help' for commands, 'exit' to quit")
		
		// Simple interactive loop
		for {
			fmt.Print("db> ")
			var input string
			fmt.Scanln(&input)
			
			if input == "exit" {
				break
			} else if input == "help" {
				fmt.Println("Available commands:")
				fmt.Println("  help     - Show this help")
				fmt.Println("  demo     - Run demonstration")
				fmt.Println("  tables   - List all tables")
				fmt.Println("  vacuum   - Run vacuum operation")
				fmt.Println("  analyze  - Run analysis operation")
				fmt.Println("  check    - Check database integrity")
				fmt.Println("  exit     - Exit interactive mode")
			} else if input == "demo" {
				if err := db.runDemo(); err != nil {
					fmt.Printf("❌ Demo failed: %v\n", err)
				}
			} else if input == "tables" {
				tables := db.ListTables()
				fmt.Printf("Tables (%d):\n", len(tables))
				for _, table := range tables {
					fmt.Printf("  - %s\n", table.Name)
				}
			} else if input == "vacuum" {
				if err := db.Vacuum(); err != nil {
					fmt.Printf("❌ Vacuum failed: %v\n", err)
				}
			} else if input == "analyze" {
				if err := db.Analyze(); err != nil {
					fmt.Printf("❌ Analysis failed: %v\n", err)
				}
			} else if input == "check" {
				if err := db.CheckIntegrity(); err != nil {
					fmt.Printf("❌ Integrity check failed: %v\n", err)
				}
			} else if input != "" {
				fmt.Printf("Unknown command: %s\n", input)
			}
		}
	} else {
		// Show help
		fmt.Println("🗄️ Database Engine")
		fmt.Println("==================")
		fmt.Println("A production-quality database engine built in Go")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  go run main.go -demo                    # Run demonstration")
		fmt.Println("  go run main.go -sql='SELECT * FROM users' # Execute SQL")
		fmt.Println("  go run main.go -interactive            # Start interactive mode")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -demo        Run demonstration")
		fmt.Println("  -sql         Execute SQL statement")
		fmt.Println("  -interactive Start interactive mode")
	}
}
