What is Pattern Searching?
Given a text and a pattern, the goal is to find all occurrences (indexes) of the pattern inside the text.

Common Algorithms for Pattern Searching:

    Naive Approach (Brute Force) — O(n*m) time, simple but inefficient.
    Rabin-Karp Algorithm — Uses hashing, average O(n + m), good for multiple pattern searches.
    Knuth-Morris-Pratt (KMP) Algorithm — Uses prefix-function (LPS array), O(n + m) time.
    Z Algorithm — Uses Z-array to find pattern occurrences, O(n + m).
    Boyer-Moore Algorithm — Uses bad character and good suffix heuristics, efficient in practice.