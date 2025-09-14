package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"time"
)

// Custom types for demonstration
type Person struct {
	XMLName xml.Name `xml:"person"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email,omitempty"`
	Age     int      `xml:"age"`
	Address Address  `xml:"address"`
}

type Address struct {
	XMLName xml.Name `xml:"address"`
	Street  string   `xml:"street"`
	City    string   `xml:"city"`
	Country string   `xml:"country"`
	ZIP     string   `xml:"zip,attr"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      int      `xml:"id,attr"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
	Year    int      `xml:"year,attr"`
	Price   float64  `xml:"price,attr"`
	Tags    []string `xml:"tags>tag"`
}

type Library struct {
	XMLName xml.Name `xml:"library"`
	Name    string   `xml:"name,attr"`
	Books   []Book   `xml:"books>book"`
}

type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(ct.Time.Format("2006-01-02 15:04:05"), start)
}

func (ct *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type Article struct {
	XMLName xml.Name `xml:"article"`
	ID      int      `xml:"id,attr"`
	Title   string   `xml:"title"`
	Content string   `xml:"content"`
	Author  string   `xml:"author"`
	Date    CustomTime `xml:"date"`
}

type Config struct {
	XMLName xml.Name `xml:"config"`
	App     string   `xml:"app,attr"`
	Version string   `xml:"version,attr"`
	Settings []Setting `xml:"settings>setting"`
	Comment  string   `xml:"comment,omitempty"`
}

type Setting struct {
	XMLName xml.Name `xml:"setting"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
	Type    string   `xml:"type,attr"`
}

type Document struct {
	XMLName xml.Name `xml:"http://example.com/schema document"`
	Title   string   `xml:"title"`
	Content string   `xml:"content"`
	Meta    Meta     `xml:"meta"`
}

type Meta struct {
	XMLName xml.Name `xml:"meta"`
	Author  string   `xml:"author"`
	Created CustomTime `xml:"created"`
	Updated CustomTime `xml:"updated"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Description string `xml:"description"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Description string `xml:"description"`
	PubDate string   `xml:"pubDate"`
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	SOAP    string   `xml:"xmlns:soap,attr"`
	Body    SOAPBody `xml:"soap:Body"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Content interface{} `xml:",innerxml"`
}

type SOAPRequest struct {
	XMLName xml.Name `xml:"GetUserRequest"`
	UserID  int      `xml:"userId"`
}

type SOAPResponse struct {
	XMLName xml.Name `xml:"GetUserResponse"`
	User    Person   `xml:"person"`
}

type XMLWithCDATA struct {
	XMLName xml.Name `xml:"document"`
	Title   string   `xml:"title"`
	Content string   `xml:"content"`
	Code    string   `xml:"code"`
}

type XMLWithPI struct {
	XMLName xml.Name `xml:"document"`
	Title   string   `xml:"title"`
	Content string   `xml:"content"`
}

type XMLWithComment struct {
	XMLName xml.Name `xml:"config"`
	Settings []Setting `xml:"setting"`
	Comment  string   `xml:"comment,omitempty"`
}

type MixedContent struct {
	XMLName xml.Name `xml:"document"`
	Title   string   `xml:"title"`
	Content string   `xml:"content"`
	Para    string   `xml:"para"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Name    string   `xml:"name,attr"`
	CEO     Person   `xml:"person"`
	Employees []Person `xml:"employees>person"`
	Address Address  `xml:"address"`
}

func main() {
	fmt.Println("🚀 Go xml Package Mastery Examples")
	fmt.Println("==================================")

	// 1. Basic XML Operations
	fmt.Println("\n1. Basic XML Operations:")
	
	person := Person{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
			ZIP:     "10001",
		},
	}
	
	// Marshal to XML
	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Marshaled XML: %s\n", string(xmlData))
	
	// Pretty print XML
	prettyXML, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pretty XML:\n%s\n", string(prettyXML))
	
	// Unmarshal from XML
	var person2 Person
	err = xml.Unmarshal(xmlData, &person2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Unmarshaled person: %+v\n", person2)

	// 2. XML with Attributes
	fmt.Println("\n2. XML with Attributes:")
	
	book := Book{
		ID:     1,
		Title:  "Go Programming",
		Author: "John Smith",
		Year:   2023,
		Price:  29.99,
		Tags:   []string{"programming", "golang", "book"},
	}
	
	bookXML, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Book XML:\n%s\n", string(bookXML))

	// 3. XML with Namespaces
	fmt.Println("\n3. XML with Namespaces:")
	
	document := Document{
		Title:   "Sample Document",
		Content: "This is a sample document with namespace.",
		Meta: Meta{
			Author:  "John Doe",
			Created: CustomTime{Time: time.Now()},
			Updated: CustomTime{Time: time.Now()},
		},
	}
	
	documentXML, err := xml.MarshalIndent(document, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Document XML:\n%s\n", string(documentXML))

	// 4. XML with Custom Types
	fmt.Println("\n4. XML with Custom Types:")
	
	article := Article{
		ID:      1,
		Title:   "Go XML Tutorial",
		Content: "This is a comprehensive tutorial on Go XML handling.",
		Author:  "Jane Smith",
		Date:    CustomTime{Time: time.Now()},
	}
	
	articleXML, err := xml.MarshalIndent(article, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Article XML:\n%s\n", string(articleXML))

	// 5. XML Configuration
	fmt.Println("\n5. XML Configuration:")
	
	config := Config{
		App:     "MyApp",
		Version: "1.0.0",
		Settings: []Setting{
			{Key: "debug", Value: "true", Type: "boolean"},
			{Key: "port", Value: "8080", Type: "integer"},
			{Key: "host", Value: "localhost", Type: "string"},
		},
		Comment: "This is a configuration file",
	}
	
	configXML, err := xml.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config XML:\n%s\n", string(configXML))

	// 6. XML with CDATA
	fmt.Println("\n6. XML with CDATA:")
	
	xmlWithCDATA := XMLWithCDATA{
		Title:   "Code Example",
		Content: "This is some content with <special> characters & symbols.",
		Code:    "func main() {\n    fmt.Println(\"Hello, World!\")\n}",
	}
	
	cdataXML, err := xml.MarshalIndent(xmlWithCDATA, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CDATA XML:\n%s\n", string(cdataXML))

	// 7. XML with Processing Instructions
	fmt.Println("\n7. XML with Processing Instructions:")
	
	xmlWithPI := XMLWithPI{
		Title:   "Document with PI",
		Content: "This document has processing instructions.",
	}
	
	piXML, err := xml.MarshalIndent(xmlWithPI, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PI XML:\n%s\n", string(piXML))

	// 8. XML with Comments
	fmt.Println("\n8. XML with Comments:")
	
	xmlWithComment := XMLWithComment{
		Settings: []Setting{
			{Key: "theme", Value: "dark", Type: "string"},
			{Key: "language", Value: "en", Type: "string"},
		},
		Comment: "This is a comment",
	}
	
	commentXML, err := xml.MarshalIndent(xmlWithComment, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Comment XML:\n%s\n", string(commentXML))

	// 9. RSS Feed
	fmt.Println("\n9. RSS Feed:")
	
	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       "My Blog",
			Link:        "https://example.com",
			Description: "A sample blog",
			Items: []Item{
				{
					Title:       "First Post",
					Link:        "https://example.com/post1",
					Description: "This is the first post",
					PubDate:     time.Now().Format(time.RFC1123Z),
				},
				{
					Title:       "Second Post",
					Link:        "https://example.com/post2",
					Description: "This is the second post",
					PubDate:     time.Now().Add(-24 * time.Hour).Format(time.RFC1123Z),
				},
			},
		},
	}
	
	rssXML, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("RSS XML:\n%s\n", string(rssXML))

	// 10. SOAP Request/Response
	fmt.Println("\n10. SOAP Request/Response:")
	
	soapRequest := SOAPEnvelope{
		SOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: SOAPBody{
			Content: SOAPRequest{
				UserID: 123,
			},
		},
	}
	
	soapRequestXML, err := xml.MarshalIndent(soapRequest, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SOAP Request XML:\n%s\n", string(soapRequestXML))
	
	soapResponse := SOAPEnvelope{
		SOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: SOAPBody{
			Content: SOAPResponse{
				User: Person{
					ID:    123,
					Name:  "John Doe",
					Email: "john@example.com",
					Age:   30,
				},
			},
		},
	}
	
	soapResponseXML, err := xml.MarshalIndent(soapResponse, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SOAP Response XML:\n%s\n", string(soapResponseXML))

	// 11. XML with Library
	fmt.Println("\n11. XML with Library:")
	
	library := Library{
		Name: "My Library",
		Books: []Book{
			{
				ID:     1,
				Title:  "Go Programming",
				Author: "John Smith",
				Year:   2023,
				Price:  29.99,
				Tags:   []string{"programming", "golang"},
			},
			{
				ID:     2,
				Title:  "Python Basics",
				Author: "Jane Doe",
				Year:   2022,
				Price:  24.99,
				Tags:   []string{"programming", "python"},
			},
		},
	}
	
	libraryXML, err := xml.MarshalIndent(library, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Library XML:\n%s\n", string(libraryXML))

	// 12. XML Validation
	fmt.Println("\n12. XML Validation:")
	
	validXML := `<person id="1"><name>John</name><age>30</age></person>`
	invalidXML := `<person id="1"><name>John</name><age>30</age></person`
	
	// Simple XML validation by trying to unmarshal
	var validPerson Person
	err = xml.Unmarshal([]byte(validXML), &validPerson)
	fmt.Printf("Valid XML: %t\n", err == nil)
	
	var invalidPerson Person
	err = xml.Unmarshal([]byte(invalidXML), &invalidPerson)
	fmt.Printf("Invalid XML: %t\n", err != nil)

	// 13. XML Escaping
	fmt.Println("\n13. XML Escaping:")
	
	text := "This text has <special> characters & symbols \"quotes\" and 'apostrophes'"
	escaped := html.EscapeString(text)
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Escaped: %s\n", escaped)
	
	unescaped := html.UnescapeString(escaped)
	fmt.Printf("Unescaped: %s\n", unescaped)

	// 14. Streaming XML
	fmt.Println("\n14. Streaming XML:")
	
	// Create encoder
	var encoderBuf bytes.Buffer
	encoder := xml.NewEncoder(&encoderBuf)
	encoder.Indent("", "  ")
	
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
	
	fmt.Printf("Streamed XML:\n%s\n", encoderBuf.String())
	
	// Create decoder
	decoder := xml.NewDecoder(&encoderBuf)
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

	// 15. XML with Mixed Content
	fmt.Println("\n15. XML with Mixed Content:")
	
	mixedContent := MixedContent{
		Title:   "Mixed Content Document",
		Content: "This is some content with <b>bold</b> text and <i>italic</i> text.",
		Para:    "This is a paragraph with <a href=\"#\">links</a> and other elements.",
	}
	
	mixedXML, err := xml.MarshalIndent(mixedContent, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mixed Content XML:\n%s\n", string(mixedXML))

	// 16. XML with Attributes and Elements
	fmt.Println("\n16. XML with Attributes and Elements:")
	
	type Product struct {
		XMLName xml.Name `xml:"product"`
		ID      int      `xml:"id,attr"`
		Name    string   `xml:"name,attr"`
		Price   float64  `xml:"price,attr"`
		Description string `xml:"description"`
		Category string   `xml:"category"`
		InStock bool      `xml:"inStock,attr"`
	}
	
	product := Product{
		ID:          1,
		Name:        "Laptop",
		Price:       999.99,
		Description: "A high-performance laptop",
		Category:    "Electronics",
		InStock:     true,
	}
	
	productXML, err := xml.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product XML:\n%s\n", string(productXML))

	// 17. XML with Nested Structures
	fmt.Println("\n17. XML with Nested Structures:")
	
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			ID:    1,
			Name:  "Jane Smith",
			Age:   45,
			Email: "jane@techcorp.com",
		},
		Employees: []Person{
			{ID: 2, Name: "Alice", Age: 30},
			{ID: 3, Name: "Bob", Age: 35},
			{ID: 4, Name: "Charlie", Age: 28},
		},
		Address: Address{
			Street:  "456 Tech Ave",
			City:    "San Francisco",
			Country: "USA",
			ZIP:     "94105",
		},
	}
	
	companyXML, err := xml.MarshalIndent(company, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Company XML:\n%s\n", string(companyXML))

	// 18. XML Performance Test
	fmt.Println("\n18. XML Performance Test:")
	
	// Test marshaling performance
	largeData := make([]Person, 1000)
	for i := 0; i < 1000; i++ {
		largeData[i] = Person{
			ID:     i,
			Name:   fmt.Sprintf("Person %d", i),
			Email:  fmt.Sprintf("person%d@example.com", i),
			Age:    20 + (i % 50),
		}
	}
	
	start := time.Now()
	largeXML, err := xml.Marshal(largeData)
	if err != nil {
		log.Fatal(err)
	}
	marshalTime := time.Since(start)
	
	fmt.Printf("Marshaled %d persons in %v\n", len(largeData), marshalTime)
	fmt.Printf("XML size: %d bytes\n", len(largeXML))
	
	// Test unmarshaling performance
	start = time.Now()
	var unmarshaledData []Person
	err = xml.Unmarshal(largeXML, &unmarshaledData)
	if err != nil {
		log.Fatal(err)
	}
	unmarshalTime := time.Since(start)
	
	fmt.Printf("Unmarshaled %d persons in %v\n", len(unmarshaledData), unmarshalTime)

	fmt.Println("\n🎉 xml Package Mastery Complete!")
}
