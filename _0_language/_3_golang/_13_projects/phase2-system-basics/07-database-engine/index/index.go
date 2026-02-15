package index

import (
	"database-engine/schema"
	"database-engine/storage"
	"database-engine/types"
	"fmt"
	"sync"
)

// IndexManager manages database indexes
type IndexManager struct {
	indexes map[string]*Index
	mutex   sync.RWMutex
}

// Index represents a database index
type Index struct {
	Name    string
	Table   string
	Columns []string
	Unique  bool
	Type    IndexType
	BTree   *storage.BTree
}

// IndexType represents the type of index
type IndexType int

const (
	BTreeIndexType IndexType = iota
	HashIndexType
	CompositeIndexType
)

// NewIndexManager creates a new index manager
func NewIndexManager() *IndexManager {
	return &IndexManager{
		indexes: make(map[string]*Index),
	}
}

// Initialize initializes the index manager
func (im *IndexManager) Initialize(se storage.StorageEngine, sm *schema.SchemaManager) error {
	// Initialize index manager
	return nil
}

// CreateIndex creates a new index
func (im *IndexManager) CreateIndex(tableName, indexName string, columns []string) error {
	im.mutex.Lock()
	defer im.mutex.Unlock()
	
	// Check if index already exists
	if _, exists := im.indexes[indexName]; exists {
		return fmt.Errorf("index %s already exists", indexName)
	}
	
	// Create index
	index := &Index{
		Name:    indexName,
		Table:   tableName,
		Columns: columns,
		Unique:  false, // Default to non-unique
		Type:    BTreeIndexType,
		BTree:   storage.NewBTree(),
	}
	
	im.indexes[indexName] = index
	
	return nil
}

// DropIndex drops an index
func (im *IndexManager) DropIndex(indexName string) error {
	im.mutex.Lock()
	defer im.mutex.Unlock()
	
	// Check if index exists
	if _, exists := im.indexes[indexName]; !exists {
		return fmt.Errorf("index %s does not exist", indexName)
	}
	
	// Remove index
	delete(im.indexes, indexName)
	
	return nil
}

// GetIndex gets an index by name
func (im *IndexManager) GetIndex(indexName string) (*Index, error) {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	
	index, exists := im.indexes[indexName]
	if !exists {
		return nil, fmt.Errorf("index %s does not exist", indexName)
	}
	
	return index, nil
}

// ListIndexes lists all indexes
func (im *IndexManager) ListIndexes() []*Index {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	
	indexes := make([]*Index, 0, len(im.indexes))
	for _, index := range im.indexes {
		indexes = append(indexes, index)
	}
	
	return indexes
}

// ListIndexesForTable lists indexes for a specific table
func (im *IndexManager) ListIndexesForTable(tableName string) []*Index {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	
	var indexes []*Index
	for _, index := range im.indexes {
		if index.Table == tableName {
			indexes = append(indexes, index)
		}
	}
	
	return indexes
}

// InsertIntoIndex inserts a key-value pair into an index
func (im *IndexManager) InsertIntoIndex(indexName string, key storage.Key, value storage.Key) error {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	
	index, exists := im.indexes[indexName]
	if !exists {
		return fmt.Errorf("index %s does not exist", indexName)
	}
	
	// Insert into B+ Tree
	record := &storage.BTreeRecord{
		Key: key,
		Values: map[string]types.Value{
			"value": value.(*storage.SimpleKey).Value,
		},
	}
	
	return index.BTree.Insert(record)
}

// DeleteFromIndex deletes a key from an index
func (im *IndexManager) DeleteFromIndex(indexName string, key storage.Key) error {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	
	index, exists := im.indexes[indexName]
	if !exists {
		return fmt.Errorf("index %s does not exist", indexName)
	}
	
	return index.BTree.Delete(key)
}

// SearchIndex searches an index
func (im *IndexManager) SearchIndex(indexName string, key storage.Key) (storage.Key, error) {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	
	index, exists := im.indexes[indexName]
	if !exists {
		return nil, fmt.Errorf("index %s does not exist", indexName)
	}
	
	record, err := index.BTree.Search(key)
	if err != nil {
		return nil, err
	}
	
	// Extract value from record
	if val, exists := record.Values["value"]; exists {
		return &storage.SimpleKey{Value: val}, nil
	}
	
	return nil, fmt.Errorf("value not found in index record")
}

// GetIndexCount gets the number of indexes
func (im *IndexManager) GetIndexCount() int {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	
	return len(im.indexes)
}

// GetIndexInfo gets information about an index
func (im *IndexManager) GetIndexInfo(indexName string) (*IndexInfo, error) {
	im.mutex.RLock()
	defer im.mutex.RUnlock()
	
	index, exists := im.indexes[indexName]
	if !exists {
		return nil, fmt.Errorf("index %s does not exist", indexName)
	}
	
	return &IndexInfo{
		Name:    index.Name,
		Table:   index.Table,
		Columns: index.Columns,
		Unique:  index.Unique,
		Type:    index.Type.String(),
	}, nil
}

// IndexInfo represents index information
type IndexInfo struct {
	Name    string   `json:"name"`
	Table   string   `json:"table"`
	Columns []string `json:"columns"`
	Unique  bool     `json:"unique"`
	Type    string   `json:"type"`
}

// String method for IndexType
func (it IndexType) String() string {
	switch it {
	case BTreeIndexType:
		return "BTREE"
	case HashIndexType:
		return "HASH"
	case CompositeIndexType:
		return "COMPOSITE"
	default:
		return "UNKNOWN"
	}
}
