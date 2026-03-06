/*
 * Sliding Window Patterns for Strings
 * Fixed and variable size window problems
 */

#include <iostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <climits>
#include <algorithm>

using namespace std;

// ==================== 1. FIXED SIZE SLIDING WINDOW ====================

// Maximum sum of subarray of size k
int maxSumSubarray(vector<int>& arr, int k) {
    int n = arr.size();
    if (n < k) return -1;
    
    int windowSum = 0;
    for (int i = 0; i < k; i++) {
        windowSum += arr[i];
    }
    
    int maxSum = windowSum;
    for (int i = k; i < n; i++) {
        windowSum = windowSum - arr[i - k] + arr[i];
        maxSum = max(maxSum, windowSum);
    }
    
    return maxSum;
}

// Maximum average subarray of size k
double maxAverageSubarray(vector<int>& arr, int k) {
    int n = arr.size();
    int windowSum = 0;
    
    for (int i = 0; i < k; i++) {
        windowSum += arr[i];
    }
    
    int maxSum = windowSum;
    for (int i = k; i < n; i++) {
        windowSum = windowSum - arr[i - k] + arr[i];
        maxSum = max(maxSum, windowSum);
    }
    
    return (double)maxSum / k;
}

// ==================== 2. VARIABLE SIZE SLIDING WINDOW ====================

// Longest substring with at most K distinct characters
int longestSubstringKDistinct(string s, int k) {
    int n = s.length();
    if (n == 0 || k == 0) return 0;
    
    unordered_map<char, int> freq;
    int left = 0, maxLen = 0;
    
    for (int right = 0; right < n; right++) {
        freq[s[right]]++;
        
        // Shrink window if distinct chars exceed k
        while (freq.size() > k) {
            freq[s[left]]--;
            if (freq[s[left]] == 0) {
                freq.erase(s[left]);
            }
            left++;
        }
        
        maxLen = max(maxLen, right - left + 1);
    }
    
    return maxLen;
}

// Longest substring with exactly K distinct characters
int longestSubstringExactlyKDistinct(string s, int k) {
    int n = s.length();
    if (n == 0 || k == 0) return 0;
    
    unordered_map<char, int> freq;
    int left = 0, maxLen = 0;
    
    for (int right = 0; right < n; right++) {
        freq[s[right]]++;
        
        // Shrink window if distinct chars exceed k
        while (freq.size() > k) {
            freq[s[left]]--;
            if (freq[s[left]] == 0) {
                freq.erase(s[left]);
            }
            left++;
        }
        
        // Check if we have exactly k distinct characters
        if (freq.size() == k) {
            maxLen = max(maxLen, right - left + 1);
        }
    }
    
    return maxLen;
}

// ==================== 3. LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS ====================

// LeetCode 3: Longest Substring Without Repeating Characters
int lengthOfLongestSubstring(string s) {
    int n = s.length();
    if (n == 0) return 0;
    
    unordered_map<char, int> lastSeen;
    int left = 0, maxLen = 0;
    
    for (int right = 0; right < n; right++) {
        // If character seen before and within current window
        if (lastSeen.find(s[right]) != lastSeen.end() && 
            lastSeen[s[right]] >= left) {
            left = lastSeen[s[right]] + 1;
        }
        
        lastSeen[s[right]] = right;
        maxLen = max(maxLen, right - left + 1);
    }
    
    return maxLen;
}

// ==================== 4. MINIMUM WINDOW SUBSTRING ====================

// LeetCode 76: Minimum Window Substring
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
        
        // Check if current character completes a requirement
        if (required.find(c) != required.end() && 
            window[c] == required[c]) {
            formed++;
        }
        
        // Try to contract window
        while (left <= right && formed == requiredCount) {
            // Update minimum window
            if (right - left + 1 < minLen) {
                minLen = right - left + 1;
                minLeft = left;
            }
            
            // Remove leftmost character
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

// ==================== 5. SUBSTRING WITH ALL WORDS ====================

// LeetCode 30: Substring with Concatenation of All Words
vector<int> findSubstring(string s, vector<string>& words) {
    vector<int> result;
    if (s.empty() || words.empty()) return result;
    
    int wordLen = words[0].length();
    int totalWords = words.size();
    int totalLen = wordLen * totalWords;
    
    if (s.length() < totalLen) return result;
    
    unordered_map<string, int> wordCount;
    for (string word : words) {
        wordCount[word]++;
    }
    
    for (int i = 0; i <= (int)s.length() - totalLen; i++) {
        unordered_map<string, int> seen;
        int j = 0;
        
        while (j < totalWords) {
            string word = s.substr(i + j * wordLen, wordLen);
            
            if (wordCount.find(word) == wordCount.end()) {
                break;
            }
            
            seen[word]++;
            if (seen[word] > wordCount[word]) {
                break;
            }
            
            j++;
        }
        
        if (j == totalWords) {
            result.push_back(i);
        }
    }
    
    return result;
}

// ==================== 6. LONGEST SUBSTRING WITH AT LEAST K REPEATING CHARACTERS ====================

// LeetCode 395: Longest Substring with At Least K Repeating Characters
int longestSubstring(string s, int k) {
    int n = s.length();
    if (n == 0 || k > n) return 0;
    if (k == 0) return n;
    
    unordered_map<char, int> freq;
    for (char c : s) {
        freq[c]++;
    }
    
    // Find first character with frequency < k
    int idx = 0;
    while (idx < n && freq[s[idx]] >= k) {
        idx++;
    }
    
    if (idx == n) return n;
    
    // Split at this character and recurse
    int left = longestSubstring(s.substr(0, idx), k);
    int right = longestSubstring(s.substr(idx + 1), k);
    
    return max(left, right);
}

// ==================== 7. REPEATED DNA SEQUENCES ====================

// LeetCode 187: Repeated DNA Sequences
vector<string> findRepeatedDnaSequences(string s) {
    vector<string> result;
    if (s.length() < 10) return result;
    
    unordered_map<string, int> seen;
    
    for (int i = 0; i <= (int)s.length() - 10; i++) {
        string sequence = s.substr(i, 10);
        seen[sequence]++;
    }
    
    for (auto& pair : seen) {
        if (pair.second > 1) {
            result.push_back(pair.first);
        }
    }
    
    return result;
}

// ==================== 8. PERMUTATION IN STRING ====================

// LeetCode 567: Permutation in String
bool checkInclusion(string s1, string s2) {
    if (s1.length() > s2.length()) return false;
    
    vector<int> s1Count(26, 0);
    vector<int> s2Count(26, 0);
    
    // Initialize frequency arrays
    for (int i = 0; i < s1.length(); i++) {
        s1Count[s1[i] - 'a']++;
        s2Count[s2[i] - 'a']++;
    }
    
    // Check first window
    int matches = 0;
    for (int i = 0; i < 26; i++) {
        if (s1Count[i] == s2Count[i]) {
            matches++;
        }
    }
    
    // Slide window
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

// ==================== 9. MAXIMUM NUMBER OF VOWELS IN SUBSTRING ====================

int maxVowels(string s, int k) {
    unordered_set<char> vowels = {'a', 'e', 'i', 'o', 'u'};
    
    int vowelCount = 0;
    // Count vowels in first window
    for (int i = 0; i < k; i++) {
        if (vowels.count(s[i])) {
            vowelCount++;
        }
    }
    
    int maxVowel = vowelCount;
    
    // Slide window
    for (int i = k; i < s.length(); i++) {
        if (vowels.count(s[i - k])) {
            vowelCount--;
        }
        if (vowels.count(s[i])) {
            vowelCount++;
        }
        maxVowel = max(maxVowel, vowelCount);
    }
    
    return maxVowel;
}

// ==================== MAIN FUNCTION (TESTING) ====================

int main() {
    cout << "=== Sliding Window Patterns ===\n\n";
    
    // Test longest substring without repeating
    string s1 = "abcabcbb";
    cout << "Longest substring without repeating in '" << s1 << "': " 
         << lengthOfLongestSubstring(s1) << "\n";
    
    // Test minimum window
    string s2 = "ADOBECODEBANC";
    string t = "ABC";
    cout << "Minimum window substring: " << minWindow(s2, t) << "\n";
    
    // Test K distinct characters
    string s3 = "eceba";
    cout << "Longest substring with 2 distinct chars in '" << s3 << "': " 
         << longestSubstringKDistinct(s3, 2) << "\n";
    
    return 0;
}
