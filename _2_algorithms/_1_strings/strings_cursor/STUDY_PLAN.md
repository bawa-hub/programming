# String Algorithms Study Plan

## Overview
This repository contains a complete guide to mastering string algorithms for DSA interviews at big tech companies. All code is implemented in C++.

## File Structure

```
strings/
├── README.md                    # Main overview and table of contents
├── QUICK_REFERENCE.md           # Quick reference guide and cheat sheet
├── STUDY_PLAN.md                # This file - study plan
├── Makefile                     # Build all programs
│
├── 01_basic_operations.cpp      # Basic string operations
├── 02_pattern_matching.cpp      # KMP, Rabin-Karp, Z-algorithm
├── 03_sliding_window.cpp        # Sliding window patterns
├── 04_two_pointers.cpp          # Two pointers technique
├── 05_hash_techniques.cpp       # Hash-based solutions
├── 06_advanced_algorithms.cpp   # Trie, Manacher's, Suffix Array
├── 07_common_patterns.cpp       # Common interview patterns
└── 08_practice_problems.cpp     # LeetCode-style problems
```

## 4-Week Study Plan

### Week 1: Foundations
**Goal**: Master basic operations and two pointers

**Day 1-2: Basic Operations**
- Study: `01_basic_operations.cpp`
- Practice:
  - Reverse string (LeetCode 344)
  - Valid palindrome (LeetCode 125)
  - Valid anagram (LeetCode 242)
  - First unique character (LeetCode 387)

**Day 3-4: Two Pointers**
- Study: `04_two_pointers.cpp`
- Practice:
  - Reverse words (LeetCode 151)
  - Valid palindrome II (LeetCode 680)
  - Remove duplicates
  - String compression (LeetCode 443)

**Day 5-7: Review & Practice**
- Solve 10 easy string problems
- Review concepts
- Time yourself on problems

---

### Week 2: Core Techniques
**Goal**: Master sliding window and hash maps

**Day 1-3: Sliding Window**
- Study: `03_sliding_window.cpp`
- Practice:
  - Longest substring without repeating (LeetCode 3)
  - Minimum window substring (LeetCode 76)
  - Longest substring with K distinct (LeetCode 340)
  - Permutation in string (LeetCode 567)

**Day 4-5: Hash Techniques**
- Study: `05_hash_techniques.cpp`
- Practice:
  - Group anagrams (LeetCode 49)
  - Isomorphic strings (LeetCode 205)
  - Word pattern (LeetCode 290)
  - Find all anagrams (LeetCode 438)

**Day 6-7: Review & Practice**
- Solve 15 medium string problems
- Focus on pattern recognition
- Optimize solutions

---

### Week 3: Pattern Matching & Advanced
**Goal**: Learn pattern matching and advanced algorithms

**Day 1-2: Pattern Matching**
- Study: `02_pattern_matching.cpp`
- Understand:
  - KMP algorithm (most important)
  - Rabin-Karp
  - Z-algorithm
- Practice: Implement KMP from scratch

**Day 3-4: Advanced Algorithms**
- Study: `06_advanced_algorithms.cpp`
- Understand:
  - Trie data structure
  - Manacher's algorithm
  - Longest Common Subsequence
  - Edit Distance

**Day 5-7: Review & Practice**
- Solve 10 medium-hard problems
- Implement algorithms from scratch
- Understand time/space complexity

---

### Week 4: Mastery & Interview Prep
**Goal**: Master all patterns and interview readiness

**Day 1-2: Common Patterns**
- Study: `07_common_patterns.cpp`
- Practice:
  - String encoding/decoding
  - Valid parentheses
  - Roman to integer
  - Zigzag conversion

**Day 3-4: Practice Problems**
- Study: `08_practice_problems.cpp`
- Solve problems across all difficulty levels
- Time yourself: 30-45 min per problem

**Day 5-7: Mock Interviews**
- Solve 3-5 problems daily
- Practice explaining solutions
- Review weak areas
- Focus on optimization

---

## Daily Practice Routine

### Morning (30-45 minutes)
1. Review one algorithm/pattern (15 min)
2. Solve 1-2 problems (30 min)

### Evening (1-2 hours)
1. Solve 2-3 problems
2. Review solutions
3. Implement algorithms from scratch

### Weekly
- Review all patterns
- Solve 20+ problems
- Mock interview (2 hours)

---

## Problem Difficulty Distribution

### Target Distribution
- **Easy**: 30% (build confidence)
- **Medium**: 50% (most interview questions)
- **Hard**: 20% (advanced companies)

### Recommended Problems by Topic

#### Two Pointers (10 problems)
- Easy: 5 problems
- Medium: 5 problems

#### Sliding Window (15 problems)
- Easy: 3 problems
- Medium: 10 problems
- Hard: 2 problems

#### Hash Maps (12 problems)
- Easy: 4 problems
- Medium: 8 problems

#### Pattern Matching (8 problems)
- Medium: 6 problems
- Hard: 2 problems

#### Advanced (10 problems)
- Medium: 5 problems
- Hard: 5 problems

---

## Key Concepts to Master

### Must Know
1. ✅ Two pointers technique
2. ✅ Sliding window (fixed & variable)
3. ✅ Hash map for frequency counting
4. ✅ KMP algorithm
5. ✅ Dynamic programming for strings
6. ✅ Trie data structure

### Should Know
1. ✅ Manacher's algorithm
2. ✅ Rabin-Karp
3. ✅ Z-algorithm
4. ✅ Suffix array basics
5. ✅ Edit distance variations

### Nice to Know
1. ✅ Aho-Corasick
2. ✅ Suffix tree
3. ✅ Advanced DP optimizations

---

## Interview Strategy

### Before Coding
1. **Clarify requirements** (2 min)
   - Case sensitivity?
   - Whitespace handling?
   - Edge cases?

2. **Think out loud** (3-5 min)
   - Explain approach
   - Discuss time/space complexity
   - Mention edge cases

3. **Code** (15-20 min)
   - Clean, readable code
   - Proper variable names
   - Comments for complex logic

4. **Test** (5 min)
   - Walk through examples
   - Test edge cases
   - Verify complexity

### Common Mistakes to Avoid
- ❌ Not handling empty strings
- ❌ Off-by-one errors
- ❌ Forgetting to reset variables
- ❌ Not optimizing after brute force
- ❌ Not explaining approach clearly

---

## Resources

### Primary Resources
- This repository (complete guide)
- LeetCode String Tag (200+ problems)
- GeeksforGeeks String section

### Additional Resources
- Competitive Programming Handbook
- Cracking the Coding Interview (String chapter)
- Algorithm Design Manual

### Practice Platforms
- LeetCode (primary)
- HackerRank
- Codeforces
- InterviewBit

---

## Success Metrics

### Week 1 Goals
- ✅ Solve 20 easy problems
- ✅ Understand two pointers
- ✅ Can implement basic operations

### Week 2 Goals
- ✅ Solve 30 medium problems
- ✅ Master sliding window
- ✅ Comfortable with hash maps

### Week 3 Goals
- ✅ Understand KMP algorithm
- ✅ Can implement Trie
- ✅ Solve 20 medium-hard problems

### Week 4 Goals
- ✅ Solve 50+ problems total
- ✅ Can explain all algorithms
- ✅ Ready for interviews

---

## Tips for Success

1. **Consistency > Intensity**
   - Practice daily, even if just 30 minutes
   - Better than cramming

2. **Understand, Don't Memorize**
   - Understand why algorithms work
   - Can derive from first principles

3. **Pattern Recognition**
   - Learn to identify which technique to use
   - Practice pattern matching

4. **Time Management**
   - Easy: 10-15 minutes
   - Medium: 20-30 minutes
   - Hard: 45-60 minutes

5. **Review Regularly**
   - Review solved problems weekly
   - Re-implement from scratch
   - Explain to someone else

---

## Final Checklist Before Interview

- [ ] Can implement KMP from scratch
- [ ] Comfortable with sliding window patterns
- [ ] Can solve edit distance problem
- [ ] Understand Trie operations
- [ ] Can explain Manacher's algorithm
- [ ] Solved 50+ string problems
- [ ] Can identify pattern quickly
- [ ] Comfortable with time constraints
- [ ] Can explain solutions clearly
- [ ] Handled all edge cases

---

## Good Luck! 🚀

Remember: Mastery comes from consistent practice. Focus on understanding patterns rather than memorizing solutions. You've got this!
