# Complete String Algorithms Guide for DSA Interviews

## Table of Contents

1. [Basic String Operations](#1-basic-string-operations)
2. [Pattern Matching Algorithms](#2-pattern-matching-algorithms)
3. [Sliding Window Patterns](#3-sliding-window-patterns)
4. [Two Pointers Technique](#4-two-pointers-technique)
5. [Hash-Based Techniques](#5-hash-based-techniques)
6. [Advanced Algorithms](#6-advanced-algorithms)
7. [Common Interview Patterns](#7-common-interview-patterns)
8. [Practice Problems](#8-practice-problems)

---

## 1. Basic String Operations

**Files:**
- `01_basic_operations.cpp` - Reverse, palindrome, anagram, substring operations

**Key Concepts:**
- String reversal
- Palindrome checking
- Anagram detection
- Substring operations
- Character frequency counting

---

## 2. Pattern Matching Algorithms

**Files:**
- `02_pattern_matching.cpp` - KMP, Rabin-Karp, Z-algorithm

**Key Concepts:**
- Naive pattern matching
- Knuth-Morris-Pratt (KMP) Algorithm
- Rabin-Karp Algorithm
- Z-Algorithm
- Boyer-Moore Algorithm

---

## 3. Sliding Window Patterns

**Files:**
- `03_sliding_window.cpp` - Fixed and variable window problems

**Key Concepts:**
- Fixed size sliding window
- Variable size sliding window
- Longest substring with K unique characters
- Minimum window substring
- Substring with all characters

---

## 4. Two Pointers Technique

**Files:**
- `04_two_pointers.cpp` - Two pointers for strings

**Key Concepts:**
- Valid palindrome
- Reverse words in string
- Remove duplicates
- Merge sorted strings
- String compression

---

## 5. Hash-Based Techniques

**Files:**
- `05_hash_techniques.cpp` - Hash maps for string problems

**Key Concepts:**
- Character frequency maps
- Anagram groups
- First unique character
- Longest substring without repeating characters
- Group anagrams

---

## 6. Advanced Algorithms

**Files:**
- `06_advanced_algorithms.cpp` - Trie, Suffix Array, Manacher's Algorithm

**Key Concepts:**
- Trie (Prefix Tree)
- Suffix Array
- Manacher's Algorithm (Longest Palindromic Substring)
- Aho-Corasick Algorithm
- Suffix Tree basics

---

## 7. Common Interview Patterns

**Files:**
- `07_common_patterns.cpp` - Frequently asked patterns

**Key Concepts:**
- String encoding/decoding
- String interleaving
- Edit distance (DP)
- Longest Common Subsequence (LCS)
- Word break problem
- String transformation

---

## 8. Practice Problems

**Files:**
- `08_practice_problems.cpp` - LeetCode-style problems with solutions

**Problems Covered:**
- Easy: Valid Anagram, First Unique Character, Reverse String
- Medium: Longest Substring Without Repeating Characters, Group Anagrams, Minimum Window Substring
- Hard: Edit Distance, Longest Palindromic Substring, Word Break II

---

## How to Use This Guide

1. **Start with Basics**: Master basic operations first
2. **Learn Patterns**: Understand each pattern category
3. **Practice**: Solve problems in each category
4. **Review**: Regularly revisit advanced algorithms
5. **Mock Interviews**: Practice explaining solutions

## Time Complexity Cheat Sheet

| Algorithm | Time Complexity | Space Complexity |
|-----------|----------------|------------------|
| Naive Pattern Matching | O(n*m) | O(1) |
| KMP | O(n+m) | O(m) |
| Rabin-Karp | O(n+m) average | O(1) |
| Z-Algorithm | O(n+m) | O(n+m) |
| Manacher's | O(n) | O(n) |
| Trie Search | O(m) | O(ALPHABET_SIZE * N * M) |

---

## Tips for Interviews

1. **Clarify**: Ask about case sensitivity, whitespace, special characters
2. **Edge Cases**: Empty strings, single character, all same characters
3. **Optimize**: Start with brute force, then optimize
4. **Explain**: Walk through your approach before coding
5. **Test**: Always test with examples
