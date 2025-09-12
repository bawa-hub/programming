package storage

import (
	"database-engine/schema"
	"database-engine/types"
	"fmt"
	"sync"
)

// StorageEngine represents the storage engine interface
type StorageEngine interface {
	// Table operations
	CreateTable(table *schema.Table) error
	DropTable(tableName string) error
	GetTable(tableName string) (*schema.Table, error)
	
	// Row operations
	InsertRow(tableName string, row *Row) error
	UpdateRow(tableName string, key Key, row *Row) error
	DeleteRow(tableName string, key Key) error
	GetRow(tableName string, key Key) (*Row, error)
	
	// Scan operations
	ScanTable(tableName string, filter Filter) (Iterator, error)
	ScanIndex(tableName, indexName string, key Key) (Iterator, error)
	
	// Transaction operations
	BeginTransaction() *Transaction
	CommitTransaction(tx *Transaction) error
	RollbackTransaction(tx *Transaction) error
	
	// Maintenance operations
	Vacuum() error
	Analyze() error
	CheckIntegrity() error
}

// BTreeStorageEngine implements a B+ Tree storage engine
type BTreeStorageEngine struct {
	tables      map[string]*BTreeTable
	transactions map[TransactionID]*Transaction
	mutex       sync.RWMutex
	nextTxID    TransactionID
}

// NewBTreeStorageEngine creates a new B+ Tree storage engine
func NewBTreeStorageEngine() *BTreeStorageEngine {
	return &BTreeStorageEngine{
		tables:       make(map[string]*BTreeTable),
		transactions: make(map[TransactionID]*Transaction),
		nextTxID:     1,
	}
}

// Initialize initializes the storage engine
func (se *BTreeStorageEngine) Initialize() error {
	// Initialize storage engine
	return nil
}

// BTreeTable represents a table in the B+ Tree storage engine
type BTreeTable struct {
	Schema    *schema.Table
	BTree     *BTree
	Indexes   map[string]*BTreeIndex
	RowCount  int64
	mutex     sync.RWMutex
}

// NewBTreeTable creates a new B+ Tree table
func NewBTreeTable(schema *schema.Table) *BTreeTable {
	return &BTreeTable{
		Schema:  schema,
		BTree:   NewBTree(),
		Indexes: make(map[string]*BTreeIndex),
	}
}

// Row represents a database row
type Row struct {
	Key    Key
	Values map[string]types.Value
}

// Key represents a row key
type Key interface {
	// Compare compares this key with another key
	Compare(other Key) int
	
	// Encode encodes the key to bytes
	Encode() []byte
	
	// Decode decodes bytes to a key
	Decode(data []byte) Key
}

// CompositeKey represents a composite key
type CompositeKey struct {
	Values []types.Value
}

func (ck *CompositeKey) Compare(other Key) int {
	otherCK := other.(*CompositeKey)
	
	// Compare each component
	for i, val := range ck.Values {
		if i >= len(otherCK.Values) {
			return 1
		}
		
		cmp := val.Type().Compare(val, otherCK.Values[i])
		if cmp != 0 {
			return cmp
		}
	}
	
	if len(ck.Values) < len(otherCK.Values) {
		return -1
	}
	
	return 0
}

func (ck *CompositeKey) Encode() []byte {
	// Simple encoding - in real implementation, use proper binary encoding
	data := make([]byte, 0)
	for _, val := range ck.Values {
		encoded := val.Type().Encode(val)
		data = append(data, encoded...)
	}
	return data
}

func (ck *CompositeKey) Decode(data []byte) Key {
	// Simple decoding - in real implementation, use proper binary decoding
	// This is a placeholder implementation
	return &CompositeKey{Values: []types.Value{}}
}

// SimpleKey represents a simple key
type SimpleKey struct {
	Value types.Value
}

func (sk *SimpleKey) Compare(other Key) int {
	otherSK := other.(*SimpleKey)
	return sk.Value.Type().Compare(sk.Value, otherSK.Value)
}

func (sk *SimpleKey) Encode() []byte {
	return sk.Value.Type().Encode(sk.Value)
}

func (sk *SimpleKey) Decode(data []byte) Key {
	// This is a placeholder - in real implementation, decode properly
	return &SimpleKey{Value: types.NullValue{}}
}

// Filter represents a row filter
type Filter interface {
	// Match returns true if the row matches the filter
	Match(row *Row) bool
}

// SimpleFilter represents a simple equality filter
type SimpleFilter struct {
	Column string
	Value  types.Value
}

func (f *SimpleFilter) Match(row *Row) bool {
	if val, exists := row.Values[f.Column]; exists {
		return val.Type().Compare(val, f.Value) == 0
	}
	return false
}

// Iterator represents a row iterator
type Iterator interface {
	// Next returns the next row, or nil if no more rows
	Next() (*Row, error)
	
	// Close closes the iterator
	Close() error
}

// BTreeIterator implements an iterator for B+ Tree
type BTreeIterator struct {
	btree   *BTree
	filter  Filter
	current *BTreeNode
	index   int
}

func (iter *BTreeIterator) Next() (*Row, error) {
	// Simple implementation - return nil for now
	// In a real implementation, this would iterate through the B+ Tree
	return nil, nil
}

func (iter *BTreeIterator) Close() error {
	// Cleanup resources
	iter.current = nil
	iter.index = 0
	return nil
}

// Transaction represents a database transaction
type Transaction struct {
	ID        TransactionID
	State     TransactionState
	Changes   []*Change
	StartTime int64
	EndTime   int64
}

// TransactionID represents a transaction ID
type TransactionID int64

// TransactionState represents the state of a transaction
type TransactionState int

const (
	Active TransactionState = iota
	Committed
	Aborted
)

// Change represents a change in a transaction
type Change struct {
	Type      ChangeType
	TableName string
	Key       Key
	OldRow    *Row
	NewRow    *Row
}

// ChangeType represents the type of change
type ChangeType int

const (
	Insert ChangeType = iota
	Update
	Delete
)

// StorageEngine implementation

// CreateTable creates a new table
func (se *BTreeStorageEngine) CreateTable(table *schema.Table) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	// Check if table already exists
	if _, exists := se.tables[table.Name]; exists {
		return fmt.Errorf("table %s already exists", table.Name)
	}
	
	// Create B+ Tree table
	btreeTable := NewBTreeTable(table)
	
	// Create indexes
	for _, index := range table.Indexes {
		btreeIndex := NewBTreeIndex(index)
		btreeTable.Indexes[index.Name] = btreeIndex
	}
	
	// Store table
	se.tables[table.Name] = btreeTable
	
	return nil
}

// DropTable drops a table
func (se *BTreeStorageEngine) DropTable(tableName string) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	// Check if table exists
	if _, exists := se.tables[tableName]; !exists {
		return fmt.Errorf("table %s does not exist", tableName)
	}
	
	// Remove table
	delete(se.tables, tableName)
	
	return nil
}

// GetTable gets a table by name
func (se *BTreeStorageEngine) GetTable(tableName string) (*schema.Table, error) {
	se.mutex.RLock()
	defer se.mutex.RUnlock()
	
	table, exists := se.tables[tableName]
	if !exists {
		return nil, fmt.Errorf("table %s does not exist", tableName)
	}
	
	return table.Schema, nil
}

// InsertRow inserts a new row
func (se *BTreeStorageEngine) InsertRow(tableName string, row *Row) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	// Get table
	table, exists := se.tables[tableName]
	if !exists {
		return fmt.Errorf("table %s does not exist", tableName)
	}
	
	// Validate row
	if err := se.validateRow(table.Schema, row); err != nil {
		return fmt.Errorf("row validation failed: %w", err)
	}
	
	// Insert into B+ Tree
	record := &BTreeRecord{
		Key:    row.Key,
		Values: row.Values,
	}
	
	if err := table.BTree.Insert(record); err != nil {
		return fmt.Errorf("failed to insert row: %w", err)
	}
	
	// Update indexes
	for _, index := range table.Indexes {
		indexKey := se.createIndexKey(row, index)
		if err := index.Insert(indexKey, row.Key); err != nil {
			return fmt.Errorf("failed to update index %s: %w", index.Name, err)
		}
	}
	
	// Update row count
	table.RowCount++
	
	return nil
}

// UpdateRow updates an existing row
func (se *BTreeStorageEngine) UpdateRow(tableName string, key Key, row *Row) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	// Get table
	table, exists := se.tables[tableName]
	if !exists {
		return fmt.Errorf("table %s does not exist", tableName)
	}
	
	// Get old row
	oldRecord, err := table.BTree.Search(key)
	if err != nil {
		return fmt.Errorf("row not found: %w", err)
	}
	
	oldRow := &Row{
		Key:    oldRecord.Key,
		Values: oldRecord.Values,
	}
	
	// Validate new row
	if err := se.validateRow(table.Schema, row); err != nil {
		return fmt.Errorf("row validation failed: %w", err)
	}
	
	// Update B+ Tree
	newRecord := &BTreeRecord{
		Key:    row.Key,
		Values: row.Values,
	}
	
	if err := table.BTree.Update(key, newRecord); err != nil {
		return fmt.Errorf("failed to update row: %w", err)
	}
	
	// Update indexes
	for _, index := range table.Indexes {
		// Remove old index entry
		oldIndexKey := se.createIndexKey(oldRow, index)
		index.Delete(oldIndexKey)
		
		// Add new index entry
		newIndexKey := se.createIndexKey(row, index)
		if err := index.Insert(newIndexKey, row.Key); err != nil {
			return fmt.Errorf("failed to update index %s: %w", index.Name, err)
		}
	}
	
	return nil
}

// DeleteRow deletes a row
func (se *BTreeStorageEngine) DeleteRow(tableName string, key Key) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	// Get table
	table, exists := se.tables[tableName]
	if !exists {
		return fmt.Errorf("table %s does not exist", tableName)
	}
	
	// Get row for index updates
	record, err := table.BTree.Search(key)
	if err != nil {
		return fmt.Errorf("row not found: %w", err)
	}
	
	row := &Row{
		Key:    record.Key,
		Values: record.Values,
	}
	
	// Delete from B+ Tree
	if err := table.BTree.Delete(key); err != nil {
		return fmt.Errorf("failed to delete row: %w", err)
	}
	
	// Update indexes
	for _, index := range table.Indexes {
		indexKey := se.createIndexKey(row, index)
		index.Delete(indexKey)
	}
	
	// Update row count
	table.RowCount--
	
	return nil
}

// GetRow gets a row by key
func (se *BTreeStorageEngine) GetRow(tableName string, key Key) (*Row, error) {
	se.mutex.RLock()
	defer se.mutex.RUnlock()
	
	// Get table
	table, exists := se.tables[tableName]
	if !exists {
		return nil, fmt.Errorf("table %s does not exist", tableName)
	}
	
	// Search B+ Tree
	record, err := table.BTree.Search(key)
	if err != nil {
		return nil, fmt.Errorf("row not found: %w", err)
	}
	
	return &Row{
		Key:    record.Key,
		Values: record.Values,
	}, nil
}

// ScanTable scans a table
func (se *BTreeStorageEngine) ScanTable(tableName string, filter Filter) (Iterator, error) {
	se.mutex.RLock()
	defer se.mutex.RUnlock()
	
	// Get table
	table, exists := se.tables[tableName]
	if !exists {
		return nil, fmt.Errorf("table %s does not exist", tableName)
	}
	
	return &BTreeIterator{
		btree:  table.BTree,
		filter: filter,
	}, nil
}

// ScanIndex scans an index
func (se *BTreeStorageEngine) ScanIndex(tableName, indexName string, key Key) (Iterator, error) {
	se.mutex.RLock()
	defer se.mutex.RUnlock()
	
	// Get table
	table, exists := se.tables[tableName]
	if !exists {
		return nil, fmt.Errorf("table %s does not exist", tableName)
	}
	
	// Get index
	index, exists := table.Indexes[indexName]
	if !exists {
		return nil, fmt.Errorf("index %s does not exist", indexName)
	}
	
	// Create index iterator
	return &BTreeIndexIterator{
		index: index,
		key:   key,
	}, nil
}

// BeginTransaction begins a new transaction
func (se *BTreeStorageEngine) BeginTransaction() *Transaction {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	tx := &Transaction{
		ID:        se.nextTxID,
		State:     Active,
		Changes:   make([]*Change, 0),
		StartTime: getCurrentTimestamp(),
	}
	
	se.nextTxID++
	se.transactions[tx.ID] = tx
	
	return tx
}

// CommitTransaction commits a transaction
func (se *BTreeStorageEngine) CommitTransaction(tx *Transaction) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	// Check transaction state
	if tx.State != Active {
		return fmt.Errorf("transaction %d is not active", tx.ID)
	}
	
	// Apply all changes
	for _, change := range tx.Changes {
		switch change.Type {
		case Insert:
			if err := se.InsertRow(change.TableName, change.NewRow); err != nil {
				return fmt.Errorf("failed to apply insert change: %w", err)
			}
		case Update:
			if err := se.UpdateRow(change.TableName, change.Key, change.NewRow); err != nil {
				return fmt.Errorf("failed to apply update change: %w", err)
			}
		case Delete:
			if err := se.DeleteRow(change.TableName, change.Key); err != nil {
				return fmt.Errorf("failed to apply delete change: %w", err)
			}
		}
	}
	
	// Mark transaction as committed
	tx.State = Committed
	tx.EndTime = getCurrentTimestamp()
	
	// Remove from active transactions
	delete(se.transactions, tx.ID)
	
	return nil
}

// RollbackTransaction rolls back a transaction
func (se *BTreeStorageEngine) RollbackTransaction(tx *Transaction) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	// Check transaction state
	if tx.State != Active {
		return fmt.Errorf("transaction %d is not active", tx.ID)
	}
	
	// Mark transaction as aborted
	tx.State = Aborted
	tx.EndTime = getCurrentTimestamp()
	
	// Remove from active transactions
	delete(se.transactions, tx.ID)
	
	return nil
}

// Vacuum performs vacuum operation
func (se *BTreeStorageEngine) Vacuum() error {
	se.mutex.Lock()
	defer se.mutex.Unlock()
	
	// Vacuum all tables
	for _, table := range se.tables {
		if err := table.BTree.Vacuum(); err != nil {
			return fmt.Errorf("failed to vacuum table %s: %w", table.Schema.Name, err)
		}
	}
	
	return nil
}

// Analyze performs analysis operation
func (se *BTreeStorageEngine) Analyze() error {
	se.mutex.RLock()
	defer se.mutex.RUnlock()
	
	// Analyze all tables
	for _, table := range se.tables {
		if err := table.BTree.Analyze(); err != nil {
			return fmt.Errorf("failed to analyze table %s: %w", table.Schema.Name, err)
		}
	}
	
	return nil
}

// CheckIntegrity checks database integrity
func (se *BTreeStorageEngine) CheckIntegrity() error {
	se.mutex.RLock()
	defer se.mutex.RUnlock()
	
	// Check all tables
	for _, table := range se.tables {
		if err := table.BTree.CheckIntegrity(); err != nil {
			return fmt.Errorf("integrity check failed for table %s: %w", table.Schema.Name, err)
		}
	}
	
	return nil
}

// Helper methods

// validateRow validates a row against table schema
func (se *BTreeStorageEngine) validateRow(tableSchema *schema.Table, row *Row) error {
	// Check if all required columns are present
	for _, col := range tableSchema.Columns {
		if !col.Nullable {
			if val, exists := row.Values[col.Name]; !exists || val.IsNull() {
				return fmt.Errorf("column %s cannot be null", col.Name)
			}
		}
	}
	
	// Validate column values
	for colName, val := range row.Values {
		// Find column definition
		var col *schema.Column
		for _, c := range tableSchema.Columns {
			if c.Name == colName {
				col = c
				break
			}
		}
		
		if col == nil {
			return fmt.Errorf("unknown column: %s", colName)
		}
		
		// Validate value type
		if err := col.Type.Validate(val); err != nil {
			return fmt.Errorf("invalid value for column %s: %w", colName, err)
		}
	}
	
	return nil
}

// createIndexKey creates an index key from a row
func (se *BTreeStorageEngine) createIndexKey(row *Row, index *BTreeIndex) Key {
	values := make([]types.Value, len(index.Columns))
	
	for i, colName := range index.Columns {
		if val, exists := row.Values[colName]; exists {
			values[i] = val
		} else {
			values[i] = types.NullValue{}
		}
	}
	
	return &CompositeKey{Values: values}
}

// getCurrentTimestamp returns the current timestamp
func getCurrentTimestamp() int64 {
	return 1234567890 // Placeholder - in real implementation, use time.Now().Unix()
}
