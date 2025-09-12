package query

import (
	"database-engine/schema"
	"database-engine/storage"
	"database-engine/types"
	"fmt"
)

// QueryProcessor represents a query processor
type QueryProcessor struct {
	storageEngine storage.StorageEngine
	schemaManager *schema.SchemaManager
}

// NewQueryProcessor creates a new query processor
func NewQueryProcessor() *QueryProcessor {
	return &QueryProcessor{}
}

// Initialize initializes the query processor
func (qp *QueryProcessor) Initialize(se storage.StorageEngine, sm *schema.SchemaManager) error {
	qp.storageEngine = se
	qp.schemaManager = sm
	return nil
}

// Execute executes a SQL statement
func (qp *QueryProcessor) Execute(statement Statement) (*Result, error) {
	switch stmt := statement.(type) {
	case *SelectStatement:
		return qp.executeSelect(stmt)
	case *InsertStatement:
		return qp.executeInsert(stmt)
	case *UpdateStatement:
		return qp.executeUpdate(stmt)
	case *DeleteStatement:
		return qp.executeDelete(stmt)
	case *CreateTableStatement:
		return qp.executeCreateTable(stmt)
	case *DropTableStatement:
		return qp.executeDropTable(stmt)
	case *AlterTableStatement:
		return qp.executeAlterTable(stmt)
	case *CreateIndexStatement:
		return qp.executeCreateIndex(stmt)
	case *DropIndexStatement:
		return qp.executeDropIndex(stmt)
	default:
		return nil, fmt.Errorf("unsupported statement type: %T", statement)
	}
}

// Result represents a query result
type Result struct {
	Columns []string
	Rows    [][]interface{}
	Count   int
	Message string
}

// executeSelect executes a SELECT statement
func (qp *QueryProcessor) executeSelect(stmt *SelectStatement) (*Result, error) {
	// Get table
	_, err := qp.schemaManager.GetTable(stmt.From.Name)
	if err != nil {
		return nil, fmt.Errorf("table not found: %w", err)
	}
	
	// Create filter
	var filter storage.Filter
	if stmt.Where != nil {
		filter = qp.createFilter(stmt.Where)
	}
	
	// Scan table
	iterator, err := qp.storageEngine.ScanTable(stmt.From.Name, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}
	defer iterator.Close()
	
	// Collect results
	var rows [][]interface{}
	for {
		row, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next row: %w", err)
		}
		if row == nil {
			break
		}
		
		// Convert row to result format
		resultRow := qp.convertRowToResult(row, stmt.Columns)
		rows = append(rows, resultRow)
	}
	
	// Apply LIMIT and OFFSET
	if stmt.Limit > 0 {
		if stmt.Offset > 0 {
			if stmt.Offset < len(rows) {
				rows = rows[stmt.Offset:]
			} else {
				rows = [][]interface{}{}
			}
		}
		if stmt.Limit < len(rows) {
			rows = rows[:stmt.Limit]
		}
	}
	
	// Get column names
	columns := make([]string, len(stmt.Columns))
	for i, col := range stmt.Columns {
		if col.Alias != "" {
			columns[i] = col.Alias
		} else {
			columns[i] = col.Column
		}
	}
	
	return &Result{
		Columns: columns,
		Rows:    rows,
		Count:   len(rows),
		Message: fmt.Sprintf("Selected %d rows", len(rows)),
	}, nil
}

// executeInsert executes an INSERT statement
func (qp *QueryProcessor) executeInsert(stmt *InsertStatement) (*Result, error) {
	// Create row
	row := &storage.Row{
		Key:    qp.generateRowKey(stmt.Table, stmt.Values),
		Values: qp.convertExpressionsToValues(stmt.Values),
	}
	
	// Insert row
	if err := qp.storageEngine.InsertRow(stmt.Table, row); err != nil {
		return nil, fmt.Errorf("failed to insert row: %w", err)
	}
	
	return &Result{
		Message: "Row inserted successfully",
		Count:   1,
	}, nil
}

// executeUpdate executes an UPDATE statement
func (qp *QueryProcessor) executeUpdate(stmt *UpdateStatement) (*Result, error) {
	// Create filter
	var filter storage.Filter
	if stmt.Where != nil {
		filter = qp.createFilter(stmt.Where)
	}
	
	// Scan table to find rows to update
	iterator, err := qp.storageEngine.ScanTable(stmt.Table, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}
	defer iterator.Close()
	
	// Update rows
	count := 0
	for {
		row, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next row: %w", err)
		}
		if row == nil {
			break
		}
		
		// Update row values
		for _, setClause := range stmt.Set {
			value := qp.evaluateExpression(setClause.Value, row)
			row.Values[setClause.Column] = value
		}
		
		// Update row in storage
		if err := qp.storageEngine.UpdateRow(stmt.Table, row.Key, row); err != nil {
			return nil, fmt.Errorf("failed to update row: %w", err)
		}
		
		count++
	}
	
	return &Result{
		Message: fmt.Sprintf("Updated %d rows", count),
		Count:   count,
	}, nil
}

// executeDelete executes a DELETE statement
func (qp *QueryProcessor) executeDelete(stmt *DeleteStatement) (*Result, error) {
	// Create filter
	var filter storage.Filter
	if stmt.Where != nil {
		filter = qp.createFilter(stmt.Where)
	}
	
	// Scan table to find rows to delete
	iterator, err := qp.storageEngine.ScanTable(stmt.Table, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}
	defer iterator.Close()
	
	// Collect keys to delete
	var keysToDelete []storage.Key
	for {
		row, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next row: %w", err)
		}
		if row == nil {
			break
		}
		keysToDelete = append(keysToDelete, row.Key)
	}
	
	// Delete rows
	for _, key := range keysToDelete {
		if err := qp.storageEngine.DeleteRow(stmt.Table, key); err != nil {
			return nil, fmt.Errorf("failed to delete row: %w", err)
		}
	}
	
	return &Result{
		Message: fmt.Sprintf("Deleted %d rows", len(keysToDelete)),
		Count:   len(keysToDelete),
	}, nil
}

// executeCreateTable executes a CREATE TABLE statement
func (qp *QueryProcessor) executeCreateTable(stmt *CreateTableStatement) (*Result, error) {
	// Convert column definitions
	columns := make([]*schema.Column, len(stmt.Columns))
	for i, colDef := range stmt.Columns {
		columnType, err := qp.getColumnType(colDef.Type)
		if err != nil {
			return nil, fmt.Errorf("invalid column type: %w", err)
		}
		
		columns[i] = &schema.Column{
			Name:     colDef.Name,
			Type:     columnType,
			Nullable: colDef.Nullable,
			Default:  qp.evaluateExpression(colDef.Default, nil),
			Unique:   colDef.Unique,
		}
	}
	
	// Create table
	table := &schema.Table{
		Name:       stmt.Table,
		Columns:    columns,
		PrimaryKey: stmt.PrimaryKey,
	}
	
	if err := qp.schemaManager.CreateTable(table); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}
	
	if err := qp.storageEngine.CreateTable(table); err != nil {
		return nil, fmt.Errorf("failed to create table storage: %w", err)
	}
	
	return &Result{
		Message: fmt.Sprintf("Table %s created successfully", stmt.Table),
		Count:   0,
	}, nil
}

// executeDropTable executes a DROP TABLE statement
func (qp *QueryProcessor) executeDropTable(stmt *DropTableStatement) (*Result, error) {
	if err := qp.schemaManager.DropTable(stmt.Table); err != nil {
		return nil, fmt.Errorf("failed to drop table: %w", err)
	}
	
	if err := qp.storageEngine.DropTable(stmt.Table); err != nil {
		return nil, fmt.Errorf("failed to drop table storage: %w", err)
	}
	
	return &Result{
		Message: fmt.Sprintf("Table %s dropped successfully", stmt.Table),
		Count:   0,
	}, nil
}

// executeAlterTable executes an ALTER TABLE statement
func (qp *QueryProcessor) executeAlterTable(stmt *AlterTableStatement) (*Result, error) {
	// Convert alter action
	alterChange := qp.convertAlterAction(stmt.Action)
	
	if err := qp.schemaManager.AlterTable(stmt.Table, []schema.AlterChange{alterChange}); err != nil {
		return nil, fmt.Errorf("failed to alter table: %w", err)
	}
	
	return &Result{
		Message: fmt.Sprintf("Table %s altered successfully", stmt.Table),
		Count:   0,
	}, nil
}

// executeCreateIndex executes a CREATE INDEX statement
func (qp *QueryProcessor) executeCreateIndex(stmt *CreateIndexStatement) (*Result, error) {
	// Create index in schema
	index := &schema.Index{
		Name:    stmt.Index,
		Columns: stmt.Columns,
		Unique:  stmt.Unique,
		Type:    schema.BTreeIndex,
	}
	
	// Add index to table
	table, err := qp.schemaManager.GetTable(stmt.Table)
	if err != nil {
		return nil, fmt.Errorf("table not found: %w", err)
	}
	
	table.Indexes = append(table.Indexes, index)
	
	return &Result{
		Message: fmt.Sprintf("Index %s created successfully", stmt.Index),
		Count:   0,
	}, nil
}

// executeDropIndex executes a DROP INDEX statement
func (qp *QueryProcessor) executeDropIndex(stmt *DropIndexStatement) (*Result, error) {
	// Find and remove index from all tables
	tables := qp.schemaManager.ListTables()
	for _, table := range tables {
		for i, index := range table.Indexes {
			if index.Name == stmt.Index {
				table.Indexes = append(table.Indexes[:i], table.Indexes[i+1:]...)
				break
			}
		}
	}
	
	return &Result{
		Message: fmt.Sprintf("Index %s dropped successfully", stmt.Index),
		Count:   0,
	}, nil
}

// Helper methods

// createFilter creates a filter from a WHERE expression
func (qp *QueryProcessor) createFilter(expr Expression) storage.Filter {
	if expr == nil {
		return nil
	}
	
	// Simple filter creation - in real implementation, this would be more sophisticated
	if binaryExpr, ok := expr.(*BinaryExpression); ok {
		if colExpr, ok := binaryExpr.Left.(*ColumnExpression); ok {
			if literalExpr, ok := binaryExpr.Right.(*LiteralExpression); ok {
				return &storage.SimpleFilter{
					Column: colExpr.Column,
					Value:  literalExpr.Value,
				}
			}
		}
	}
	
	return nil
}

// convertRowToResult converts a storage row to result format
func (qp *QueryProcessor) convertRowToResult(row *storage.Row, columns []ColumnExpression) []interface{} {
	result := make([]interface{}, len(columns))
	
	for i, col := range columns {
		if val, exists := row.Values[col.Column]; exists {
			result[i] = val
		} else {
			result[i] = nil
		}
	}
	
	return result
}

// generateRowKey generates a row key
func (qp *QueryProcessor) generateRowKey(tableName string, values []Expression) storage.Key {
	// Simple key generation - in real implementation, use primary key
	if len(values) > 0 {
		if literalExpr, ok := values[0].(*LiteralExpression); ok {
			return &storage.SimpleKey{Value: literalExpr.Value}
		}
	}
	
	// Generate a simple key
	return &storage.SimpleKey{Value: types.IntValue(1)}
}

// convertExpressionsToValues converts expressions to values
func (qp *QueryProcessor) convertExpressionsToValues(expressions []Expression) map[string]types.Value {
	values := make(map[string]types.Value)
	
	for i, expr := range expressions {
		if literalExpr, ok := expr.(*LiteralExpression); ok {
			values[fmt.Sprintf("col_%d", i)] = literalExpr.Value
		}
	}
	
	return values
}

// evaluateExpression evaluates an expression
func (qp *QueryProcessor) evaluateExpression(expr Expression, row *storage.Row) types.Value {
	if expr == nil {
		return types.NullValue{}
	}
	
	if literalExpr, ok := expr.(*LiteralExpression); ok {
		return literalExpr.Value
	}
	
	// In real implementation, this would handle more complex expressions
	return types.NullValue{}
}

// getColumnType gets the column type from string
func (qp *QueryProcessor) getColumnType(typeStr string) (types.Type, error) {
	switch typeStr {
	case "INT", "INTEGER":
		return types.Int64Type, nil
	case "VARCHAR":
		return types.VarcharType, nil
	case "CHAR":
		return types.CharType, nil
	case "FLOAT", "REAL":
		return types.FloatTypeVar, nil
	case "DOUBLE":
		return types.DoubleTypeVar, nil
	case "BOOLEAN", "BOOL":
		return types.BoolTypeVar, nil
	case "TIMESTAMP":
		return types.TimestampTypeVar, nil
	case "BYTES", "BLOB":
		return types.BytesTypeVar, nil
	default:
		return nil, fmt.Errorf("unknown type: %s", typeStr)
	}
}

// convertAlterAction converts alter action
func (qp *QueryProcessor) convertAlterAction(action AlterAction) schema.AlterChange {
	switch act := action.(type) {
	case *AddColumnAction:
		columnType, _ := qp.getColumnType(act.Column.Type)
		return schema.AlterChange{
			Type: schema.AddColumn,
			Column: &schema.Column{
				Name:     act.Column.Name,
				Type:     columnType,
				Nullable: act.Column.Nullable,
				Default:  qp.evaluateExpression(act.Column.Default, nil),
				Unique:   act.Column.Unique,
			},
		}
	case *DropColumnAction:
		return schema.AlterChange{
			Type:       schema.DropColumn,
			ColumnName: act.Column,
		}
	case *ModifyColumnAction:
		columnType, _ := qp.getColumnType(act.Column.Type)
		return schema.AlterChange{
			Type: schema.ModifyColumn,
			Column: &schema.Column{
				Name:     act.Column.Name,
				Type:     columnType,
				Nullable: act.Column.Nullable,
				Default:  qp.evaluateExpression(act.Column.Default, nil),
				Unique:   act.Column.Unique,
			},
		}
	case *AddConstraintAction:
		return schema.AlterChange{
			Type: schema.AddConstraint,
			Constraint: &schema.Constraint{
				Name:       act.Constraint.Name,
				Type:       schema.ConstraintType(act.Constraint.Type),
				Columns:    act.Constraint.Columns,
				Expression: fmt.Sprintf("%v", act.Constraint.Expression),
			},
		}
	case *DropConstraintAction:
		return schema.AlterChange{
			Type:           schema.DropConstraint,
			ConstraintName: act.Constraint,
		}
	default:
		return schema.AlterChange{}
	}
}
