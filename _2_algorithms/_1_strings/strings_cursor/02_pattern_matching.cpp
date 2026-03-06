/*
 * Pattern Matching Algorithms
 * Essential algorithms for finding patterns in strings
 */

#include <iostream>
#include <string>
#include <vector>
#include <unordered_set>

using namespace std;

// ==================== 1. NAIVE PATTERN MATCHING ====================

// Time: O(n*m), Space: O(1)
vector<int> naivePatternMatching(string text, string pattern) {
    vector<int> positions;
    int n = text.length();
    int m = pattern.length();
    
    for (int i = 0; i <= n - m; i++) {
        int j;
        for (j = 0; j < m; j++) {
            if (text[i + j] != pattern[j]) {
                break;
            }
        }
        if (j == m) {
            positions.push_back(i);
        }
    }
    
    return positions;
}

// ==================== 2. KMP (KNUTH-MORRIS-PRATT) ALGORITHM ====================

// Build Longest Prefix Suffix (LPS) array
vector<int> buildLPS(string pattern) {
    int m = pattern.length();
    vector<int> lps(m, 0);
    int len = 0;  // Length of previous longest prefix suffix
    int i = 1;
    
    while (i < m) {
        if (pattern[i] == pattern[len]) {
            len++;
            lps[i] = len;
            i++;
        } else {
            if (len != 0) {
                len = lps[len - 1];
            } else {
                lps[i] = 0;
                i++;
            }
        }
    }
    
    return lps;
}

// KMP Pattern Matching
// Time: O(n + m), Space: O(m)
vector<int> kmpPatternMatching(string text, string pattern) {
    vector<int> positions;
    int n = text.length();
    int m = pattern.length();
    
    if (m == 0) return positions;
    
    vector<int> lps = buildLPS(pattern);
    
    int i = 0;  // Index for text
    int j = 0;  // Index for pattern
    
    while (i < n) {
        if (text[i] == pattern[j]) {
            i++;
            j++;
        }
        
        if (j == m) {
            positions.push_back(i - j);
            j = lps[j - 1];
        } else if (i < n && text[i] != pattern[j]) {
            if (j != 0) {
                j = lps[j - 1];
            } else {
                i++;
            }
        }
    }
    
    return positions;
}

// ==================== 3. RABIN-KARP ALGORITHM ====================

// Using rolling hash for pattern matching
// Time: O(n+m) average, O(n*m) worst case, Space: O(1)

const int BASE = 256;
const int MOD = 101;  // Prime number for modulo

long long calculateHash(string str, int start, int end) {
    long long hash = 0;
    for (int i = start; i <= end; i++) {
        hash = (hash * BASE + str[i]) % MOD;
    }
    return hash;
}

// Helper function to calculate power (for rolling hash)
long long powerMod(long long base, int exp, long long mod) {
    long long result = 1;
    for (int i = 0; i < exp; i++) {
        result = (result * base) % mod;
    }
    return result;
}

vector<int> rabinKarpPatternMatching(string text, string pattern) {
    vector<int> positions;
    int n = text.length();
    int m = pattern.length();
    
    if (m == 0 || m > n) return positions;
    
    // Calculate hash of pattern and first window of text
    long long patternHash = calculateHash(pattern, 0, m - 1);
    long long textHash = calculateHash(text, 0, m - 1);
    
    // Precompute BASE^(m-1) for rolling hash
    long long power = 1;
    for (int i = 0; i < m - 1; i++) {
        power = (power * BASE) % MOD;
    }
    
    // Slide the pattern over text one by one
    for (int i = 0; i <= n - m; i++) {
        // Check hash values
        if (patternHash == textHash) {
            // If hash matches, check characters one by one
            bool match = true;
            for (int j = 0; j < m; j++) {
                if (text[i + j] != pattern[j]) {
                    match = false;
                    break;
                }
            }
            if (match) {
                positions.push_back(i);
            }
        }
        
        // Calculate hash for next window
        if (i < n - m) {
            textHash = (BASE * (textHash - text[i] * power) + text[i + m]) % MOD;
            if (textHash < 0) {
                textHash += MOD;
            }
        }
    }
    
    return positions;
}

// ==================== 4. Z-ALGORITHM ====================

// Build Z-array: Z[i] = length of longest substring starting from i
// which is also a prefix of the string
vector<int> buildZArray(string str) {
    int n = str.length();
    vector<int> z(n, 0);
    
    int l = 0, r = 0;  // Left and right boundaries of Z-box
    
    for (int i = 1; i < n; i++) {
        if (i <= r) {
            z[i] = min(r - i + 1, z[i - l]);
        }
        
        while (i + z[i] < n && str[z[i]] == str[i + z[i]]) {
            z[i]++;
        }
        
        if (i + z[i] - 1 > r) {
            l = i;
            r = i + z[i] - 1;
        }
    }
    
    return z;
}

// Z-Algorithm for pattern matching
// Time: O(n + m), Space: O(n + m)
vector<int> zAlgorithmPatternMatching(string text, string pattern) {
    vector<int> positions;
    int n = text.length();
    int m = pattern.length();
    
    if (m == 0) return positions;
    
    // Create combined string: pattern + $ + text
    string combined = pattern + "$" + text;
    vector<int> z = buildZArray(combined);
    
    // Find positions where Z[i] == pattern length
    for (int i = m + 1; i < combined.length(); i++) {
        if (z[i] == m) {
            positions.push_back(i - m - 1);
        }
    }
    
    return positions;
}

// ==================== 5. BOYER-MOORE ALGORITHM (SIMPLIFIED) ====================

// Bad character heuristic
vector<int> buildBadCharTable(string pattern) {
    vector<int> badChar(256, -1);
    for (int i = 0; i < pattern.length(); i++) {
        badChar[pattern[i]] = i;
    }
    return badChar;
}

// Simplified Boyer-Moore (using only bad character rule)
// Time: O(n*m) worst case, O(n/m) best case, Space: O(1)
vector<int> boyerMoorePatternMatching(string text, string pattern) {
    vector<int> positions;
    int n = text.length();
    int m = pattern.length();
    
    if (m == 0) return positions;
    
    vector<int> badChar = buildBadCharTable(pattern);
    
    int shift = 0;
    while (shift <= n - m) {
        int j = m - 1;
        
        // Match pattern from right to left
        while (j >= 0 && pattern[j] == text[shift + j]) {
            j--;
        }
        
        if (j < 0) {
            // Pattern found
            positions.push_back(shift);
            // Shift by at least 1
            shift += (shift + m < n) ? m - badChar[text[shift + m]] : 1;
        } else {
            // Shift based on bad character
            shift += max(1, j - badChar[text[shift + j]]);
        }
    }
    
    return positions;
}

// ==================== 6. MULTIPLE PATTERN MATCHING ====================

// Find all patterns in text using KMP for each
vector<vector<int>> findMultiplePatterns(string text, vector<string> patterns) {
    vector<vector<int>> results;
    
    for (string pattern : patterns) {
        results.push_back(kmpPatternMatching(text, pattern));
    }
    
    return results;
}

// ==================== MAIN FUNCTION (TESTING) ====================

int main() {
    cout << "=== Pattern Matching Algorithms ===\n\n";
    
    string text = "ABABDABACDABABCABCABAB";
    string pattern = "ABABCABAB";
    
    cout << "Text: " << text << "\n";
    cout << "Pattern: " << pattern << "\n\n";
    
    // Naive
    vector<int> naive = naivePatternMatching(text, pattern);
    cout << "Naive Algorithm: ";
    for (int pos : naive) cout << pos << " ";
    cout << "\n";
    
    // KMP
    vector<int> kmp = kmpPatternMatching(text, pattern);
    cout << "KMP Algorithm: ";
    for (int pos : kmp) cout << pos << " ";
    cout << "\n";
    
    // Rabin-Karp
    vector<int> rk = rabinKarpPatternMatching(text, pattern);
    cout << "Rabin-Karp Algorithm: ";
    for (int pos : rk) cout << pos << " ";
    cout << "\n";
    
    // Z-Algorithm
    vector<int> z = zAlgorithmPatternMatching(text, pattern);
    cout << "Z-Algorithm: ";
    for (int pos : z) cout << pos << " ";
    cout << "\n";
    
    // Boyer-Moore
    vector<int> bm = boyerMoorePatternMatching(text, pattern);
    cout << "Boyer-Moore Algorithm: ";
    for (int pos : bm) cout << pos << " ";
    cout << "\n";
    
    return 0;
}
