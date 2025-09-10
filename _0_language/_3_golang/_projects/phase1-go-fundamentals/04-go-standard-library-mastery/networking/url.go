package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("🚀 Go url Package Mastery Examples")
	fmt.Println("===================================")

	// 1. Basic URL Parsing
	fmt.Println("\n1. Basic URL Parsing:")
	
	// Parse various URL formats
	urls := []string{
		"https://example.com/path?key=value&name=John",
		"http://user:pass@example.com:8080/path?query=value#fragment",
		"ftp://files.example.com/pub/file.txt",
		"mailto:user@example.com",
		"tel:+1234567890",
		"file:///path/to/file.txt",
	}
	
	for _, urlStr := range urls {
		u, err := url.Parse(urlStr)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", urlStr, err)
			continue
		}
		
		fmt.Printf("URL: %s\n", urlStr)
		fmt.Printf("  Scheme: %s\n", u.Scheme)
		fmt.Printf("  Host: %s\n", u.Host)
		fmt.Printf("  Path: %s\n", u.Path)
		fmt.Printf("  RawQuery: %s\n", u.RawQuery)
		fmt.Printf("  Fragment: %s\n", u.Fragment)
		fmt.Printf("  User: %s\n", u.User)
		fmt.Printf("  IsAbs: %t\n", u.IsAbs())
		fmt.Println()
	}

	// 2. Query Parameter Handling
	fmt.Println("\n2. Query Parameter Handling:")
	
	// Parse query string
	queryString := "name=John&age=30&hobby=reading&hobby=coding&city=New+York"
	values, err := url.ParseQuery(queryString)
	if err != nil {
		log.Printf("Error parsing query: %v", err)
	} else {
		fmt.Printf("Query string: %s\n", queryString)
		fmt.Printf("Parsed values:\n")
		for key, vals := range values {
			fmt.Printf("  %s: %v\n", key, vals)
		}
		
		// Get specific values
		fmt.Printf("Name: %s\n", values.Get("name"))
		fmt.Printf("Age: %s\n", values.Get("age"))
		fmt.Printf("Hobbies: %v\n", values["hobby"])
		fmt.Printf("City: %s\n", values.Get("city"))
	}

	// 3. Building Query Parameters
	fmt.Println("\n3. Building Query Parameters:")
	
	// Create query parameters
	params := url.Values{}
	params.Set("name", "Alice")
	params.Set("age", "25")
	params.Add("skill", "Go")
	params.Add("skill", "Python")
	params.Add("skill", "JavaScript")
	params.Set("location", "San Francisco")
	
	fmt.Printf("Built query string: %s\n", params.Encode())
	
	// Modify parameters
	params.Set("age", "26")
	params.Del("location")
	params.Add("experience", "5 years")
	
	fmt.Printf("Modified query string: %s\n", params.Encode())

	// 4. URL Building and Manipulation
	fmt.Println("\n4. URL Building and Manipulation:")
	
	// Build URL from components
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "api.example.com",
		Path:   "/v1/users",
	}
	
	fmt.Printf("Base URL: %s\n", baseURL.String())
	
	// Add query parameters
	baseURL.RawQuery = "page=1&limit=10"
	fmt.Printf("With query: %s\n", baseURL.String())
	
	// Add fragment
	baseURL.Fragment = "section1"
	fmt.Printf("With fragment: %s\n", baseURL.String())
	
	// Add user info
	baseURL.User = url.UserPassword("admin", "secret")
	fmt.Printf("With auth: %s\n", baseURL.String())

	// 5. Relative URL Resolution
	fmt.Println("\n5. Relative URL Resolution:")
	
	// Base URL
	base, err := url.Parse("https://example.com/api/v1/")
	if err != nil {
		log.Printf("Error parsing base URL: %v", err)
	} else {
		// Relative URLs
		relativeURLs := []string{
			"users",
			"../v2/users",
			"../../admin",
			"/absolute/path",
			"users/123",
			"users/123?active=true",
		}
		
		for _, rel := range relativeURLs {
			relURL, err := url.Parse(rel)
			if err != nil {
				fmt.Printf("Error parsing relative URL %s: %v\n", rel, err)
				continue
			}
			
			resolved := base.ResolveReference(relURL)
			fmt.Printf("Relative: %s -> %s\n", rel, resolved.String())
		}
	}

	// 6. URL Encoding and Decoding
	fmt.Println("\n6. URL Encoding and Decoding:")
	
	// Test strings for encoding
	testStrings := []string{
		"Hello World",
		"Hello, World!",
		"Hello & World",
		"Hello = World",
		"Hello + World",
		"Hello/World",
		"Hello?World",
		"Hello#World",
		"Hello World with spaces",
		"Special chars: !@#$%^&*()",
	}
	
	for _, str := range testStrings {
		encoded := url.QueryEscape(str)
		decoded, err := url.QueryUnescape(encoded)
		if err != nil {
			fmt.Printf("Error decoding %s: %v\n", encoded, err)
		} else {
			fmt.Printf("Original: %s\n", str)
			fmt.Printf("Encoded:  %s\n", encoded)
			fmt.Printf("Decoded:  %s\n", decoded)
			fmt.Printf("Match: %t\n", str == decoded)
			fmt.Println()
		}
	}

	// 7. Path Encoding and Decoding
	fmt.Println("\n7. Path Encoding and Decoding:")
	
	// Test path encoding
	paths := []string{
		"/simple/path",
		"/path with spaces",
		"/path/with/special/chars",
		"/path/with/unicode/测试",
		"/path/with/query?param=value",
		"/path/with/fragment#section",
	}
	
	for _, path := range paths {
		encoded := url.PathEscape(path)
		decoded, err := url.PathUnescape(encoded)
		if err != nil {
			fmt.Printf("Error decoding path %s: %v\n", encoded, err)
		} else {
			fmt.Printf("Original: %s\n", path)
			fmt.Printf("Encoded:  %s\n", encoded)
			fmt.Printf("Decoded:  %s\n", decoded)
			fmt.Printf("Match: %t\n", path == decoded)
			fmt.Println()
		}
	}

	// 8. URL Validation
	fmt.Println("\n8. URL Validation:")
	
	// Test URL validation
	testURLs := []string{
		"https://example.com",
		"http://example.com:8080",
		"ftp://files.example.com",
		"mailto:user@example.com",
		"tel:+1234567890",
		"file:///path/to/file",
		"invalid-url",
		"://example.com",
		"https://",
		"",
		"not-a-url",
	}
	
	for _, urlStr := range testURLs {
		u, err := url.Parse(urlStr)
		isValid := err == nil && u.Scheme != "" && u.Host != ""
		
		fmt.Printf("URL: %s\n", urlStr)
		fmt.Printf("  Valid: %t\n", isValid)
		if !isValid && err != nil {
			fmt.Printf("  Error: %v\n", err)
		}
		fmt.Println()
	}

	// 9. URL Component Extraction
	fmt.Println("\n9. URL Component Extraction:")
	
	// Complex URL for component extraction
	complexURL := "https://user:pass@api.example.com:8080/v1/users/123?page=1&limit=10&sort=name#profile"
	
	u, err := url.Parse(complexURL)
	if err != nil {
		log.Printf("Error parsing complex URL: %v", err)
	} else {
		fmt.Printf("Complex URL: %s\n", complexURL)
		fmt.Printf("  Scheme: %s\n", u.Scheme)
		fmt.Printf("  Host: %s\n", u.Host)
		fmt.Printf("  Hostname: %s\n", u.Hostname())
		fmt.Printf("  Port: %s\n", u.Port())
		fmt.Printf("  Path: %s\n", u.Path)
		fmt.Printf("  RawQuery: %s\n", u.RawQuery)
		fmt.Printf("  Fragment: %s\n", u.Fragment)
		fmt.Printf("  User: %s\n", u.User)
		if u.User != nil {
			fmt.Printf("  Username: %s\n", u.User.Username())
			password, _ := u.User.Password()
			fmt.Printf("  Password: %s\n", password)
		}
		fmt.Printf("  IsAbs: %t\n", u.IsAbs())
		fmt.Printf("  RequestURI: %s\n", u.RequestURI())
	}

	// 10. URL Template Building
	fmt.Println("\n10. URL Template Building:")
	
	// Build URL templates
	baseTemplate := "https://api.example.com/v1"
	endpoints := []string{
		"users",
		"users/{id}",
		"users/{id}/posts",
		"posts",
		"posts/{id}",
		"posts/{id}/comments",
		"search",
	}
	
	for _, endpoint := range endpoints {
		fullURL := baseTemplate + "/" + endpoint
		fmt.Printf("Template: %s\n", fullURL)
		
		// Replace template variables
		if strings.Contains(endpoint, "{id}") {
			replaced := strings.ReplaceAll(endpoint, "{id}", "123")
			fullURL = baseTemplate + "/" + replaced
			fmt.Printf("  With ID: %s\n", fullURL)
		}
		fmt.Println()
	}

	// 11. URL Comparison
	fmt.Println("\n11. URL Comparison:")
	
	// Compare URLs
	url1 := "https://example.com/path?param=value"
	url2 := "https://example.com/path?param=value"
	url3 := "https://example.com/path?param=other"
	url4 := "http://example.com/path?param=value"
	
	u1, _ := url.Parse(url1)
	u2, _ := url.Parse(url2)
	u3, _ := url.Parse(url3)
	u4, _ := url.Parse(url4)
	
	fmt.Printf("URL1: %s\n", url1)
	fmt.Printf("URL2: %s\n", url2)
	fmt.Printf("URL3: %s\n", url3)
	fmt.Printf("URL4: %s\n", url4)
	
	fmt.Printf("URL1 == URL2: %t\n", u1.String() == u2.String())
	fmt.Printf("URL1 == URL3: %t\n", u1.String() == u3.String())
	fmt.Printf("URL1 == URL4: %t\n", u1.String() == u4.String())

	// 12. URL Normalization
	fmt.Println("\n12. URL Normalization:")
	
	// Normalize URLs
	urlsToNormalize := []string{
		"https://EXAMPLE.COM/path",
		"https://example.com:80/path",
		"https://example.com/path/",
		"https://example.com/path/../other",
		"https://example.com/path/./file",
	}
	
	for _, urlStr := range urlsToNormalize {
		u, err := url.Parse(urlStr)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", urlStr, err)
			continue
		}
		
		// Normalize by reconstructing
		normalized := &url.URL{
			Scheme: strings.ToLower(u.Scheme),
			Host:   strings.ToLower(u.Host),
			Path:   u.Path,
		}
		
		fmt.Printf("Original: %s\n", urlStr)
		fmt.Printf("Normalized: %s\n", normalized.String())
		fmt.Println()
	}

	// 13. URL Query Parameter Manipulation
	fmt.Println("\n13. URL Query Parameter Manipulation:")
	
	// Start with a URL
	baseURLStr := "https://api.example.com/search"
	u, err = url.Parse(baseURLStr)
	if err != nil {
		log.Printf("Error parsing base URL: %v", err)
	} else {
		// Add query parameters
		query := u.Query()
		query.Set("q", "golang")
		query.Set("type", "repositories")
		query.Add("sort", "stars")
		query.Add("sort", "updated")
		query.Set("per_page", "10")
		query.Set("page", "1")
		
		u.RawQuery = query.Encode()
		fmt.Printf("URL with query: %s\n", u.String())
		
		// Modify parameters
		query.Set("per_page", "20")
		query.Del("sort")
		query.Set("sort", "forks")
		query.Set("page", "2")
		
		u.RawQuery = query.Encode()
		fmt.Printf("Modified URL: %s\n", u.String())
		
		// Extract specific parameters
		fmt.Printf("Search query: %s\n", query.Get("q"))
		fmt.Printf("Type: %s\n", query.Get("type"))
		fmt.Printf("Sort: %s\n", query.Get("sort"))
		fmt.Printf("Per page: %s\n", query.Get("per_page"))
		fmt.Printf("Page: %s\n", query.Get("page"))
	}

	// 14. URL Fragment Handling
	fmt.Println("\n14. URL Fragment Handling:")
	
	// Test fragment handling
	fragmentURLs := []string{
		"https://example.com/page#section1",
		"https://example.com/page#section2",
		"https://example.com/page#top",
		"https://example.com/page#bottom",
		"https://example.com/page#middle",
	}
	
	for _, urlStr := range fragmentURLs {
		u, err := url.Parse(urlStr)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", urlStr, err)
			continue
		}
		
		fmt.Printf("URL: %s\n", urlStr)
		fmt.Printf("  Fragment: %s\n", u.Fragment)
		fmt.Printf("  Without fragment: %s\n", u.Scheme+"://"+u.Host+u.Path)
		fmt.Println()
	}

	// 15. Advanced URL Operations
	fmt.Println("\n15. Advanced URL Operations:")
	
	// Complex URL operations
	complexURLStr := "https://user:pass@api.example.com:8080/v1/users/123/posts?page=1&limit=10&sort=created_at&order=desc#recent"
	
	u, err = url.Parse(complexURLStr)
	if err != nil {
		log.Printf("Error parsing complex URL: %v", err)
	} else {
		fmt.Printf("Complex URL: %s\n", complexURLStr)
		fmt.Println()
		
		// Extract and modify components
		fmt.Printf("Original components:\n")
		fmt.Printf("  Scheme: %s\n", u.Scheme)
		fmt.Printf("  Host: %s\n", u.Host)
		fmt.Printf("  Path: %s\n", u.Path)
		fmt.Printf("  Query: %s\n", u.RawQuery)
		fmt.Printf("  Fragment: %s\n", u.Fragment)
		fmt.Println()
		
		// Modify URL
		u.Scheme = "http"
		u.Host = "localhost:3000"
		u.Path = "/api/v2/users/456"
		
		// Modify query parameters
		query := u.Query()
		query.Set("page", "2")
		query.Set("limit", "20")
		query.Del("sort")
		query.Del("order")
		query.Set("filter", "active")
		u.RawQuery = query.Encode()
		
		u.Fragment = "all"
		
		fmt.Printf("Modified URL: %s\n", u.String())
		fmt.Println()
		
		// Extract final components
		fmt.Printf("Modified components:\n")
		fmt.Printf("  Scheme: %s\n", u.Scheme)
		fmt.Printf("  Host: %s\n", u.Host)
		fmt.Printf("  Path: %s\n", u.Path)
		fmt.Printf("  Query: %s\n", u.RawQuery)
		fmt.Printf("  Fragment: %s\n", u.Fragment)
	}

	fmt.Println("\n🎉 url Package Mastery Complete!")
}
