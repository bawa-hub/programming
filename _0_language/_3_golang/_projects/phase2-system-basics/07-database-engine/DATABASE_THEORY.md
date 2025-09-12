# Database Engine Theory - Complete Guide 🗄️

## 📚 **Table of Contents**

1. [Introduction to Database Engines](#introduction-to-database-engines)
2. [Storage Engines](#storage-engines)
3. [Query Processing](#query-processing)
4. [Transaction Management](#transaction-management)
5. [Indexing Systems](#indexing-systems)
6. [Concurrency Control](#concurrency-control)
7. [Data Types and Schema](#data-types-and-schema)
8. [Performance Optimization](#performance-optimization)
9. [Recovery and Durability](#recovery-and-durability)
10. [Real-World Examples](#real-world-examples)

---

## 1. Introduction to Database Engines

### **What is a Database Engine?**

A database engine (also called a storage engine) is the underlying software component that a database management system (DBMS) uses to create, read, update, and delete (CRUD) data from a database. It's the heart of any database system.

### **Key Components of a Database Engine:**

```
┌─────────────────────────────────────────────────────────────┐
│                    Database Engine                          │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐   │
│  │   Storage   │ │   Query     │ │   Transaction       │   │
│  │   Engine    │ │  Processor  │ │    Manager          │   │
│  └─────────────┘ └─────────────┘ └─────────────────────┘   │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐   │
│  │   Index     │ │Concurrency  │ │    Recovery         │   │
│  │  Manager    │ │  Control    │ │    Manager          │   │
│  └─────────────┘ └─────────────┘ └─────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

### **Database Engine Responsibilities:**

1. **Data Storage**: Efficiently store data on disk
2. **Data Retrieval**: Quickly find and retrieve data
3. **Data Integrity**: Ensure data consistency and validity
4. **Concurrency**: Handle multiple users simultaneously
5. **Recovery**: Recover from crashes and failures
6. **Performance**: Optimize for speed and efficiency

---

## 2. Storage Engines

### **2.1 B+ Tree Storage Engine**

The B+ Tree is the most common storage engine used in databases. It provides excellent performance for both read and write operations.

#### **B+ Tree Structure:**

```
                    [Internal Node]
                   /      |      \
              [Leaf]   [Leaf]   [Leaf]
              [Data]   [Data]   [Data]
```

#### **Key Properties:**

- **Balanced Tree**: All leaf nodes at same level
- **Sorted Keys**: Keys are sorted within each node
- **Leaf Nodes**: Contain actual data records
- **Internal Nodes**: Contain only keys and pointers
- **Fanout**: High fanout reduces tree height

#### **Operations:**

**Search (O(log n)):**
```go
func (bt *BTree) Search(key Key) (*Record, error) {
    node := bt.root
    for !node.isLeaf {
        // Find appropriate child
        childIndex := node.findChild(key)
        node = node.children[childIndex]
    }
    // Search within leaf node
    return node.findRecord(key)
}
```

**Insert (O(log n)):**
```go
func (bt *BTree) Insert(key Key, value Value) error {
    // Find insertion point
    leaf := bt.findLeaf(key)
    
    // Insert if space available
    if leaf.hasSpace() {
        leaf.insert(key, value)
    } else {
        // Split leaf and propagate up
        bt.splitLeaf(leaf, key, value)
    }
}
```

**Delete (O(log n)):**
```go
func (bt *BTree) Delete(key Key) error {
    leaf := bt.findLeaf(key)
    
    // Delete record
    if leaf.delete(key) {
        // Check if underflow
        if leaf.isUnderflow() {
            bt.rebalance(leaf)
        }
    }
}
```

### **2.2 Page Management**

Databases organize data into fixed-size pages for efficient disk I/O.

#### **Page Structure:**

```
┌─────────────────────────────────────────────────────────────┐
│                    Page Header (32 bytes)                   │
├─────────────────────────────────────────────────────────────┤
│                    Page Data (4096 bytes)                   │
├─────────────────────────────────────────────────────────────┤
│                    Page Footer (32 bytes)                   │
└─────────────────────────────────────────────────────────────┘
```

#### **Page Types:**

1. **Data Pages**: Store actual records
2. **Index Pages**: Store B+ Tree nodes
3. **Free Space Pages**: Track available space
4. **System Pages**: Store metadata

#### **Page Management Operations:**

```go
type Page struct {
    ID       PageID
    Data     []byte
    Dirty    bool
    PinCount int
    LSN      LogSequenceNumber
}

func (pm *PageManager) ReadPage(pageID PageID) (*Page, error) {
    // Check buffer pool first
    if page := pm.bufferPool.Get(pageID); page != nil {
        return page, nil
    }
    
    // Read from disk
    data := make([]byte, PAGE_SIZE)
    if err := pm.disk.Read(pageID, data); err != nil {
        return nil, err
    }
    
    page := &Page{
        ID:   pageID,
        Data: data,
    }
    
    // Add to buffer pool
    pm.bufferPool.Put(page)
    return page, nil
}
```

### **2.3 Buffer Pool**

The buffer pool caches frequently accessed pages in memory to reduce disk I/O.

#### **Buffer Pool Structure:**

```
┌─────────────────────────────────────────────────────────────┐
│                    Buffer Pool                              │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐   │
│  │   Page 1    │ │   Page 2    │ │      Page 3         │   │
│  │  (Pinned)   │ │ (Unpinned)  │ │    (Dirty)          │   │
│  └─────────────┘ └─────────────┘ └─────────────────────┘   │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐   │
│  │   Page 4    │ │   Page 5    │ │      Page 6         │   │
│  │ (Unpinned)  │ │  (Pinned)   │ │   (Unpinned)        │   │
│  └─────────────┘ └─────────────┘ └─────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

#### **Buffer Pool Operations:**

```go
type BufferPool struct {
    pages    map[PageID]*Page
    capacity int
    lru      *LRUList
}

func (bp *BufferPool) Get(pageID PageID) *Page {
    if page, exists := bp.pages[pageID]; exists {
        // Update LRU
        bp.lru.MoveToFront(page)
        return page
    }
    return nil
}

func (bp *BufferPool) Put(page *Page) {
    if len(bp.pages) >= bp.capacity {
        // Evict least recently used page
        evicted := bp.lru.RemoveBack()
        if evicted.Dirty {
            bp.flushPage(evicted)
        }
        delete(bp.pages, evicted.ID)
    }
    
    bp.pages[page.ID] = page
    bp.lru.AddToFront(page)
}
```

### **2.4 Write-Ahead Logging (WAL)**

WAL ensures durability by writing changes to a log before applying them to the database.

#### **WAL Structure:**

```
┌─────────────────────────────────────────────────────────────┐
│                    WAL File                                 │
├─────────────────────────────────────────────────────────────┤
│  LSN 1: INSERT INTO users (id, name) VALUES (1, 'Alice')   │
├─────────────────────────────────────────────────────────────┤
│  LSN 2: UPDATE users SET name='Bob' WHERE id=1             │
├─────────────────────────────────────────────────────────────┤
│  LSN 3: DELETE FROM users WHERE id=1                       │
├─────────────────────────────────────────────────────────────┤
│  LSN 4: COMMIT                                              │
└─────────────────────────────────────────────────────────────┘
```

#### **WAL Operations:**

```go
type WALEntry struct {
    LSN     LogSequenceNumber
    Type    LogEntryType
    Data    []byte
    Checksum uint32
}

func (wal *WAL) WriteEntry(entry *WALEntry) error {
    // Write to log file
    if err := wal.logFile.Write(entry); err != nil {
        return err
    }
    
    // Force to disk
    if err := wal.logFile.Sync(); err != nil {
        return err
    }
    
    return nil
}

func (wal *WAL) Recover() error {
    // Read all log entries
    entries, err := wal.logFile.ReadAll()
    if err != nil {
        return err
    }
    
    // Apply entries to database
    for _, entry := range entries {
        if err := wal.applyEntry(entry); err != nil {
            return err
        }
    }
    
    return nil
}
```

---

## 3. Query Processing

### **3.1 SQL Parser**

The parser converts SQL text into an Abstract Syntax Tree (AST).

#### **AST Structure:**

```
                    SELECT
                   /        \
              Projection    FROM
              /      \        |
            *        users   WHERE
                           /      \
                        id = 1   AND
                                 |
                              name = 'Alice'
```

#### **Parser Implementation:**

```go
type SelectStatement struct {
    Columns    []Column
    From       Table
    Where      Expression
    GroupBy    []Column
    Having     Expression
    OrderBy    []OrderByClause
    Limit      int
    Offset     int
}

func (p *Parser) ParseSelect(query string) (*SelectStatement, error) {
    tokens := p.tokenize(query)
    ast, err := p.parseSelectStatement(tokens)
    if err != nil {
        return nil, err
    }
    return ast, nil
}
```

### **3.2 Query Planner**

The planner optimizes queries by choosing the best execution plan.

#### **Query Planning Process:**

1. **Parse Query**: Convert SQL to AST
2. **Analyze Query**: Determine required operations
3. **Generate Plans**: Create multiple execution plans
4. **Cost Estimation**: Estimate cost of each plan
5. **Select Plan**: Choose the cheapest plan

#### **Cost Estimation:**

```go
type CostEstimator struct {
    statistics *Statistics
}

func (ce *CostEstimator) EstimateCost(plan *ExecutionPlan) float64 {
    cost := 0.0
    
    for _, op := range plan.Operations {
        switch op.Type {
        case TableScan:
            cost += ce.estimateTableScanCost(op)
        case IndexScan:
            cost += ce.estimateIndexScanCost(op)
        case NestedLoopJoin:
            cost += ce.estimateNestedLoopJoinCost(op)
        case HashJoin:
            cost += ce.estimateHashJoinCost(op)
        }
    }
    
    return cost
}
```

### **3.3 Query Executor**

The executor executes the optimized query plan.

#### **Execution Operators:**

```go
type ExecutionPlan struct {
    Operations []Operator
}

type Operator interface {
    Execute() (Iterator, error)
    Next() (*Record, error)
    Close() error
}

// Table Scan Operator
type TableScan struct {
    table   *Table
    filter  Expression
    iterator Iterator
}

func (ts *TableScan) Execute() (Iterator, error) {
    ts.iterator = ts.table.Scan()
    return ts.iterator, nil
}

func (ts *TableScan) Next() (*Record, error) {
    for {
        record, err := ts.iterator.Next()
        if err != nil {
            return nil, err
        }
        
        if ts.filter == nil || ts.filter.Evaluate(record) {
            return record, nil
        }
    }
}
```

### **3.4 Join Algorithms**

#### **Nested Loop Join:**

```go
func (nlj *NestedLoopJoin) Execute() (Iterator, error) {
    outer := nlj.outer.Execute()
    inner := nlj.inner.Execute()
    
    return &NestedLoopIterator{
        outer: outer,
        inner: inner,
        joinCondition: nlj.condition,
    }, nil
}

func (nli *NestedLoopIterator) Next() (*Record, error) {
    for {
        outerRecord, err := nli.outer.Next()
        if err != nil {
            return nil, err
        }
        
        for {
            innerRecord, err := nli.inner.Next()
            if err != nil {
                break
            }
            
            if nli.joinCondition.Evaluate(outerRecord, innerRecord) {
                return nli.combineRecords(outerRecord, innerRecord), nil
            }
        }
        
        nli.inner.Reset()
    }
}
```

#### **Hash Join:**

```go
func (hj *HashJoin) Execute() (Iterator, error) {
    // Build hash table from inner relation
    hashTable := make(map[Key][]*Record)
    
    inner := hj.inner.Execute()
    for {
        record, err := inner.Next()
        if err != nil {
            break
        }
        
        key := hj.innerKey(record)
        hashTable[key] = append(hashTable[key], record)
    }
    
    return &HashJoinIterator{
        outer:     hj.outer.Execute(),
        hashTable: hashTable,
        condition: hj.condition,
    }, nil
}
```

---

## 4. Transaction Management

### **4.1 ACID Properties**

#### **Atomicity:**
All operations in a transaction succeed or all fail.

```go
func (tm *TransactionManager) BeginTransaction() *Transaction {
    return &Transaction{
        ID:        generateTransactionID(),
        StartTime: time.Now(),
        State:     Active,
        Locks:     make(map[LockID]*Lock),
    }
}

func (tm *TransactionManager) Commit(tx *Transaction) error {
    // Write all changes to WAL
    for _, change := range tx.Changes {
        if err := tm.wal.Write(change); err != nil {
            return err
        }
    }
    
    // Release all locks
    for _, lock := range tx.Locks {
        tm.lockManager.Release(lock)
    }
    
    tx.State = Committed
    return nil
}

func (tm *TransactionManager) Rollback(tx *Transaction) error {
    // Undo all changes
    for _, change := range tx.Changes {
        if err := tm.undo(change); err != nil {
            return err
        }
    }
    
    // Release all locks
    for _, lock := range tx.Locks {
        tm.lockManager.Release(lock)
    }
    
    tx.State = Aborted
    return nil
}
```

#### **Consistency:**
Database remains in a valid state after transaction.

```go
func (tm *TransactionManager) ValidateConsistency(tx *Transaction) error {
    // Check all constraints
    for _, constraint := range tm.constraints {
        if err := constraint.Validate(tx.Changes); err != nil {
            return err
        }
    }
    
    return nil
}
```

#### **Isolation:**
Concurrent transactions don't interfere with each other.

```go
func (tm *TransactionManager) Read(tx *Transaction, key Key) (*Record, error) {
    // Acquire shared lock
    lock, err := tm.lockManager.AcquireShared(tx.ID, key)
    if err != nil {
        return nil, err
    }
    
    // Read record
    record, err := tm.storage.Read(key)
    if err != nil {
        tm.lockManager.Release(lock)
        return nil, err
    }
    
    return record, nil
}

func (tm *TransactionManager) Write(tx *Transaction, key Key, value Value) error {
    // Acquire exclusive lock
    lock, err := tm.lockManager.AcquireExclusive(tx.ID, key)
    if err != nil {
        return err
    }
    
    // Write record
    if err := tm.storage.Write(key, value); err != nil {
        tm.lockManager.Release(lock)
        return err
    }
    
    // Record change for rollback
    tx.Changes = append(tx.Changes, &Change{
        Type:  Write,
        Key:   key,
        Value: value,
    })
    
    return nil
}
```

#### **Durability:**
Committed changes survive system failures.

```go
func (tm *TransactionManager) EnsureDurability(tx *Transaction) error {
    // Force WAL to disk
    if err := tm.wal.Sync(); err != nil {
        return err
    }
    
    // Update transaction log
    if err := tm.transactionLog.Commit(tx.ID); err != nil {
        return err
    }
    
    return nil
}
```

### **4.2 Multi-Version Concurrency Control (MVCC)**

MVCC allows readers to see consistent snapshots without blocking writers.

#### **MVCC Implementation:**

```go
type MVCCRecord struct {
    Key       Key
    Value     Value
    Version   Version
    Created   TransactionID
    Deleted   TransactionID
    Next      *MVCCRecord
}

func (mvcc *MVCC) Read(tx *Transaction, key Key) (*Record, error) {
    // Find visible version
    version := mvcc.findVisibleVersion(key, tx.StartTime)
    if version == nil {
        return nil, ErrRecordNotFound
    }
    
    return &Record{
        Key:   key,
        Value: version.Value,
    }, nil
}

func (mvcc *MVCC) Write(tx *Transaction, key Key, value Value) error {
    // Create new version
    newVersion := &MVCCRecord{
        Key:     key,
        Value:   value,
        Version: mvcc.nextVersion(),
        Created: tx.ID,
    }
    
    // Add to version chain
    mvcc.addVersion(key, newVersion)
    
    return nil
}
```

### **4.3 Lock Manager**

The lock manager handles locking and deadlock detection.

#### **Lock Types:**

```go
type LockType int

const (
    SharedLock LockType = iota
    ExclusiveLock
    IntentSharedLock
    IntentExclusiveLock
)

type Lock struct {
    ID        LockID
    Type      LockType
    Holder    TransactionID
    Waiters   []TransactionID
    Resource  ResourceID
}
```

#### **Lock Manager Operations:**

```go
func (lm *LockManager) AcquireShared(txID TransactionID, resource ResourceID) (*Lock, error) {
    lock := lm.getLock(resource)
    
    if lock == nil {
        // Create new lock
        lock = &Lock{
            ID:       generateLockID(),
            Type:     SharedLock,
            Holder:   txID,
            Resource: resource,
        }
        lm.locks[resource] = lock
        return lock, nil
    }
    
    if lock.Type == SharedLock || lock.Holder == txID {
        // Can acquire shared lock
        lock.addWaiter(txID)
        return lock, nil
    }
    
    // Must wait for exclusive lock
    return lm.waitForLock(txID, resource)
}

func (lm *LockManager) AcquireExclusive(txID TransactionID, resource ResourceID) (*Lock, error) {
    lock := lm.getLock(resource)
    
    if lock == nil {
        // Create new lock
        lock = &Lock{
            ID:       generateLockID(),
            Type:     ExclusiveLock,
            Holder:   txID,
            Resource: resource,
        }
        lm.locks[resource] = lock
        return lock, nil
    }
    
    if lock.Holder == txID {
        // Upgrade to exclusive
        lock.Type = ExclusiveLock
        return lock, nil
    }
    
    // Must wait for lock
    return lm.waitForLock(txID, resource)
}
```

---

## 5. Indexing Systems

### **5.1 B+ Tree Indexes**

B+ Tree indexes provide efficient range queries and sorted access.

#### **Index Structure:**

```
                    [Internal Node]
                   /      |      \
              [Leaf]   [Leaf]   [Leaf]
              [Key1]   [Key2]   [Key3]
              [Ptr1]   [Ptr2]   [Ptr3]
```

#### **Index Operations:**

```go
type BTreeIndex struct {
    root    *BTreeNode
    keyType Type
    unique  bool
}

func (bti *BTreeIndex) Insert(key Key, recordID RecordID) error {
    // Find insertion point
    leaf := bti.findLeaf(key)
    
    // Check for duplicates if unique
    if bti.unique && leaf.contains(key) {
        return ErrDuplicateKey
    }
    
    // Insert key-value pair
    if err := leaf.insert(key, recordID); err != nil {
        return err
    }
    
    // Check if split needed
    if leaf.isFull() {
        return bti.splitLeaf(leaf)
    }
    
    return nil
}

func (bti *BTreeIndex) Search(key Key) ([]RecordID, error) {
    leaf := bti.findLeaf(key)
    return leaf.search(key)
}

func (bti *BTreeIndex) RangeSearch(start, end Key) ([]RecordID, error) {
    var results []RecordID
    
    // Find starting leaf
    leaf := bti.findLeaf(start)
    
    // Scan leaves until end key
    for leaf != nil {
        records := leaf.rangeSearch(start, end)
        results = append(results, records...)
        
        if leaf.maxKey() >= end {
            break
        }
        
        leaf = leaf.next
    }
    
    return results, nil
}
```

### **5.2 Hash Indexes**

Hash indexes provide O(1) equality lookups.

#### **Hash Index Implementation:**

```go
type HashIndex struct {
    buckets []*HashBucket
    hashFunc HashFunction
    size     int
}

type HashBucket struct {
    entries []*HashEntry
}

type HashEntry struct {
    key      Key
    recordID RecordID
    next     *HashEntry
}

func (hi *HashIndex) Insert(key Key, recordID RecordID) error {
    bucket := hi.getBucket(key)
    
    // Check for existing entry
    if entry := bucket.find(key); entry != nil {
        entry.recordID = recordID
        return nil
    }
    
    // Add new entry
    entry := &HashEntry{
        key:      key,
        recordID: recordID,
    }
    
    bucket.add(entry)
    hi.size++
    
    // Check if rehashing needed
    if hi.size > len(hi.buckets)*2 {
        return hi.rehash()
    }
    
    return nil
}

func (hi *HashIndex) Search(key Key) (RecordID, error) {
    bucket := hi.getBucket(key)
    
    if entry := bucket.find(key); entry != nil {
        return entry.recordID, nil
    }
    
    return 0, ErrRecordNotFound
}
```

### **5.3 Composite Indexes**

Composite indexes support multi-column queries.

#### **Composite Index Implementation:**

```go
type CompositeIndex struct {
    columns []Column
    index   *BTreeIndex
}

func (ci *CompositeIndex) Insert(record *Record) error {
    // Extract key values
    keyValues := make([]Value, len(ci.columns))
    for i, col := range ci.columns {
        keyValues[i] = record.GetValue(col)
    }
    
    // Create composite key
    compositeKey := ci.createCompositeKey(keyValues)
    
    // Insert into underlying index
    return ci.index.Insert(compositeKey, record.ID)
}

func (ci *CompositeIndex) Search(conditions map[Column]Value) ([]RecordID, error) {
    // Create search key
    searchKey := ci.createSearchKey(conditions)
    
    // Search in index
    return ci.index.Search(searchKey)
}
```

---

## 6. Concurrency Control

### **6.1 Isolation Levels**

#### **Read Uncommitted:**
- No locks on reads
- Locks on writes
- Can read uncommitted data

#### **Read Committed:**
- Shared locks on reads (released immediately)
- Exclusive locks on writes
- Cannot read uncommitted data

#### **Repeatable Read:**
- Shared locks held until transaction end
- Exclusive locks on writes
- Prevents phantom reads

#### **Serializable:**
- Strictest isolation level
- Prevents all anomalies
- Highest consistency

### **6.2 Two-Phase Locking (2PL)**

#### **Growing Phase:**
- Acquire locks
- Cannot release locks

#### **Shrinking Phase:**
- Release locks
- Cannot acquire locks

```go
func (tx *Transaction) Begin2PL() {
    tx.phase = GrowingPhase
    tx.locks = make(map[LockID]*Lock)
}

func (tx *Transaction) AcquireLock(lockID LockID, lockType LockType) error {
    if tx.phase != GrowingPhase {
        return ErrLockAcquisitionAfterShrinking
    }
    
    lock, err := tx.lockManager.Acquire(lockID, lockType)
    if err != nil {
        return err
    }
    
    tx.locks[lockID] = lock
    return nil
}

func (tx *Transaction) ReleaseLocks() {
    tx.phase = ShrinkingPhase
    
    for lockID, lock := range tx.locks {
        tx.lockManager.Release(lockID)
        delete(tx.locks, lockID)
    }
}
```

### **6.3 Deadlock Detection**

#### **Wait-for Graph:**
```
T1 -> T2 (T1 waits for T2)
T2 -> T3 (T2 waits for T3)
T3 -> T1 (T3 waits for T1)
```

#### **Deadlock Detection Algorithm:**

```go
func (dd *DeadlockDetector) DetectDeadlock() []TransactionID {
    // Build wait-for graph
    graph := dd.buildWaitForGraph()
    
    // Find cycles
    cycles := dd.findCycles(graph)
    
    // Select victim (youngest transaction)
    victims := dd.selectVictims(cycles)
    
    return victims
}

func (dd *DeadlockDetector) buildWaitForGraph() map[TransactionID][]TransactionID {
    graph := make(map[TransactionID][]TransactionID)
    
    for _, lock := range dd.lockManager.GetAllLocks() {
        for _, waiter := range lock.Waiters {
            graph[waiter] = append(graph[waiter], lock.Holder)
        }
    }
    
    return graph
}
```

---

## 7. Data Types and Schema

### **7.1 Type System**

#### **Primitive Types:**

```go
type Type interface {
    Size() int
    Encode(value Value) []byte
    Decode(data []byte) Value
    Compare(a, b Value) int
}

type IntegerType struct {
    size int // 1, 2, 4, 8 bytes
}

type StringType struct {
    maxLength int
    variable  bool
}

type FloatType struct {
    precision int
    scale     int
}

type DateTimeType struct {
    precision int // seconds, milliseconds, microseconds
}
```

#### **Type Operations:**

```go
func (it *IntegerType) Encode(value Value) []byte {
    data := make([]byte, it.size)
    binary.BigEndian.PutUint64(data, uint64(value.(int64)))
    return data
}

func (it *IntegerType) Decode(data []byte) Value {
    return int64(binary.BigEndian.Uint64(data))
}

func (it *IntegerType) Compare(a, b Value) int {
    aVal := a.(int64)
    bVal := b.(int64)
    
    if aVal < bVal {
        return -1
    } else if aVal > bVal {
        return 1
    }
    return 0
}
```

### **7.2 Schema Management**

#### **Table Definition:**

```go
type Table struct {
    Name       string
    Columns    []Column
    PrimaryKey []Column
    Indexes    []Index
    Constraints []Constraint
}

type Column struct {
    Name     string
    Type     Type
    Nullable bool
    Default  Value
    Unique   bool
}

type Constraint struct {
    Name       string
    Type       ConstraintType
    Columns    []Column
    Expression Expression
}
```

#### **Schema Operations:**

```go
func (sm *SchemaManager) CreateTable(table *Table) error {
    // Validate table definition
    if err := sm.validateTable(table); err != nil {
        return err
    }
    
    // Create table metadata
    if err := sm.metadata.CreateTable(table); err != nil {
        return err
    }
    
    // Create storage for table
    if err := sm.storage.CreateTable(table); err != nil {
        return err
    }
    
    return nil
}

func (sm *SchemaManager) AlterTable(tableName string, changes []AlterChange) error {
    // Get current table
    table, err := sm.metadata.GetTable(tableName)
    if err != nil {
        return err
    }
    
    // Apply changes
    for _, change := range changes {
        switch change.Type {
        case AddColumn:
            table.Columns = append(table.Columns, change.Column)
        case DropColumn:
            table.Columns = sm.removeColumn(table.Columns, change.ColumnName)
        case ModifyColumn:
            table.Columns = sm.modifyColumn(table.Columns, change.Column)
        }
    }
    
    // Update metadata
    return sm.metadata.UpdateTable(table)
}
```

---

## 8. Performance Optimization

### **8.1 Query Optimization**

#### **Cost-Based Optimization:**

```go
type Optimizer struct {
    statistics *Statistics
    rules      []OptimizationRule
}

func (opt *Optimizer) Optimize(query *Query) (*ExecutionPlan, error) {
    // Generate all possible plans
    plans := opt.generatePlans(query)
    
    // Estimate cost for each plan
    for _, plan := range plans {
        plan.Cost = opt.estimateCost(plan)
    }
    
    // Select best plan
    bestPlan := opt.selectBestPlan(plans)
    
    return bestPlan, nil
}
```

#### **Index Selection:**

```go
func (opt *Optimizer) selectIndexes(query *Query) []Index {
    var selectedIndexes []Index
    
    for _, condition := range query.WhereConditions {
        if index := opt.findBestIndex(condition); index != nil {
            selectedIndexes = append(selectedIndexes, index)
        }
    }
    
    return selectedIndexes
}
```

### **8.2 Buffer Pool Optimization**

#### **LRU Replacement:**

```go
type LRUBufferPool struct {
    pages    map[PageID]*Page
    lruList  *DoublyLinkedList
    capacity int
}

func (bp *LRUBufferPool) GetPage(pageID PageID) *Page {
    if page, exists := bp.pages[pageID]; exists {
        // Move to front of LRU list
        bp.lruList.MoveToFront(page)
        return page
    }
    
    // Page not in buffer pool
    return nil
}

func (bp *LRUBufferPool) EvictPage() *Page {
    // Remove least recently used page
    page := bp.lruList.RemoveBack()
    delete(bp.pages, page.ID)
    
    // Flush if dirty
    if page.Dirty {
        bp.flushPage(page)
    }
    
    return page
}
```

### **8.3 Index Optimization**

#### **Index Maintenance:**

```go
func (im *IndexManager) MaintainIndexes() error {
    for _, index := range im.indexes {
        if err := im.maintainIndex(index); err != nil {
            return err
        }
    }
    return nil
}

func (im *IndexManager) maintainIndex(index Index) error {
    // Check if index needs rebuilding
    if index.needsRebuild() {
        return im.rebuildIndex(index)
    }
    
    // Check if index needs compaction
    if index.needsCompaction() {
        return im.compactIndex(index)
    }
    
    return nil
}
```

---

## 9. Recovery and Durability

### **9.1 Crash Recovery**

#### **Recovery Process:**

1. **Analysis Phase**: Determine which transactions were active
2. **Redo Phase**: Replay committed transactions
3. **Undo Phase**: Rollback uncommitted transactions

```go
func (rm *RecoveryManager) Recover() error {
    // Analysis phase
    activeTransactions, err := rm.analyzeLog()
    if err != nil {
        return err
    }
    
    // Redo phase
    if err := rm.redo(activeTransactions); err != nil {
        return err
    }
    
    // Undo phase
    if err := rm.undo(activeTransactions); err != nil {
        return err
    }
    
    return nil
}
```

#### **Checkpointing:**

```go
func (rm *RecoveryManager) Checkpoint() error {
    // Write checkpoint record
    checkpoint := &CheckpointRecord{
        LSN:     rm.nextLSN(),
        Time:    time.Now(),
        ActiveTransactions: rm.getActiveTransactions(),
    }
    
    // Write to log
    if err := rm.wal.Write(checkpoint); err != nil {
        return err
    }
    
    // Force to disk
    if err := rm.wal.Sync(); err != nil {
        return err
    }
    
    return nil
}
```

### **9.2 Durability Guarantees**

#### **WAL Durability:**

```go
func (db *Database) Commit(tx *Transaction) error {
    // Write commit record to WAL
    commitRecord := &CommitRecord{
        TransactionID: tx.ID,
        LSN:          db.nextLSN(),
    }
    
    if err := db.wal.Write(commitRecord); err != nil {
        return err
    }
    
    // Force WAL to disk
    if err := db.wal.Sync(); err != nil {
        return err
    }
    
    // Mark transaction as committed
    tx.State = Committed
    
    return nil
}
```

---

## 10. Real-World Examples

### **10.1 E-Commerce Database**

#### **Schema Design:**

```sql
-- Users table
CREATE TABLE users (
    id INT PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    created_at TIMESTAMP
);

-- Products table
CREATE TABLE products (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    price DECIMAL(10,2),
    category_id INT,
    stock_quantity INT
);

-- Orders table
CREATE TABLE orders (
    id INT PRIMARY KEY,
    user_id INT,
    total_amount DECIMAL(10,2),
    status VARCHAR(50),
    created_at TIMESTAMP
);

-- Order items table
CREATE TABLE order_items (
    id INT PRIMARY KEY,
    order_id INT,
    product_id INT,
    quantity INT,
    price DECIMAL(10,2)
);
```

#### **Indexes:**

```sql
-- User email index
CREATE INDEX idx_users_email ON users(email);

-- Product category index
CREATE INDEX idx_products_category ON products(category_id);

-- Order user index
CREATE INDEX idx_orders_user ON orders(user_id);

-- Order status index
CREATE INDEX idx_orders_status ON orders(status);
```

### **10.2 Query Examples**

#### **Complex Queries:**

```sql
-- Find top 10 customers by total spending
SELECT u.name, SUM(o.total_amount) as total_spent
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE o.status = 'completed'
GROUP BY u.id, u.name
ORDER BY total_spent DESC
LIMIT 10;

-- Find products with low stock
SELECT p.name, p.stock_quantity
FROM products p
WHERE p.stock_quantity < 10
ORDER BY p.stock_quantity ASC;

-- Find monthly sales
SELECT 
    DATE_TRUNC('month', created_at) as month,
    COUNT(*) as order_count,
    SUM(total_amount) as total_revenue
FROM orders
WHERE status = 'completed'
GROUP BY month
ORDER BY month;
```

### **10.3 Performance Tuning**

#### **Query Optimization:**

```sql
-- Use covering index
CREATE INDEX idx_orders_user_status_covering 
ON orders(user_id, status) 
INCLUDE (total_amount, created_at);

-- Use partial index
CREATE INDEX idx_orders_active 
ON orders(user_id) 
WHERE status = 'active';

-- Use composite index
CREATE INDEX idx_products_category_price 
ON products(category_id, price);
```

#### **Monitoring Queries:**

```sql
-- Find slow queries
SELECT query, avg_time, calls
FROM pg_stat_statements
ORDER BY avg_time DESC
LIMIT 10;

-- Find missing indexes
SELECT schemaname, tablename, attname, n_distinct, correlation
FROM pg_stats
WHERE n_distinct > 100 AND correlation < 0.1;
```

---

## 🎯 **Summary**

This comprehensive database theory covers all the essential concepts needed to understand and implement a real database engine:

1. **Storage Engines**: B+ Trees, page management, buffer pools, WAL
2. **Query Processing**: Parsing, planning, execution, joins
3. **Transaction Management**: ACID properties, MVCC, locking
4. **Indexing**: B+ Tree indexes, hash indexes, composite indexes
5. **Concurrency Control**: Isolation levels, 2PL, deadlock detection
6. **Data Types**: Type system, schema management
7. **Performance**: Query optimization, buffer pool optimization
8. **Recovery**: Crash recovery, durability guarantees
9. **Real-World Examples**: E-commerce database, performance tuning

**This is the foundation of how all modern databases work internally!** 🗄️
