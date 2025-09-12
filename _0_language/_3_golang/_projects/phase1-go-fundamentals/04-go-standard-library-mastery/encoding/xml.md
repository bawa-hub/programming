# xml Package - XML Encoding/Decoding 📄

The `xml` package provides functionality for encoding and decoding XML data. It's essential for configuration files, data exchange, and legacy system integration.

## 🎯 Key Concepts

### 1. **Basic Operations**
- `Marshal()` - Convert Go value to XML
- `Unmarshal()` - Convert XML to Go value
- `MarshalIndent()` - Pretty-print XML
- `Valid()` - Validate XML syntax
- `Escape()` - Escape XML special characters
- `Unescape()` - Unescape XML special characters

### 2. **Struct Tags**
- `xml:"tagname"` - Custom tag name
- `xml:"-"` - Skip field
- `xml:",omitempty"` - Skip if empty
- `xml:",chardata"` - Character data
- `xml:",innerxml"` - Raw XML content
- `xml:",attr"` - XML attribute
- `xml:"tagname,attr"` - Named attribute

### 3. **Custom Types**
- `Marshaler` - Custom marshaling
- `Unmarshaler` - Custom unmarshaling
- `TextMarshaler` - Text-based marshaling
- `TextUnmarshaler` - Text-based unmarshaling
- `Name` - XML name with namespace
- `StartElement` - XML start element
- `EndElement` - XML end element

### 4. **Streaming**
- `Encoder` - Stream XML encoding
- `Decoder` - Stream XML decoding
- `NewEncoder()` - Create encoder
- `NewDecoder()` - Create decoder
- `Encode()` - Encode to stream
- `Decode()` - Decode from stream

### 5. **Error Handling**
- `SyntaxError` - XML syntax errors
- `UnmarshalError` - Unmarshal errors
- `TagPathError` - Tag path errors
- `AttrError` - Attribute errors

### 6. **Advanced Features**
- `NameSpace` - XML namespace handling
- `CharData` - Character data
- `Comment` - XML comments
- `ProcInst` - Processing instructions
- `Directive` - XML directives

## 🚀 Common Patterns

### Basic XML Operations
```go
type Person struct {
    Name string `xml:"name"`
    Age  int    `xml:"age"`
}

// Marshal
person := Person{Name: "John", Age: 30}
xmlData, err := xml.Marshal(person)

// Unmarshal
var person2 Person
err := xml.Unmarshal(xmlData, &person2)
```

### XML with Attributes
```go
type Book struct {
    ID    int    `xml:"id,attr"`
    Title string `xml:"title"`
    Author string `xml:"author"`
}
```

### XML with Namespaces
```go
type Document struct {
    XMLName xml.Name `xml:"http://example.com/schema document"`
    Title   string   `xml:"title"`
    Content string   `xml:"content"`
}
```

### Custom Marshaling
```go
type CustomTime struct {
    time.Time
}

func (ct CustomTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return e.EncodeElement(ct.Time.Format("2006-01-02"), start)
}
```

## ⚠️ Common Pitfalls

1. **XML Names** - XML names are case sensitive
2. **Namespaces** - Complex namespace handling
3. **Attributes vs Elements** - Choose appropriate representation
4. **Character Data** - Handling text content
5. **Empty Elements** - Self-closing vs empty elements

## 🎯 Best Practices

1. **Use struct tags** - Control XML structure
2. **Handle namespaces** - Use proper namespace declarations
3. **Choose attributes** - For simple values
4. **Use elements** - For complex data
5. **Validate XML** - Check for well-formed XML

## 🔍 Advanced Features

### XML with CDATA
```go
type Article struct {
    Title   string `xml:"title"`
    Content string `xml:"content,chardata"`
}
```

### XML with Processing Instructions
```go
type Document struct {
    XMLName xml.Name `xml:"document"`
    Title   string   `xml:"title"`
    PI      string   `xml:"pi,omitempty"`
}
```

### XML with Comments
```go
type Config struct {
    XMLName xml.Name `xml:"config"`
    Settings []Setting `xml:"setting"`
    Comment  string   `xml:"comment,omitempty"`
}
```

## 📚 Real-world Applications

1. **Configuration Files** - App configuration
2. **Data Exchange** - Legacy system integration
3. **Web Services** - SOAP and XML-RPC
4. **Document Processing** - XML document handling
5. **RSS/Atom Feeds** - Syndication formats

## 🧠 Memory Tips

- **xml** = **X**ML **M**arkup **L**anguage
- **Marshal** = **M**arshal to XML
- **Unmarshal** = **U**nmarshal from XML
- **Encoder** = **E**ncoder for streaming
- **Decoder** = **D**ecoder for streaming
- **Tags** = **T**ags for control
- **Namespace** = **N**amespace handling

Remember: The xml package is your gateway to XML processing in Go! 🎯
