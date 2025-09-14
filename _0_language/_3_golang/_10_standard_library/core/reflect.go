package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Sample structs for demonstration
type Person struct {
	Name    string `json:"name" db:"full_name" validate:"required"`
	Age     int    `json:"age" db:"age" validate:"min=0,max=120"`
	Email   string `json:"email" db:"email" validate:"email"`
	Address Address `json:"address" db:"address"`
}

type Address struct {
	Street string `json:"street" db:"street"`
	City   string `json:"city" db:"city"`
	Zip    string `json:"zip" db:"zip_code"`
}

type Calculator struct {
	Result float64
}

func (c *Calculator) Add(a, b float64) float64 {
	c.Result = a + b
	return c.Result
}

func (c *Calculator) Multiply(a, b float64) float64 {
	c.Result = a * b
	return c.Result
}

func main() {
	fmt.Println("🚀 Go reflect Package Mastery Examples")
	fmt.Println("======================================")

	// 1. Basic Type and Value Operations
	fmt.Println("\n1. Basic Type and Value Operations:")
	
	// Get type and value
	value := 42
	val := reflect.ValueOf(value)
	typ := reflect.TypeOf(value)
	
	fmt.Printf("Value: %v\n", val.Interface())
	fmt.Printf("Type: %s\n", typ)
	fmt.Printf("Kind: %s\n", val.Kind())
	fmt.Printf("Can set: %t\n", val.CanSet())
	fmt.Printf("Can address: %t\n", val.CanAddr())

	// 2. Pointer Operations
	fmt.Println("\n2. Pointer Operations:")
	
	ptr := &value
	ptrVal := reflect.ValueOf(ptr)
	ptrTyp := reflect.TypeOf(ptr)
	
	fmt.Printf("Pointer value: %v\n", ptrVal.Interface())
	fmt.Printf("Pointer type: %s\n", ptrTyp)
	fmt.Printf("Pointer kind: %s\n", ptrVal.Kind())
	fmt.Printf("Element type: %s\n", ptrTyp.Elem())
	fmt.Printf("Element value: %v\n", ptrVal.Elem().Interface())
	
	// Modify value through pointer
	if ptrVal.Elem().CanSet() {
		ptrVal.Elem().SetInt(100)
		fmt.Printf("Modified value: %v\n", value)
	}

	// 3. Slice Operations
	fmt.Println("\n3. Slice Operations:")
	
	slice := []int{1, 2, 3, 4, 5}
	sliceVal := reflect.ValueOf(slice)
	sliceTyp := reflect.TypeOf(slice)
	
	fmt.Printf("Slice: %v\n", sliceVal.Interface())
	fmt.Printf("Slice type: %s\n", sliceTyp)
	fmt.Printf("Slice kind: %s\n", sliceVal.Kind())
	fmt.Printf("Element type: %s\n", sliceTyp.Elem())
	fmt.Printf("Length: %d\n", sliceVal.Len())
	fmt.Printf("Capacity: %d\n", sliceVal.Cap())
	
	// Access slice elements
	for i := 0; i < sliceVal.Len(); i++ {
		fmt.Printf("  [%d] = %v\n", i, sliceVal.Index(i).Interface())
	}

	// 4. Map Operations
	fmt.Println("\n4. Map Operations:")
	
	myMap := map[string]int{"apple": 5, "banana": 3, "orange": 8}
	mapVal := reflect.ValueOf(myMap)
	mapTyp := reflect.TypeOf(myMap)
	
	fmt.Printf("Map: %v\n", mapVal.Interface())
	fmt.Printf("Map type: %s\n", mapTyp)
	fmt.Printf("Map kind: %s\n", mapVal.Kind())
	fmt.Printf("Key type: %s\n", mapTyp.Key())
	fmt.Printf("Value type: %s\n", mapTyp.Elem())
	
	// Iterate over map
	mapKeys := mapVal.MapKeys()
	for _, key := range mapKeys {
		value := mapVal.MapIndex(key)
		fmt.Printf("  %v: %v\n", key.Interface(), value.Interface())
	}

	// 5. Struct Operations
	fmt.Println("\n5. Struct Operations:")
	
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
			Zip:    "10001",
		},
	}
	
	personVal := reflect.ValueOf(person)
	personTyp := reflect.TypeOf(person)
	
	fmt.Printf("Person: %+v\n", personVal.Interface())
	fmt.Printf("Person type: %s\n", personTyp)
	fmt.Printf("Person kind: %s\n", personVal.Kind())
	fmt.Printf("Number of fields: %d\n", personVal.NumField())
	
	// Iterate over struct fields
	for i := 0; i < personVal.NumField(); i++ {
		field := personVal.Field(i)
		fieldType := personTyp.Field(i)
		
		fmt.Printf("  Field %d: %s (%s) = %v\n", 
			i, fieldType.Name, fieldType.Type, field.Interface())
	}

	// 6. Struct Field Access by Name
	fmt.Println("\n6. Struct Field Access by Name:")
	
	// Get field by name
	nameField := personVal.FieldByName("Name")
	if nameField.IsValid() {
		fmt.Printf("Name field: %v\n", nameField.Interface())
	}
	
	// Get field by index path
	streetField := personVal.FieldByIndex([]int{3, 0}) // Address.Street
	if streetField.IsValid() {
		fmt.Printf("Street field: %v\n", streetField.Interface())
	}

	// 7. Struct Tags
	fmt.Println("\n7. Struct Tags:")
	
	for i := 0; i < personTyp.NumField(); i++ {
		field := personTyp.Field(i)
		fmt.Printf("Field: %s\n", field.Name)
		fmt.Printf("  JSON tag: %s\n", field.Tag.Get("json"))
		fmt.Printf("  DB tag: %s\n", field.Tag.Get("db"))
		fmt.Printf("  Validate tag: %s\n", field.Tag.Get("validate"))
		fmt.Printf("  All tags: %s\n", field.Tag)
		fmt.Println()
	}

	// 8. Function Operations
	fmt.Println("\n8. Function Operations:")
	
	calc := &Calculator{}
	calcVal := reflect.ValueOf(calc)
	calcTyp := reflect.TypeOf(calc)
	
	fmt.Printf("Calculator type: %s\n", calcTyp)
	fmt.Printf("Calculator kind: %s\n", calcVal.Kind())
	
	// Get method by name
	addMethod := calcVal.MethodByName("Add")
	if addMethod.IsValid() {
		fmt.Printf("Add method: %s\n", addMethod.Type())
		fmt.Printf("Number of inputs: %d\n", addMethod.Type().NumIn())
		fmt.Printf("Number of outputs: %d\n", addMethod.Type().NumOut())
		
		// Call method
		args := []reflect.Value{reflect.ValueOf(10.0), reflect.ValueOf(20.0)}
		results := addMethod.Call(args)
		fmt.Printf("Add(10, 20) = %v\n", results[0].Interface())
	}

	// 9. Channel Operations
	fmt.Println("\n9. Channel Operations:")
	
	ch := make(chan int, 3)
	chVal := reflect.ValueOf(ch)
	chTyp := reflect.TypeOf(ch)
	
	fmt.Printf("Channel type: %s\n", chTyp)
	fmt.Printf("Channel kind: %s\n", chVal.Kind())
	fmt.Printf("Channel direction: %s\n", chTyp.ChanDir())
	fmt.Printf("Channel element type: %s\n", chTyp.Elem())
	
	// Send to channel
	chVal.Send(reflect.ValueOf(42))
	fmt.Printf("Sent 42 to channel\n")
	
	// Receive from channel
	received, ok := chVal.TryRecv()
	if ok {
		fmt.Printf("Received: %v\n", received.Interface())
	}

	// 10. Type Creation
	fmt.Println("\n10. Type Creation:")
	
	// Create new value of type
	newVal := reflect.New(reflect.TypeOf(0))
	newVal.Elem().SetInt(999)
	fmt.Printf("New int value: %v\n", newVal.Elem().Interface())
	
	// Create slice
	sliceType := reflect.SliceOf(reflect.TypeOf(""))
	newSlice := reflect.MakeSlice(sliceType, 3, 3)
	for i := 0; i < 3; i++ {
		newSlice.Index(i).SetString(fmt.Sprintf("item%d", i))
	}
	fmt.Printf("New slice: %v\n", newSlice.Interface())
	
	// Create map
	mapType := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	newMap := reflect.MakeMap(mapType)
	newMap.SetMapIndex(reflect.ValueOf("key1"), reflect.ValueOf(100))
	newMap.SetMapIndex(reflect.ValueOf("key2"), reflect.ValueOf(200))
	fmt.Printf("New map: %v\n", newMap.Interface())

	// 11. Type Conversion
	fmt.Println("\n11. Type Conversion:")
	
	// Convert between types
	intVal := reflect.ValueOf(42)
	floatVal := intVal.Convert(reflect.TypeOf(0.0))
	fmt.Printf("Int 42 converted to float: %v\n", floatVal.Interface())
	
	// Check if conversion is possible
	if intVal.Type().ConvertibleTo(reflect.TypeOf("")) {
		strVal := intVal.Convert(reflect.TypeOf(""))
		fmt.Printf("Int 42 converted to string: %v\n", strVal.Interface())
	}

	// 12. Interface Operations
	fmt.Println("\n12. Interface Operations:")
	
	var iface interface{} = "Hello, World!"
	ifaceVal := reflect.ValueOf(iface)
	ifaceTyp := reflect.TypeOf(iface)
	
	fmt.Printf("Interface value: %v\n", ifaceVal.Interface())
	fmt.Printf("Interface type: %s\n", ifaceTyp)
	fmt.Printf("Interface kind: %s\n", ifaceVal.Kind())
	
	// Type assertion using reflection
	if ifaceVal.Kind() == reflect.String {
		fmt.Printf("It's a string: %s\n", ifaceVal.String())
	}

	// 13. Deep Inspection
	fmt.Println("\n13. Deep Inspection:")
	
	// Inspect nested struct
	inspectValue(personVal, 0)

	// 14. Dynamic Function Call
	fmt.Println("\n14. Dynamic Function Call:")
	
	// Get function by name and call it
	multiplyMethod := calcVal.MethodByName("Multiply")
	if multiplyMethod.IsValid() {
		args := []reflect.Value{reflect.ValueOf(3.0), reflect.ValueOf(4.0)}
		results := multiplyMethod.Call(args)
		fmt.Printf("Multiply(3, 4) = %v\n", results[0].Interface())
	}

	// 15. Type Assertion Helpers
	fmt.Println("\n15. Type Assertion Helpers:")
	
	// Safe type assertion
	value, ok := safeTypeAssert[int](iface)
	if ok {
		fmt.Printf("Type assertion successful: %v\n", value)
	} else {
		fmt.Printf("Type assertion failed\n")
	}

	// 16. JSON-like Serialization
	fmt.Println("\n16. JSON-like Serialization:")
	
	// Convert struct to map
	personMap := structToMap(person)
	fmt.Printf("Person as map: %+v\n", personMap)
	
	// Convert map back to struct
	newPerson := mapToStruct[Person](personMap)
	fmt.Printf("Map as person: %+v\n", newPerson)

	// 17. Validation using Tags
	fmt.Println("\n17. Validation using Tags:")
	
	// Validate struct fields
	errors := validateStruct(person)
	if len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Printf("  %s\n", err)
		}
	} else {
		fmt.Println("Validation passed!")
	}

	// 18. Performance Comparison
	fmt.Println("\n18. Performance Comparison:")
	
	// Direct access vs reflection
	directAccess := person.Name
	reflectionAccess := personVal.FieldByName("Name").Interface().(string)
	
	fmt.Printf("Direct access: %s\n", directAccess)
	fmt.Printf("Reflection access: %s\n", reflectionAccess)

	fmt.Println("\n🎉 reflect Package Mastery Complete!")
}

// Helper function to inspect values recursively
func inspectValue(val reflect.Value, depth int) {
	indent := strings.Repeat("  ", depth)
	
	if !val.IsValid() {
		fmt.Printf("%s<invalid>\n", indent)
		return
	}
	
	typ := val.Type()
	kind := val.Kind()
	
	fmt.Printf("%sType: %s, Kind: %s\n", indent, typ, kind)
	
	switch kind {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := typ.Field(i)
			fmt.Printf("%sField %s:\n", indent, fieldType.Name)
			inspectValue(field, depth+1)
		}
	case reflect.Slice, reflect.Array:
		fmt.Printf("%sLength: %d\n", indent, val.Len())
		if val.Len() > 0 {
			fmt.Printf("%sFirst element:\n", indent)
			inspectValue(val.Index(0), depth+1)
		}
	case reflect.Map:
		keys := val.MapKeys()
		fmt.Printf("%sKeys: %d\n", indent, len(keys))
		if len(keys) > 0 {
			fmt.Printf("%sFirst key-value:\n", indent)
			inspectValue(keys[0], depth+1)
			inspectValue(val.MapIndex(keys[0]), depth+1)
		}
	case reflect.Ptr:
		if val.IsNil() {
			fmt.Printf("%s<nil>\n", indent)
		} else {
			fmt.Printf("%sPoints to:\n", indent)
			inspectValue(val.Elem(), depth+1)
		}
	default:
		fmt.Printf("%sValue: %v\n", indent, val.Interface())
	}
}

// Safe type assertion helper
func safeTypeAssert[T any](value interface{}) (T, bool) {
	if v, ok := value.(T); ok {
		return v, true
	}
	var zero T
	return zero, false
}

// Convert struct to map
func structToMap(v interface{}) map[string]interface{} {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	result := make(map[string]interface{})
	
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		if field.CanInterface() {
			result[fieldType.Name] = field.Interface()
		}
	}
	
	return result
}

// Convert map to struct
func mapToStruct[T any](m map[string]interface{}) T {
	var result T
	val := reflect.ValueOf(&result).Elem()
	typ := reflect.TypeOf(result)
	
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		if value, exists := m[fieldType.Name]; exists && field.CanSet() {
			fieldVal := reflect.ValueOf(value)
			if fieldVal.Type().ConvertibleTo(field.Type()) {
				field.Set(fieldVal.Convert(field.Type()))
			}
		}
	}
	
	return result
}

// Validate struct using tags
func validateStruct(v interface{}) []string {
	var errors []string
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		validateTag := fieldType.Tag.Get("validate")
		if validateTag == "" {
			continue
		}
		
		fieldName := fieldType.Name
		fieldValue := field.Interface()
		
		// Simple validation rules
		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			rule = strings.TrimSpace(rule)
			
			switch {
			case rule == "required":
				if isEmpty(fieldValue) {
					errors = append(errors, fmt.Sprintf("%s is required", fieldName))
				}
			case strings.HasPrefix(rule, "min="):
				minStr := strings.TrimPrefix(rule, "min=")
				if min, err := strconv.Atoi(minStr); err == nil {
					if intVal, ok := fieldValue.(int); ok && intVal < min {
						errors = append(errors, fmt.Sprintf("%s must be at least %d", fieldName, min))
					}
				}
			case strings.HasPrefix(rule, "max="):
				maxStr := strings.TrimPrefix(rule, "max=")
				if max, err := strconv.Atoi(maxStr); err == nil {
					if intVal, ok := fieldValue.(int); ok && intVal > max {
						errors = append(errors, fmt.Sprintf("%s must be at most %d", fieldName, max))
					}
				}
			case rule == "email":
				if strVal, ok := fieldValue.(string); ok && !strings.Contains(strVal, "@") {
					errors = append(errors, fmt.Sprintf("%s must be a valid email", fieldName))
				}
			}
		}
	}
	
	return errors
}

// Check if value is empty
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String:
		return val.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.Slice, reflect.Array, reflect.Map:
		return val.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	default:
		return false
	}
}
