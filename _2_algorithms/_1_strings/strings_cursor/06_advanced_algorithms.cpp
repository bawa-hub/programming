/*
 * Advanced String Algorithms
 * Trie, Suffix Array, Manacher's Algorithm, etc.
 */

#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <algorithm>
#include <climits>

using namespace std;

// ==================== 1. TRIE (PREFIX TREE) ====================

class TrieNode {
public:
    unordered_map<char, TrieNode*> children;
    bool isEndOfWord;
    
    TrieNode() {
        isEndOfWord = false;
    }
};

class Trie {
private:
    TrieNode* root;
    
public:
    Trie() {
        root = new TrieNode();
    }
    
    // Insert a word into the trie
    void insert(string word) {
        TrieNode* current = root;
        
        for (char c : word) {
            if (current->children.find(c) == current->children.end()) {
                current->children[c] = new TrieNode();
            }
            current = current->children[c];
        }
        
        current->isEndOfWord = true;
    }
    
    // Search for a word in the trie
    bool search(string word) {
        TrieNode* current = root;
        
        for (char c : word) {
            if (current->children.find(c) == current->children.end()) {
                return false;
            }
            current = current->children[c];
        }
        
        return current->isEndOfWord;
    }
    
    // Check if any word starts with the given prefix
    bool startsWith(string prefix) {
        TrieNode* current = root;
        
        for (char c : prefix) {
            if (current->children.find(c) == current->children.end()) {
                return false;
            }
            current = current->children[c];
        }
        
        return true;
    }
    
    // Delete a word from the trie
    bool deleteWord(string word) {
        return deleteHelper(root, word, 0);
    }
    
private:
    bool deleteHelper(TrieNode* node, string word, int index) {
        if (index == word.length()) {
            if (!node->isEndOfWord) {
                return false;
            }
            node->isEndOfWord = false;
            return node->children.empty();
        }
        
        char c = word[index];
        if (node->children.find(c) == node->children.end()) {
            return false;
        }
        
        TrieNode* child = node->children[c];
        bool shouldDelete = deleteHelper(child, word, index + 1);
        
        if (shouldDelete) {
            node->children.erase(c);
            delete child;
            return node->children.empty() && !node->isEndOfWord;
        }
        
        return false;
    }
};

// ==================== 2. MANACHER'S ALGORITHM ====================

// Find longest palindromic substring
// Time: O(n), Space: O(n)
string longestPalindromicSubstring(string s) {
    if (s.empty()) return "";
    
    // Transform string: "abc" -> "^#a#b#c#$"
    string transformed = "^";
    for (char c : s) {
        transformed += "#";
        transformed += c;
    }
    transformed += "#$";
    
    int n = transformed.length();
    vector<int> P(n, 0);
    int center = 0, right = 0;
    int maxLen = 0, centerIndex = 0;
    
    for (int i = 1; i < n - 1; i++) {
        int mirror = 2 * center - i;
        
        if (i < right) {
            P[i] = min(right - i, P[mirror]);
        }
        
        // Expand around center i
        while (transformed[i + (1 + P[i])] == transformed[i - (1 + P[i])]) {
            P[i]++;
        }
        
        // Update center and right boundary
        if (i + P[i] > right) {
            center = i;
            right = i + P[i];
        }
        
        // Update longest palindrome
        if (P[i] > maxLen) {
            maxLen = P[i];
            centerIndex = i;
        }
    }
    
    int start = (centerIndex - maxLen) / 2;
    return s.substr(start, maxLen);
}

// Count palindromic substrings
int countPalindromicSubstrings(string s) {
    if (s.empty()) return 0;
    
    string transformed = "^";
    for (char c : s) {
        transformed += "#";
        transformed += c;
    }
    transformed += "#$";
    
    int n = transformed.length();
    vector<int> P(n, 0);
    int center = 0, right = 0;
    int count = 0;
    
    for (int i = 1; i < n - 1; i++) {
        int mirror = 2 * center - i;
        
        if (i < right) {
            P[i] = min(right - i, P[mirror]);
        }
        
        while (transformed[i + (1 + P[i])] == transformed[i - (1 + P[i])]) {
            P[i]++;
        }
        
        if (i + P[i] > right) {
            center = i;
            right = i + P[i];
        }
        
        // Count palindromes (each palindrome centered at i)
        count += (P[i] + 1) / 2;
    }
    
    return count;
}

// ==================== 3. SUFFIX ARRAY (BASIC) ====================

// Build suffix array - returns sorted indices of suffixes
vector<int> buildSuffixArray(string s) {
    int n = s.length();
    vector<pair<string, int>> suffixes;
    
    for (int i = 0; i < n; i++) {
        suffixes.push_back({s.substr(i), i});
    }
    
    sort(suffixes.begin(), suffixes.end());
    
    vector<int> suffixArray;
    for (auto& p : suffixes) {
        suffixArray.push_back(p.second);
    }
    
    return suffixArray;
}

// Find longest common prefix of two strings
int longestCommonPrefix(string s1, string s2) {
    int i = 0;
    while (i < s1.length() && i < s2.length() && s1[i] == s2[i]) {
        i++;
    }
    return i;
}

// Find longest repeated substring using suffix array
string longestRepeatedSubstring(string s) {
    vector<int> suffixArray = buildSuffixArray(s);
    string longest = "";
    
    for (int i = 0; i < suffixArray.size() - 1; i++) {
        string suffix1 = s.substr(suffixArray[i]);
        string suffix2 = s.substr(suffixArray[i + 1]);
        int lcp = longestCommonPrefix(suffix1, suffix2);
        
        if (lcp > longest.length()) {
            longest = suffix1.substr(0, lcp);
        }
    }
    
    return longest;
}

// ==================== 4. LONGEST COMMON SUBSEQUENCE (LCS) ====================

// Using Dynamic Programming
// Time: O(m*n), Space: O(m*n)
int longestCommonSubsequence(string text1, string text2) {
    int m = text1.length();
    int n = text2.length();
    
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));
    
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            if (text1[i - 1] == text2[j - 1]) {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]);
            }
        }
    }
    
    return dp[m][n];
}

// Get actual LCS string
string getLCSString(string text1, string text2) {
    int m = text1.length();
    int n = text2.length();
    
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));
    
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            if (text1[i - 1] == text2[j - 1]) {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]);
            }
        }
    }
    
    // Reconstruct LCS
    string lcs = "";
    int i = m, j = n;
    
    while (i > 0 && j > 0) {
        if (text1[i - 1] == text2[j - 1]) {
            lcs = text1[i - 1] + lcs;
            i--;
            j--;
        } else if (dp[i - 1][j] > dp[i][j - 1]) {
            i--;
        } else {
            j--;
        }
    }
    
    return lcs;
}

// ==================== 5. EDIT DISTANCE (LEVENSHTEIN DISTANCE) ====================

// LeetCode 72: Edit Distance
int minDistance(string word1, string word2) {
    int m = word1.length();
    int n = word2.length();
    
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));
    
    // Base cases
    for (int i = 0; i <= m; i++) {
        dp[i][0] = i;
    }
    for (int j = 0; j <= n; j++) {
        dp[0][j] = j;
    }
    
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            if (word1[i - 1] == word2[j - 1]) {
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                dp[i][j] = 1 + min({
                    dp[i - 1][j],      // Delete
                    dp[i][j - 1],      // Insert
                    dp[i - 1][j - 1]   // Replace
                });
            }
        }
    }
    
    return dp[m][n];
}

// ==================== 6. WORD BREAK PROBLEM ====================

// LeetCode 139: Word Break
bool wordBreak(string s, vector<string>& wordDict) {
    unordered_set<string> wordSet(wordDict.begin(), wordDict.end());
    int n = s.length();
    vector<bool> dp(n + 1, false);
    dp[0] = true;  // Empty string
    
    for (int i = 1; i <= n; i++) {
        for (int j = 0; j < i; j++) {
            if (dp[j] && wordSet.find(s.substr(j, i - j)) != wordSet.end()) {
                dp[i] = true;
                break;
            }
        }
    }
    
    return dp[n];
}

// ==================== 7. LONGEST INCREASING SUBSEQUENCE (STRING ANALOGY) ====================

// Find longest increasing subsequence in string (lexicographically)
int longestIncreasingSubsequence(string s) {
    int n = s.length();
    vector<int> dp(n, 1);
    
    for (int i = 1; i < n; i++) {
        for (int j = 0; j < i; j++) {
            if (s[j] < s[i]) {
                dp[i] = max(dp[i], dp[j] + 1);
            }
        }
    }
    
    return *max_element(dp.begin(), dp.end());
}

// ==================== 8. RABIN-KARP FOR MULTIPLE PATTERNS ====================

// Find all patterns in text using rolling hash
vector<vector<int>> findMultiplePatternsRabinKarp(string text, vector<string>& patterns) {
    vector<vector<int>> results;
    const int BASE = 256;
    const int MOD = 101;
    
    for (string pattern : patterns) {
        vector<int> positions;
        int n = text.length();
        int m = pattern.length();
        
        if (m == 0 || m > n) {
            results.push_back(positions);
            continue;
        }
        
        long long patternHash = 0;
        long long textHash = 0;
        long long power = 1;
        
        for (int i = 0; i < m; i++) {
            patternHash = (patternHash * BASE + pattern[i]) % MOD;
            textHash = (textHash * BASE + text[i]) % MOD;
            if (i < m - 1) {
                power = (power * BASE) % MOD;
            }
        }
        
        for (int i = 0; i <= n - m; i++) {
            if (patternHash == textHash) {
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
            
            if (i < n - m) {
                textHash = (BASE * (textHash - text[i] * power) + text[i + m]) % MOD;
                if (textHash < 0) textHash += MOD;
            }
        }
        
        results.push_back(positions);
    }
    
    return results;
}

// ==================== MAIN FUNCTION (TESTING) ====================

int main() {
    cout << "=== Advanced String Algorithms ===\n\n";
    
    // Test Trie
    Trie trie;
    trie.insert("apple");
    trie.insert("app");
    trie.insert("application");
    
    cout << "Trie search 'app': " << (trie.search("app") ? "Found" : "Not found") << "\n";
    cout << "Trie startsWith 'app': " << (trie.startsWith("app") ? "Yes" : "No") << "\n";
    
    // Test Manacher's
    string s1 = "babad";
    cout << "\nLongest palindromic substring in '" << s1 << "': " 
         << longestPalindromicSubstring(s1) << "\n";
    cout << "Number of palindromic substrings: " << countPalindromicSubstrings(s1) << "\n";
    
    // Test LCS
    string s2 = "abcde";
    string s3 = "ace";
    cout << "\nLCS of '" << s2 << "' and '" << s3 << "': " 
         << longestCommonSubsequence(s2, s3) << "\n";
    cout << "LCS string: " << getLCSString(s2, s3) << "\n";
    
    // Test Edit Distance
    string s4 = "horse";
    string s5 = "ros";
    cout << "\nEdit distance between '" << s4 << "' and '" << s5 << "': " 
         << minDistance(s4, s5) << "\n";
    
    return 0;
}
