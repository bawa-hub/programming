package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

// Custom types for demonstration
type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
	Admin    bool   `json:"admin,string"`
	Age      int    `json:"age"`
	Created  time.Time `json:"created_at"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	ZIP     string `json:"zip_code"`
}

type User struct {
	Person
	Address Address `json:"address"`
	Tags    []string `json:"tags,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time.Format("2006-01-02 15:04:05"))
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type CustomNumber struct {
	Value float64
}

func (cn CustomNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%.2f", cn.Value))
}

func (cn *CustomNumber) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	cn.Value = val
	return nil
}

type JSONNumber json.Number

func (jn JSONNumber) Int64() (int64, error) {
	return strconv.ParseInt(string(jn), 10, 64)
}

func (jn JSONNumber) Float64() (float64, error) {
	return strconv.ParseFloat(string(jn), 64)
}

type CustomJSONNumber JSONNumber

func (cjn CustomJSONNumber) String() string {
	return string(cjn)
}

type CustomData struct {
	ID     int              `json:"id"`
	Value  CustomJSONNumber `json:"value"`
	Active bool             `json:"active"`
}

type Product struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Price       CustomNumber `json:"price"`
	Available   bool        `json:"available"`
	Tags        []string    `json:"tags,omitempty"`
	CreatedAt   CustomTime  `json:"created_at"`
	UpdatedAt   *CustomTime `json:"updated_at,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResponse struct {
	Valid   bool              `json:"valid"`
	Errors  []ValidationError `json:"errors,omitempty"`
	Message string            `json:"message"`
}

func main() {
	fmt.Println("🚀 Go json Package Mastery Examples")
	fmt.Println("===================================")

	// 1. Basic JSON Operations
	fmt.Println("\n1. Basic JSON Operations:")
	
	person := Person{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "secret123",
		Admin:    true,
		Age:      30,
		Created:  time.Now(),
	}
	
	// Marshal to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Marshaled JSON: %s\n", string(jsonData))
	
	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pretty JSON:\n%s\n", string(prettyJSON))
	
	// Unmarshal from JSON
	var person2 Person
	err = json.Unmarshal(jsonData, &person2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Unmarshaled person: %+v\n", person2)

	// 2. JSON with Struct Tags
	fmt.Println("\n2. JSON with Struct Tags:")
	
	user := User{
		Person: Person{
			ID:      1,
			Name:    "Alice Smith",
			Email:   "alice@example.com",
			Admin:   false,
			Age:     25,
			Created: time.Now(),
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
			ZIP:     "10001",
		},
		Tags: []string{"developer", "golang"},
		Meta: map[string]interface{}{
			"last_login": time.Now().Add(-24 * time.Hour),
			"preferences": map[string]string{
				"theme": "dark",
				"lang":  "en",
			},
		},
	}
	
	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User JSON:\n%s\n", string(userJSON))

	// 3. Custom Marshaling
	fmt.Println("\n3. Custom Marshaling:")
	
	product := Product{
		ID:        1,
		Name:      "Go Programming Book",
		Price:     CustomNumber{Value: 29.99},
		Available: true,
		Tags:      []string{"programming", "golang", "book"},
		CreatedAt: CustomTime{Time: time.Now()},
		UpdatedAt: &CustomTime{Time: time.Now().Add(-1 * time.Hour)},
		Metadata: map[string]interface{}{
			"isbn":    "978-0134190440",
			"pages":   380,
			"edition": "1st",
		},
	}
	
	productJSON, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product JSON:\n%s\n", string(productJSON))

	// 4. JSON Validation
	fmt.Println("\n4. JSON Validation:")
	
	validJSON := `{"name": "John", "age": 30}`
	invalidJSON := `{"name": "John", "age": 30,}`
	
	fmt.Printf("Valid JSON: %t\n", json.Valid([]byte(validJSON)))
	fmt.Printf("Invalid JSON: %t\n", json.Valid([]byte(invalidJSON)))
	
	// Compact JSON
	compactJSON := `{
		"name": "John",
		"age": 30
	}`
	
	var buf bytes.Buffer
	err = json.Compact(&buf, []byte(compactJSON))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Compacted JSON: %s\n", buf.String())

	// 5. JSON Indentation
	fmt.Println("\n5. JSON Indentation:")
	
	indentJSON := `{"name":"John","age":30,"city":"New York"}`
	
	var indentBuf bytes.Buffer
	err = json.Indent(&indentBuf, []byte(indentJSON), "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Indented JSON:\n%s\n", indentBuf.String())

	// 6. Streaming JSON
	fmt.Println("\n6. Streaming JSON:")
	
	// Create encoder
	var encoderBuf bytes.Buffer
	encoder := json.NewEncoder(&encoderBuf)
	encoder.SetIndent("", "  ")
	
	// Encode multiple objects
	objects := []interface{}{
		Person{ID: 1, Name: "Alice", Age: 25},
		Person{ID: 2, Name: "Bob", Age: 30},
		Person{ID: 3, Name: "Charlie", Age: 35},
	}
	
	for _, obj := range objects {
		err := encoder.Encode(obj)
		if err != nil {
			log.Fatal(err)
		}
	}
	
	fmt.Printf("Streamed JSON:\n%s\n", encoderBuf.String())
	
	// Create decoder
	decoder := json.NewDecoder(&encoderBuf)
	var decodedPersons []Person
	
	for {
		var person Person
		err := decoder.Decode(&person)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		decodedPersons = append(decodedPersons, person)
	}
	
	fmt.Printf("Decoded persons: %+v\n", decodedPersons)

	// 7. JSON with RawMessage
	fmt.Println("\n7. JSON with RawMessage:")
	
	rawJSON := `{
		"id": 1,
		"name": "John",
		"data": {
			"type": "user",
			"permissions": ["read", "write"],
			"settings": {
				"theme": "dark",
				"notifications": true
			}
		}
	}`
	
	type UserWithRawData struct {
		ID   int             `json:"id"`
		Name string          `json:"name"`
		Data json.RawMessage `json:"data"`
	}
	
	var userWithRaw UserWithRawData
	err = json.Unmarshal([]byte(rawJSON), &userWithRaw)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("User ID: %d, Name: %s\n", userWithRaw.ID, userWithRaw.Name)
	fmt.Printf("Raw Data: %s\n", string(userWithRaw.Data))
	
	// Parse raw data
	var rawData map[string]interface{}
	err = json.Unmarshal(userWithRaw.Data, &rawData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parsed raw data: %+v\n", rawData)

	// 8. JSON Number Type
	fmt.Println("\n8. JSON Number Type:")
	
	numberJSON := `{
		"int_value": 42,
		"float_value": 3.14159,
		"string_number": "123"
	}`
	
	var numberData struct {
		IntValue     json.Number `json:"int_value"`
		FloatValue   json.Number `json:"float_value"`
		StringNumber json.Number `json:"string_number"`
	}
	
	err = json.Unmarshal([]byte(numberJSON), &numberData)
	if err != nil {
		log.Fatal(err)
	}
	
	intVal, _ := numberData.IntValue.Int64()
	floatVal, _ := numberData.FloatValue.Float64()
	stringVal, _ := numberData.StringNumber.Int64()
	
	fmt.Printf("Int value: %d\n", intVal)
	fmt.Printf("Float value: %.5f\n", floatVal)
	fmt.Printf("String number: %d\n", stringVal)

	// 9. JSON Error Handling
	fmt.Println("\n9. JSON Error Handling:")
	
	// Test different error types
	testCases := []struct {
		name string
		json string
	}{
		{"Valid JSON", `{"name": "John", "age": 30}`},
		{"Invalid JSON", `{"name": "John", "age": 30,}`},
		{"Type Mismatch", `{"name": "John", "age": "thirty"}`},
		{"Invalid Target", `{"name": "John"}`},
	}
	
	for _, tc := range testCases {
		fmt.Printf("\nTesting %s:\n", tc.name)
		
		var person Person
		err := json.Unmarshal([]byte(tc.json), &person)
		if err != nil {
			switch e := err.(type) {
			case *json.SyntaxError:
				fmt.Printf("  Syntax Error: %s at offset %d\n", e.Error(), e.Offset)
			case *json.UnmarshalTypeError:
				fmt.Printf("  Type Error: %s for field %s at offset %d\n", e.Error(), e.Field, e.Offset)
			case *json.InvalidUnmarshalError:
				fmt.Printf("  Invalid Unmarshal Error: %s\n", e.Error())
			default:
				fmt.Printf("  Other Error: %s\n", err.Error())
			}
		} else {
			fmt.Printf("  Success: %+v\n", person)
		}
	}

	// 10. JSON Transformation
	fmt.Println("\n10. JSON Transformation:")
	
	transformJSON := `{
		"user_id": 123,
		"user_name": "john_doe",
		"user_email": "john@example.com",
		"user_age": 30
	}`
	
	// Transform snake_case to camelCase
	var rawData2 map[string]interface{}
	err = json.Unmarshal([]byte(transformJSON), &rawData2)
	if err != nil {
		log.Fatal(err)
	}
	
	// Transform keys
	transformed := make(map[string]interface{})
	for key, value := range rawData2 {
		camelKey := strings.ReplaceAll(key, "_", "")
		camelKey = strings.ToUpper(camelKey[:1]) + camelKey[1:]
		transformed[camelKey] = value
	}
	
	transformedJSON, err := json.MarshalIndent(transformed, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transformed JSON:\n%s\n", string(transformedJSON))

	// 11. JSON API Response
	fmt.Println("\n11. JSON API Response:")
	
	// Success response
	successResponse := APIResponse{
		Success: true,
		Data:    person,
		Code:    200,
	}
	
	successJSON, err := json.MarshalIndent(successResponse, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Success Response:\n%s\n", string(successJSON))
	
	// Error response
	errorResponse := APIResponse{
		Success: false,
		Error:   "User not found",
		Code:    404,
	}
	
	errorJSON, err := json.MarshalIndent(errorResponse, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Error Response:\n%s\n", string(errorJSON))

	// 12. JSON Validation Response
	fmt.Println("\n12. JSON Validation Response:")
	
	validationResponse := ValidationResponse{
		Valid: false,
		Errors: []ValidationError{
			{Field: "email", Message: "Invalid email format"},
			{Field: "age", Message: "Age must be between 18 and 100"},
		},
		Message: "Validation failed",
	}
	
	validationJSON, err := json.MarshalIndent(validationResponse, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Validation Response:\n%s\n", string(validationJSON))

	// 13. JSON with Nested Structures
	fmt.Println("\n13. JSON with Nested Structures:")
	
	type Company struct {
		Name    string `json:"name"`
		Address Address `json:"address"`
		CEO     Person  `json:"ceo"`
		Employees []Person `json:"employees"`
	}
	
	company := Company{
		Name: "Tech Corp",
		Address: Address{
			Street:  "456 Tech Ave",
			City:    "San Francisco",
			Country: "USA",
			ZIP:     "94105",
		},
		CEO: Person{
			ID:   1,
			Name: "Jane Smith",
			Age:  45,
		},
		Employees: []Person{
			{ID: 2, Name: "Alice", Age: 30},
			{ID: 3, Name: "Bob", Age: 35},
			{ID: 4, Name: "Charlie", Age: 28},
		},
	}
	
	companyJSON, err := json.MarshalIndent(company, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Company JSON:\n%s\n", string(companyJSON))

	// 14. JSON with Anonymous Fields
	fmt.Println("\n14. JSON with Anonymous Fields:")
	
	type Base struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	
	type Post struct {
		Base
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
	}
	
	post := Post{
		Base: Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:   "Go JSON Tutorial",
		Content: "This is a comprehensive tutorial on Go JSON handling.",
		Author:  "John Doe",
	}
	
	postJSON, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Post JSON:\n%s\n", string(postJSON))

	// 15. JSON with Custom Types
	fmt.Println("\n15. JSON with Custom Types:")
	
	customData := CustomData{
		ID:     1,
		Value:  CustomJSONNumber("42.5"),
		Active: true,
	}
	
	customJSON, err := json.MarshalIndent(customData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Custom Data JSON:\n%s\n", string(customJSON))

	// 16. JSON with Interface{}
	fmt.Println("\n16. JSON with Interface{}:")
	
	interfaceData := map[string]interface{}{
		"string":  "hello",
		"number":  42,
		"float":   3.14,
		"boolean": true,
		"array":   []int{1, 2, 3},
		"object": map[string]interface{}{
			"nested": "value",
			"count":  5,
		},
		"null": nil,
	}
	
	interfaceJSON, err := json.MarshalIndent(interfaceData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Interface Data JSON:\n%s\n", string(interfaceJSON))

	// 17. JSON with Omitempty
	fmt.Println("\n17. JSON with Omitempty:")
	
	type OmitemptyExample struct {
		Name     string `json:"name"`
		Email    string `json:"email,omitempty"`
		Age      int    `json:"age,omitempty"`
		Active   bool   `json:"active,omitempty"`
		Tags     []string `json:"tags,omitempty"`
		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}
	
	// With all fields
	fullExample := OmitemptyExample{
		Name:     "John",
		Email:    "john@example.com",
		Age:      30,
		Active:   true,
		Tags:     []string{"user", "admin"},
		Metadata: map[string]interface{}{"role": "admin"},
	}
	
	fullJSON, err := json.MarshalIndent(fullExample, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Full Example JSON:\n%s\n", string(fullJSON))
	
	// With empty fields
	emptyExample := OmitemptyExample{
		Name: "Jane",
	}
	
	emptyJSON, err := json.MarshalIndent(emptyExample, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Empty Example JSON:\n%s\n", string(emptyJSON))

	// 18. JSON Performance Test
	fmt.Println("\n18. JSON Performance Test:")
	
	// Test marshaling performance
	largeData := make([]Person, 1000)
	for i := 0; i < 1000; i++ {
		largeData[i] = Person{
			ID:      i,
			Name:    fmt.Sprintf("Person %d", i),
			Email:   fmt.Sprintf("person%d@example.com", i),
			Age:     20 + (i % 50),
			Created: time.Now(),
		}
	}
	
	start := time.Now()
	largeJSON, err := json.Marshal(largeData)
	if err != nil {
		log.Fatal(err)
	}
	marshalTime := time.Since(start)
	
	fmt.Printf("Marshaled %d persons in %v\n", len(largeData), marshalTime)
	fmt.Printf("JSON size: %d bytes\n", len(largeJSON))
	
	// Test unmarshaling performance
	start = time.Now()
	var unmarshaledData []Person
	err = json.Unmarshal(largeJSON, &unmarshaledData)
	if err != nil {
		log.Fatal(err)
	}
	unmarshalTime := time.Since(start)
	
	fmt.Printf("Unmarshaled %d persons in %v\n", len(unmarshaledData), unmarshalTime)

	fmt.Println("\n🎉 json Package Mastery Complete!")
}
