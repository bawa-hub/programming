package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Custom string functions for demonstration
func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && 
		   strings.Contains(email, ".") &&
		   !strings.HasPrefix(email, "@") &&
		   !strings.HasSuffix(email, "@")
}

func IsValidPhone(phone string) bool {
	// Remove all non-digit characters
	digits := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, phone)
	
	// Check if we have 10 digits
	return len(digits) == 10
}

func CountWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func CountLines(text string) int {
	lines := strings.Split(text, "\n")
	return len(lines)
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RemoveDuplicates(s string) string {
	seen := make(map[rune]bool)
	var result strings.Builder
	
	for _, r := range s {
		if !seen[r] {
			seen[r] = true
			result.WriteRune(r)
		}
	}
	
	return result.String()
}

func CapitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

func ExtractNumbers(s string) []int {
	var numbers []int
	words := strings.Fields(s)
	
	for _, word := range words {
		if num, err := strconv.Atoi(word); err == nil {
			numbers = append(numbers, num)
		}
	}
	
	return numbers
}

func IsPalindrome(s string) bool {
	// Remove spaces and convert to lowercase
	cleaned := strings.ReplaceAll(strings.ToLower(s), " ", "")
	return cleaned == ReverseString(cleaned)
}

func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func PadString(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	
	padding := strings.Repeat(string(padChar), length-len(s))
	return s + padding
}

func main() {
	fmt.Println("🚀 Go strings Package Mastery Examples")
	fmt.Println("=====================================")

	// 1. Basic String Operations
	fmt.Println("\n1. Basic String Operations:")
	
	text := "  Hello, World!  "
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(text))
	fmt.Printf("Upper: '%s'\n", strings.ToUpper(text))
	fmt.Printf("Lower: '%s'\n", strings.ToLower(text))
	fmt.Printf("Title: '%s'\n", strings.Title(text))

	// 2. String Comparison
	fmt.Println("\n2. String Comparison:")
	
	str1 := "Hello"
	str2 := "hello"
	str3 := "HELLO"
	
	fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)
	fmt.Printf("'%s' == '%s' (case-insensitive): %t\n", str1, str2, strings.EqualFold(str1, str2))
	fmt.Printf("'%s' == '%s' (case-insensitive): %t\n", str1, str3, strings.EqualFold(str1, str3))
	fmt.Printf("'%s' starts with 'He': %t\n", str1, strings.HasPrefix(str1, "He"))
	fmt.Printf("'%s' ends with 'lo': %t\n", str1, strings.HasSuffix(str1, "lo"))

	// 3. String Searching
	fmt.Println("\n3. String Searching:")
	
	text = "The quick brown fox jumps over the lazy dog"
	fmt.Printf("Text: '%s'\n", text)
	fmt.Printf("Contains 'fox': %t\n", strings.Contains(text, "fox"))
	fmt.Printf("Contains 'cat': %t\n", strings.Contains(text, "cat"))
	fmt.Printf("Index of 'fox': %d\n", strings.Index(text, "fox"))
	fmt.Printf("Last index of 'o': %d\n", strings.LastIndex(text, "o"))
	fmt.Printf("Count of 'o': %d\n", strings.Count(text, "o"))
	fmt.Printf("Contains any 'aeiou': %t\n", strings.ContainsAny(text, "aeiou"))

	// 4. String Splitting
	fmt.Println("\n4. String Splitting:")
	
	csv := "apple,banana,cherry,date"
	fmt.Printf("CSV: '%s'\n", csv)
	
	parts := strings.Split(csv, ",")
	fmt.Printf("Split by ',': %v\n", parts)
	
	parts2 := strings.SplitN(csv, ",", 3)
	fmt.Printf("Split by ',' (max 3): %v\n", parts2)
	
	words := strings.Fields("  hello   world   go  ")
	fmt.Printf("Fields: %v\n", words)

	// 5. String Joining
	fmt.Println("\n5. String Joining:")
	
	items := []string{"apple", "banana", "cherry"}
	fmt.Printf("Items: %v\n", items)
	
	joined := strings.Join(items, " | ")
	fmt.Printf("Joined with ' | ': '%s'\n", joined)
	
	repeated := strings.Repeat("Go ", 3)
	fmt.Printf("Repeated 'Go ' 3 times: '%s'\n", repeated)

	// 6. String Replacement
	fmt.Println("\n6. String Replacement:")
	
	text = "Hello World Hello"
	fmt.Printf("Original: '%s'\n", text)
	
	replaced := strings.Replace(text, "Hello", "Hi", 1)
	fmt.Printf("Replace first 'Hello': '%s'\n", replaced)
	
	replacedAll := strings.ReplaceAll(text, "Hello", "Hi")
	fmt.Printf("Replace all 'Hello': '%s'\n", replacedAll)

	// 7. String Trimming
	fmt.Println("\n7. String Trimming:")
	
	text = "!!!Hello World!!!"
	fmt.Printf("Original: '%s'\n", text)
	
	trimmed := strings.Trim(text, "!")
	fmt.Printf("Trim '!': '%s'\n", trimmed)
	
	trimmedLeft := strings.TrimLeft(text, "!")
	fmt.Printf("Trim left '!': '%s'\n", trimmedLeft)
	
	trimmedRight := strings.TrimRight(text, "!")
	fmt.Printf("Trim right '!': '%s'\n", trimmedRight)
	
	trimmedPrefix := strings.TrimPrefix(text, "!!!")
	fmt.Printf("Trim prefix '!!!': '%s'\n", trimmedPrefix)
	
	trimmedSuffix := strings.TrimSuffix(text, "!!!")
	fmt.Printf("Trim suffix '!!!': '%s'\n", trimmedSuffix)

	// 8. String Building
	fmt.Println("\n8. String Building:")
	
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	builder.WriteRune('!')
	builder.WriteByte('\n')
	
	result := builder.String()
	fmt.Printf("Built string: '%s'\n", result)
	
	// Reset builder
	builder.Reset()
	builder.WriteString("New content")
	fmt.Printf("Reset and new content: '%s'\n", builder.String())

	// 9. String Mapping
	fmt.Println("\n9. String Mapping:")
	
	text = "Hello World 123"
	fmt.Printf("Original: '%s'\n", text)
	
	// Convert to uppercase
	upper := strings.Map(unicode.ToUpper, text)
	fmt.Printf("To upper: '%s'\n", upper)
	
	// Remove digits
	noDigits := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return -1
		}
		return r
	}, text)
	fmt.Printf("Remove digits: '%s'\n", noDigits)
	
	// Keep only letters
	onlyLetters := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, text)
	fmt.Printf("Only letters: '%s'\n", onlyLetters)

	// 10. String Validation
	fmt.Println("\n10. String Validation:")
	
	emails := []string{"user@example.com", "invalid-email", "@example.com", "user@"}
	for _, email := range emails {
		fmt.Printf("'%s' is valid email: %t\n", email, IsValidEmail(email))
	}
	
	phones := []string{"1234567890", "(123) 456-7890", "123-456-7890", "123456789"}
	for _, phone := range phones {
		fmt.Printf("'%s' is valid phone: %t\n", phone, IsValidPhone(phone))
	}

	// 11. String Analysis
	fmt.Println("\n11. String Analysis:")
	
	text = "The quick brown fox jumps over the lazy dog"
	fmt.Printf("Text: '%s'\n", text)
	fmt.Printf("Word count: %d\n", CountWords(text))
	fmt.Printf("Character count: %d\n", len(text))
	fmt.Printf("Rune count: %d\n", len([]rune(text)))
	
	multilineText := "Line 1\nLine 2\nLine 3"
	fmt.Printf("Multiline text:\n%s\n", multilineText)
	fmt.Printf("Line count: %d\n", CountLines(multilineText))

	// 12. String Transformation
	fmt.Println("\n12. String Transformation:")
	
	text = "hello world"
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("Reversed: '%s'\n", ReverseString(text))
	fmt.Printf("Capitalized: '%s'\n", CapitalizeWords(text))
	
	textWithDups := "hello world"
	fmt.Printf("Original with duplicates: '%s'\n", textWithDups)
	fmt.Printf("Remove duplicates: '%s'\n", RemoveDuplicates(textWithDups))

	// 13. String Extraction
	fmt.Println("\n13. String Extraction:")
	
	text = "I have 5 apples and 3 oranges"
	fmt.Printf("Text: '%s'\n", text)
	numbers := ExtractNumbers(text)
	fmt.Printf("Extracted numbers: %v\n", numbers)

	// 14. String Palindrome Check
	fmt.Println("\n14. String Palindrome Check:")
	
	palindromes := []string{"racecar", "hello", "level", "world", "madam"}
	for _, word := range palindromes {
		fmt.Printf("'%s' is palindrome: %t\n", word, IsPalindrome(word))
	}

	// 15. String Truncation and Padding
	fmt.Println("\n15. String Truncation and Padding:")
	
	text = "This is a very long string that needs to be truncated"
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("Truncated to 20 chars: '%s'\n", TruncateString(text, 20))
	
	shortText := "Hi"
	fmt.Printf("Original: '%s'\n", shortText)
	fmt.Printf("Padded to 10 chars with '*': '%s'\n", PadString(shortText, 10, '*'))

	// 16. String Comparison Functions
	fmt.Println("\n16. String Comparison Functions:")
	
	str1Comp := "apple"
	str2Comp := "banana"
	str3Comp := "apple"
	
	fmt.Printf("Compare('%s', '%s'): %d\n", str1Comp, str2Comp, strings.Compare(str1Comp, str2Comp))
	fmt.Printf("Compare('%s', '%s'): %d\n", str2Comp, str1Comp, strings.Compare(str2Comp, str1Comp))
	fmt.Printf("Compare('%s', '%s'): %d\n", str1Comp, str3Comp, strings.Compare(str1Comp, str3Comp))

	// 17. String Index Functions
	fmt.Println("\n17. String Index Functions:")
	
	text = "Hello, World!"
	fmt.Printf("Text: '%s'\n", text)
	
	fmt.Printf("Index of 'o': %d\n", strings.Index(text, "o"))
	fmt.Printf("Last index of 'o': %d\n", strings.LastIndex(text, "o"))
	fmt.Printf("Index of any 'aeiou': %d\n", strings.IndexAny(text, "aeiou"))
	fmt.Printf("Last index of any 'aeiou': %d\n", strings.LastIndexAny(text, "aeiou"))
	fmt.Printf("Index of rune 'o': %d\n", strings.IndexRune(text, 'o'))

	// 18. String Split Functions
	fmt.Println("\n18. String Split Functions:")
	
	text = "a,b,c,d,e"
	fmt.Printf("Text: '%s'\n", text)
	
	split := strings.Split(text, ",")
	fmt.Printf("Split: %v\n", split)
	
	splitAfter := strings.SplitAfter(text, ",")
	fmt.Printf("Split after: %v\n", splitAfter)
	
	splitN := strings.SplitN(text, ",", 3)
	fmt.Printf("Split N (3): %v\n", splitN)
	
	splitAfterN := strings.SplitAfterN(text, ",", 3)
	fmt.Printf("Split after N (3): %v\n", splitAfterN)

	// 19. String Fields Functions
	fmt.Println("\n19. String Fields Functions:")
	
	text = "  hello   world   go  "
	fmt.Printf("Text: '%s'\n", text)
	
	fields := strings.Fields(text)
	fmt.Printf("Fields: %v\n", fields)
	
	fieldsFunc := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == '\t'
	})
	fmt.Printf("Fields func: %v\n", fieldsFunc)

	// 20. String Performance Test
	fmt.Println("\n20. String Performance Test:")
	
	// Test string concatenation methods
	wordsPerf := []string{"Hello", "World", "Go", "Programming", "Language"}
	
	// Method 1: Using + operator
	start := time.Now()
	result1 := ""
	for _, word := range wordsPerf {
		result1 += word + " "
	}
	time1 := time.Since(start)
	
	// Method 2: Using strings.Builder
	start2 := time.Now()
	var builder2 strings.Builder
	for _, word := range wordsPerf {
		builder2.WriteString(word)
		builder2.WriteString(" ")
	}
	result2 := builder2.String()
	time2 := time.Since(start2)
	
	// Method 3: Using strings.Join
	start3 := time.Now()
	result3 := strings.Join(wordsPerf, " ") + " "
	time3 := time.Since(start3)
	
	fmt.Printf("Result 1 (+ operator): '%s' (time: %v)\n", result1, time1)
	fmt.Printf("Result 2 (Builder): '%s' (time: %v)\n", result2, time2)
	fmt.Printf("Result 3 (Join): '%s' (time: %v)\n", result3, time3)
	
	if time2 < time1 {
		fmt.Printf("Builder is %.2fx faster than + operator\n", float64(time1)/float64(time2))
	}
	if time3 < time1 {
		fmt.Printf("Join is %.2fx faster than + operator\n", float64(time1)/float64(time3))
	}

	// 21. String Unicode Handling
	fmt.Println("\n21. String Unicode Handling:")
	
	unicodeText := "Hello 世界 🌍"
	fmt.Printf("Text: '%s'\n", unicodeText)
	fmt.Printf("Byte length: %d\n", len(unicodeText))
	fmt.Printf("Rune length: %d\n", len([]rune(unicodeText)))
	
	// Count runes
	runeCount := 0
	for range unicodeText {
		runeCount++
	}
	fmt.Printf("Rune count (manual): %d\n", runeCount)
	
	// Count specific runes
	spaceCount := strings.Count(unicodeText, " ")
	fmt.Printf("Space count: %d\n", spaceCount)

	// 22. String Case Functions
	fmt.Println("\n22. String Case Functions:")
	
	text = "hello world"
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("ToUpper: '%s'\n", strings.ToUpper(text))
	fmt.Printf("ToLower: '%s'\n", strings.ToLower(text))
	fmt.Printf("ToTitle: '%s'\n", strings.ToTitle(text))
	
	// Test with mixed case
	mixedText := "hELLo WoRLd"
	fmt.Printf("Mixed case: '%s'\n", mixedText)
	fmt.Printf("ToUpper: '%s'\n", strings.ToUpper(mixedText))
	fmt.Printf("ToLower: '%s'\n", strings.ToLower(mixedText))

	// 23. String Contains Functions
	fmt.Println("\n23. String Contains Functions:")
	
	text = "Hello, World!"
	fmt.Printf("Text: '%s'\n", text)
	
	fmt.Printf("Contains 'World': %t\n", strings.Contains(text, "World"))
	fmt.Printf("Contains 'world': %t\n", strings.Contains(text, "world"))
	fmt.Printf("Contains any 'aeiou': %t\n", strings.ContainsAny(text, "aeiou"))
	fmt.Printf("Contains rune '!': %t\n", strings.ContainsRune(text, '!'))
	fmt.Printf("Contains rune '?': %t\n", strings.ContainsRune(text, '?'))

	// 24. String Replace Functions
	fmt.Println("\n24. String Replace Functions:")
	
	text = "Hello Hello Hello"
	fmt.Printf("Original: '%s'\n", text)
	
	replace1 := strings.Replace(text, "Hello", "Hi", 1)
	fmt.Printf("Replace 1: '%s'\n", replace1)
	
	replace2 := strings.Replace(text, "Hello", "Hi", 2)
	fmt.Printf("Replace 2: '%s'\n", replace2)
	
	replaceAll := strings.ReplaceAll(text, "Hello", "Hi")
	fmt.Printf("Replace all: '%s'\n", replaceAll)

	// 25. String Trim Functions
	fmt.Println("\n25. String Trim Functions:")
	
	text = "!!!Hello World!!!"
	fmt.Printf("Original: '%s'\n", text)
	
	trim := strings.Trim(text, "!")
	fmt.Printf("Trim '!': '%s'\n", trim)
	
	trimLeft := strings.TrimLeft(text, "!")
	fmt.Printf("Trim left '!': '%s'\n", trimLeft)
	
	trimRight := strings.TrimRight(text, "!")
	fmt.Printf("Trim right '!': '%s'\n", trimRight)
	
	trimSpace := strings.TrimSpace("  hello world  ")
	fmt.Printf("Trim space: '%s'\n", trimSpace)

	fmt.Println("\n🎉 strings Package Mastery Complete!")
}
