// 🪞 REFLECTION & METAPROGRAMMING DEMONSTRATION
// Advanced reflection techniques and dynamic programming in Go
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ============================================================================
// DYNAMIC STRUCT BUILDER
// ============================================================================

type DynamicStructBuilder struct {
	fields []reflect.StructField
	tags   map[string]map[string]string
}

func NewDynamicStructBuilder() *DynamicStructBuilder {
	return &DynamicStructBuilder{
		fields: make([]reflect.StructField, 0),
		tags:   make(map[string]map[string]string),
	}
}

func (dsb *DynamicStructBuilder) AddField(name string, fieldType reflect.Type, jsonTag string) *DynamicStructBuilder {
	field := reflect.StructField{
		Name: name,
		Type: fieldType,
		Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, jsonTag)),
	}
	
	dsb.fields = append(dsb.fields, field)
	return dsb
}

func (dsb *DynamicStructBuilder) AddFieldWithTags(name string, fieldType reflect.Type, tags map[string]string) *DynamicStructBuilder {
	tagParts := make([]string, 0, len(tags))
	for key, value := range tags {
		tagParts = append(tagParts, fmt.Sprintf(`%s:"%s"`, key, value))
	}
	
	field := reflect.StructField{
		Name: name,
		Type: fieldType,
		Tag:  reflect.StructTag(strings.Join(tagParts, " ")),
	}
	
	dsb.fields = append(dsb.fields, field)
	return dsb
}

func (dsb *DynamicStructBuilder) Build() reflect.Type {
	return reflect.StructOf(dsb.fields)
}

func (dsb *DynamicStructBuilder) CreateInstance() interface{} {
	structType := dsb.Build()
	return reflect.New(structType).Interface()
}

// ============================================================================
// DYNAMIC METHOD INVOKER
// ============================================================================

type MethodInvoker struct {
	obj        interface{}
	methodCache map[string]reflect.Method
}

func NewMethodInvoker(obj interface{}) *MethodInvoker {
	invoker := &MethodInvoker{
		obj:        obj,
		methodCache: make(map[string]reflect.Method),
	}
	
	// Cache all methods
	objType := reflect.TypeOf(obj)
	for i := 0; i < objType.NumMethod(); i++ {
		method := objType.Method(i)
		invoker.methodCache[method.Name] = method
	}
	
	return invoker
}

func (mi *MethodInvoker) Call(methodName string, args ...interface{}) ([]interface{}, error) {
	method, exists := mi.methodCache[methodName]
	if !exists {
		return nil, fmt.Errorf("method %s not found", methodName)
	}
	
	// Convert arguments to reflect.Values
	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}
	
	// Call the method (include receiver as first argument)
	allArgs := append([]reflect.Value{reflect.ValueOf(mi.obj)}, argValues...)
	results := method.Func.Call(allArgs)
	
	// Convert results back to interface{}
	resultInterfaces := make([]interface{}, len(results))
	for i, result := range results {
		resultInterfaces[i] = result.Interface()
	}
	
	return resultInterfaces, nil
}

func (mi *MethodInvoker) GetMethodNames() []string {
	names := make([]string, 0, len(mi.methodCache))
	for name := range mi.methodCache {
		names = append(names, name)
	}
	return names
}

// ============================================================================
// DYNAMIC INTERFACE BUILDER
// ============================================================================

type InterfaceBuilder struct {
	methods []reflect.Method
}

func NewInterfaceBuilder() *InterfaceBuilder {
	return &InterfaceBuilder{
		methods: make([]reflect.Method, 0),
	}
}

func (ib *InterfaceBuilder) AddMethod(name string, inTypes []reflect.Type, outTypes []reflect.Type) *InterfaceBuilder {
	method := reflect.Method{
		Name: name,
		Type: reflect.FuncOf(inTypes, outTypes, false),
	}
	
	ib.methods = append(ib.methods, method)
	return ib
}

func (ib *InterfaceBuilder) Build() reflect.Type {
	// Note: Go doesn't support creating interfaces at runtime
	// This is a conceptual demonstration
	return nil
}

// ============================================================================
// DYNAMIC PROXY
// ============================================================================

type DynamicProxy struct {
	target    interface{}
	handlers  map[string]func([]interface{}) ([]interface{}, error)
	invoker   *MethodInvoker
}

func NewDynamicProxy(target interface{}) *DynamicProxy {
	return &DynamicProxy{
		target:   target,
		handlers: make(map[string]func([]interface{}) ([]interface{}, error)),
		invoker:  NewMethodInvoker(target),
	}
}

func (dp *DynamicProxy) AddHandler(methodName string, handler func([]interface{}) ([]interface{}, error)) {
	dp.handlers[methodName] = handler
}

func (dp *DynamicProxy) Call(methodName string, args ...interface{}) ([]interface{}, error) {
	// Check if we have a custom handler
	if handler, exists := dp.handlers[methodName]; exists {
		return handler(args)
	}
	
	// Fall back to direct method call
	return dp.invoker.Call(methodName, args...)
}

// ============================================================================
// DYNAMIC SERIALIZER
// ============================================================================

type DynamicSerializer struct {
	typeCache map[reflect.Type]reflect.Type
}

func NewDynamicSerializer() *DynamicSerializer {
	return &DynamicSerializer{
		typeCache: make(map[reflect.Type]reflect.Type),
	}
}

func (ds *DynamicSerializer) Serialize(obj interface{}) ([]byte, error) {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()
	
	// Handle pointers
	if objType.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
		objType = objType.Elem()
	}
	
	// Create a map to hold the serialized data
	result := make(map[string]interface{})
	
	// Iterate through struct fields
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)
		
		// Get JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		
		// Serialize field value
		serializedValue := ds.serializeValue(fieldValue)
		result[jsonTag] = serializedValue
	}
	
	return json.Marshal(result)
}

func (ds *DynamicSerializer) Deserialize(data []byte, targetType reflect.Type) (interface{}, error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}
	
	// Create new instance of target type
	targetValue := reflect.New(targetType).Elem()
	
	// Populate fields
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		fieldValue := targetValue.Field(i)
		
		// Get JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		
		// Get value from JSON data
		if jsonValue, exists := jsonData[jsonTag]; exists {
			ds.setFieldValue(fieldValue, jsonValue)
		}
	}
	
	return targetValue.Interface(), nil
}

func (ds *DynamicSerializer) serializeValue(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint()
	case reflect.Float32, reflect.Float64:
		return value.Float()
	case reflect.Bool:
		return value.Bool()
	case reflect.Slice:
		slice := make([]interface{}, value.Len())
		for i := 0; i < value.Len(); i++ {
			slice[i] = ds.serializeValue(value.Index(i))
		}
		return slice
	case reflect.Map:
		mapResult := make(map[string]interface{})
		for _, key := range value.MapKeys() {
			mapResult[key.String()] = ds.serializeValue(value.MapIndex(key))
		}
		return mapResult
	case reflect.Struct:
		structResult := make(map[string]interface{})
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			fieldValue := value.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}
			structResult[jsonTag] = ds.serializeValue(fieldValue)
		}
		return structResult
	default:
		return value.Interface()
	}
}

func (ds *DynamicSerializer) setFieldValue(fieldValue reflect.Value, jsonValue interface{}) {
	if !fieldValue.CanSet() {
		return
	}
	
	jsonValueType := reflect.TypeOf(jsonValue)
	fieldType := fieldValue.Type()
	
	// Try direct assignment first
	if jsonValueType.AssignableTo(fieldType) {
		fieldValue.Set(reflect.ValueOf(jsonValue))
		return
	}
	
	// Try conversion
	if jsonValueType.ConvertibleTo(fieldType) {
		fieldValue.Set(reflect.ValueOf(jsonValue).Convert(fieldType))
		return
	}
	
	// Handle special cases
	switch fieldType.Kind() {
	case reflect.String:
		fieldValue.SetString(fmt.Sprintf("%v", jsonValue))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intVal, ok := jsonValue.(float64); ok {
			fieldValue.SetInt(int64(intVal))
		}
	case reflect.Float32, reflect.Float64:
		if floatVal, ok := jsonValue.(float64); ok {
			fieldValue.SetFloat(floatVal)
		}
	case reflect.Bool:
		if boolVal, ok := jsonValue.(bool); ok {
			fieldValue.SetBool(boolVal)
		}
	}
}

// ============================================================================
// DYNAMIC VALIDATOR
// ============================================================================

type ValidationRule struct {
	FieldName string
	Rule      func(interface{}) error
	Message   string
}

type DynamicValidator struct {
	rules []ValidationRule
}

func NewDynamicValidator() *DynamicValidator {
	return &DynamicValidator{
		rules: make([]ValidationRule, 0),
	}
}

func (dv *DynamicValidator) AddRule(fieldName string, rule func(interface{}) error, message string) {
	dv.rules = append(dv.rules, ValidationRule{
		FieldName: fieldName,
		Rule:      rule,
		Message:   message,
	})
}

func (dv *DynamicValidator) Validate(obj interface{}) []error {
	var errors []error
	
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()
	
	// Handle pointers
	if objType.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
		objType = objType.Elem()
	}
	
	// Create field map
	fieldMap := make(map[string]reflect.Value)
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)
		
		// Get JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		
		fieldMap[jsonTag] = fieldValue
	}
	
	// Apply validation rules
	for _, rule := range dv.rules {
		if fieldValue, exists := fieldMap[rule.FieldName]; exists {
			if err := rule.Rule(fieldValue.Interface()); err != nil {
				errors = append(errors, fmt.Errorf("%s: %s", rule.FieldName, rule.Message))
			}
		}
	}
	
	return errors
}

// ============================================================================
// DYNAMIC ORM
// ============================================================================

type DynamicORM struct {
	tables map[string]reflect.Type
}

func NewDynamicORM() *DynamicORM {
	return &DynamicORM{
		tables: make(map[string]reflect.Type),
	}
}

func (orm *DynamicORM) RegisterTable(name string, modelType reflect.Type) {
	orm.tables[name] = modelType
}

func (orm *DynamicORM) CreateTable(name string) string {
	modelType, exists := orm.tables[name]
	if !exists {
		return fmt.Sprintf("Table %s not registered", name)
	}
	
	var columns []string
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		
		dbType := orm.getDBType(field.Type)
		columns = append(columns, fmt.Sprintf("%s %s", jsonTag, dbType))
	}
	
	return fmt.Sprintf("CREATE TABLE %s (%s)", name, strings.Join(columns, ", "))
}

func (orm *DynamicORM) InsertQuery(name string, obj interface{}) string {
	_, exists := orm.tables[name]
	if !exists {
		return fmt.Sprintf("Table %s not registered", name)
	}
	
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()
	
	// Handle pointers
	if objType.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
		objType = objType.Elem()
	}
	
	var columns []string
	var placeholders []string
	var values []interface{}
	
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)
		
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		
		columns = append(columns, jsonTag)
		placeholders = append(placeholders, "?")
		values = append(values, fieldValue.Interface())
	}
	
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", 
		name, 
		strings.Join(columns, ", "), 
		strings.Join(placeholders, ", "))
}

func (orm *DynamicORM) SelectQuery(name string, conditions map[string]interface{}) string {
	modelType, exists := orm.tables[name]
	if !exists {
		return fmt.Sprintf("Table %s not registered", name)
	}
	
	var columns []string
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		columns = append(columns, jsonTag)
	}
	
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), name)
	
	if len(conditions) > 0 {
		var whereClauses []string
		for key, value := range conditions {
			whereClauses = append(whereClauses, fmt.Sprintf("%s = '%v'", key, value))
		}
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}
	
	return query
}

func (orm *DynamicORM) getDBType(goType reflect.Type) string {
	switch goType.Kind() {
	case reflect.String:
		return "VARCHAR(255)"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "INT"
	case reflect.Float32, reflect.Float64:
		return "FLOAT"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Struct:
		if goType == reflect.TypeOf(time.Time{}) {
			return "TIMESTAMP"
		}
		return "TEXT"
	default:
		return "TEXT"
	}
}

// ============================================================================
// DYNAMIC MIDDLEWARE
// ============================================================================

type MiddlewareFunc func(interface{}, []interface{}) ([]interface{}, error)

type DynamicMiddleware struct {
	middlewares []MiddlewareFunc
}

func NewDynamicMiddleware() *DynamicMiddleware {
	return &DynamicMiddleware{
		middlewares: make([]MiddlewareFunc, 0),
	}
}

func (dm *DynamicMiddleware) AddMiddleware(middleware MiddlewareFunc) {
	dm.middlewares = append(dm.middlewares, middleware)
}

func (dm *DynamicMiddleware) Execute(target interface{}, methodName string, args []interface{}) ([]interface{}, error) {
	// Get the target method
	targetType := reflect.TypeOf(target)
	method, exists := targetType.MethodByName(methodName)
	if !exists {
		return nil, fmt.Errorf("method %s not found", methodName)
	}
	
	// Execute middlewares in order
	for _, middleware := range dm.middlewares {
		var err error
		args, err = middleware(target, args)
		if err != nil {
			return nil, err
		}
	}
	
	// Execute the actual method (include receiver as first argument)
	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}
	
	allArgs := append([]reflect.Value{reflect.ValueOf(target)}, argValues...)
	results := method.Func.Call(allArgs)
	
	// Convert results back to interface{}
	resultInterfaces := make([]interface{}, len(results))
	for i, result := range results {
		resultInterfaces[i] = result.Interface()
	}
	
	return resultInterfaces, nil
}

// ============================================================================
// SAMPLE DATA MODELS
// ============================================================================

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	InStock     bool    `json:"in_stock"`
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

// ============================================================================
// SAMPLE SERVICE CLASSES
// ============================================================================

type UserService struct {
	users []User
}

func NewUserService() *UserService {
	return &UserService{
		users: make([]User, 0),
	}
}

func (us *UserService) CreateUser(name, email string, age int) User {
	user := User{
		ID:       len(us.users) + 1,
		Name:     name,
		Email:    email,
		Age:      age,
		IsActive: true,
	}
	us.users = append(us.users, user)
	return user
}

func (us *UserService) GetUser(id int) (User, error) {
	for _, user := range us.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}

func (us *UserService) GetAllUsers() []User {
	return us.users
}

func (us *UserService) UpdateUser(id int, name, email string, age int) error {
	for i, user := range us.users {
		if user.ID == id {
			us.users[i].Name = name
			us.users[i].Email = email
			us.users[i].Age = age
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (us *UserService) DeleteUser(id int) error {
	for i, user := range us.users {
		if user.ID == id {
			us.users = append(us.users[:i], us.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

// ============================================================================
// MAIN DEMONSTRATION
// ============================================================================

func main() {
	fmt.Println("🪞 REFLECTION & METAPROGRAMMING DEMONSTRATION")
	fmt.Println("==============================================")
	fmt.Println()
	
	// Demonstrate dynamic struct builder
	demonstrateDynamicStructBuilder()
	
	// Demonstrate method invoker
	demonstrateMethodInvoker()
	
	// Demonstrate dynamic proxy
	demonstrateDynamicProxy()
	
	// Demonstrate dynamic serializer
	demonstrateDynamicSerializer()
	
	// Demonstrate dynamic validator
	demonstrateDynamicValidator()
	
	// Demonstrate dynamic ORM
	demonstrateDynamicORM()
	
	// Demonstrate dynamic middleware
	demonstrateDynamicMiddleware()
	
	fmt.Println("\n🎉 Reflection & Metaprogramming Demonstration Complete!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Dynamic struct creation")
	fmt.Println("✅ Method invocation and reflection")
	fmt.Println("✅ Dynamic proxy patterns")
	fmt.Println("✅ Dynamic serialization/deserialization")
	fmt.Println("✅ Dynamic validation")
	fmt.Println("✅ Dynamic ORM generation")
	fmt.Println("✅ Dynamic middleware execution")
}

func demonstrateDynamicStructBuilder() {
	fmt.Println("🏗️ Dynamic Struct Builder:")
	
	builder := NewDynamicStructBuilder()
	
	// Build a dynamic struct
	dynamicType := builder.
		AddField("ID", reflect.TypeOf(0), "id").
		AddField("Name", reflect.TypeOf(""), "name").
		AddField("Email", reflect.TypeOf(""), "email").
		AddFieldWithTags("CreatedAt", reflect.TypeOf(time.Time{}), map[string]string{
			"json": "created_at",
			"db":   "created_at",
		}).
		Build()
	
	fmt.Printf("Dynamic struct type: %s\n", dynamicType.String())
	
	// Create an instance
	instance := reflect.New(dynamicType).Elem()
	instance.FieldByName("ID").SetInt(1)
	instance.FieldByName("Name").SetString("John Doe")
	instance.FieldByName("Email").SetString("john@example.com")
	instance.FieldByName("CreatedAt").Set(reflect.ValueOf(time.Now()))
	
	fmt.Printf("Dynamic struct instance: %+v\n", instance.Interface())
	fmt.Println()
}

func demonstrateMethodInvoker() {
	fmt.Println("🔧 Method Invoker:")
	
	userService := NewUserService()
	invoker := NewMethodInvoker(userService)
	
	// Get available methods
	methods := invoker.GetMethodNames()
	fmt.Printf("Available methods: %v\n", methods)
	
	// Call methods dynamically
	results, err := invoker.Call("CreateUser", "Alice", "alice@example.com", 25)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("CreateUser result: %v\n", results[0])
	}
	
	results, err = invoker.Call("GetAllUsers")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("GetAllUsers result: %v\n", results[0])
	}
	
	fmt.Println()
}

func demonstrateDynamicProxy() {
	fmt.Println("🎭 Dynamic Proxy:")
	
	userService := NewUserService()
	proxy := NewDynamicProxy(userService)
	
	// Add logging handler
	proxy.AddHandler("CreateUser", func(args []interface{}) ([]interface{}, error) {
		fmt.Printf("🔍 Logging: Creating user with name=%s, email=%s, age=%d\n", 
			args[0], args[1], args[2])
		
		// Call original method
		invoker := NewMethodInvoker(userService)
		return invoker.Call("CreateUser", args...)
	})
	
	// Add validation handler
	proxy.AddHandler("CreateUser", func(args []interface{}) ([]interface{}, error) {
		name := args[0].(string)
		email := args[1].(string)
		age := args[2].(int)
		
		if name == "" {
			return nil, fmt.Errorf("name cannot be empty")
		}
		if email == "" {
			return nil, fmt.Errorf("email cannot be empty")
		}
		if age < 0 {
			return nil, fmt.Errorf("age cannot be negative")
		}
		
		// Call original method
		invoker := NewMethodInvoker(userService)
		return invoker.Call("CreateUser", args...)
	})
	
	// Use proxy
	results, err := proxy.Call("CreateUser", "Bob", "bob@example.com", 30)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Proxy result: %v\n", results[0])
	}
	
	fmt.Println()
}

func demonstrateDynamicSerializer() {
	fmt.Println("📦 Dynamic Serializer:")
	
	serializer := NewDynamicSerializer()
	
	// Create a user
	user := User{
		ID:       1,
		Name:     "Charlie",
		Email:    "charlie@example.com",
		Age:      35,
		IsActive: true,
	}
	
	// Serialize
	data, err := serializer.Serialize(user)
	if err != nil {
		fmt.Printf("Serialization error: %v\n", err)
	} else {
		fmt.Printf("Serialized data: %s\n", string(data))
	}
	
	// Deserialize
	deserialized, err := serializer.Deserialize(data, reflect.TypeOf(User{}))
	if err != nil {
		fmt.Printf("Deserialization error: %v\n", err)
	} else {
		fmt.Printf("Deserialized user: %+v\n", deserialized)
	}
	
	fmt.Println()
}

func demonstrateDynamicValidator() {
	fmt.Println("✅ Dynamic Validator:")
	
	validator := NewDynamicValidator()
	
	// Add validation rules
	validator.AddRule("name", func(value interface{}) error {
		if value.(string) == "" {
			return fmt.Errorf("name is required")
		}
		return nil
	}, "name is required")
	
	validator.AddRule("email", func(value interface{}) error {
		email := value.(string)
		if email == "" {
			return fmt.Errorf("email is required")
		}
		if !strings.Contains(email, "@") {
			return fmt.Errorf("email must contain @")
		}
		return nil
	}, "email must be valid")
	
	validator.AddRule("age", func(value interface{}) error {
		age := value.(int)
		if age < 0 {
			return fmt.Errorf("age must be positive")
		}
		if age > 150 {
			return fmt.Errorf("age must be less than 150")
		}
		return nil
	}, "age must be between 0 and 150")
	
	// Test validation
	validUser := User{
		Name:  "David",
		Email: "david@example.com",
		Age:   28,
	}
	
	invalidUser := User{
		Name:  "",
		Email: "invalid-email",
		Age:   -5,
	}
	
	// Validate valid user
	errors := validator.Validate(validUser)
	if len(errors) == 0 {
		fmt.Println("✅ Valid user passed validation")
	} else {
		for _, err := range errors {
			fmt.Printf("❌ Validation error: %v\n", err)
		}
	}
	
	// Validate invalid user
	errors = validator.Validate(invalidUser)
	if len(errors) == 0 {
		fmt.Println("✅ Invalid user passed validation")
	} else {
		fmt.Println("❌ Invalid user failed validation:")
		for _, err := range errors {
			fmt.Printf("   - %v\n", err)
		}
	}
	
	fmt.Println()
}

func demonstrateDynamicORM() {
	fmt.Println("🗄️ Dynamic ORM:")
	
	orm := NewDynamicORM()
	
	// Register models
	orm.RegisterTable("users", reflect.TypeOf(User{}))
	orm.RegisterTable("products", reflect.TypeOf(Product{}))
	orm.RegisterTable("orders", reflect.TypeOf(Order{}))
	
	// Generate CREATE TABLE statements
	fmt.Println("CREATE TABLE statements:")
	fmt.Println(orm.CreateTable("users"))
	fmt.Println(orm.CreateTable("products"))
	fmt.Println(orm.CreateTable("orders"))
	
	// Generate INSERT statements
	user := User{
		ID:       1,
		Name:     "Eve",
		Email:    "eve@example.com",
		Age:      32,
		IsActive: true,
	}
	
	fmt.Println("\nINSERT statements:")
	fmt.Println(orm.InsertQuery("users", user))
	
	// Generate SELECT statements
	fmt.Println("\nSELECT statements:")
	fmt.Println(orm.SelectQuery("users", map[string]interface{}{"is_active": true}))
	fmt.Println(orm.SelectQuery("users", map[string]interface{}{"age": 25}))
	
	fmt.Println()
}

func demonstrateDynamicMiddleware() {
	fmt.Println("🔗 Dynamic Middleware:")
	
	userService := NewUserService()
	middleware := NewDynamicMiddleware()
	
	// Add logging middleware
	middleware.AddMiddleware(func(target interface{}, args []interface{}) ([]interface{}, error) {
		fmt.Printf("🔍 Logging: Calling method with args: %v\n", args)
		return args, nil
	})
	
	// Add timing middleware
	middleware.AddMiddleware(func(target interface{}, args []interface{}) ([]interface{}, error) {
		start := time.Now()
		fmt.Printf("⏱️ Timing: Method started at %v\n", start)
		
		// This is a simplified example - in practice, you'd need to track timing across calls
		return args, nil
	})
	
	// Add authentication middleware
	middleware.AddMiddleware(func(target interface{}, args []interface{}) ([]interface{}, error) {
		fmt.Printf("🔐 Auth: Checking permissions for method call\n")
		// Simulate auth check
		return args, nil
	})
	
	// Execute method through middleware
	results, err := middleware.Execute(userService, "CreateUser", []interface{}{"Frank", "frank@example.com", 40})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Middleware result: %v\n", results[0])
	}
	
	fmt.Println()
}
