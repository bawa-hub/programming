# url Package - URL Parsing and Manipulation 🔗

The `url` package provides functions for parsing URLs and querying their components. It's essential for web development and URL handling.

## 🎯 Key Concepts

### 1. **URL Structure**
- `URL` - URL representation
- `Scheme` - Protocol (http, https, ftp, etc.)
- `Host` - Hostname and port
- `Path` - URL path
- `RawQuery` - Query string
- `Fragment` - URL fragment
- `User` - User information
- `Opaque` - Opaque data

### 2. **URL Parsing**
- `Parse()` - Parse URL string
- `ParseRequestURI()` - Parse absolute URI
- `ParseQuery()` - Parse query string
- `Query()` - Get query values
- `QueryEscape()` - Escape query string
- `QueryUnescape()` - Unescape query string

### 3. **URL Building**
- `URL.String()` - Convert URL to string
- `ResolveReference()` - Resolve relative URL
- `RequestURI()` - Get request URI
- `Hostname()` - Get hostname
- `Port()` - Get port
- `IsAbs()` - Check if absolute URL

### 4. **Query Parameters**
- `Values` - Query parameter values
- `Set()` - Set parameter value
- `Get()` - Get parameter value
- `Add()` - Add parameter value
- `Del()` - Delete parameter
- `Encode()` - Encode parameters

### 5. **URL Encoding**
- `PathEscape()` - Escape URL path
- `PathUnescape()` - Unescape URL path
- `QueryEscape()` - Escape query string
- `QueryUnescape()` - Unescape query string
- `User()` - Parse user info
- `Password()` - Get password

### 6. **URL Validation**
- `IsAbs()` - Check if absolute
- `Hostname()` - Get hostname
- `Port()` - Get port
- `Scheme` - Get scheme
- `Path` - Get path
- `Fragment` - Get fragment

## 🚀 Common Patterns

### Basic URL Parsing
```go
u, err := url.Parse("https://example.com/path?key=value")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Scheme:", u.Scheme)
fmt.Println("Host:", u.Host)
fmt.Println("Path:", u.Path)
fmt.Println("Query:", u.RawQuery)
```

### Query Parameter Handling
```go
values := url.Values{}
values.Set("name", "John")
values.Set("age", "30")
values.Add("hobby", "reading")
values.Add("hobby", "coding")

queryString := values.Encode()
fmt.Println(queryString) // name=John&age=30&hobby=reading&hobby=coding
```

### URL Building
```go
baseURL, _ := url.Parse("https://api.example.com")
relPath, _ := url.Parse("/users")
fullURL := baseURL.ResolveReference(relPath)
fmt.Println(fullURL.String()) // https://api.example.com/users
```

## ⚠️ Common Pitfalls

1. **Not handling errors** - Always check parsing errors
2. **Not encoding special characters** - Use proper encoding
3. **Not validating URLs** - Validate URLs before use
4. **Not handling relative URLs** - Use ResolveReference
5. **Not escaping query parameters** - Always escape parameters

## 🎯 Best Practices

1. **Handle errors** - Always check parsing errors
2. **Use proper encoding** - Escape special characters
3. **Validate URLs** - Check URL validity
4. **Use ResolveReference** - For relative URLs
5. **Escape parameters** - Always escape query parameters

## 🔍 Advanced Features

### Custom URL Parsing
```go
func parseCustomURL(urlStr string) (*url.URL, error) {
    u, err := url.Parse(urlStr)
    if err != nil {
        return nil, err
    }
    
    if u.Scheme == "" {
        u.Scheme = "https"
    }
    
    return u, nil
}
```

### URL Validation
```go
func isValidURL(urlStr string) bool {
    u, err := url.Parse(urlStr)
    return err == nil && u.Scheme != "" && u.Host != ""
}
```

## 📚 Real-world Applications

1. **Web Scraping** - URL parsing and manipulation
2. **API Development** - Query parameter handling
3. **URL Shortening** - URL validation and processing
4. **Web Crawling** - URL resolution and validation
5. **Form Processing** - Query parameter extraction

## 🧠 Memory Tips

- **url** = **U**RL **L**anguage **R**esource
- **Parse** = **P**arse URL
- **Query** = **Q**uery parameters
- **Encode** = **E**ncode URL
- **Decode** = **D**ecode URL
- **Resolve** = **R**esolve reference
- **Values** = **V**alue collection
- **String** = **S**tring conversion

Remember: The url package is your gateway to URL manipulation in Go! 🎯
