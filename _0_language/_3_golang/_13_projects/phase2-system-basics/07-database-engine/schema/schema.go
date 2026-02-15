package schema

import (
	"database-engine/types"
	"fmt"
	"sync"
)

// SchemaManager manages database schema
type SchemaManager struct {
	tables map[string]*Table
	mutex  sync.RWMutex
}

// NewSchemaManager creates a new schema manager
func NewSchemaManager() *SchemaManager {
	return &SchemaManager{
		tables: make(map[string]*Table),
	}
}

// Table represents a database table
type Table struct {
	Name        string
	Columns     []*Column
	PrimaryKey  []string
	Indexes     []*Index
	Constraints []*Constraint
	RowCount    int64
	CreatedAt   int64
	UpdatedAt   int64
}

// Column represents a table column
type Column struct {
	Name     string
	Type     types.Type
	Nullable bool
	Default  types.Value
	Unique   bool
	Index    bool
}

// Index represents a table index
type Index struct {
	Name    string
	Columns []string
	Unique  bool
	Type    IndexType
}

// IndexType represents the type of index
type IndexType int

const (
	BTreeIndex IndexType = iota
	HashIndex
	CompositeIndex
)

// Constraint represents a table constraint
type Constraint struct {
	Name       string
	Type       ConstraintType
	Columns    []string
	Expression string
}

// ConstraintType represents the type of constraint
type ConstraintType int

const (
	PrimaryKeyConstraint ConstraintType = iota
	ForeignKeyConstraint
	UniqueConstraint
	CheckConstraint
	NotNullConstraint
)

// CreateTable creates a new table
func (sm *SchemaManager) CreateTable(table *Table) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	// Validate table
	if err := sm.validateTable(table); err != nil {
		return fmt.Errorf("table validation failed: %w", err)
	}
	
	// Check if table already exists
	if _, exists := sm.tables[table.Name]; exists {
		return fmt.Errorf("table %s already exists", table.Name)
	}
	
	// Set metadata
	table.CreatedAt = getCurrentTimestamp()
	table.UpdatedAt = table.CreatedAt
	
	// Store table
	sm.tables[table.Name] = table
	
	return nil
}

// DropTable drops a table
func (sm *SchemaManager) DropTable(tableName string) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	// Check if table exists
	if _, exists := sm.tables[tableName]; !exists {
		return fmt.Errorf("table %s does not exist", tableName)
	}
	
	// Remove table
	delete(sm.tables, tableName)
	
	return nil
}

// GetTable gets a table by name
func (sm *SchemaManager) GetTable(tableName string) (*Table, error) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	table, exists := sm.tables[tableName]
	if !exists {
		return nil, fmt.Errorf("table %s does not exist", tableName)
	}
	
	return table, nil
}

// ListTables lists all tables
func (sm *SchemaManager) ListTables() []*Table {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	tables := make([]*Table, 0, len(sm.tables))
	for _, table := range sm.tables {
		tables = append(tables, table)
	}
	
	return tables
}

// AlterTable alters a table
func (sm *SchemaManager) AlterTable(tableName string, changes []AlterChange) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	// Get table
	table, exists := sm.tables[tableName]
	if !exists {
		return fmt.Errorf("table %s does not exist", tableName)
	}
	
	// Apply changes
	for _, change := range changes {
		if err := sm.applyAlterChange(table, change); err != nil {
			return fmt.Errorf("failed to apply change: %w", err)
		}
	}
	
	// Update timestamp
	table.UpdatedAt = getCurrentTimestamp()
	
	return nil
}

// AlterChange represents a table alteration
type AlterChange struct {
	Type        AlterType
	Column      *Column
	ColumnName  string
	Index       *Index
	IndexName   string
	Constraint  *Constraint
	ConstraintName string
}

// AlterType represents the type of alteration
type AlterType int

const (
	AddColumn AlterType = iota
	DropColumn
	ModifyColumn
	AddIndex
	DropIndex
	AddConstraint
	DropConstraint
)

// applyAlterChange applies a single alteration to a table
func (sm *SchemaManager) applyAlterChange(table *Table, change AlterChange) error {
	switch change.Type {
	case AddColumn:
		// Check if column already exists
		for _, col := range table.Columns {
			if col.Name == change.Column.Name {
				return fmt.Errorf("column %s already exists", change.Column.Name)
			}
		}
		table.Columns = append(table.Columns, change.Column)
		
	case DropColumn:
		// Find and remove column
		found := false
		for i, col := range table.Columns {
			if col.Name == change.ColumnName {
				table.Columns = append(table.Columns[:i], table.Columns[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("column %s does not exist", change.ColumnName)
		}
		
	case ModifyColumn:
		// Find and modify column
		found := false
		for i, col := range table.Columns {
			if col.Name == change.Column.Name {
				table.Columns[i] = change.Column
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("column %s does not exist", change.Column.Name)
		}
		
	case AddIndex:
		// Check if index already exists
		for _, idx := range table.Indexes {
			if idx.Name == change.Index.Name {
				return fmt.Errorf("index %s already exists", change.Index.Name)
			}
		}
		table.Indexes = append(table.Indexes, change.Index)
		
	case DropIndex:
		// Find and remove index
		found := false
		for i, idx := range table.Indexes {
			if idx.Name == change.IndexName {
				table.Indexes = append(table.Indexes[:i], table.Indexes[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("index %s does not exist", change.IndexName)
		}
		
	case AddConstraint:
		// Check if constraint already exists
		for _, constraint := range table.Constraints {
			if constraint.Name == change.Constraint.Name {
				return fmt.Errorf("constraint %s already exists", change.Constraint.Name)
			}
		}
		table.Constraints = append(table.Constraints, change.Constraint)
		
	case DropConstraint:
		// Find and remove constraint
		found := false
		for i, constraint := range table.Constraints {
			if constraint.Name == change.ConstraintName {
				table.Constraints = append(table.Constraints[:i], table.Constraints[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("constraint %s does not exist", change.ConstraintName)
		}
		
	default:
		return fmt.Errorf("unknown alter type: %v", change.Type)
	}
	
	return nil
}

// validateTable validates a table definition
func (sm *SchemaManager) validateTable(table *Table) error {
	// Check table name
	if table.Name == "" {
		return fmt.Errorf("table name cannot be empty")
	}
	
	// Check columns
	if len(table.Columns) == 0 {
		return fmt.Errorf("table must have at least one column")
	}
	
	// Validate columns
	columnNames := make(map[string]bool)
	for i, col := range table.Columns {
		if err := sm.validateColumn(col, i); err != nil {
			return fmt.Errorf("column %d validation failed: %w", i, err)
		}
		
		// Check for duplicate column names
		if columnNames[col.Name] {
			return fmt.Errorf("duplicate column name: %s", col.Name)
		}
		columnNames[col.Name] = true
	}
	
	// Validate primary key
	if err := sm.validatePrimaryKey(table); err != nil {
		return fmt.Errorf("primary key validation failed: %w", err)
	}
	
	// Validate indexes
	for i, idx := range table.Indexes {
		if err := sm.validateIndex(idx, table); err != nil {
			return fmt.Errorf("index %d validation failed: %w", i, err)
		}
	}
	
	// Validate constraints
	for i, constraint := range table.Constraints {
		if err := sm.validateConstraint(constraint, table); err != nil {
			return fmt.Errorf("constraint %d validation failed: %w", i, err)
		}
	}
	
	return nil
}

// validateColumn validates a column definition
func (sm *SchemaManager) validateColumn(col *Column, index int) error {
	// Check column name
	if col.Name == "" {
		return fmt.Errorf("column name cannot be empty")
	}
	
	// Check column type
	if col.Type == nil {
		return fmt.Errorf("column type cannot be nil")
	}
	
	// Validate default value
	if col.Default != nil && !col.Default.IsNull() {
		if err := col.Type.Validate(col.Default); err != nil {
			return fmt.Errorf("default value validation failed: %w", err)
		}
	}
	
	return nil
}

// validatePrimaryKey validates the primary key
func (sm *SchemaManager) validatePrimaryKey(table *Table) error {
	if len(table.PrimaryKey) == 0 {
		return nil // No primary key is allowed
	}
	
	// Check if all primary key columns exist
	columnNames := make(map[string]bool)
	for _, col := range table.Columns {
		columnNames[col.Name] = true
	}
	
	for _, pkCol := range table.PrimaryKey {
		if !columnNames[pkCol] {
			return fmt.Errorf("primary key column %s does not exist", pkCol)
		}
	}
	
	return nil
}

// validateIndex validates an index definition
func (sm *SchemaManager) validateIndex(idx *Index, table *Table) error {
	// Check index name
	if idx.Name == "" {
		return fmt.Errorf("index name cannot be empty")
	}
	
	// Check columns
	if len(idx.Columns) == 0 {
		return fmt.Errorf("index must have at least one column")
	}
	
	// Check if all index columns exist
	columnNames := make(map[string]bool)
	for _, col := range table.Columns {
		columnNames[col.Name] = true
	}
	
	for _, idxCol := range idx.Columns {
		if !columnNames[idxCol] {
			return fmt.Errorf("index column %s does not exist", idxCol)
		}
	}
	
	return nil
}

// validateConstraint validates a constraint definition
func (sm *SchemaManager) validateConstraint(constraint *Constraint, table *Table) error {
	// Check constraint name
	if constraint.Name == "" {
		return fmt.Errorf("constraint name cannot be empty")
	}
	
	// Check columns
	if len(constraint.Columns) == 0 {
		return fmt.Errorf("constraint must have at least one column")
	}
	
	// Check if all constraint columns exist
	columnNames := make(map[string]bool)
	for _, col := range table.Columns {
		columnNames[col.Name] = true
	}
	
	for _, constraintCol := range constraint.Columns {
		if !columnNames[constraintCol] {
			return fmt.Errorf("constraint column %s does not exist", constraintCol)
		}
	}
	
	return nil
}

// GetColumn gets a column by name
func (sm *SchemaManager) GetColumn(tableName, columnName string) (*Column, error) {
	table, err := sm.GetTable(tableName)
	if err != nil {
		return nil, err
	}
	
	for _, col := range table.Columns {
		if col.Name == columnName {
			return col, nil
		}
	}
	
	return nil, fmt.Errorf("column %s does not exist in table %s", columnName, tableName)
}

// GetIndex gets an index by name
func (sm *SchemaManager) GetIndex(tableName, indexName string) (*Index, error) {
	table, err := sm.GetTable(tableName)
	if err != nil {
		return nil, err
	}
	
	for _, idx := range table.Indexes {
		if idx.Name == indexName {
			return idx, nil
		}
	}
	
	return nil, fmt.Errorf("index %s does not exist in table %s", indexName, tableName)
}

// GetConstraint gets a constraint by name
func (sm *SchemaManager) GetConstraint(tableName, constraintName string) (*Constraint, error) {
	table, err := sm.GetTable(tableName)
	if err != nil {
		return nil, err
	}
	
	for _, constraint := range table.Constraints {
		if constraint.Name == constraintName {
			return constraint, nil
		}
	}
	
	return nil, fmt.Errorf("constraint %s does not exist in table %s", constraintName, tableName)
}

// UpdateRowCount updates the row count for a table
func (sm *SchemaManager) UpdateRowCount(tableName string, count int64) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	table, exists := sm.tables[tableName]
	if !exists {
		return fmt.Errorf("table %s does not exist", tableName)
	}
	
	table.RowCount = count
	table.UpdatedAt = getCurrentTimestamp()
	
	return nil
}

// GetRowCount gets the row count for a table
func (sm *SchemaManager) GetRowCount(tableName string) (int64, error) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	table, exists := sm.tables[tableName]
	if !exists {
		return 0, fmt.Errorf("table %s does not exist", tableName)
	}
	
	return table.RowCount, nil
}

// Helper functions

// getCurrentTimestamp returns the current timestamp
func getCurrentTimestamp() int64 {
	return 1234567890 // Placeholder - in real implementation, use time.Now().Unix()
}

// TableInfo represents table information for display
type TableInfo struct {
	Name        string
	Columns     []ColumnInfo
	PrimaryKey  []string
	Indexes     []IndexInfo
	Constraints []ConstraintInfo
	RowCount    int64
	CreatedAt   int64
	UpdatedAt   int64
}

// ColumnInfo represents column information for display
type ColumnInfo struct {
	Name     string
	Type     string
	Nullable bool
	Default  string
	Unique   bool
	Index    bool
}

// IndexInfo represents index information for display
type IndexInfo struct {
	Name    string
	Columns []string
	Unique  bool
	Type    string
}

// ConstraintInfo represents constraint information for display
type ConstraintInfo struct {
	Name       string
	Type       string
	Columns    []string
	Expression string
}

// GetTableInfo gets table information for display
func (sm *SchemaManager) GetTableInfo(tableName string) (*TableInfo, error) {
	table, err := sm.GetTable(tableName)
	if err != nil {
		return nil, err
	}
	
	if table == nil {
		return nil, fmt.Errorf("table %s not found", tableName)
	}
	
	// Convert columns
	columns := make([]ColumnInfo, len(table.Columns))
	for i, col := range table.Columns {
		defaultValue := ""
		if col.Default != nil {
			defaultValue = col.Default.String()
		}
		
		columns[i] = ColumnInfo{
			Name:     col.Name,
			Type:     col.Type.String(),
			Nullable: col.Nullable,
			Default:  defaultValue,
			Unique:   col.Unique,
			Index:    col.Index,
		}
	}
	
	// Convert indexes
	indexes := make([]IndexInfo, len(table.Indexes))
	for i, idx := range table.Indexes {
		indexes[i] = IndexInfo{
			Name:    idx.Name,
			Columns: idx.Columns,
			Unique:  idx.Unique,
			Type:    idx.Type.String(),
		}
	}
	
	// Convert constraints
	constraints := make([]ConstraintInfo, len(table.Constraints))
	for i, constraint := range table.Constraints {
		constraints[i] = ConstraintInfo{
			Name:       constraint.Name,
			Type:       constraint.Type.String(),
			Columns:    constraint.Columns,
			Expression: constraint.Expression,
		}
	}
	
	return &TableInfo{
		Name:        table.Name,
		Columns:     columns,
		PrimaryKey:  table.PrimaryKey,
		Indexes:     indexes,
		Constraints: constraints,
		RowCount:    table.RowCount,
		CreatedAt:   table.CreatedAt,
		UpdatedAt:   table.UpdatedAt,
	}, nil
}

// String methods for display

func (it IndexType) String() string {
	switch it {
	case BTreeIndex:
		return "BTREE"
	case HashIndex:
		return "HASH"
	case CompositeIndex:
		return "COMPOSITE"
	default:
		return "UNKNOWN"
	}
}

func (ct ConstraintType) String() string {
	switch ct {
	case PrimaryKeyConstraint:
		return "PRIMARY KEY"
	case ForeignKeyConstraint:
		return "FOREIGN KEY"
	case UniqueConstraint:
		return "UNIQUE"
	case CheckConstraint:
		return "CHECK"
	case NotNullConstraint:
		return "NOT NULL"
	default:
		return "UNKNOWN"
	}
}
