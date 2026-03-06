# Quick Reference Guide - String Algorithms

## Pattern Recognition Cheat Sheet

### When to use which technique?

#### 1. **Two Pointers**
- ✅ Valid palindrome
- ✅ Reverse string/words
- ✅ Remove duplicates
- ✅ Merge sorted arrays
- ✅ Container with most water
- **Time Complexity**: Usually O(n)
- **Space Complexity**: Usually O(1)

#### 2. **Sliding Window**
- ✅ Longest substring with K distinct characters
- ✅ Minimum window substring
- ✅ Longest substring without repeating characters
- ✅ Permutation in string
- ✅ Maximum average subarray
- **Time Complexity**: Usually O(n)
- **Space Complexity**: Usually O(k) where k is window size

#### 3. **Hash Map/Set**
- ✅ Group anagrams
- ✅ First unique character
- ✅ Isomorphic strings
- ✅ Word pattern
- ✅ Character frequency
- **Time Complexity**: Usually O(n)
- **Space Complexity**: Usually O(n)

#### 4. **Dynamic Programming**
- ✅ Edit distance
- ✅ Longest common subsequence
- ✅ Word break
- ✅ Interleaving string
- ✅ Scramble string
- **Time Complexity**: Usually O(m*n)
- **Space Complexity**: Usually O(m*n) or O(min(m,n))

#### 5. **Pattern Matching**
- ✅ KMP: Find pattern in text
- ✅ Rabin-Karp: Multiple pattern matching
- ✅ Z-Algorithm: Pattern matching
- ✅ Boyer-Moore: Fast pattern matching
- **Time Complexity**: O(n+m) for KMP, Z-Algorithm

#### 6. **Advanced Algorithms**
- ✅ Trie: Prefix matching, autocomplete
- ✅ Manacher's: Longest palindromic substring
- ✅ Suffix Array: Longest repeated substring
- **Time Complexity**: Varies

---

## Common String Operations

### Basic Operations
```cpp
// Reverse
reverse(s.begin(), s.end());

// Substring
string sub = s.substr(start, length);

// Find
size_t pos = s.find(pattern);

// Replace
s.replace(start, length, newStr);

// Convert case
transform(s.begin(), s.end(), s.begin(), ::tolower);
```

### Character Checks
```cpp
isalnum(c)  // Alphanumeric
isalpha(c)  // Alphabet
isdigit(c)  // Digit
islower(c)  // Lowercase
isupper(c)  // Uppercase
```

---

## Template Patterns

### Sliding Window Template
```cpp
int left = 0, right = 0;
while (right < n) {
    // Expand window
    // Update state
    
    while (/* condition to shrink */) {
        // Shrink window
        left++;
    }
    
    // Update answer
    right++;
}
```

### Two Pointers Template
```cpp
int left = 0, right = n - 1;
while (left < right) {
    if (/* condition */) {
        left++;
    } else {
        right--;
    }
}
```

### Hash Map Frequency Template
```cpp
unordered_map<char, int> freq;
for (char c : s) {
    freq[c]++;
}
```

---

## Time Complexity Reference

| Operation | Time Complexity |
|-----------|----------------|
| String comparison | O(n) |
| Substring | O(n) |
| Find | O(n*m) worst case |
| KMP | O(n+m) |
| Rabin-Karp | O(n+m) average |
| Trie insert/search | O(m) where m is word length |
| Manacher's | O(n) |
| Edit Distance (DP) | O(m*n) |
| LCS (DP) | O(m*n) |

---

## Space Complexity Reference

| Algorithm | Space Complexity |
|-----------|------------------|
| Two Pointers | O(1) |
| Sliding Window | O(k) where k is window size |
| Hash Map | O(n) |
| KMP | O(m) |
| Trie | O(ALPHABET_SIZE * N * M) |
| DP (2D) | O(m*n) |
| DP (optimized) | O(min(m,n)) |

---

## Interview Tips

1. **Always clarify**:
   - Case sensitivity?
   - Whitespace handling?
   - Special characters?
   - Empty string edge case?

2. **Common edge cases**:
   - Empty string
   - Single character
   - All same characters
   - Very long string
   - Special characters

3. **Optimization strategy**:
   - Start with brute force
   - Identify bottlenecks
   - Use appropriate data structure
   - Consider space-time tradeoff

4. **Common mistakes**:
   - Off-by-one errors
   - Not handling empty strings
   - Forgetting to reset variables
   - Index out of bounds

---

## Problem Classification

### Easy (Most Common)
- Valid palindrome
- Reverse string
- Valid anagram
- First unique character
- Isomorphic strings

### Medium (Frequently Asked)
- Longest substring without repeating
- Group anagrams
- Minimum window substring
- Longest palindromic substring
- Find all anagrams

### Hard (Advanced)
- Edit distance
- Word break II
- Scramble string
- Minimum window subsequence
- Interleaving string

---

## Practice Strategy

1. **Week 1**: Master basic operations and two pointers
2. **Week 2**: Focus on sliding window and hash maps
3. **Week 3**: Learn pattern matching (KMP, etc.)
4. **Week 4**: Advanced algorithms (Trie, Manacher's)
5. **Week 5**: DP problems on strings
6. **Week 6**: Mixed practice and mock interviews

---

## Resources

- LeetCode String Tag: 200+ problems
- GeeksforGeeks: String algorithms
- Competitive Programming Handbook: String section
- Practice daily: 2-3 problems minimum
