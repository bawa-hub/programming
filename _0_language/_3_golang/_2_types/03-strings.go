package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	// ----- STRINGS -----
	// Strings are arrays of bytes []byte
	// Escape Sequences : \n \t \" \\
	sV1 := "A word"

	// this will print the bytes (unicode values)
	for i := 0; i < len(sV1); i++ {
		fmt.Println(sV1[i])
	}
	// 65
	// 32
	// 119
	// 111
	// 114
	// 100

	// ----- RUNES -----
	// In Go characters are called Runes
	// Runes are unicodes that represent characters
	// An alias for int32
	rStr := "abcdefg"

	// Runes in string
	pl("Rune Count :", utf8.RuneCountInString(rStr))

	// Print runes in string
	for i, runeVal := range rStr {
		// Get index, Rune unicode and character
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

	// print as char
	for _, ch := range rStr {
		fmt.Println(ch, string(ch))
	}

	// 	🧠 4️⃣ When To Use What?

	// ✅ Use byte when:
	// Problem says lowercase letters only
	// ASCII characters only
	// Competitive programming
	// Frequency array of size 26

	// count := make([]int, 26)
	// count[s[i]-'a']++

	// ✅ Use rune when:
	// Unicode involved
	// Multi-language input
	// Real world applications
	// Safe string reversal

	// runes := []rune(s)

	// 🔵 CASE 1 — ASCII Only (Use byte)
	// Problem: Count lowercase letters.
	s := "vikram"

	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	fmt.Println("Frequency of 'v':", count['v'-'a'])

	// 🔵 CASE 2 — Unicode String (byte breaks)
	s = "你好"

	fmt.Println("Length:", len(s))
	// Length = 6
	// Because each Chinese character = 3 bytes

	for i := 0; i < len(s); i++ {
		fmt.Println("Byte:", s[i])
	}

	// 🔵 CASE 3 — Correct Unicode Iteration (Use rune)
	fmt.Println("Length in bytes:", len(s))
	fmt.Println("Length in runes:", len([]rune(s)))

	for _, ch := range s {
		fmt.Println("Rune:", ch, "Character:", string(ch))
	}

	// 🔵 CASE 4 — Reversing String (Wrong vs Correct)
	pl(reverse(s))

	// 🔵 CASE 5 — When range Automatically Uses Rune
	for i, ch := range s {
		fmt.Println("Index:", i, "Rune:", ch, "Char:", string(ch))
	}

	// yad rkhne yogya batein
	s = "你好"
	bytes := []byte(s)
	runes := []rune(s)
	for i := 0; i < len(bytes); i++ {
		fmt.Println(string(bytes[i])) // this will not work
	}
	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i])) // this will work
	}
}

func reverse(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

func methods() {
	sV1 := "A word"

	// Replacer that can be used on multiple strings to replace one string with another
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	// Get length
	pl("Length : ", len(sV2))

	// Contains string
	pl("Contains Another :", strings.Contains(sV2, "Another"))

	// Get first index match
	pl("o index :", strings.Index(sV2, "o"))

	// Replace all matches with 0
	// If -1 was 2 it would replace the 1st 2 matches
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))

	// Remove whitespace characters from beginning and end of string
	sV3 := "\nSome words\n"
	sV3 = strings.TrimSpace(sV3)

	// Split at delimiter
	pl("Split :", strings.Split("a-b-c-d", "-"))

	// Upper and lowercase string
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))

	// Prefix or suffix
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))
}
