package query

import (
	"database-engine/types"
	"fmt"
	"strconv"
	"strings"
)

// Parser represents a SQL parser
type Parser struct {
	tokens []Token
	pos    int
}

// Token represents a SQL token
type Token struct {
	Type  TokenType
	Value string
	Pos   int
}

// TokenType represents the type of token
type TokenType int

const (
	// Keywords
	SELECT TokenType = iota
	FROM
	WHERE
	INSERT
	INTO
	VALUES
	UPDATE
	SET
	DELETE
	CREATE
	TABLE
	DATABASE
	DROP
	ALTER
	INDEX
	PRIMARY
	KEY
	FOREIGN
	REFERENCES
	UNIQUE
	NOT
	NULL
	DEFAULT
	CHECK
	CONSTRAINT
	AND
	OR
	IN
	LIKE
	BETWEEN
	IS
	ORDER
	BY
	GROUP
	HAVING
	LIMIT
	OFFSET
	JOIN
	LEFT
	RIGHT
	INNER
	OUTER
	ON
	AS
	ASC
	DESC
	COUNT
	SUM
	AVG
	MIN
	MAX
	DISTINCT
	ALL
	EXISTS
	UNION
	INTERSECT
	EXCEPT
	ADD
	COLUMN
	MODIFY
	
	// Operators
	EQUALS
	NOT_EQUALS
	LESS_THAN
	LESS_THAN_EQUAL
	GREATER_THAN
	GREATER_THAN_EQUAL
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULO
	
	// Delimiters
	COMMA
	SEMICOLON
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACKET
	RIGHT_BRACKET
	DOT
	
	// Literals
	IDENTIFIER
	STRING_LITERAL
	NUMBER_LITERAL
	BOOLEAN_LITERAL
	NULL_LITERAL
	
	// Other
	EOF
	UNKNOWN
)

// Statement represents a SQL statement
type Statement interface {
	Type() StatementType
}

// StatementType represents the type of statement
type StatementType int

const (
	SelectStatementType StatementType = iota
	InsertStatementType
	UpdateStatementType
	DeleteStatementType
	CreateTableStatementType
	DropTableStatementType
	AlterTableStatementType
	CreateIndexStatementType
	DropIndexStatementType
)

// SelectStatement represents a SELECT statement
type SelectStatement struct {
	Columns    []ColumnExpression
	From       TableExpression
	Where      Expression
	GroupBy    []ColumnExpression
	Having     Expression
	OrderBy    []OrderByClause
	Limit      int
	Offset     int
	Distinct   bool
}

func (s *SelectStatement) Type() StatementType {
	return SelectStatementType
}

// InsertStatement represents an INSERT statement
type InsertStatement struct {
	Table   string
	Columns []string
	Values  []Expression
}

func (s *InsertStatement) Type() StatementType {
	return InsertStatementType
}

// UpdateStatement represents an UPDATE statement
type UpdateStatement struct {
	Table  string
	Set    []SetClause
	Where  Expression
}

func (s *UpdateStatement) Type() StatementType {
	return UpdateStatementType
}

// DeleteStatement represents a DELETE statement
type DeleteStatement struct {
	Table  string
	Where  Expression
}

func (s *DeleteStatement) Type() StatementType {
	return DeleteStatementType
}

// CreateTableStatement represents a CREATE TABLE statement
type CreateTableStatement struct {
	Table       string
	Columns     []ColumnDefinition
	PrimaryKey  []string
	Constraints []ConstraintDefinition
}

func (s *CreateTableStatement) Type() StatementType {
	return CreateTableStatementType
}

// CreateDatabaseStatement represents a CREATE DATABASE statement
type CreateDatabaseStatement struct {
	Database string
}

func (s *CreateDatabaseStatement) Type() StatementType {
	return CreateTableStatementType // Reuse the same type for now
}

// DropTableStatement represents a DROP TABLE statement
type DropTableStatement struct {
	Table string
}

func (s *DropTableStatement) Type() StatementType {
	return DropTableStatementType
}

// AlterTableStatement represents an ALTER TABLE statement
type AlterTableStatement struct {
	Table  string
	Action AlterAction
}

func (s *AlterTableStatement) Type() StatementType {
	return AlterTableStatementType
}

// CreateIndexStatement represents a CREATE INDEX statement
type CreateIndexStatement struct {
	Index   string
	Table   string
	Columns []string
	Unique  bool
}

func (s *CreateIndexStatement) Type() StatementType {
	return CreateIndexStatementType
}

// DropIndexStatement represents a DROP INDEX statement
type DropIndexStatement struct {
	Index string
}

func (s *DropIndexStatement) Type() StatementType {
	return DropIndexStatementType
}

// Expression represents a SQL expression
type Expression interface {
	Type() ExpressionType
}

// ExpressionType represents the type of expression
type ExpressionType int

const (
	ColumnExpressionType ExpressionType = iota
	LiteralExpressionType
	BinaryExpressionType
	UnaryExpressionType
	FunctionExpressionType
	SubqueryExpressionType
)

// ColumnExpression represents a column reference
type ColumnExpression struct {
	Table  string
	Column string
	Alias  string
}

func (e ColumnExpression) Type() ExpressionType {
	return ColumnExpressionType
}

// LiteralExpression represents a literal value
type LiteralExpression struct {
	Value types.Value
}

func (e *LiteralExpression) Type() ExpressionType {
	return LiteralExpressionType
}

// BinaryExpression represents a binary operation
type BinaryExpression struct {
	Left     Expression
	Operator TokenType
	Right    Expression
}

func (e *BinaryExpression) Type() ExpressionType {
	return BinaryExpressionType
}

// UnaryExpression represents a unary operation
type UnaryExpression struct {
	Operator TokenType
	Operand  Expression
}

func (e *UnaryExpression) Type() ExpressionType {
	return UnaryExpressionType
}

// FunctionExpression represents a function call
type FunctionExpression struct {
	Name      string
	Arguments []Expression
}

func (e *FunctionExpression) Type() ExpressionType {
	return FunctionExpressionType
}

// SubqueryExpression represents a subquery
type SubqueryExpression struct {
	Query *SelectStatement
}

func (e *SubqueryExpression) Type() ExpressionType {
	return SubqueryExpressionType
}

// TableExpression represents a table reference
type TableExpression struct {
	Name  string
	Alias string
	Join  *JoinClause
}

// JoinClause represents a JOIN clause
type JoinClause struct {
	Type      JoinType
	Table     TableExpression
	Condition Expression
}

// JoinType represents the type of join
type JoinType int

const (
	InnerJoin JoinType = iota
	LeftJoin
	RightJoin
	FullOuterJoin
)

// OrderByClause represents an ORDER BY clause
type OrderByClause struct {
	Column ColumnExpression
	Order  OrderDirection
}

// OrderDirection represents the sort order
type OrderDirection int

const (
	AscOrder OrderDirection = iota
	DescOrder
)

// SetClause represents a SET clause in UPDATE
type SetClause struct {
	Column string
	Value  Expression
}

// ColumnDefinition represents a column definition
type ColumnDefinition struct {
	Name     string
	Type     string
	Nullable bool
	Default  Expression
	Unique   bool
}

// ConstraintDefinition represents a constraint definition
type ConstraintDefinition struct {
	Name       string
	Type       ConstraintType
	Columns    []string
	Expression Expression
}

// ConstraintType represents the type of constraint
type ConstraintType int

const (
	PrimaryKeyConstraintType ConstraintType = iota
	ForeignKeyConstraintType
	UniqueConstraintType
	CheckConstraintType
	NotNullConstraintType
)

// AlterAction represents an ALTER TABLE action
type AlterAction interface {
	Type() AlterActionType
}

// AlterActionType represents the type of alter action
type AlterActionType int

const (
	AddColumnActionType AlterActionType = iota
	DropColumnActionType
	ModifyColumnActionType
	AddConstraintActionType
	DropConstraintActionType
)

// AddColumnAction represents an ADD COLUMN action
type AddColumnAction struct {
	Column ColumnDefinition
}

func (a *AddColumnAction) Type() AlterActionType {
	return AddColumnActionType
}

// DropColumnAction represents a DROP COLUMN action
type DropColumnAction struct {
	Column string
}

func (a *DropColumnAction) Type() AlterActionType {
	return DropColumnActionType
}

// ModifyColumnAction represents a MODIFY COLUMN action
type ModifyColumnAction struct {
	Column ColumnDefinition
}

func (a *ModifyColumnAction) Type() AlterActionType {
	return ModifyColumnActionType
}

// AddConstraintAction represents an ADD CONSTRAINT action
type AddConstraintAction struct {
	Constraint ConstraintDefinition
}

func (a *AddConstraintAction) Type() AlterActionType {
	return AddConstraintActionType
}

// DropConstraintAction represents a DROP CONSTRAINT action
type DropConstraintAction struct {
	Constraint string
}

func (a *DropConstraintAction) Type() AlterActionType {
	return DropConstraintActionType
}

// NewParser creates a new SQL parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses a SQL statement
func (p *Parser) Parse(sql string) (Statement, error) {
	// Tokenize the SQL
	if err := p.tokenize(sql); err != nil {
		return nil, fmt.Errorf("tokenization failed: %w", err)
	}
	
	// Debug output removed for cleaner interface
	
	// Parse the statement
	return p.parseStatement()
}

// tokenize tokenizes the SQL string
func (p *Parser) tokenize(sql string) error {
	p.tokens = make([]Token, 0)
	p.pos = 0
	
	// Simple tokenizer implementation
	// In a real implementation, this would be more sophisticated
	
	// Remove trailing semicolon
	sql = strings.TrimSuffix(sql, ";")
	
	// Split by spaces but handle parentheses and commas properly
	words := strings.Fields(sql)
	for i, word := range words {
		// Handle parentheses and commas
		if strings.Contains(word, "(") || strings.Contains(word, ")") || strings.Contains(word, ",") {
			// Split by special characters
			parts := []string{}
			current := ""
			for _, char := range word {
				if char == '(' || char == ')' || char == ',' {
					if current != "" {
						parts = append(parts, current)
						current = ""
					}
					parts = append(parts, string(char))
				} else {
					current += string(char)
				}
			}
			if current != "" {
				parts = append(parts, current)
			}
			
			// Add each part as a token
			for _, part := range parts {
				token := Token{
					Value: part,
					Pos:   i,
				}
				// Determine token type immediately
				switch strings.ToUpper(part) {
				case "SELECT":
					token.Type = SELECT
				case "FROM":
					token.Type = FROM
				case "WHERE":
					token.Type = WHERE
				case "INSERT":
					token.Type = INSERT
				case "INTO":
					token.Type = INTO
				case "VALUES":
					token.Type = VALUES
				case "UPDATE":
					token.Type = UPDATE
				case "SET":
					token.Type = SET
				case "DELETE":
					token.Type = DELETE
				case "CREATE":
					token.Type = CREATE
				case "TABLE":
					token.Type = TABLE
				case "DATABASE":
					token.Type = DATABASE
				case "DROP":
					token.Type = DROP
				case "ALTER":
					token.Type = ALTER
				case "INDEX":
					token.Type = INDEX
				case "ON":
					token.Type = ON
				case "(":
					token.Type = LEFT_PAREN
				case ")":
					token.Type = RIGHT_PAREN
				case ",":
					token.Type = COMMA
				default:
					// Check if it's a literal
					if p.isStringLiteral(part) {
						token.Type = STRING_LITERAL
					} else if p.isNumberLiteral(part) {
						token.Type = NUMBER_LITERAL
					} else if p.isBooleanLiteral(part) {
						token.Type = BOOLEAN_LITERAL
					} else if p.isNullLiteral(part) {
						token.Type = NULL_LITERAL
					} else {
						token.Type = IDENTIFIER
					}
				}
				p.tokens = append(p.tokens, token)
			}
		} else {
			token := Token{
				Value: word,
				Pos:   i,
			}
			// Determine token type immediately
			switch strings.ToUpper(word) {
			case "SELECT":
				token.Type = SELECT
			case "FROM":
				token.Type = FROM
			case "WHERE":
				token.Type = WHERE
			case "INSERT":
				token.Type = INSERT
			case "INTO":
				token.Type = INTO
			case "VALUES":
				token.Type = VALUES
			case "UPDATE":
				token.Type = UPDATE
			case "SET":
				token.Type = SET
			case "DELETE":
				token.Type = DELETE
			case "CREATE":
				token.Type = CREATE
			case "TABLE":
				token.Type = TABLE
			case "DATABASE":
				token.Type = DATABASE
			case "DROP":
				token.Type = DROP
			case "ALTER":
				token.Type = ALTER
			case "INDEX":
				token.Type = INDEX
			case "ON":
				token.Type = ON
			default:
				// Check if it's a literal
				if p.isStringLiteral(word) {
					token.Type = STRING_LITERAL
				} else if p.isNumberLiteral(word) {
					token.Type = NUMBER_LITERAL
				} else if p.isBooleanLiteral(word) {
					token.Type = BOOLEAN_LITERAL
				} else if p.isNullLiteral(word) {
					token.Type = NULL_LITERAL
				} else {
					token.Type = IDENTIFIER
				}
			}
			p.tokens = append(p.tokens, token)
		}
	}
	
	// Add EOF token
	p.tokens = append(p.tokens, Token{Type: EOF, Value: "", Pos: len(words)})
	
	return nil
}

// isStringLiteral checks if a word is a string literal
func (p *Parser) isStringLiteral(word string) bool {
	return len(word) >= 2 && word[0] == '\'' && word[len(word)-1] == '\''
}

// isNumberLiteral checks if a word is a number literal
func (p *Parser) isNumberLiteral(word string) bool {
	_, err := strconv.ParseFloat(word, 64)
	return err == nil
}

// isBooleanLiteral checks if a word is a boolean literal
func (p *Parser) isBooleanLiteral(word string) bool {
	return word == "TRUE" || word == "FALSE"
}

// isNullLiteral checks if a word is a null literal
func (p *Parser) isNullLiteral(word string) bool {
	return word == "NULL"
}

// parseStatement parses a SQL statement
func (p *Parser) parseStatement() (Statement, error) {
	if p.pos >= len(p.tokens) {
		return nil, fmt.Errorf("unexpected end of input")
	}
	
	token := p.tokens[p.pos]
	
	switch token.Type {
	case SELECT:
		p.pos++
		return p.parseSelectStatement()
	case INSERT:
		p.pos++
		return p.parseInsertStatement()
	case UPDATE:
		p.pos++
		return p.parseUpdateStatement()
	case DELETE:
		p.pos++
		return p.parseDeleteStatement()
	case CREATE:
		p.pos++
		return p.parseCreateStatement()
	case DROP:
		p.pos++
		return p.parseDropStatement()
	case ALTER:
		p.pos++
		return p.parseAlterStatement()
	default:
		return nil, fmt.Errorf("unexpected token: %s", token.Value)
	}
}

// parseSelectStatement parses a SELECT statement
func (p *Parser) parseSelectStatement() (*SelectStatement, error) {
	stmt := &SelectStatement{}
	
	// SELECT already consumed in main parser
	
	// Parse DISTINCT
	if p.peek().Type == DISTINCT {
		p.consume()
		stmt.Distinct = true
	}
	
	// Parse columns
	columns, err := p.parseColumnList()
	if err != nil {
		return nil, err
	}
	stmt.Columns = columns
	
	// Parse FROM
	if err := p.expect(FROM); err != nil {
		return nil, err
	}
	
	// Parse table
	table, err := p.parseTableExpression()
	if err != nil {
		return nil, err
	}
	stmt.From = table
	
	// Parse WHERE
	if p.peek().Type == WHERE {
		p.consume()
		where, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		stmt.Where = where
	}
	
	// Parse GROUP BY
	if p.peek().Type == GROUP {
		p.consume()
		if err := p.expect(BY); err != nil {
			return nil, err
		}
		groupBy, err := p.parseColumnList()
		if err != nil {
			return nil, err
		}
		stmt.GroupBy = groupBy
	}
	
	// Parse HAVING
	if p.peek().Type == HAVING {
		p.consume()
		having, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		stmt.Having = having
	}
	
	// Parse ORDER BY
	if p.peek().Type == ORDER {
		p.consume()
		if err := p.expect(BY); err != nil {
			return nil, err
		}
		orderBy, err := p.parseOrderByList()
		if err != nil {
			return nil, err
		}
		stmt.OrderBy = orderBy
	}
	
	// Parse LIMIT
	if p.peek().Type == LIMIT {
		p.consume()
		limit, err := p.parseNumber()
		if err != nil {
			return nil, err
		}
		stmt.Limit = limit
	}
	
	// Parse OFFSET
	if p.peek().Type == OFFSET {
		p.consume()
		offset, err := p.parseNumber()
		if err != nil {
			return nil, err
		}
		stmt.Offset = offset
	}
	
	return stmt, nil
}

// parseInsertStatement parses an INSERT statement
func (p *Parser) parseInsertStatement() (*InsertStatement, error) {
	stmt := &InsertStatement{}
	
	// Parse INTO (INSERT already consumed)
	if err := p.expect(INTO); err != nil {
		return nil, err
	}
	
	// Parse table name
	table, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Table = table
	
	// Parse columns
	if p.peek().Type == LEFT_PAREN {
		p.consume()
		columns, err := p.parseIdentifierList()
		if err != nil {
			return nil, err
		}
		stmt.Columns = columns
		if err := p.expect(RIGHT_PAREN); err != nil {
			return nil, err
		}
	}
	
	// Parse VALUES
	if err := p.expect(VALUES); err != nil {
		return nil, err
	}
	
	// Parse values in parentheses
	if err := p.expect(LEFT_PAREN); err != nil {
		return nil, err
	}
	
	values, err := p.parseExpressionList()
	if err != nil {
		return nil, err
	}
	stmt.Values = values
	
	if err := p.expect(RIGHT_PAREN); err != nil {
		return nil, err
	}
	
	return stmt, nil
}

// parseUpdateStatement parses an UPDATE statement
func (p *Parser) parseUpdateStatement() (*UpdateStatement, error) {
	stmt := &UpdateStatement{}
	
	// UPDATE already consumed
	
	// Parse table name
	table, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Table = table
	
	// Parse SET
	if err := p.expect(SET); err != nil {
		return nil, err
	}
	
	// Parse set clauses
	setClauses, err := p.parseSetClauseList()
	if err != nil {
		return nil, err
	}
	stmt.Set = setClauses
	
	// Parse WHERE
	if p.peek().Type == WHERE {
		p.consume()
		where, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		stmt.Where = where
	}
	
	return stmt, nil
}

// parseDeleteStatement parses a DELETE statement
func (p *Parser) parseDeleteStatement() (*DeleteStatement, error) {
	stmt := &DeleteStatement{}
	
	// Parse FROM (DELETE already consumed)
	if err := p.expect(FROM); err != nil {
		return nil, err
	}
	
	// Parse table name
	table, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Table = table
	
	// Parse WHERE
	if p.peek().Type == WHERE {
		p.consume()
		where, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		stmt.Where = where
	}
	
	return stmt, nil
}

// parseCreateStatement parses a CREATE statement
func (p *Parser) parseCreateStatement() (Statement, error) {
	if p.peek().Type == TABLE {
		return p.parseCreateTableStatement()
	} else if p.peek().Type == INDEX {
		return p.parseCreateIndexStatement()
	} else if p.peek().Type == DATABASE {
		return p.parseCreateDatabaseStatement()
	}
	
	return nil, fmt.Errorf("unexpected CREATE statement type")
}

// parseCreateTableStatement parses a CREATE TABLE statement
func (p *Parser) parseCreateTableStatement() (*CreateTableStatement, error) {
	stmt := &CreateTableStatement{}
	
	// Parse TABLE (CREATE already consumed)
	if err := p.expect(TABLE); err != nil {
		return nil, err
	}
	
	// Parse table name
	table, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Table = table
	
	// Parse column definitions
	if err := p.expect(LEFT_PAREN); err != nil {
		return nil, err
	}
	
	columns, err := p.parseColumnDefinitionList()
	if err != nil {
		return nil, err
	}
	stmt.Columns = columns
	
	if err := p.expect(RIGHT_PAREN); err != nil {
		return nil, err
	}
	
	return stmt, nil
}

// parseCreateDatabaseStatement parses a CREATE DATABASE statement
func (p *Parser) parseCreateDatabaseStatement() (*CreateDatabaseStatement, error) {
	stmt := &CreateDatabaseStatement{}
	
	// Parse DATABASE (CREATE already consumed)
	if err := p.expect(DATABASE); err != nil {
		return nil, err
	}
	
	// Parse database name
	if p.peek().Type != IDENTIFIER {
		return nil, fmt.Errorf("expected database name")
	}
	stmt.Database = p.peek().Value
	p.pos++
	
	return stmt, nil
}

// parseCreateIndexStatement parses a CREATE INDEX statement
func (p *Parser) parseCreateIndexStatement() (*CreateIndexStatement, error) {
	stmt := &CreateIndexStatement{}
	
	// Parse INDEX (CREATE already consumed)
	if err := p.expect(INDEX); err != nil {
		return nil, err
	}
	
	// Parse index name
	index, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Index = index
	
	// Parse ON table
	if err := p.expect(ON); err != nil {
		return nil, err
	}
	
	table, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Table = table
	
	// Parse columns
	if err := p.expect(LEFT_PAREN); err != nil {
		return nil, err
	}
	
	columns, err := p.parseIdentifierList()
	if err != nil {
		return nil, err
	}
	stmt.Columns = columns
	
	if err := p.expect(RIGHT_PAREN); err != nil {
		return nil, err
	}
	
	return stmt, nil
}

// parseDropStatement parses a DROP statement
func (p *Parser) parseDropStatement() (Statement, error) {
	if p.peek().Type == TABLE {
		return p.parseDropTableStatement()
	} else if p.peek().Type == INDEX {
		return p.parseDropIndexStatement()
	}
	
	return nil, fmt.Errorf("unexpected DROP statement type")
}

// parseDropTableStatement parses a DROP TABLE statement
func (p *Parser) parseDropTableStatement() (*DropTableStatement, error) {
	stmt := &DropTableStatement{}
	
	// Parse DROP TABLE
	if err := p.expect(DROP); err != nil {
		return nil, err
	}
	if err := p.expect(TABLE); err != nil {
		return nil, err
	}
	
	// Parse table name
	table, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Table = table
	
	return stmt, nil
}

// parseDropIndexStatement parses a DROP INDEX statement
func (p *Parser) parseDropIndexStatement() (*DropIndexStatement, error) {
	stmt := &DropIndexStatement{}
	
	// Parse DROP INDEX
	if err := p.expect(DROP); err != nil {
		return nil, err
	}
	if err := p.expect(INDEX); err != nil {
		return nil, err
	}
	
	// Parse index name
	index, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Index = index
	
	return stmt, nil
}

// parseAlterStatement parses an ALTER statement
func (p *Parser) parseAlterStatement() (*AlterTableStatement, error) {
	stmt := &AlterTableStatement{}
	
	// Parse ALTER TABLE
	if err := p.expect(ALTER); err != nil {
		return nil, err
	}
	if err := p.expect(TABLE); err != nil {
		return nil, err
	}
	
	// Parse table name
	table, err := p.expectIdentifier()
	if err != nil {
		return nil, err
	}
	stmt.Table = table
	
	// Parse action
	action, err := p.parseAlterAction()
	if err != nil {
		return nil, err
	}
	stmt.Action = action
	
	return stmt, nil
}

// Helper methods for parsing

func (p *Parser) peek() Token {
	if p.pos >= len(p.tokens) {
		return Token{Type: EOF}
	}
	return p.tokens[p.pos]
}

func (p *Parser) consume() Token {
	if p.pos >= len(p.tokens) {
		return Token{Type: EOF}
	}
	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) expect(expected TokenType) error {
	token := p.consume()
	if token.Type != expected {
		return fmt.Errorf("expected %v, got %v", expected, token.Type)
	}
	return nil
}

func (p *Parser) expectIdentifier() (string, error) {
	token := p.consume()
	if token.Type != IDENTIFIER {
		return "", fmt.Errorf("expected identifier, got %v", token.Type)
	}
	return token.Value, nil
}

func (p *Parser) parseColumnList() ([]ColumnExpression, error) {
	columns := make([]ColumnExpression, 0)
	
	for {
		column, err := p.parseColumnExpression()
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
		
		if p.peek().Type != COMMA {
			break
		}
		p.consume()
	}
	
	return columns, nil
}

func (p *Parser) parseColumnExpression() (ColumnExpression, error) {
	column := ColumnExpression{}
	
	// Parse table.column or just column
	if p.peek().Type == IDENTIFIER {
		table, err := p.expectIdentifier()
		if err != nil {
			return column, err
		}
		
		if p.peek().Type == DOT {
			p.consume()
			col, err := p.expectIdentifier()
			if err != nil {
				return column, err
			}
			column.Table = table
			column.Column = col
		} else {
			column.Column = table
		}
	}
	
	// Parse alias
	if p.peek().Type == AS {
		p.consume()
		alias, err := p.expectIdentifier()
		if err != nil {
			return column, err
		}
		column.Alias = alias
	}
	
	return column, nil
}

func (p *Parser) parseTableExpression() (TableExpression, error) {
	table := TableExpression{}
	
	// Parse table name
	name, err := p.expectIdentifier()
	if err != nil {
		return table, err
	}
	table.Name = name
	
	// Parse alias
	if p.peek().Type == AS {
		p.consume()
		alias, err := p.expectIdentifier()
		if err != nil {
			return table, err
		}
		table.Alias = alias
	}
	
	return table, nil
}

func (p *Parser) parseExpression() (Expression, error) {
	// Simple expression parsing
	// In a real implementation, this would handle operator precedence
	
	token := p.peek()
	
	switch token.Type {
	case IDENTIFIER:
		// Check if it's a column reference or a literal
		// For VALUES clause, identifiers should be treated as literals
		return p.parseColumnExpression()
	case STRING_LITERAL:
		p.consume()
		return &LiteralExpression{
			Value: types.StringValue(token.Value[1 : len(token.Value)-1]),
		}, nil
	case NUMBER_LITERAL:
		p.consume()
		// Check if it's an integer or float
		if strings.Contains(token.Value, ".") {
			// Float value
			val, _ := strconv.ParseFloat(token.Value, 64)
			return &LiteralExpression{
				Value: types.FloatValue(val),
			}, nil
		} else {
			// Integer value
			val, _ := strconv.ParseInt(token.Value, 10, 64)
			return &LiteralExpression{
				Value: types.IntValue(val),
			}, nil
		}
	case BOOLEAN_LITERAL:
		p.consume()
		val := token.Value == "TRUE"
		return &LiteralExpression{
			Value: types.BoolValue(val),
		}, nil
	case NULL_LITERAL:
		p.consume()
		return &LiteralExpression{
			Value: types.NullValue{},
		}, nil
	default:
		return nil, fmt.Errorf("unexpected token in expression: %v", token.Type)
	}
}

func (p *Parser) parseNumber() (int, error) {
	token := p.consume()
	if token.Type != NUMBER_LITERAL {
		return 0, fmt.Errorf("expected number, got %v", token.Type)
	}
	
	val, err := strconv.Atoi(token.Value)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", token.Value)
	}
	
	return val, nil
}

func (p *Parser) parseOrderByList() ([]OrderByClause, error) {
	clauses := make([]OrderByClause, 0)
	
	for {
		column, err := p.parseColumnExpression()
		if err != nil {
			return nil, err
		}
		
		clause := OrderByClause{
			Column: column,
			Order:  AscOrder,
		}
		
		// Parse order direction
		if p.peek().Type == ASC {
			p.consume()
			clause.Order = AscOrder
		} else if p.peek().Type == DESC {
			p.consume()
			clause.Order = DescOrder
		}
		
		clauses = append(clauses, clause)
		
		if p.peek().Type != COMMA {
			break
		}
		p.consume()
	}
	
	return clauses, nil
}

func (p *Parser) parseIdentifierList() ([]string, error) {
	identifiers := make([]string, 0)
	
	for {
		identifier, err := p.expectIdentifier()
		if err != nil {
			return nil, err
		}
		identifiers = append(identifiers, identifier)
		
		if p.peek().Type != COMMA {
			break
		}
		p.consume()
	}
	
	return identifiers, nil
}

func (p *Parser) parseExpressionList() ([]Expression, error) {
	expressions := make([]Expression, 0)
	
	for {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, expr)
		
		if p.peek().Type != COMMA {
			break
		}
		p.consume()
	}
	
	return expressions, nil
}

func (p *Parser) parseSetClauseList() ([]SetClause, error) {
	clauses := make([]SetClause, 0)
	
	for {
		column, err := p.expectIdentifier()
		if err != nil {
			return nil, err
		}
		
		if err := p.expect(EQUALS); err != nil {
			return nil, err
		}
		
		value, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		
		clauses = append(clauses, SetClause{
			Column: column,
			Value:  value,
		})
		
		if p.peek().Type != COMMA {
			break
		}
		p.consume()
	}
	
	return clauses, nil
}

func (p *Parser) parseColumnDefinitionList() ([]ColumnDefinition, error) {
	columns := make([]ColumnDefinition, 0)
	
	for {
		column, err := p.parseColumnDefinition()
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
		
		if p.peek().Type != COMMA {
			break
		}
		p.consume()
	}
	
	return columns, nil
}

func (p *Parser) parseColumnType() (string, error) {
	// Get the base type
	baseType, err := p.expectIdentifier()
	if err != nil {
		return "", err
	}
	
	// Check if it has parameters like VARCHAR(100) or DECIMAL(10,2)
	if p.peek().Type == LEFT_PAREN {
		p.consume() // consume (
		
		// Parse parameters
		params := make([]string, 0)
		for {
			if p.peek().Type == IDENTIFIER || p.peek().Type == NUMBER_LITERAL {
				params = append(params, p.peek().Value)
				p.consume()
			} else {
				break
			}
			
			if p.peek().Type == COMMA {
				p.consume()
			} else {
				break
			}
		}
		
		// Expect closing parenthesis
		if err := p.expect(RIGHT_PAREN); err != nil {
			return "", err
		}
		
		// Build type string with parameters
		if len(params) > 0 {
			paramStr := strings.Join(params, ",")
			return fmt.Sprintf("%s(%s)", baseType, paramStr), nil
		}
	}
	
	return baseType, nil
}

func (p *Parser) parseColumnDefinition() (ColumnDefinition, error) {
	column := ColumnDefinition{}
	
	// Parse column name
	name, err := p.expectIdentifier()
	if err != nil {
		return column, err
	}
	column.Name = name
	
	// Parse column type
	typeStr, err := p.parseColumnType()
	if err != nil {
		return column, err
	}
	column.Type = typeStr
	
	// Parse column constraints
	for {
		token := p.peek()
		switch token.Type {
		case NOT:
			p.consume()
			if err := p.expect(NULL); err != nil {
				return column, err
			}
			column.Nullable = false
		case NULL:
			p.consume()
			column.Nullable = true
		case DEFAULT:
			p.consume()
			defaultVal, err := p.parseExpression()
			if err != nil {
				return column, err
			}
			column.Default = defaultVal
		case UNIQUE:
			p.consume()
			column.Unique = true
		default:
			return column, nil
		}
	}
}

func (p *Parser) parseAlterAction() (AlterAction, error) {
	token := p.peek()
	
	switch token.Type {
	case ADD:
		p.consume()
		if p.peek().Type == COLUMN {
			p.consume()
			column, err := p.parseColumnDefinition()
			if err != nil {
				return nil, err
			}
			return &AddColumnAction{Column: column}, nil
		}
	case DROP:
		p.consume()
		if p.peek().Type == COLUMN {
			p.consume()
			column, err := p.expectIdentifier()
			if err != nil {
				return nil, err
			}
			return &DropColumnAction{Column: column}, nil
		}
	case MODIFY:
		p.consume()
		if p.peek().Type == COLUMN {
			p.consume()
			column, err := p.parseColumnDefinition()
			if err != nil {
				return nil, err
			}
			return &ModifyColumnAction{Column: column}, nil
		}
	}
	
	return nil, fmt.Errorf("unexpected alter action")
}
