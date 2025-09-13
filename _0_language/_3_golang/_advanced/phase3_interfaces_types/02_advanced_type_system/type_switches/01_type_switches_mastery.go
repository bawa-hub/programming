package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// 🔄 TYPE SWITCHES MASTERY
// Understanding type switches and runtime type checking in Go

// Define interfaces for type switching
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Stringer interface {
	String() string
}

func main() {
	fmt.Println("🔄 TYPE SWITCHES MASTERY")
	fmt.Println("========================")
	fmt.Println()

	// 1. Basic Type Switches
	basicTypeSwitches()
	fmt.Println()

	// 2. Advanced Type Switches
	advancedTypeSwitches()
	fmt.Println()

	// 3. Type Safety Patterns
	typeSafetyPatterns()
	fmt.Println()

	// 4. Interface Satisfaction
	interfaceSatisfaction()
	fmt.Println()

	// 5. Type Guard Functions
	typeGuardFunctions()
	fmt.Println()

	// 6. Runtime Type Checking
	runtimeTypeChecking()
	fmt.Println()

	// 7. Type Validation
	typeValidation()
	fmt.Println()

	// 8. Best Practices
	typeSwitchesBestPractices()
}

// 1. Basic Type Switches
func basicTypeSwitches() {
	fmt.Println("1. Basic Type Switches:")
	fmt.Println("Understanding basic type switch syntax...")

	// Demonstrate basic type switches
	basicTypeSwitchExample()
	
	// Show type assertions in switches
	typeAssertionsInSwitches()
	
	// Demonstrate default cases
	defaultCases()
}

func basicTypeSwitchExample() {
	fmt.Println("  📊 Basic type switch example:")
	
	// Test different types
	testTypes := []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1, "b": 2},
	}
	
	for _, value := range testTypes {
		processValue(value)
	}
}

func processValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("    Integer: %d\n", v)
	case string:
		fmt.Printf("    String: %s\n", v)
	case float64:
		fmt.Printf("    Float: %.2f\n", v)
	case bool:
		fmt.Printf("    Boolean: %t\n", v)
	case []int:
		fmt.Printf("    Slice: %v\n", v)
	case map[string]int:
		fmt.Printf("    Map: %v\n", v)
	default:
		fmt.Printf("    Unknown type: %T\n", v)
	}
}

func typeAssertionsInSwitches() {
	fmt.Println("  📊 Type assertions in switches:")
	
	// Test type assertions
	testValues := []interface{}{
		"123",
		456,
		"hello world",
		789.123,
	}
	
	for _, value := range testValues {
		processWithTypeAssertion(value)
	}
}

func processWithTypeAssertion(value interface{}) {
	switch v := value.(type) {
	case string:
		if len(v) > 5 {
			fmt.Printf("    Long string: %s\n", v)
		} else {
			fmt.Printf("    Short string: %s\n", v)
		}
	case int:
		if v > 100 {
			fmt.Printf("    Large integer: %d\n", v)
		} else {
			fmt.Printf("    Small integer: %d\n", v)
		}
	case float64:
		fmt.Printf("    Float: %.3f\n", v)
	default:
		fmt.Printf("    Other type: %T\n", v)
	}
}

func defaultCases() {
	fmt.Println("  📊 Default cases:")
	
	// Test with default case
	testValues := []interface{}{
		"test",
		42,
		[]string{"a", "b"},
		complex(1, 2),
	}
	
	for _, value := range testValues {
		processWithDefault(value)
	}
}

func processWithDefault(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("    String: %s\n", v)
	case int:
		fmt.Printf("    Integer: %d\n", v)
	default:
		fmt.Printf("    Default case for type %T: %v\n", v, v)
	}
}

// 2. Advanced Type Switches
func advancedTypeSwitches() {
	fmt.Println("2. Advanced Type Switches:")
	fmt.Println("Understanding advanced type switch patterns...")

	// Demonstrate multiple type cases
	multipleTypeCases()
	
	// Show type switch with interfaces
	typeSwitchWithInterfaces()
	
	// Demonstrate complex type patterns
	complexTypePatterns()
}

func multipleTypeCases() {
	fmt.Println("  📊 Multiple type cases:")
	
	// Test with multiple type cases
	testValues := []interface{}{
		int8(1),
		int16(2),
		int32(3),
		int64(4),
		uint8(5),
		uint16(6),
		uint32(7),
		uint64(8),
	}
	
	for _, value := range testValues {
		processNumericTypes(value)
	}
}

func processNumericTypes(value interface{}) {
	switch v := value.(type) {
	case int8, int16, int32, int64:
		fmt.Printf("    Signed integer: %d (type: %T)\n", v, v)
	case uint8, uint16, uint32, uint64:
		fmt.Printf("    Unsigned integer: %d (type: %T)\n", v, v)
	default:
		fmt.Printf("    Non-integer type: %T\n", v)
	}
}

func typeSwitchWithInterfaces() {
	fmt.Println("  📊 Type switch with interfaces:")
	
	// Define interfaces
	type Reader interface {
		Read([]byte) (int, error)
	}
	
	type Writer interface {
		Write([]byte) (int, error)
	}
	
	type ReadWriter interface {
		Reader
		Writer
	}
	
	// Test with different interface types
	testInterfaces := []interface{}{
		&StringReader{},
		&StringWriter{},
		&StringReadWriter{},
		&StringProcessor{},
	}
	
	for _, value := range testInterfaces {
		processInterfaceTypes(value)
	}
}

type StringReader struct{}

func (sr *StringReader) Read(data []byte) (int, error) {
	return 0, nil
}

type StringWriter struct{}

func (sw *StringWriter) Write(data []byte) (int, error) {
	return 0, nil
}

type StringReadWriter struct{}

func (srw *StringReadWriter) Read(data []byte) (int, error) {
	return 0, nil
}

func (srw *StringReadWriter) Write(data []byte) (int, error) {
	return 0, nil
}

type StringProcessor struct{}

func (sp *StringProcessor) Process(data string) string {
	return data
}

func processInterfaceTypes(value interface{}) {
	switch v := value.(type) {
	case Reader:
		fmt.Printf("    Reader interface: %T\n", v)
	case Writer:
		fmt.Printf("    Writer interface: %T\n", v)
	case ReadWriter:
		fmt.Printf("    ReadWriter interface: %T\n", v)
	default:
		fmt.Printf("    Other interface: %T\n", v)
	}
}

func complexTypePatterns() {
	fmt.Println("  📊 Complex type patterns:")
	
	// Test with complex types
	testValues := []interface{}{
		[]int{1, 2, 3},
		[]string{"a", "b", "c"},
		map[string]int{"x": 1, "y": 2},
		map[int]string{1: "a", 2: "b"},
		[3]int{1, 2, 3},
		[2]string{"a", "b"},
	}
	
	for _, value := range testValues {
		processComplexTypes(value)
	}
}

func processComplexTypes(value interface{}) {
	switch v := value.(type) {
	case []int:
		fmt.Printf("    Int slice: %v (len: %d)\n", v, len(v))
	case []string:
		fmt.Printf("    String slice: %v (len: %d)\n", v, len(v))
	case map[string]int:
		fmt.Printf("    String->Int map: %v (len: %d)\n", v, len(v))
	case map[int]string:
		fmt.Printf("    Int->String map: %v (len: %d)\n", v, len(v))
	case [3]int:
		fmt.Printf("    Int array[3]: %v\n", v)
	case [2]string:
		fmt.Printf("    String array[2]: %v\n", v)
	default:
		fmt.Printf("    Complex type: %T\n", v)
	}
}

// 3. Type Safety Patterns
func typeSafetyPatterns() {
	fmt.Println("3. Type Safety Patterns:")
	fmt.Println("Understanding type safety patterns...")

	// Demonstrate safe type assertions
	safeTypeAssertions()
	
	// Show type guard functions
	typeGuardFunctions()
	
	// Demonstrate runtime type checking
	runtimeTypeChecking()
}

func safeTypeAssertions() {
	fmt.Println("  📊 Safe type assertions:")
	
	// Test safe type assertions
	testValues := []interface{}{
		"hello",
		42,
		"123",
		3.14,
		"456",
	}
	
	for _, value := range testValues {
		processWithSafeAssertion(value)
	}
}

func processWithSafeAssertion(value interface{}) {
	// Safe type assertion with ok
	if str, ok := value.(string); ok {
		fmt.Printf("    String value: %s\n", str)
		
		// Try to convert to int
		if intVal, err := strconv.Atoi(str); err == nil {
			fmt.Printf("    Converted to int: %d\n", intVal)
		} else {
			fmt.Printf("    Cannot convert to int: %s\n", str)
		}
	} else if intVal, ok := value.(int); ok {
		fmt.Printf("    Integer value: %d\n", intVal)
	} else if floatVal, ok := value.(float64); ok {
		fmt.Printf("    Float value: %.2f\n", floatVal)
	} else {
		fmt.Printf("    Other type: %T\n", value)
	}
}

func typeGuardFunctions() {
	fmt.Println("  📊 Type guard functions:")
	
	// Test type guard functions
	testValues := []interface{}{
		"hello",
		42,
		"123",
		3.14,
		[]int{1, 2, 3},
	}
	
	for _, value := range testValues {
		processWithTypeGuards(value)
	}
}

func processWithTypeGuards(value interface{}) {
	if isString(value) {
		fmt.Printf("    String: %s\n", value.(string))
	} else if isInt(value) {
		fmt.Printf("    Integer: %d\n", value.(int))
	} else if isFloat(value) {
		fmt.Printf("    Float: %.2f\n", value.(float64))
	} else if isSlice(value) {
		fmt.Printf("    Slice: %v\n", value)
	} else {
		fmt.Printf("    Unknown type: %T\n", value)
	}
}

func isString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func isInt(value interface{}) bool {
	_, ok := value.(int)
	return ok
}

func isFloat(value interface{}) bool {
	_, ok := value.(float64)
	return ok
}

func isSlice(value interface{}) bool {
	_, ok := value.([]int)
	return ok
}

func runtimeTypeChecking() {
	fmt.Println("  📊 Runtime type checking:")
	
	// Test runtime type checking
	testValues := []interface{}{
		"hello",
		42,
		"123",
		3.14,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}
	
	for _, value := range testValues {
		processWithRuntimeChecking(value)
	}
}

func processWithRuntimeChecking(value interface{}) {
	// Use reflect package for runtime type checking
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)
	
	fmt.Printf("    Type: %s, Kind: %s, Value: %v\n", 
		valueType.String(), valueType.Kind(), valueValue.Interface())
	
	// Check specific types at runtime
	switch valueType.Kind() {
	case reflect.String:
		fmt.Printf("      String length: %d\n", valueValue.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("      Integer value: %d\n", valueValue.Int())
	case reflect.Float32, reflect.Float64:
		fmt.Printf("      Float value: %.2f\n", valueValue.Float())
	case reflect.Slice:
		fmt.Printf("      Slice length: %d\n", valueValue.Len())
	case reflect.Map:
		fmt.Printf("      Map length: %d\n", valueValue.Len())
	}
}

// 4. Interface Satisfaction
func interfaceSatisfaction() {
	fmt.Println("4. Interface Satisfaction:")
	fmt.Println("Understanding interface satisfaction verification...")

	// Demonstrate interface satisfaction verification
	interfaceSatisfactionVerification()
	
	// Show dynamic interface checking
	dynamicInterfaceChecking()
	
	// Demonstrate type compatibility
	typeCompatibility()
}

func interfaceSatisfactionVerification() {
	fmt.Println("  📊 Interface satisfaction verification:")
	
	// Define interfaces
	type Stringer interface {
		String() string
	}
	
	type Reader interface {
		Read([]byte) (int, error)
	}
	
	type Writer interface {
		Write([]byte) (int, error)
	}
	
	// Test interface satisfaction
	testValues := []interface{}{
		&StringReader{},
		&StringWriter{},
		&StringReadWriter{},
		&StringProcessor{},
	}
	
	for _, value := range testValues {
		checkInterfaceSatisfaction(value)
	}
}

func checkInterfaceSatisfaction(value interface{}) {
	fmt.Printf("    Checking %T:\n", value)
	
	// Check if value satisfies Stringer interface
	if _, ok := value.(Stringer); ok {
		fmt.Printf("      ✓ Satisfies Stringer interface\n")
	} else {
		fmt.Printf("      ✗ Does not satisfy Stringer interface\n")
	}
	
	// Check if value satisfies Reader interface
	if _, ok := value.(Reader); ok {
		fmt.Printf("      ✓ Satisfies Reader interface\n")
	} else {
		fmt.Printf("      ✗ Does not satisfy Reader interface\n")
	}
	
	// Check if value satisfies Writer interface
	if _, ok := value.(Writer); ok {
		fmt.Printf("      ✓ Satisfies Writer interface\n")
	} else {
		fmt.Printf("      ✗ Does not satisfy Writer interface\n")
	}
}

func dynamicInterfaceChecking() {
	fmt.Println("  📊 Dynamic interface checking:")
	
	// Test dynamic interface checking
	testValues := []interface{}{
		&StringReader{},
		&StringWriter{},
		&StringReadWriter{},
		&StringProcessor{},
	}
	
	for _, value := range testValues {
		checkDynamicInterfaces(value)
	}
}

func checkDynamicInterfaces(value interface{}) {
	fmt.Printf("    Checking %T dynamically:\n", value)
	
	// Use reflect to check interface satisfaction
	valueType := reflect.TypeOf(value)
	
	// Check if type implements Stringer interface
	stringerType := reflect.TypeOf((*Stringer)(nil)).Elem()
	if valueType.Implements(stringerType) {
		fmt.Printf("      ✓ Implements Stringer interface\n")
	} else {
		fmt.Printf("      ✗ Does not implement Stringer interface\n")
	}
	
	// Check if type implements Reader interface
	readerType := reflect.TypeOf((*Reader)(nil)).Elem()
	if valueType.Implements(readerType) {
		fmt.Printf("      ✓ Implements Reader interface\n")
	} else {
		fmt.Printf("      ✗ Does not implement Reader interface\n")
	}
	
	// Check if type implements Writer interface
	writerType := reflect.TypeOf((*Writer)(nil)).Elem()
	if valueType.Implements(writerType) {
		fmt.Printf("      ✓ Implements Writer interface\n")
	} else {
		fmt.Printf("      ✗ Does not implement Writer interface\n")
	}
}

func typeCompatibility() {
	fmt.Println("  📊 Type compatibility:")
	
	// Test type compatibility
	testValues := []interface{}{
		int(42),
		int32(42),
		int64(42),
		float32(42.0),
		float64(42.0),
		string("42"),
	}
	
	for _, value := range testValues {
		checkTypeCompatibility(value)
	}
}

func checkTypeCompatibility(value interface{}) {
	fmt.Printf("    Checking %T (%v):\n", value, value)
	
	// Check if value can be converted to int
	if intVal, ok := value.(int); ok {
		fmt.Printf("      ✓ Can be converted to int: %d\n", intVal)
	} else {
		fmt.Printf("      ✗ Cannot be converted to int\n")
	}
	
	// Check if value can be converted to float64
	if floatVal, ok := value.(float64); ok {
		fmt.Printf("      ✓ Can be converted to float64: %.2f\n", floatVal)
	} else {
		fmt.Printf("      ✗ Cannot be converted to float64\n")
	}
	
	// Check if value can be converted to string
	if strVal, ok := value.(string); ok {
		fmt.Printf("      ✓ Can be converted to string: %s\n", strVal)
	} else {
		fmt.Printf("      ✗ Cannot be converted to string\n")
	}
}

func typeGuardFunctionExamples() {
	fmt.Println("  📊 Type guard function examples:")
	
	// Test type guard functions
	testValues := []interface{}{
		"hello",
		42,
		"123",
		3.14,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}
	
	for _, value := range testValues {
		processWithTypeGuards(value)
	}
}

func genericTypeGuards() {
	fmt.Println("  📊 Generic type guards:")
	
	// Test generic type guards
	testValues := []interface{}{
		"hello",
		42,
		"123",
		3.14,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}
	
	for _, value := range testValues {
		processWithGenericGuards(value)
	}
}

func processWithGenericGuards(value interface{}) {
	// Generic type guard for numeric types
	if isNumeric(value) {
		fmt.Printf("    Numeric value: %v\n", value)
	} else if isTextual(value) {
		fmt.Printf("    Textual value: %v\n", value)
	} else if isCollection(value) {
		fmt.Printf("    Collection value: %v\n", value)
	} else {
		fmt.Printf("    Other value: %v\n", value)
	}
}

func isNumeric(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64,
		 uint, uint8, uint16, uint32, uint64,
		 float32, float64:
		return true
	default:
		return false
	}
}

func isTextual(value interface{}) bool {
	switch value.(type) {
	case string, []byte, []rune:
		return true
	default:
		return false
	}
}

func isCollection(value interface{}) bool {
	switch value.(type) {
	case []int, []string, []interface{},
		 map[string]int, map[string]string, map[string]interface{}:
		return true
	default:
		return false
	}
}

func typeGuardComposition() {
	fmt.Println("  📊 Type guard composition:")
	
	// Test composed type guards
	testValues := []interface{}{
		"hello",
		42,
		"123",
		3.14,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}
	
	for _, value := range testValues {
		processWithComposedGuards(value)
	}
}

func processWithComposedGuards(value interface{}) {
	// Compose type guards
	if isNumeric(value) && isPositive(value) {
		fmt.Printf("    Positive numeric: %v\n", value)
	} else if isTextual(value) && isNotEmpty(value) {
		fmt.Printf("    Non-empty textual: %v\n", value)
	} else if isCollection(value) && hasElements(value) {
		fmt.Printf("    Non-empty collection: %v\n", value)
	} else {
		fmt.Printf("    Other value: %v\n", value)
	}
}

func isPositive(value interface{}) bool {
	switch v := value.(type) {
	case int:
		return v > 0
	case int8:
		return v > 0
	case int16:
		return v > 0
	case int32:
		return v > 0
	case int64:
		return v > 0
	case uint, uint8, uint16, uint32, uint64:
		return true // unsigned types are always positive
	case float32:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}

func isNotEmpty(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return len(v) > 0
	case []byte:
		return len(v) > 0
	case []rune:
		return len(v) > 0
	default:
		return false
	}
}

func hasElements(value interface{}) bool {
	switch v := value.(type) {
	case []int:
		return len(v) > 0
	case []string:
		return len(v) > 0
	case []interface{}:
		return len(v) > 0
	case map[string]int:
		return len(v) > 0
	case map[string]string:
		return len(v) > 0
	case map[string]interface{}:
		return len(v) > 0
	default:
		return false
	}
}

func runtimeTypeCheckingExamples() {
	fmt.Println("  📊 Runtime type checking examples:")
	
	// Test runtime type checking
	testValues := []interface{}{
		"hello",
		42,
		"123",
		3.14,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}
	
	for _, value := range testValues {
		processWithRuntimeTypeChecking(value)
	}
}

func processWithRuntimeTypeChecking(value interface{}) {
	// Use reflect for runtime type checking
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)
	
	fmt.Printf("    Type: %s, Kind: %s\n", valueType.String(), valueType.Kind())
	
	// Check specific properties
	switch valueType.Kind() {
	case reflect.String:
		fmt.Printf("      String length: %d\n", valueValue.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("      Integer value: %d\n", valueValue.Int())
	case reflect.Float32, reflect.Float64:
		fmt.Printf("      Float value: %.2f\n", valueValue.Float())
	case reflect.Slice:
		fmt.Printf("      Slice length: %d, Element type: %s\n", 
			valueValue.Len(), valueType.Elem().String())
	case reflect.Map:
		fmt.Printf("      Map length: %d, Key type: %s, Value type: %s\n", 
			valueValue.Len(), valueType.Key().String(), valueType.Elem().String())
	}
	
	// Use valueValue in the switch statement above
}

func typeIntrospection() {
	fmt.Println("  📊 Type introspection:")
	
	// Test type introspection
	testValues := []interface{}{
		"hello",
		42,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}
	
	for _, value := range testValues {
		introspectType(value)
	}
}

func introspectType(value interface{}) {
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)
	
	fmt.Printf("    Introspecting %T:\n", value)
	fmt.Printf("      Type: %s\n", valueType.String())
	fmt.Printf("      Kind: %s\n", valueType.Kind())
	fmt.Printf("      PkgPath: %s\n", valueType.PkgPath())
	fmt.Printf("      Name: %s\n", valueType.Name())
	fmt.Printf("      Size: %d bytes\n", valueType.Size())
	fmt.Printf("      Align: %d bytes\n", valueType.Align())
	fmt.Printf("      FieldAlign: %d bytes\n", valueType.FieldAlign())
	fmt.Printf("      NumMethod: %d\n", valueType.NumMethod())
	
	// Show methods
	for i := 0; i < valueType.NumMethod(); i++ {
		method := valueType.Method(i)
		fmt.Printf("      Method %d: %s\n", i, method.Name)
	}
	
	// Use valueValue to avoid unused variable error
	_ = valueValue
}

func typeConversion() {
	fmt.Println("  📊 Type conversion:")
	
	// Test type conversion
	testValues := []interface{}{
		"123",
		456,
		"789.123",
		3.14,
		"hello",
	}
	
	for _, value := range testValues {
		convertType(value)
	}
}

func convertType(value interface{}) {
	fmt.Printf("    Converting %T (%v):\n", value, value)
	
	// Try to convert to int
	if intVal, err := strconv.Atoi(fmt.Sprintf("%v", value)); err == nil {
		fmt.Printf("      ✓ Converted to int: %d\n", intVal)
	} else {
		fmt.Printf("      ✗ Cannot convert to int: %v\n", err)
	}
	
	// Try to convert to float64
	if floatVal, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64); err == nil {
		fmt.Printf("      ✓ Converted to float64: %.2f\n", floatVal)
	} else {
		fmt.Printf("      ✗ Cannot convert to float64: %v\n", err)
	}
	
	// Try to convert to string
	fmt.Printf("      ✓ Converted to string: %s\n", fmt.Sprintf("%v", value))
}

// 7. Type Validation
func typeValidation() {
	fmt.Println("7. Type Validation:")
	fmt.Println("Understanding type validation...")

	// Demonstrate type validation
	typeValidationExamples()
	
	// Show custom type validation
	customTypeValidation()
	
	// Demonstrate validation errors
	validationErrors()
}

func typeValidationExamples() {
	fmt.Println("  📊 Type validation examples:")
	
	// Test type validation
	testValues := []interface{}{
		"hello",
		42,
		"123",
		3.14,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}
	
	for _, value := range testValues {
		validateType(value)
	}
}

func validateType(value interface{}) {
	fmt.Printf("    Validating %T (%v):\n", value, value)
	
	// Validate string
	if str, ok := value.(string); ok {
		if len(str) > 0 {
			fmt.Printf("      ✓ Valid string (length: %d)\n", len(str))
		} else {
			fmt.Printf("      ✗ Invalid string (empty)\n")
		}
	}
	
	// Validate int
	if intVal, ok := value.(int); ok {
		if intVal >= 0 {
			fmt.Printf("      ✓ Valid int (non-negative: %d)\n", intVal)
		} else {
			fmt.Printf("      ✗ Invalid int (negative: %d)\n", intVal)
		}
	}
	
	// Validate float
	if floatVal, ok := value.(float64); ok {
		if !isNaN(floatVal) && !isInf(floatVal) {
			fmt.Printf("      ✓ Valid float (%.2f)\n", floatVal)
		} else {
			fmt.Printf("      ✗ Invalid float (NaN or Inf)\n")
		}
	}
	
	// Validate slice
	if slice, ok := value.([]int); ok {
		if len(slice) > 0 {
			fmt.Printf("      ✓ Valid slice (length: %d)\n", len(slice))
		} else {
			fmt.Printf("      ✗ Invalid slice (empty)\n")
		}
	}
	
	// Validate map
	if m, ok := value.(map[string]int); ok {
		if len(m) > 0 {
			fmt.Printf("      ✓ Valid map (length: %d)\n", len(m))
		} else {
			fmt.Printf("      ✗ Invalid map (empty)\n")
		}
	}
}

func isNaN(f float64) bool {
	return f != f
}

func isInf(f float64) bool {
	return math.IsInf(f, 0)
}

func customTypeValidation() {
	fmt.Println("  📊 Custom type validation:")
	
	// Test custom type validation
	testValues := []interface{}{
		"hello",
		"",
		"123",
		"hello world",
		[]int{1, 2, 3},
		[]int{},
		map[string]int{"a": 1},
		map[string]int{},
	}
	
	for _, value := range testValues {
		validateWithCustomRules(value)
	}
}

func validateWithCustomRules(value interface{}) {
	fmt.Printf("    Custom validation for %T (%v):\n", value, value)
	
	// Custom validation rules
	switch v := value.(type) {
	case string:
		if len(v) >= 3 && len(v) <= 10 {
			fmt.Printf("      ✓ Valid string (length 3-10)\n")
		} else {
			fmt.Printf("      ✗ Invalid string (length %d, expected 3-10)\n", len(v))
		}
	case []int:
		if len(v) >= 1 && len(v) <= 5 {
			fmt.Printf("      ✓ Valid slice (length 1-5)\n")
		} else {
			fmt.Printf("      ✗ Invalid slice (length %d, expected 1-5)\n", len(v))
		}
	case map[string]int:
		if len(v) >= 1 && len(v) <= 3 {
			fmt.Printf("      ✓ Valid map (length 1-3)\n")
		} else {
			fmt.Printf("      ✗ Invalid map (length %d, expected 1-3)\n", len(v))
		}
	default:
		fmt.Printf("      ✗ Unsupported type for custom validation\n")
	}
}

func validationErrors() {
	fmt.Println("  📊 Validation errors:")
	
	// Test validation errors
	testValues := []interface{}{
		"",
		"a",
		"very long string that exceeds limit",
		[]int{},
		[]int{1, 2, 3, 4, 5, 6},
		map[string]int{},
		map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
	}
	
	for _, value := range testValues {
		validateWithErrors(value)
	}
}

func validateWithErrors(value interface{}) {
	fmt.Printf("    Validation with errors for %T (%v):\n", value, value)
	
	// Validation with error reporting
	switch v := value.(type) {
	case string:
		if len(v) < 3 {
			fmt.Printf("      ✗ Error: string too short (length %d, minimum 3)\n", len(v))
		} else if len(v) > 10 {
			fmt.Printf("      ✗ Error: string too long (length %d, maximum 10)\n", len(v))
		} else {
			fmt.Printf("      ✓ Valid string\n")
		}
	case []int:
		if len(v) < 1 {
			fmt.Printf("      ✗ Error: slice too short (length %d, minimum 1)\n", len(v))
		} else if len(v) > 5 {
			fmt.Printf("      ✗ Error: slice too long (length %d, maximum 5)\n", len(v))
		} else {
			fmt.Printf("      ✓ Valid slice\n")
		}
	case map[string]int:
		if len(v) < 1 {
			fmt.Printf("      ✗ Error: map too short (length %d, minimum 1)\n", len(v))
		} else if len(v) > 3 {
			fmt.Printf("      ✗ Error: map too long (length %d, maximum 3)\n", len(v))
		} else {
			fmt.Printf("      ✓ Valid map\n")
		}
	default:
		fmt.Printf("      ✗ Error: unsupported type\n")
	}
}

// 8. Best Practices
func typeSwitchesBestPractices() {
	fmt.Println("8. Type Switches Best Practices:")
	fmt.Println("Best practices for type switches...")

	fmt.Println("  📝 Best Practice 1: Use type switches for type checking")
	fmt.Println("    - Use type switches instead of multiple type assertions")
	fmt.Println("    - Handle all possible types explicitly")
	fmt.Println("    - Use default case for unexpected types")
	
	fmt.Println("  📝 Best Practice 2: Prefer type switches over type assertions")
	fmt.Println("    - Type switches are safer than type assertions")
	fmt.Println("    - Type switches handle multiple types efficiently")
	fmt.Println("    - Type switches provide better error handling")
	
	fmt.Println("  📝 Best Practice 3: Use type guards for complex logic")
	fmt.Println("    - Create type guard functions for reusable logic")
	fmt.Println("    - Compose type guards for complex conditions")
	fmt.Println("    - Use type guards to improve code readability")
	
	fmt.Println("  📝 Best Practice 4: Handle interface satisfaction properly")
	fmt.Println("    - Check interface satisfaction at runtime")
	fmt.Println("    - Use reflect package for dynamic checking")
	fmt.Println("    - Verify interface compatibility")
	
	fmt.Println("  📝 Best Practice 5: Use type validation for data integrity")
	fmt.Println("    - Validate types before processing")
	fmt.Println("    - Provide clear error messages for invalid types")
	fmt.Println("    - Use custom validation rules when needed")
	
	fmt.Println("  📝 Best Practice 6: Avoid excessive type switching")
	fmt.Println("    - Use type switches only when necessary")
	fmt.Println("    - Consider using interfaces instead of type switching")
	fmt.Println("    - Refactor code to reduce type switching")
	
	fmt.Println("  📝 Best Practice 7: Document type switch behavior")
	fmt.Println("    - Document expected types and behavior")
	fmt.Println("    - Add comments for complex type logic")
	fmt.Println("    - Use descriptive variable names in type switches")
}
