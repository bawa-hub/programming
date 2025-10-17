package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// StringProcessor handles various string operations
type StringProcessor struct {
	Text        string
	Words       []string
	Lines       []string
	Characters  []rune
	Bytes       []byte
	Patterns    map[string]*regexp.Regexp
	Statistics  StringStats
}

// StringStats holds statistical information about strings
type StringStats struct {
	Length        int
	RuneCount     int
	WordCount     int
	LineCount     int
	ByteCount     int
	VowelCount    int
	ConsonantCount int
	DigitCount    int
	SpaceCount    int
	PunctuationCount int
}

// NewStringProcessor creates a new string processor
func NewStringProcessor(text string) *StringProcessor {
	sp := &StringProcessor{
		Text:     text,
		Patterns: make(map[string]*regexp.Regexp),
	}
	
	// Compile common regex patterns
	sp.Patterns["email"] = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	sp.Patterns["phone"] = regexp.MustCompile(`^\+?[\d\s\-\(\)]{10,}$`)
	sp.Patterns["url"] = regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	sp.Patterns["ip"] = regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)
	sp.Patterns["word"] = regexp.MustCompile(`\b\w+\b`)
	sp.Patterns["number"] = regexp.MustCompile(`\d+`)
	
	sp.analyze()
	return sp
}

// CRUD Operations for Strings

// Create - Initialize and build strings
func (sp *StringProcessor) Create() {
	fmt.Println("🔧 Creating and building strings...")
	
	// String creation methods
	var emptyString string
	fmt.Printf("Empty string: '%s' (len: %d)\n", emptyString, len(emptyString))
	
	// String literals
	literal := "Hello, World!"
	fmt.Printf("Literal: '%s'\n", literal)
	
	// Raw strings (backticks)
	rawString := `This is a raw string
	with multiple lines
	and "quotes" and \backslashes\`
	fmt.Printf("Raw string:\n%s\n", rawString)
	
	// String concatenation
	concatenated := "Hello" + ", " + "World" + "!"
	fmt.Printf("Concatenated: '%s'\n", concatenated)
	
	// Using fmt.Sprintf
	formatted := fmt.Sprintf("Number: %d, Float: %.2f, String: %s", 42, 3.14159, "Go")
	fmt.Printf("Formatted: '%s'\n", formatted)
	
	// Using strings.Builder for efficient concatenation
	var builder strings.Builder
	builder.WriteString("Efficient")
	builder.WriteString(" ")
	builder.WriteString("concatenation")
	builder.WriteString("!")
	efficientString := builder.String()
	fmt.Printf("Builder result: '%s'\n", efficientString)
	
	// Using bytes.Buffer
	var buffer bytes.Buffer
	buffer.WriteString("Buffer")
	buffer.WriteString(" ")
	buffer.WriteString("method")
	buffer.WriteString("!")
	bufferString := buffer.String()
	fmt.Printf("Buffer result: '%s'\n", bufferString)
	
	// String from bytes
	bytes := []byte{72, 101, 108, 108, 111} // "Hello"
	fromBytes := string(bytes)
	fmt.Printf("From bytes: '%s'\n", fromBytes)
	
	// String from runes
	runes := []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界'}
	fromRunes := string(runes)
	fmt.Printf("From runes: '%s'\n", fromRunes)
	
	fmt.Println("✅ Strings created successfully")
}

// Read - Display string information and analysis
func (sp *StringProcessor) Read() {
	fmt.Println("\n📖 READING STRING INFORMATION:")
	fmt.Println("==============================")
	
	// Basic string information
	fmt.Printf("Text: '%s'\n", sp.Text)
	fmt.Printf("Length (bytes): %d\n", len(sp.Text))
	fmt.Printf("Length (runes): %d\n", utf8.RuneCountInString(sp.Text))
	
	// Display statistics
	fmt.Printf("\nStatistics:\n")
	fmt.Printf("  Words: %d\n", sp.Statistics.WordCount)
	fmt.Printf("  Lines: %d\n", sp.Statistics.LineCount)
	fmt.Printf("  Vowels: %d\n", sp.Statistics.VowelCount)
	fmt.Printf("  Consonants: %d\n", sp.Statistics.ConsonantCount)
	fmt.Printf("  Digits: %d\n", sp.Statistics.DigitCount)
	fmt.Printf("  Spaces: %d\n", sp.Statistics.SpaceCount)
	fmt.Printf("  Punctuation: %d\n", sp.Statistics.PunctuationCount)
	
	// Display words
	fmt.Printf("\nWords: %v\n", sp.Words)
	
	// Display lines
	fmt.Printf("\nLines:\n")
	for i, line := range sp.Lines {
		fmt.Printf("  %d: '%s'\n", i+1, line)
	}
	
	// Display characters
	fmt.Printf("\nCharacters (first 20): ")
	for i, char := range sp.Characters {
		if i >= 20 {
			fmt.Printf("...")
			break
		}
		fmt.Printf("'%c' ", char)
	}
	fmt.Println()
	
	// Display bytes
	fmt.Printf("\nBytes (first 20): ")
	for i, b := range sp.Bytes {
		if i >= 20 {
			fmt.Printf("...")
			break
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
}

// Update - Modify string content
func (sp *StringProcessor) Update() {
	fmt.Println("\n🔄 UPDATING STRING CONTENT:")
	fmt.Println("===========================")
	
	// String replacement
	original := sp.Text
	sp.Text = strings.ReplaceAll(sp.Text, " ", "_")
	fmt.Printf("Replace spaces with underscores: '%s'\n", sp.Text)
	
	// Restore original
	sp.Text = original
	
	// Case transformations
	sp.Text = strings.ToUpper(sp.Text)
	fmt.Printf("Uppercase: '%s'\n", sp.Text)
	
	sp.Text = strings.ToLower(sp.Text)
	fmt.Printf("Lowercase: '%s'\n", sp.Text)
	
	sp.Text = strings.Title(sp.Text)
	fmt.Printf("Title case: '%s'\n", sp.Text)
	
	// Trim operations
	sp.Text = strings.TrimSpace(sp.Text)
	fmt.Printf("Trimmed: '%s'\n", sp.Text)
	
	sp.Text = strings.Trim(sp.Text, "!.,")
	fmt.Printf("Trimmed punctuation: '%s'\n", sp.Text)
	
	// Padding
	sp.Text = strings.Repeat(" ", 5) + sp.Text + strings.Repeat(" ", 5)
	fmt.Printf("Padded: '%s'\n", sp.Text)
	
	// Restore original for further operations
	sp.Text = original
	
	// String splitting and joining
	words := strings.Fields(sp.Text)
	fmt.Printf("Words: %v\n", words)
	
	joined := strings.Join(words, "-")
	fmt.Printf("Joined with '-': '%s'\n", joined)
	
	// Update internal data
	sp.analyze()
	fmt.Println("✅ String content updated successfully")
}

// Delete - Remove content from strings
func (sp *StringProcessor) Delete() {
	fmt.Println("\n🗑️  DELETING FROM STRINGS:")
	fmt.Println("==========================")
	
	// Remove specific characters
	original := sp.Text
	sp.Text = strings.ReplaceAll(sp.Text, " ", "")
	fmt.Printf("Remove spaces: '%s'\n", sp.Text)
	
	sp.Text = strings.ReplaceAll(sp.Text, "a", "")
	fmt.Printf("Remove 'a': '%s'\n", sp.Text)
	
	sp.Text = strings.ReplaceAll(sp.Text, "e", "")
	fmt.Printf("Remove 'e': '%s'\n", sp.Text)
	
	// Remove using regex
	sp.Text = sp.Patterns["number"].ReplaceAllString(sp.Text, "")
	fmt.Printf("Remove numbers: '%s'\n", sp.Text)
	
	// Remove punctuation
	sp.Text = strings.ReplaceAll(sp.Text, ".", "")
	sp.Text = strings.ReplaceAll(sp.Text, ",", "")
	sp.Text = strings.ReplaceAll(sp.Text, "!", "")
	sp.Text = strings.ReplaceAll(sp.Text, "?", "")
	fmt.Printf("Remove punctuation: '%s'\n", sp.Text)
	
	// Restore original
	sp.Text = original
	
	// Remove by position
	if len(sp.Text) > 5 {
		sp.Text = sp.Text[:5] + sp.Text[6:] // Remove character at index 5
		fmt.Printf("Remove character at index 5: '%s'\n", sp.Text)
	}
	
	// Remove leading/trailing characters
	sp.Text = strings.TrimLeft(sp.Text, " ")
	sp.Text = strings.TrimRight(sp.Text, " ")
	fmt.Printf("Trimmed: '%s'\n", sp.Text)
	
	// Update internal data
	sp.analyze()
	fmt.Println("✅ String content deleted successfully")
}

// analyze performs comprehensive string analysis
func (sp *StringProcessor) analyze() {
	sp.Statistics.Length = len(sp.Text)
	sp.Statistics.RuneCount = utf8.RuneCountInString(sp.Text)
	
	// Split into words and lines
	sp.Words = strings.Fields(sp.Text)
	sp.Lines = strings.Split(sp.Text, "\n")
	sp.Statistics.WordCount = len(sp.Words)
	sp.Statistics.LineCount = len(sp.Lines)
	
	// Convert to runes and bytes
	sp.Characters = []rune(sp.Text)
	sp.Bytes = []byte(sp.Text)
	sp.Statistics.ByteCount = len(sp.Bytes)
	
	// Count character types
	for _, char := range sp.Characters {
		switch {
		case unicode.IsLetter(char):
			if isVowel(char) {
				sp.Statistics.VowelCount++
			} else {
				sp.Statistics.ConsonantCount++
			}
		case unicode.IsDigit(char):
			sp.Statistics.DigitCount++
		case unicode.IsSpace(char):
			sp.Statistics.SpaceCount++
		case unicode.IsPunct(char):
			sp.Statistics.PunctuationCount++
		}
	}
}

// isVowel checks if a character is a vowel
func isVowel(char rune) bool {
	char = unicode.ToLower(char)
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

// Advanced String Operations

// DemonstrateStringSearching shows various string search operations
func (sp *StringProcessor) DemonstrateStringSearching() {
	fmt.Println("\n🔍 STRING SEARCHING OPERATIONS:")
	fmt.Println("===============================")
	
	text := "The quick brown fox jumps over the lazy dog"
	searchText := "fox"
	
	// Contains
	contains := strings.Contains(text, searchText)
	fmt.Printf("Contains '%s': %t\n", searchText, contains)
	
	// ContainsAny
	containsAny := strings.ContainsAny(text, "aeiou")
	fmt.Printf("Contains any vowel: %t\n", containsAny)
	
	// ContainsRune
	containsRune := strings.ContainsRune(text, 'x')
	fmt.Printf("Contains rune 'x': %t\n", containsRune)
	
	// Index operations
	index := strings.Index(text, searchText)
	fmt.Printf("Index of '%s': %d\n", searchText, index)
	
	lastIndex := strings.LastIndex(text, "o")
	fmt.Printf("Last index of 'o': %d\n", lastIndex)
	
	indexAny := strings.IndexAny(text, "aeiou")
	fmt.Printf("Index of any vowel: %d\n", indexAny)
	
	indexRune := strings.IndexRune(text, 'x')
	fmt.Printf("Index of rune 'x': %d\n", indexRune)
	
	// HasPrefix and HasSuffix
	hasPrefix := strings.HasPrefix(text, "The")
	fmt.Printf("Has prefix 'The': %t\n", hasPrefix)
	
	hasSuffix := strings.HasSuffix(text, "dog")
	fmt.Printf("Has suffix 'dog': %t\n", hasSuffix)
	
	// Count
	count := strings.Count(text, "o")
	fmt.Printf("Count of 'o': %d\n", count)
}

// DemonstrateStringTransformation shows string transformation operations
func (sp *StringProcessor) DemonstrateStringTransformation() {
	fmt.Println("\n🔄 STRING TRANSFORMATION OPERATIONS:")
	fmt.Println("====================================")
	
	text := "  hello, world!  "
	fmt.Printf("Original: '%s'\n", text)
	
	// Case transformations
	fmt.Printf("ToUpper: '%s'\n", strings.ToUpper(text))
	fmt.Printf("ToLower: '%s'\n", strings.ToLower(text))
	fmt.Printf("Title: '%s'\n", strings.Title(text))
	
	// Trim operations
	fmt.Printf("TrimSpace: '%s'\n", strings.TrimSpace(text))
	fmt.Printf("TrimLeft: '%s'\n", strings.TrimLeft(text, " "))
	fmt.Printf("TrimRight: '%s'\n", strings.TrimRight(text, " "))
	fmt.Printf("Trim: '%s'\n", strings.Trim(text, " "))
	
	// TrimFunc
	trimmed := strings.TrimFunc(text, func(r rune) bool {
		return r == ' ' || r == '!'
	})
	fmt.Printf("TrimFunc: '%s'\n", trimmed)
	
	// Map transformation
	mapped := strings.Map(func(r rune) rune {
		if r == 'o' {
			return '0'
		}
		return r
	}, text)
	fmt.Printf("Map 'o' to '0': '%s'\n", mapped)
	
	// Replace operations
	replaced := strings.Replace(text, "world", "Go", 1)
	fmt.Printf("Replace 'world' with 'Go': '%s'\n", replaced)
	
	replacedAll := strings.ReplaceAll(text, "l", "L")
	fmt.Printf("Replace all 'l' with 'L': '%s'\n", replacedAll)
}

// DemonstrateStringSplitting shows string splitting operations
func (sp *StringProcessor) DemonstrateStringSplitting() {
	fmt.Println("\n✂️  STRING SPLITTING OPERATIONS:")
	fmt.Println("===============================")
	
	text := "apple,banana,cherry,date,elderberry"
	fmt.Printf("Original: '%s'\n", text)
	
	// Split by comma
	parts := strings.Split(text, ",")
	fmt.Printf("Split by ',': %v\n", parts)
	
	// Split with limit
	partsLimited := strings.SplitN(text, ",", 3)
	fmt.Printf("Split with limit 3: %v\n", partsLimited)
	
	// Split after
	partsAfter := strings.SplitAfter(text, ",")
	fmt.Printf("SplitAfter: %v\n", partsAfter)
	
	// Split after with limit
	partsAfterLimited := strings.SplitAfterN(text, ",", 3)
	fmt.Printf("SplitAfterN with limit 3: %v\n", partsAfterLimited)
	
	// Fields (split by whitespace)
	textWithSpaces := "  apple   banana  cherry  "
	fields := strings.Fields(textWithSpaces)
	fmt.Printf("Fields of '%s': %v\n", textWithSpaces, fields)
	
	// FieldsFunc
	fieldsFunc := strings.FieldsFunc(text, func(r rune) bool {
		return r == ',' || r == ' '
	})
	fmt.Printf("FieldsFunc: %v\n", fieldsFunc)
}

// DemonstrateRegexOperations shows regex operations
func (sp *StringProcessor) DemonstrateRegexOperations() {
	fmt.Println("\n🔍 REGEX OPERATIONS:")
	fmt.Println("===================")
	
	text := "Contact us at john@example.com or call +1-555-123-4567 or visit https://example.com"
	fmt.Printf("Text: %s\n\n", text)
	
	// Email validation
	emailMatch := sp.Patterns["email"].FindString(text)
	fmt.Printf("Email found: %s\n", emailMatch)
	
	// Phone validation
	phoneMatch := sp.Patterns["phone"].FindString(text)
	fmt.Printf("Phone found: %s\n", phoneMatch)
	
	// URL validation
	urlMatch := sp.Patterns["url"].FindString(text)
	fmt.Printf("URL found: %s\n", urlMatch)
	
	// Find all matches
	allEmails := sp.Patterns["email"].FindAllString(text, -1)
	fmt.Printf("All emails: %v\n", allEmails)
	
	// Find all words
	allWords := sp.Patterns["word"].FindAllString(text, -1)
	fmt.Printf("All words: %v\n", allWords)
	
	// Find all numbers
	allNumbers := sp.Patterns["number"].FindAllString(text, -1)
	fmt.Printf("All numbers: %v\n", allNumbers)
	
	// Replace with regex
	replaced := sp.Patterns["email"].ReplaceAllString(text, "[EMAIL]")
	fmt.Printf("Replace emails: %s\n", replaced)
	
	// Split with regex
	splitByWords := sp.Patterns["word"].Split(text, -1)
	fmt.Printf("Split by words: %v\n", splitByWords)
}

// DemonstrateStringFormatting shows string formatting operations
func (sp *StringProcessor) DemonstrateStringFormatting() {
	fmt.Println("\n📝 STRING FORMATTING OPERATIONS:")
	fmt.Println("================================")
	
	// Basic formatting
	name := "Alice"
	age := 30
	height := 5.6
	
	// Sprintf
	formatted := fmt.Sprintf("Name: %s, Age: %d, Height: %.1f", name, age, height)
	fmt.Printf("Sprintf: %s\n", formatted)
	
	// Different format verbs
	number := 42
	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Octal: %o\n", number)
	fmt.Printf("Hex: %x\n", number)
	fmt.Printf("Hex (uppercase): %X\n", number)
	
	// String formatting
	text := "Hello"
	fmt.Printf("String: %s\n", text)
	fmt.Printf("String (quoted): %q\n", text)
	fmt.Printf("String (hex): %x\n", text)
	
	// Width and precision
	pi := 3.14159265359
	fmt.Printf("Float (default): %f\n", pi)
	fmt.Printf("Float (2 decimals): %.2f\n", pi)
	fmt.Printf("Float (width 10): %10f\n", pi)
	fmt.Printf("Float (width 10, 2 decimals): %10.2f\n", pi)
	
	// Padding
	fmt.Printf("Padded right: %-10s|\n", text)
	fmt.Printf("Padded left:  %10s|\n", text)
	fmt.Printf("Zero padded:  %010d\n", number)
}

// DemonstrateStringConversion shows string conversion operations
func (sp *StringProcessor) DemonstrateStringConversion() {
	fmt.Println("\n🔄 STRING CONVERSION OPERATIONS:")
	fmt.Println("===============================")
	
	// String to number conversions
	str := "12345"
	
	// String to int
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("String '%s' to int: %d\n", str, num)
	}
	
	// String to int64
	if num, err := strconv.ParseInt(str, 10, 64); err == nil {
		fmt.Printf("String '%s' to int64: %d\n", str, num)
	}
	
	// String to float
	floatStr := "3.14159"
	if num, err := strconv.ParseFloat(floatStr, 64); err == nil {
		fmt.Printf("String '%s' to float64: %f\n", floatStr, num)
	}
	
	// String to bool
	boolStr := "true"
	if b, err := strconv.ParseBool(boolStr); err == nil {
		fmt.Printf("String '%s' to bool: %t\n", boolStr, b)
	}
	
	// Number to string conversions
	num := 42
	fmt.Printf("Int %d to string: %s\n", num, strconv.Itoa(num))
	fmt.Printf("Int64 %d to string: %s\n", int64(num), strconv.FormatInt(int64(num), 10))
	
	// Float to string
	pi := 3.14159
	fmt.Printf("Float %.2f to string: %s\n", pi, strconv.FormatFloat(pi, 'f', 2, 64))
	
	// Bool to string
	flag := true
	fmt.Printf("Bool %t to string: %s\n", flag, strconv.FormatBool(flag))
	
	// Base conversions
	number := 255
	fmt.Printf("Decimal %d in different bases:\n", number)
	fmt.Printf("  Binary: %s\n", strconv.FormatInt(int64(number), 2))
	fmt.Printf("  Octal: %s\n", strconv.FormatInt(int64(number), 8))
	fmt.Printf("  Hex: %s\n", strconv.FormatInt(int64(number), 16))
}

// DemonstrateStringBuilder shows efficient string building
func (sp *StringProcessor) DemonstrateStringBuilder() {
	fmt.Println("\n🏗️  STRING BUILDER OPERATIONS:")
	fmt.Println("==============================")
	
	// Using strings.Builder
	var builder strings.Builder
	
	// Write methods
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	builder.WriteString("!")
	
	// Write bytes
	builder.Write([]byte(" from Go"))
	
	// Write rune
	builder.WriteRune('!')
	
	// Write byte
	builder.WriteByte('\n')
	
	// Get the result
	result := builder.String()
	fmt.Printf("Builder result: %s", result)
	
	// Reset builder
	builder.Reset()
	builder.WriteString("Reset and rebuild")
	fmt.Printf("After reset: %s\n", builder.String())
	
	// Check capacity and length
	fmt.Printf("Length: %d, Capacity: %d\n", builder.Len(), builder.Cap())
	
	// Grow capacity
	builder.Grow(100)
	fmt.Printf("After grow: Length: %d, Capacity: %d\n", builder.Len(), builder.Cap())
}

// DemonstrateStringComparison shows string comparison operations
func (sp *StringProcessor) DemonstrateStringComparison() {
	fmt.Println("\n⚖️  STRING COMPARISON OPERATIONS:")
	fmt.Println("=================================")
	
	str1 := "apple"
	str2 := "banana"
	str3 := "apple"
	
	// Equal comparison
	fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)
	fmt.Printf("'%s' == '%s': %t\n", str1, str3, str1 == str3)
	
	// Compare function
	compare := strings.Compare(str1, str2)
	fmt.Printf("Compare('%s', '%s'): %d\n", str1, str2, compare)
	
	compare = strings.Compare(str1, str3)
	fmt.Printf("Compare('%s', '%s'): %d\n", str1, str3, compare)
	
	// EqualFold (case-insensitive)
	str4 := "APPLE"
	fmt.Printf("EqualFold('%s', '%s'): %t\n", str1, str4, strings.EqualFold(str1, str4))
	
	// Sorting
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Original: %v\n", words)
	
	sort.Strings(words)
	fmt.Printf("Sorted: %v\n", words)
	
	// Custom sorting
	words = []string{"banana", "apple", "cherry", "date"}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j]) // Sort by length
	})
	fmt.Printf("Sorted by length: %v\n", words)
}