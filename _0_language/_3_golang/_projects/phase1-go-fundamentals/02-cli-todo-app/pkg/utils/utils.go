package utils

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

// Color constants for terminal output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

// Colorize applies color to text
func Colorize(text string, color string) string {
	return color + text + ColorReset
}

// Bold makes text bold
func Bold(text string) string {
	return Colorize(text, ColorBold)
}

// Red makes text red
func Red(text string) string {
	return Colorize(text, ColorRed)
}

// Green makes text green
func Green(text string) string {
	return Colorize(text, ColorGreen)
}

// Yellow makes text yellow
func Yellow(text string) string {
	return Colorize(text, ColorYellow)
}

// Blue makes text blue
func Blue(text string) string {
	return Colorize(text, ColorBlue)
}

// Purple makes text purple
func Purple(text string) string {
	return Colorize(text, ColorPurple)
}

// Cyan makes text cyan
func Cyan(text string) string {
	return Colorize(text, ColorCyan)
}

// String utilities demonstrating string operations

// TruncateString truncates a string to the specified length
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// PadString pads a string to the specified length
func PadString(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(" ", length-len(s))
}

// CenterString centers a string within the specified width
func CenterString(s string, width int) string {
	if len(s) >= width {
		return s
	}
	
	padding := width - len(s)
	leftPadding := padding / 2
	rightPadding := padding - leftPadding
	
	return strings.Repeat(" ", leftPadding) + s + strings.Repeat(" ", rightPadding)
}

// ReverseString reverses a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return s == ReverseString(s)
}

// CountWords counts the number of words in a string
func CountWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// CountCharacters counts the number of characters in a string
func CountCharacters(s string) int {
	return len([]rune(s))
}

// CountBytes counts the number of bytes in a string
func CountBytes(s string) int {
	return len(s)
}

// Array and slice utilities demonstrating collection operations

// FilterStrings filters a slice of strings based on a predicate
func FilterStrings(slice []string, predicate func(string) bool) []string {
	var result []string
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// MapStrings applies a function to each string in a slice
func MapStrings(slice []string, mapper func(string) string) []string {
	result := make([]string, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

// ReduceStrings reduces a slice of strings to a single value
func ReduceStrings(slice []string, reducer func(string, string) string, initial string) string {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

// ContainsString checks if a slice contains a string
func ContainsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// RemoveString removes all occurrences of a string from a slice
func RemoveString(slice []string, item string) []string {
	var result []string
	for _, s := range slice {
		if s != item {
			result = append(result, s)
		}
	}
	return result
}

// UniqueStrings returns a slice with unique strings
func UniqueStrings(slice []string) []string {
	seen := make(map[string]bool)
	var result []string
	
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// SortStrings sorts a slice of strings
func SortStrings(slice []string) []string {
	result := make([]string, len(slice))
	copy(result, slice)
	sort.Strings(result)
	return result
}

// ReverseStrings reverses a slice of strings
func ReverseStrings(slice []string) []string {
	result := make([]string, len(slice))
	for i, item := range slice {
		result[len(slice)-1-i] = item
	}
	return result
}

// ChunkStrings splits a slice into chunks of specified size
func ChunkStrings(slice []string, chunkSize int) [][]string {
	var result [][]string
	
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	
	return result
}

// Numeric utilities demonstrating math operations

// SumInts calculates the sum of integers
func SumInts(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// AverageInts calculates the average of integers
func AverageInts(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(SumInts(numbers)) / float64(len(numbers))
}

// MaxInt finds the maximum integer in a slice
func MaxInt(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// MinInt finds the minimum integer in a slice
func MinInt(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

// Factorial calculates the factorial of a number
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fibonacci calculates the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Time utilities demonstrating time operations

// FormatDuration formats a duration in a human-readable way
func FormatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%.0fs", d.Seconds())
	} else if d < time.Hour {
		return fmt.Sprintf("%.0fm", d.Minutes())
	} else if d < 24*time.Hour {
		return fmt.Sprintf("%.1fh", d.Hours())
	} else {
		days := d.Hours() / 24
		return fmt.Sprintf("%.1fd", days)
	}
}

// TimeAgo returns a human-readable time ago string
func TimeAgo(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)
	
	if duration < time.Minute {
		return "just now"
	} else if duration < time.Hour {
		minutes := int(duration.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	} else if duration < 24*time.Hour {
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else {
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}
}

// IsToday checks if a time is today
func IsToday(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() && t.YearDay() == now.YearDay()
}

// IsYesterday checks if a time is yesterday
func IsYesterday(t time.Time) bool {
	yesterday := time.Now().AddDate(0, 0, -1)
	return t.Year() == yesterday.Year() && t.YearDay() == yesterday.YearDay()
}

// IsThisWeek checks if a time is this week
func IsThisWeek(t time.Time) bool {
	now := time.Now()
	year, week := now.ISOWeek()
	tYear, tWeek := t.ISOWeek()
	return year == tYear && week == tWeek
}

// Interface utilities demonstrating interface usage

// Stringer interface for custom string representation
type Stringer interface {
	String() string
}

// PrintStringers prints a slice of Stringers
func PrintStringers(items []Stringer) {
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item.String())
	}
}

// Comparable interface for items that can be compared
type Comparable interface {
	Compare(other Comparable) int
}

// SortComparables sorts a slice of Comparables
func SortComparables(items []Comparable) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Compare(items[j]) < 0
	})
}

// Pointer utilities demonstrating pointer operations

// SwapInts swaps two integers using pointers
func SwapInts(a, b *int) {
	*a, *b = *b, *a
}

// SwapStrings swaps two strings using pointers
func SwapStrings(a, b *string) {
	*a, *b = *b, *a
}

// SafeDereference safely dereferences a pointer
func SafeDereference(ptr *int) (int, bool) {
	if ptr == nil {
		return 0, false
	}
	return *ptr, true
}

// CreateIntPointer creates a pointer to an integer
func CreateIntPointer(value int) *int {
	return &value
}

// CreateStringPointer creates a pointer to a string
func CreateStringPointer(value string) *string {
	return &value
}

// Struct utilities demonstrating struct operations

// Point represents a 2D point
type Point struct {
	X, Y float64
}

// Distance calculates the distance between two points
func (p Point) Distance(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Move moves a point by the given offset
func (p *Point) Move(dx, dy float64) {
	p.X += dx
	p.Y += dy
}

// String returns the string representation of a point
func (p Point) String() string {
	return fmt.Sprintf("Point(%.2f, %.2f)", p.X, p.Y)
}

// Rectangle represents a rectangle
type Rectangle struct {
	TopLeft, BottomRight Point
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	width := r.BottomRight.X - r.TopLeft.X
	height := r.BottomRight.Y - r.TopLeft.Y
	return width * height
}

// Perimeter calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	width := r.BottomRight.X - r.TopLeft.X
	height := r.BottomRight.Y - r.TopLeft.Y
	return 2 * (width + height)
}

// String returns the string representation of a rectangle
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.2f, %.2f) to (%.2f, %.2f)", 
		r.TopLeft.X, r.TopLeft.Y, r.BottomRight.X, r.BottomRight.Y)
}

// Validation utilities

// IsValidEmail checks if a string is a valid email
func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// IsValidPhone checks if a string is a valid phone number
func IsValidPhone(phone string) bool {
	// Simple validation - just check if it contains only digits and common phone characters
	for _, char := range phone {
		if char < '0' || char > '9' {
			if char != '-' && char != '(' && char != ')' && char != ' ' && char != '+' {
				return false
			}
		}
	}
	return len(phone) >= 10
}

// IsValidURL checks if a string is a valid URL
func IsValidURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

// File utilities

// GetFileExtension returns the file extension
func GetFileExtension(filename string) string {
	lastDot := strings.LastIndex(filename, ".")
	if lastDot == -1 {
		return ""
	}
	return filename[lastDot+1:]
}

// GetFileName returns the filename without extension
func GetFileName(filename string) string {
	lastDot := strings.LastIndex(filename, ".")
	if lastDot == -1 {
		return filename
	}
	return filename[:lastDot]
}

// GetDirectory returns the directory path
func GetDirectory(filepath string) string {
	lastSlash := strings.LastIndex(filepath, "/")
	if lastSlash == -1 {
		return "."
	}
	return filepath[:lastSlash]
}

// TagCount represents a tag with its usage count
type TagCount struct {
	Tag   string
	Count int
}

// String returns the string representation of TagCount
func (tc TagCount) String() string {
	return fmt.Sprintf("%s (%d)", tc.Tag, tc.Count)
}
