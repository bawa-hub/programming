# json Package - JSON Encoding/Decoding 📄

The `json` package provides functionality for encoding and decoding JSON data. It's essential for API communication, configuration files, and data persistence.

## 🎯 Key Concepts

### 1. **Basic Operations**
- `Marshal()` - Convert Go value to JSON
- `Unmarshal()` - Convert JSON to Go value
- `MarshalIndent()` - Pretty-print JSON
- `Valid()` - Validate JSON syntax
- `Compact()` - Remove whitespace from JSON
- `Indent()` - Add indentation to JSON

### 2. **Struct Tags**
- `json:"fieldname"` - Custom field name
- `json:"-"` - Skip field
- `json:",omitempty"` - Skip if empty
- `json:",string"` - Encode as string
- `json:",inline"` - Inline embedded struct
- `json:"fieldname,omitempty"` - Combined tags

### 3. **Custom Types**
- `Marshaler` - Custom marshaling
- `Unmarshaler` - Custom unmarshaling
- `TextMarshaler` - Text-based marshaling
- `TextUnmarshaler` - Text-based unmarshaling
- `RawMessage` - Raw JSON data

### 4. **Streaming**
- `Encoder` - Stream JSON encoding
- `Decoder` - Stream JSON decoding
- `NewEncoder()` - Create encoder
- `NewDecoder()` - Create decoder
- `Encode()` - Encode to stream
- `Decode()` - Decode from stream

### 5. **Error Handling**
- `SyntaxError` - JSON syntax errors
- `UnmarshalTypeError` - Type mismatch errors
- `InvalidUnmarshalError` - Invalid unmarshal target
- `MarshalerError` - Custom marshaler errors

### 6. **Advanced Features**
- `Number` - JSON number type
- `RawMessage` - Raw JSON message
- `Token` - JSON token types
- `Delim` - JSON delimiters
- `Token` - JSON token interface

## 🚀 Common Patterns

### Basic JSON Operations
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Marshal
person := Person{Name: "John", Age: 30}
jsonData, err := json.Marshal(person)

// Unmarshal
var person2 Person
err := json.Unmarshal(jsonData, &person2)
```

### Custom Marshaling
```go
type CustomTime struct {
    time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
    return json.Marshal(ct.Time.Format("2006-01-02"))
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }
    t, err := time.Parse("2006-01-02", s)
    if err != nil {
        return err
    }
    ct.Time = t
    return nil
}
```

### Streaming JSON
```go
// Encoding
var buf bytes.Buffer
encoder := json.NewEncoder(&buf)
err := encoder.Encode(data)

// Decoding
decoder := json.NewDecoder(&buf)
var data MyStruct
err := decoder.Decode(&data)
```

### JSON with Tags
```go
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email,omitempty"`
    Password string `json:"-"`
    Admin    bool   `json:"admin,string"`
}
```

## ⚠️ Common Pitfalls

1. **Pointer vs Value** - Unmarshal needs pointer
2. **Case sensitivity** - JSON field names are case sensitive
3. **Type conversion** - Automatic type conversion limitations
4. **Numeric precision** - JSON numbers are float64
5. **Empty values** - nil vs empty string vs omitempty

## 🎯 Best Practices

1. **Use struct tags** - Control JSON field names
2. **Handle errors** - Always check for errors
3. **Use pointers** - For optional fields
4. **Custom types** - For special formatting
5. **Streaming** - For large data sets

## 🔍 Advanced Features

### Custom JSON Types
```go
type JSONNumber json.Number

func (jn JSONNumber) Int64() (int64, error) {
    return strconv.ParseInt(string(jn), 10, 64)
}
```

### JSON Validation
```go
func ValidateJSON(data []byte) error {
    if !json.Valid(data) {
        return fmt.Errorf("invalid JSON")
    }
    return nil
}
```

### JSON Transformation
```go
func TransformJSON(input []byte) ([]byte, error) {
    var data map[string]interface{}
    if err := json.Unmarshal(input, &data); err != nil {
        return nil, err
    }
    
    // Transform data
    data["transformed"] = true
    
    return json.MarshalIndent(data, "", "  ")
}
```

## 📚 Real-world Applications

1. **API Communication** - REST API requests/responses
2. **Configuration Files** - App configuration
3. **Data Persistence** - Save/load data
4. **Message Queues** - Serialize messages
5. **Web Services** - JSON over HTTP

## 🧠 Memory Tips

- **json** = **J**SON **S**erialization **O**perations **N**etwork
- **Marshal** = **M**arshal to JSON
- **Unmarshal** = **U**nmarshal from JSON
- **Encoder** = **E**ncoder for streaming
- **Decoder** = **D**ecoder for streaming
- **Tags** = **T**ags for control
- **Raw** = **R**aw JSON data

Remember: The json package is your gateway to data serialization in Go! 🎯
