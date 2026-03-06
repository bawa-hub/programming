/*
 * Practice Problems - LeetCode Style
 * Solutions to common string problems from easy to hard
 */

#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <algorithm>
#include <stack>
#include <queue>
#include <climits>

using namespace std;

// ==================== EASY PROBLEMS ====================

// Problem 1: Valid Anagram (LeetCode 242)
bool isAnagram(string s, string t) {
    if (s.length() != t.length()) return false;
    
    vector<int> count(26, 0);
    for (int i = 0; i < s.length(); i++) {
        count[s[i] - 'a']++;
        count[t[i] - 'a']--;
    }
    
    for (int c : count) {
        if (c != 0) return false;
    }
    return true;
}

// Problem 2: First Unique Character (LeetCode 387)
int firstUniqChar(string s) {
    unordered_map<char, int> freq;
    for (char c : s) {
        freq[c]++;
    }
    
    for (int i = 0; i < s.length(); i++) {
        if (freq[s[i]] == 1) {
            return i;
        }
    }
    return -1;
}

// Problem 3: Reverse String (LeetCode 344)
void reverseString(vector<char>& s) {
    int left = 0, right = s.size() - 1;
    while (left < right) {
        swap(s[left], s[right]);
        left++;
        right--;
    }
}

// Problem 4: Valid Palindrome (LeetCode 125)
bool isPalindrome(string s) {
    int left = 0, right = s.length() - 1;
    
    while (left < right) {
        while (left < right && !isalnum(s[left])) left++;
        while (left < right && !isalnum(s[right])) right--;
        
        if (tolower(s[left]) != tolower(s[right])) {
            return false;
        }
        left++;
        right--;
    }
    return true;
}

// Problem 5: Reverse Words in String (LeetCode 151)
string reverseWords(string s) {
    // Remove extra spaces
    string cleaned = "";
    for (int i = 0; i < s.length(); i++) {
        if (s[i] != ' ') {
            cleaned += s[i];
        } else if (!cleaned.empty() && cleaned.back() != ' ') {
            cleaned += ' ';
        }
    }
    
    if (!cleaned.empty() && cleaned.back() == ' ') {
        cleaned.pop_back();
    }
    
    // Reverse entire string
    reverse(cleaned.begin(), cleaned.end());
    
    // Reverse each word
    int start = 0;
    for (int i = 0; i <= cleaned.length(); i++) {
        if (i == cleaned.length() || cleaned[i] == ' ') {
            reverse(cleaned.begin() + start, cleaned.begin() + i);
            start = i + 1;
        }
    }
    
    return cleaned;
}

// Problem 6: Isomorphic Strings (LeetCode 205)
bool isIsomorphic(string s, string t) {
    if (s.length() != t.length()) return false;
    
    unordered_map<char, char> sToT;
    unordered_map<char, char> tToS;
    
    for (int i = 0; i < s.length(); i++) {
        if (sToT.find(s[i]) != sToT.end()) {
            if (sToT[s[i]] != t[i]) return false;
        } else {
            sToT[s[i]] = t[i];
        }
        
        if (tToS.find(t[i]) != tToS.end()) {
            if (tToS[t[i]] != s[i]) return false;
        } else {
            tToS[t[i]] = s[i];
        }
    }
    
    return true;
}

// ==================== MEDIUM PROBLEMS ====================

// Problem 7: Longest Substring Without Repeating Characters (LeetCode 3)
int lengthOfLongestSubstring(string s) {
    unordered_map<char, int> lastSeen;
    int maxLen = 0;
    int start = 0;
    
    for (int end = 0; end < s.length(); end++) {
        if (lastSeen.find(s[end]) != lastSeen.end() && 
            lastSeen[s[end]] >= start) {
            start = lastSeen[s[end]] + 1;
        }
        
        lastSeen[s[end]] = end;
        maxLen = max(maxLen, end - start + 1);
    }
    
    return maxLen;
}

// Problem 8: Group Anagrams (LeetCode 49)
vector<vector<string>> groupAnagrams(vector<string>& strs) {
    unordered_map<string, vector<string>> groups;
    
    for (string str : strs) {
        string key = str;
        sort(key.begin(), key.end());
        groups[key].push_back(str);
    }
    
    vector<vector<string>> result;
    for (auto& pair : groups) {
        result.push_back(pair.second);
    }
    
    return result;
}

// Problem 9: Minimum Window Substring (LeetCode 76)
string minWindow(string s, string t) {
    if (s.length() < t.length()) return "";
    
    unordered_map<char, int> required;
    for (char c : t) {
        required[c]++;
    }
    
    int requiredCount = required.size();
    int left = 0, right = 0;
    int formed = 0;
    int minLen = INT_MAX;
    int minLeft = 0;
    
    unordered_map<char, int> window;
    
    while (right < s.length()) {
        char c = s[right];
        window[c]++;
        
        if (required.find(c) != required.end() && 
            window[c] == required[c]) {
            formed++;
        }
        
        while (left <= right && formed == requiredCount) {
            if (right - left + 1 < minLen) {
                minLen = right - left + 1;
                minLeft = left;
            }
            
            char leftChar = s[left];
            window[leftChar]--;
            
            if (required.find(leftChar) != required.end() && 
                window[leftChar] < required[leftChar]) {
                formed--;
            }
            
            left++;
        }
        
        right++;
    }
    
    return minLen == INT_MAX ? "" : s.substr(minLeft, minLen);
}

// Problem 10: Longest Palindromic Substring (LeetCode 5)
string longestPalindrome(string s) {
    if (s.empty()) return "";
    
    int n = s.length();
    int start = 0, maxLen = 1;
    
    // Expand around center
    for (int i = 0; i < n; i++) {
        // Odd length palindromes
        int left = i, right = i;
        while (left >= 0 && right < n && s[left] == s[right]) {
            if (right - left + 1 > maxLen) {
                start = left;
                maxLen = right - left + 1;
            }
            left--;
            right++;
        }
        
        // Even length palindromes
        left = i;
        right = i + 1;
        while (left >= 0 && right < n && s[left] == s[right]) {
            if (right - left + 1 > maxLen) {
                start = left;
                maxLen = right - left + 1;
            }
            left--;
            right++;
        }
    }
    
    return s.substr(start, maxLen);
}

// Problem 11: Find All Anagrams (LeetCode 438)
vector<int> findAnagrams(string s, string p) {
    vector<int> result;
    if (s.length() < p.length()) return result;
    
    vector<int> pFreq(26, 0);
    vector<int> windowFreq(26, 0);
    
    for (int i = 0; i < p.length(); i++) {
        pFreq[p[i] - 'a']++;
        windowFreq[s[i] - 'a']++;
    }
    
    if (pFreq == windowFreq) {
        result.push_back(0);
    }
    
    for (int i = p.length(); i < s.length(); i++) {
        windowFreq[s[i - p.length()] - 'a']--;
        windowFreq[s[i] - 'a']++;
        
        if (pFreq == windowFreq) {
            result.push_back(i - p.length() + 1);
        }
    }
    
    return result;
}

// Problem 12: Permutation in String (LeetCode 567)
bool checkInclusion(string s1, string s2) {
    if (s1.length() > s2.length()) return false;
    
    vector<int> s1Count(26, 0);
    vector<int> s2Count(26, 0);
    
    for (int i = 0; i < s1.length(); i++) {
        s1Count[s1[i] - 'a']++;
        s2Count[s2[i] - 'a']++;
    }
    
    int matches = 0;
    for (int i = 0; i < 26; i++) {
        if (s1Count[i] == s2Count[i]) {
            matches++;
        }
    }
    
    for (int i = s1.length(); i < s2.length(); i++) {
        if (matches == 26) return true;
        
        int right = s2[i] - 'a';
        int left = s2[i - s1.length()] - 'a';
        
        s2Count[right]++;
        if (s1Count[right] == s2Count[right]) {
            matches++;
        } else if (s1Count[right] + 1 == s2Count[right]) {
            matches--;
        }
        
        s2Count[left]--;
        if (s1Count[left] == s2Count[left]) {
            matches++;
        } else if (s1Count[left] - 1 == s2Count[left]) {
            matches--;
        }
    }
    
    return matches == 26;
}

// Problem 13: Longest Repeating Character Replacement (LeetCode 424)
int characterReplacement(string s, int k) {
    vector<int> count(26, 0);
    int maxCount = 0;
    int maxLen = 0;
    int left = 0;
    
    for (int right = 0; right < s.length(); right++) {
        count[s[right] - 'A']++;
        maxCount = max(maxCount, count[s[right] - 'A']);
        
        if (right - left + 1 - maxCount > k) {
            count[s[left] - 'A']--;
            left++;
        }
        
        maxLen = max(maxLen, right - left + 1);
    }
    
    return maxLen;
}

// ==================== HARD PROBLEMS ====================

// Problem 14: Edit Distance (LeetCode 72)
int minDistance(string word1, string word2) {
    int m = word1.length();
    int n = word2.length();
    
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));
    
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
                    dp[i - 1][j],
                    dp[i][j - 1],
                    dp[i - 1][j - 1]
                });
            }
        }
    }
    
    return dp[m][n];
}

// Problem 15: Word Break II (LeetCode 140)
vector<string> wordBreak(string s, vector<string>& wordDict) {
    unordered_set<string> wordSet(wordDict.begin(), wordDict.end());
    unordered_map<int, vector<string>> memo;
    
    return wordBreakHelper(s, wordSet, 0, memo);
}

vector<string> wordBreakHelper(string& s, unordered_set<string>& wordSet, 
                               int start, unordered_map<int, vector<string>>& memo) {
    if (memo.find(start) != memo.end()) {
        return memo[start];
    }
    
    vector<string> result;
    
    if (start == s.length()) {
        result.push_back("");
        return result;
    }
    
    for (int end = start + 1; end <= s.length(); end++) {
        string word = s.substr(start, end - start);
        if (wordSet.find(word) != wordSet.end()) {
            vector<string> suffixes = wordBreakHelper(s, wordSet, end, memo);
            for (string suffix : suffixes) {
                result.push_back(suffix.empty() ? word : word + " " + suffix);
            }
        }
    }
    
    memo[start] = result;
    return result;
}

// Problem 16: Scramble String (LeetCode 87)
bool isScramble(string s1, string s2) {
    if (s1 == s2) return true;
    if (s1.length() != s2.length()) return false;
    
    int n = s1.length();
    vector<int> count(26, 0);
    
    for (int i = 0; i < n; i++) {
        count[s1[i] - 'a']++;
        count[s2[i] - 'a']--;
    }
    
    for (int i = 0; i < 26; i++) {
        if (count[i] != 0) return false;
    }
    
    for (int i = 1; i < n; i++) {
        if ((isScramble(s1.substr(0, i), s2.substr(0, i)) && 
             isScramble(s1.substr(i), s2.substr(i))) ||
            (isScramble(s1.substr(0, i), s2.substr(n - i)) && 
             isScramble(s1.substr(i), s2.substr(0, n - i)))) {
            return true;
        }
    }
    
    return false;
}

// Problem 17: Minimum Window Subsequence (LeetCode 727)
string minWindowSubsequence(string s, string t) {
    int m = s.length();
    int n = t.length();
    
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, -1));
    
    for (int i = 0; i <= m; i++) {
        dp[i][0] = i;
    }
    
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            if (s[i - 1] == t[j - 1]) {
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                dp[i][j] = dp[i - 1][j];
            }
        }
    }
    
    int minLen = INT_MAX;
    int start = -1;
    
    for (int i = 1; i <= m; i++) {
        if (dp[i][n] != -1) {
            int len = i - dp[i][n];
            if (len < minLen) {
                minLen = len;
                start = dp[i][n];
            }
        }
    }
    
    return start == -1 ? "" : s.substr(start, minLen);
}

// ==================== MAIN FUNCTION (TESTING) ====================

int main() {
    cout << "=== Practice Problems ===\n\n";
    
    // Test easy problems
    cout << "Easy Problems:\n";
    cout << "Is 'anagram' and 'nagaram' anagrams? " 
         << (isAnagram("anagram", "nagaram") ? "Yes" : "No") << "\n";
    
    cout << "First unique char in 'leetcode': " 
         << firstUniqChar("leetcode") << "\n";
    
    // Test medium problems
    cout << "\nMedium Problems:\n";
    cout << "Longest substring without repeating in 'abcabcbb': " 
         << lengthOfLongestSubstring("abcabcbb") << "\n";
    
    string s = "ADOBECODEBANC";
    string t = "ABC";
    cout << "Minimum window substring: " << minWindow(s, t) << "\n";
    
    return 0;
}
