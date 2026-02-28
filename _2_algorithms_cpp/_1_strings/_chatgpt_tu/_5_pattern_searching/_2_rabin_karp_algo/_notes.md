What is Rabin-Karp?
Rabin-Karp is a string searching algorithm that uses hashing to find a pattern inside text efficiently.

Idea:

    Calculate hash of the pattern and hash of each substring of text with the same length.
    If hashes match, compare characters to avoid false positives (due to collisions).
    Using rolling hash, we can compute substring hashes in O(1) after initial calculation.

Time Complexity:

    Average and best case: O(n + m)
    Worst case: O(n*m) (due to collisions but rare with good hashing)    