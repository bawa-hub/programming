package main

import (
	"fmt"
	"strconv"
	"strings"
)

// =============================================================================
// BASIC INTERPRETER PATTERN
// =============================================================================

// Context interface
type Context interface {
	GetVariable(name string) interface{}
	SetVariable(name string, value interface{})
	GetFunction(name string) func([]interface{}) interface{}
	SetFunction(name string, fn func([]interface{}) interface{})
}

// Concrete Context
type ExpressionContext struct {
	variables map[string]interface{}
	functions map[string]func([]interface{}) interface{}
}

func NewExpressionContext() *ExpressionContext {
	return &ExpressionContext{
		variables: make(map[string]interface{}),
		functions: make(map[string]func([]interface{}) interface{}),
	}
}

func (ec *ExpressionContext) GetVariable(name string) interface{} {
	return ec.variables[name]
}

func (ec *ExpressionContext) SetVariable(name string, value interface{}) {
	ec.variables[name] = value
}

func (ec *ExpressionContext) GetFunction(name string) func([]interface{}) interface{} {
	return ec.functions[name]
}

func (ec *ExpressionContext) SetFunction(name string, fn func([]interface{}) interface{}) {
	ec.functions[name] = fn
}

// Abstract Expression
type Expression interface {
	Interpret(context Context) interface{}
	Evaluate() interface{}
}

// Terminal Expressions
type NumberExpression struct {
	value float64
}

func NewNumberExpression(value float64) *NumberExpression {
	return &NumberExpression{value: value}
}

func (ne *NumberExpression) Interpret(context Context) interface{} {
	return ne.value
}

func (ne *NumberExpression) Evaluate() interface{} {
	return ne.value
}

type VariableExpression struct {
	name string
}

func NewVariableExpression(name string) *VariableExpression {
	return &VariableExpression{name: name}
}

func (ve *VariableExpression) Interpret(context Context) interface{} {
	return context.GetVariable(ve.name)
}

func (ve *VariableExpression) Evaluate() interface{} {
	return nil // Cannot evaluate without context
}

// Non-Terminal Expressions
type AddExpression struct {
	left  Expression
	right Expression
}

func NewAddExpression(left, right Expression) *AddExpression {
	return &AddExpression{left: left, right: right}
}

func (ae *AddExpression) Interpret(context Context) interface{} {
	leftValue := ae.left.Interpret(context)
	rightValue := ae.right.Interpret(context)
	
	leftNum, leftOk := leftValue.(float64)
	rightNum, rightOk := rightValue.(float64)
	
	if leftOk && rightOk {
		return leftNum + rightNum
	}
	return nil
}

func (ae *AddExpression) Evaluate() interface{} {
	return nil // Cannot evaluate without context
}

type SubtractExpression struct {
	left  Expression
	right Expression
}

func NewSubtractExpression(left, right Expression) *SubtractExpression {
	return &SubtractExpression{left: left, right: right}
}

func (se *SubtractExpression) Interpret(context Context) interface{} {
	leftValue := se.left.Interpret(context)
	rightValue := se.right.Interpret(context)
	
	leftNum, leftOk := leftValue.(float64)
	rightNum, rightOk := rightValue.(float64)
	
	if leftOk && rightOk {
		return leftNum - rightNum
	}
	return nil
}

func (se *SubtractExpression) Evaluate() interface{} {
	return nil // Cannot evaluate without context
}

type MultiplyExpression struct {
	left  Expression
	right Expression
}

func NewMultiplyExpression(left, right Expression) *MultiplyExpression {
	return &MultiplyExpression{left: left, right: right}
}

func (me *MultiplyExpression) Interpret(context Context) interface{} {
	leftValue := me.left.Interpret(context)
	rightValue := me.right.Interpret(context)
	
	leftNum, leftOk := leftValue.(float64)
	rightNum, rightOk := rightValue.(float64)
	
	if leftOk && rightOk {
		return leftNum * rightNum
	}
	return nil
}

func (me *MultiplyExpression) Evaluate() interface{} {
	return nil // Cannot evaluate without context
}

type DivideExpression struct {
	left  Expression
	right Expression
}

func NewDivideExpression(left, right Expression) *DivideExpression {
	return &DivideExpression{left: left, right: right}
}

func (de *DivideExpression) Interpret(context Context) interface{} {
	leftValue := de.left.Interpret(context)
	rightValue := de.right.Interpret(context)
	
	leftNum, leftOk := leftValue.(float64)
	rightNum, rightOk := rightValue.(float64)
	
	if leftOk && rightOk {
		if rightNum != 0 {
			return leftNum / rightNum
		}
	}
	return nil
}

func (de *DivideExpression) Evaluate() interface{} {
	return nil // Cannot evaluate without context
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. SQL-LIKE QUERY INTERPRETER
type QueryContext struct {
	tables map[string][]map[string]interface{}
}

func NewQueryContext() *QueryContext {
	return &QueryContext{
		tables: make(map[string][]map[string]interface{}),
	}
}

func (qc *QueryContext) AddTable(name string, data []map[string]interface{}) {
	qc.tables[name] = data
}

func (qc *QueryContext) GetTable(name string) []map[string]interface{} {
	return qc.tables[name]
}

type QueryExpression interface {
	Execute(context *QueryContext) []map[string]interface{}
}

type SelectExpression struct {
	columns []string
	from    string
	where   WhereExpression
}

func NewSelectExpression(columns []string, from string, where WhereExpression) *SelectExpression {
	return &SelectExpression{
		columns: columns,
		from:    from,
		where:   where,
	}
}

func (se *SelectExpression) Execute(context *QueryContext) []map[string]interface{} {
	table := context.GetTable(se.from)
	var result []map[string]interface{}
	
	for _, row := range table {
		if se.where == nil || se.where.Evaluate(row) {
			selectedRow := make(map[string]interface{})
			for _, column := range se.columns {
				if value, exists := row[column]; exists {
					selectedRow[column] = value
				}
			}
			result = append(result, selectedRow)
		}
	}
	
	return result
}

type WhereExpression interface {
	Evaluate(row map[string]interface{}) bool
}

type EqualsExpression struct {
	column string
	value  interface{}
}

func NewEqualsExpression(column string, value interface{}) *EqualsExpression {
	return &EqualsExpression{column: column, value: value}
}

func (ee *EqualsExpression) Evaluate(row map[string]interface{}) bool {
	rowValue, exists := row[ee.column]
	return exists && rowValue == ee.value
}

type AndExpression struct {
	left  WhereExpression
	right WhereExpression
}

func NewAndExpression(left, right WhereExpression) *AndExpression {
	return &AndExpression{left: left, right: right}
}

func (ae *AndExpression) Evaluate(row map[string]interface{}) bool {
	return ae.left.Evaluate(row) && ae.right.Evaluate(row)
}

type OrExpression struct {
	left  WhereExpression
	right WhereExpression
}

func NewOrExpression(left, right WhereExpression) *OrExpression {
	return &OrExpression{left: left, right: right}
}

func (oe *OrExpression) Evaluate(row map[string]interface{}) bool {
	return oe.left.Evaluate(row) || oe.right.Evaluate(row)
}

// 2. CONFIGURATION INTERPRETER
type ConfigContext struct {
	values map[string]interface{}
}

func NewConfigContext() *ConfigContext {
	return &ConfigContext{
		values: make(map[string]interface{}),
	}
}

func (cc *ConfigContext) SetValue(key string, value interface{}) {
	cc.values[key] = value
}

func (cc *ConfigContext) GetValue(key string) interface{} {
	return cc.values[key]
}

type ConfigExpression interface {
	Evaluate(context *ConfigContext) interface{}
}

type StringExpression struct {
	value string
}

func NewStringExpression(value string) *StringExpression {
	return &StringExpression{value: value}
}

func (se *StringExpression) Evaluate(context *ConfigContext) interface{} {
	return se.value
}

type NumberConfigExpression struct {
	value float64
}

func NewNumberConfigExpression(value float64) *NumberConfigExpression {
	return &NumberConfigExpression{value: value}
}

func (nce *NumberConfigExpression) Evaluate(context *ConfigContext) interface{} {
	return nce.value
}

type VariableConfigExpression struct {
	name string
}

func NewVariableConfigExpression(name string) *VariableConfigExpression {
	return &VariableConfigExpression{name: name}
}

func (vce *VariableConfigExpression) Evaluate(context *ConfigContext) interface{} {
	return context.GetValue(vce.name)
}

type InterpolationExpression struct {
	template string
	context  *ConfigContext
}

func NewInterpolationExpression(template string, context *ConfigContext) *InterpolationExpression {
	return &InterpolationExpression{template: template, context: context}
}

func (ie *InterpolationExpression) Evaluate(context *ConfigContext) interface{} {
	result := ie.template
	for key, value := range context.values {
		placeholder := fmt.Sprintf("${%s}", key)
		valueStr := fmt.Sprintf("%v", value)
		result = strings.ReplaceAll(result, placeholder, valueStr)
	}
	return result
}

// 3. RULE ENGINE INTERPRETER
type RuleContext struct {
	facts map[string]interface{}
}

func NewRuleContext() *RuleContext {
	return &RuleContext{
		facts: make(map[string]interface{}),
	}
}

func (rc *RuleContext) SetFact(name string, value interface{}) {
	rc.facts[name] = value
}

func (rc *RuleContext) GetFact(name string) interface{} {
	return rc.facts[name]
}

type RuleExpression interface {
	Evaluate(context *RuleContext) bool
}

type FactExpression struct {
	name  string
	value interface{}
}

func NewFactExpression(name string, value interface{}) *FactExpression {
	return &FactExpression{name: name, value: value}
}

func (fe *FactExpression) Evaluate(context *RuleContext) bool {
	factValue := context.GetFact(fe.name)
	return factValue == fe.value
}

type GreaterThanExpression struct {
	left  string
	right interface{}
}

func NewGreaterThanExpression(left string, right interface{}) *GreaterThanExpression {
	return &GreaterThanExpression{left: left, right: right}
}

func (gte *GreaterThanExpression) Evaluate(context *RuleContext) bool {
	factValue := context.GetFact(gte.left)
	if factValue == nil {
		return false
	}
	
	// Simple comparison for numbers
	if factNum, ok := factValue.(float64); ok {
		if rightNum, ok := gte.right.(float64); ok {
			return factNum > rightNum
		}
	}
	return false
}

type LessThanExpression struct {
	left  string
	right interface{}
}

func NewLessThanExpression(left string, right interface{}) *LessThanExpression {
	return &LessThanExpression{left: left, right: right}
}

func (lte *LessThanExpression) Evaluate(context *RuleContext) bool {
	factValue := context.GetFact(lte.left)
	if factValue == nil {
		return false
	}
	
	// Simple comparison for numbers
	if factNum, ok := factValue.(float64); ok {
		if rightNum, ok := lte.right.(float64); ok {
			return factNum < rightNum
		}
	}
	return false
}

type RuleAndExpression struct {
	left  RuleExpression
	right RuleExpression
}

func NewRuleAndExpression(left, right RuleExpression) *RuleAndExpression {
	return &RuleAndExpression{left: left, right: right}
}

func (rae *RuleAndExpression) Evaluate(context *RuleContext) bool {
	return rae.left.Evaluate(context) && rae.right.Evaluate(context)
}

type RuleOrExpression struct {
	left  RuleExpression
	right RuleExpression
}

func NewRuleOrExpression(left, right RuleExpression) *RuleOrExpression {
	return &RuleOrExpression{left: left, right: right}
}

func (roe *RuleOrExpression) Evaluate(context *RuleContext) bool {
	return roe.left.Evaluate(context) || roe.right.Evaluate(context)
}

// =============================================================================
// EXPRESSION PARSER
// =============================================================================

type ExpressionParser struct{}

func NewExpressionParser() *ExpressionParser {
	return &ExpressionParser{}
}

func (ep *ExpressionParser) Parse(expression string) Expression {
	// Simple parser for mathematical expressions
	// This is a simplified version - real parsers would be more complex
	tokens := ep.tokenize(expression)
	return ep.parseExpression(tokens)
}

func (ep *ExpressionParser) tokenize(expression string) []string {
	// Simple tokenization - split by spaces and operators
	expression = strings.ReplaceAll(expression, "+", " + ")
	expression = strings.ReplaceAll(expression, "-", " - ")
	expression = strings.ReplaceAll(expression, "*", " * ")
	expression = strings.ReplaceAll(expression, "/", " / ")
	expression = strings.ReplaceAll(expression, "(", " ( ")
	expression = strings.ReplaceAll(expression, ")", " ) ")
	
	return strings.Fields(expression)
}

func (ep *ExpressionParser) parseExpression(tokens []string) Expression {
	// Very simplified parser - just handles basic arithmetic
	if len(tokens) == 1 {
		if num, err := strconv.ParseFloat(tokens[0], 64); err == nil {
			return NewNumberExpression(num)
		}
		return NewVariableExpression(tokens[0])
	}
	
	// Look for operators in order of precedence
	for i, token := range tokens {
		switch token {
		case "+":
			if i > 0 && i < len(tokens)-1 {
				left := ep.parseExpression(tokens[:i])
				right := ep.parseExpression(tokens[i+1:])
				return NewAddExpression(left, right)
			}
		case "-":
			if i > 0 && i < len(tokens)-1 {
				left := ep.parseExpression(tokens[:i])
				right := ep.parseExpression(tokens[i+1:])
				return NewSubtractExpression(left, right)
			}
		case "*":
			if i > 0 && i < len(tokens)-1 {
				left := ep.parseExpression(tokens[:i])
				right := ep.parseExpression(tokens[i+1:])
				return NewMultiplyExpression(left, right)
			}
		case "/":
			if i > 0 && i < len(tokens)-1 {
				left := ep.parseExpression(tokens[:i])
				right := ep.parseExpression(tokens[i+1:])
				return NewDivideExpression(left, right)
			}
		}
	}
	
	// If no operator found, return the first token as a number or variable
	if len(tokens) > 0 {
		if num, err := strconv.ParseFloat(tokens[0], 64); err == nil {
			return NewNumberExpression(num)
		}
		return NewVariableExpression(tokens[0])
	}
	
	return NewNumberExpression(0)
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== INTERPRETER PATTERN DEMONSTRATION ===\n")

	// 1. BASIC MATHEMATICAL EXPRESSION INTERPRETER
	fmt.Println("1. BASIC MATHEMATICAL EXPRESSION INTERPRETER:")
	context := NewExpressionContext()
	context.SetVariable("x", 10.0)
	context.SetVariable("y", 5.0)
	
	// Create expression: (x + y) * 2
	expression := NewMultiplyExpression(
		NewAddExpression(
			NewVariableExpression("x"),
			NewVariableExpression("y"),
		),
		NewNumberExpression(2.0),
	)
	
	result := expression.Interpret(context)
	fmt.Printf("Expression: (x + y) * 2 where x=10, y=5\n")
	fmt.Printf("Result: %v\n", result)
	
	// Test with different values
	context.SetVariable("x", 20.0)
	context.SetVariable("y", 3.0)
	result = expression.Interpret(context)
	fmt.Printf("Expression: (x + y) * 2 where x=20, y=3\n")
	fmt.Printf("Result: %v\n", result)
	fmt.Println()

	// 2. SQL-LIKE QUERY INTERPRETER
	fmt.Println("2. SQL-LIKE QUERY INTERPRETER:")
	queryContext := NewQueryContext()
	
	// Add sample data
	users := []map[string]interface{}{
		{"id": 1, "name": "Alice", "age": 25, "city": "New York"},
		{"id": 2, "name": "Bob", "age": 30, "city": "Los Angeles"},
		{"id": 3, "name": "Charlie", "age": 35, "city": "New York"},
		{"id": 4, "name": "David", "age": 28, "city": "Chicago"},
	}
	queryContext.AddTable("users", users)
	
	// Create query: SELECT name, age FROM users WHERE city = 'New York'
	selectQuery := NewSelectExpression(
		[]string{"name", "age"},
		"users",
		NewEqualsExpression("city", "New York"),
	)
	
	result = selectQuery.Execute(queryContext)
	fmt.Println("Query: SELECT name, age FROM users WHERE city = 'New York'")
	fmt.Printf("Result: %v\n", result)
	
	// Create query with AND condition
	andQuery := NewSelectExpression(
		[]string{"name", "age"},
		"users",
		NewAndExpression(
			NewEqualsExpression("city", "New York"),
			NewGreaterThanExpression("age", 25),
		),
	)
	
	result = andQuery.Execute(queryContext)
	fmt.Println("Query: SELECT name, age FROM users WHERE city = 'New York' AND age > 25")
	fmt.Printf("Result: %v\n", result)
	fmt.Println()

	// 3. CONFIGURATION INTERPRETER
	fmt.Println("3. CONFIGURATION INTERPRETER:")
	configContext := NewConfigContext()
	configContext.SetValue("app_name", "MyApp")
	configContext.SetValue("version", "1.0.0")
	configContext.SetValue("port", 8080)
	
	// Create interpolation expression
	interpolation := NewInterpolationExpression(
		"${app_name} version ${version} running on port ${port}",
		configContext,
	)
	
	result = interpolation.Evaluate(configContext)
	fmt.Printf("Interpolation: ${app_name} version ${version} running on port ${port}\n")
	fmt.Printf("Result: %v\n", result)
	fmt.Println()

	// 4. RULE ENGINE INTERPRETER
	fmt.Println("4. RULE ENGINE INTERPRETER:")
	ruleContext := NewRuleContext()
	ruleContext.SetFact("user_age", 25.0)
	ruleContext.SetFact("user_role", "admin")
	ruleContext.SetFact("user_balance", 1000.0)
	
	// Create rule: user_role = 'admin' AND user_balance > 500
	rule := NewRuleAndExpression(
		NewFactExpression("user_role", "admin"),
		NewGreaterThanExpression("user_balance", 500.0),
	)
	
	result = rule.Evaluate(ruleContext)
	fmt.Println("Rule: user_role = 'admin' AND user_balance > 500")
	fmt.Printf("Result: %v\n", result)
	
	// Test with different facts
	ruleContext.SetFact("user_balance", 300.0)
	result = rule.Evaluate(ruleContext)
	fmt.Println("Rule: user_role = 'admin' AND user_balance > 500 (with balance 300)")
	fmt.Printf("Result: %v\n", result)
	fmt.Println()

	// 5. EXPRESSION PARSER
	fmt.Println("5. EXPRESSION PARSER:")
	parser := NewExpressionParser()
	
	// Parse and evaluate expressions
	expressions := []string{
		"10 + 5",
		"20 - 8",
		"6 * 7",
		"15 / 3",
		"x + y",
	}
	
	for _, exprStr := range expressions {
		expr := parser.Parse(exprStr)
		result := expr.Interpret(context)
		fmt.Printf("Expression: %s\n", exprStr)
		fmt.Printf("Result: %v\n", result)
	}
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
