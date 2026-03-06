/*
 * Hash-Based Techniques for Strings
 * Using hash maps and sets for efficient string operations
 */

#include <iostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <algorithm>
#include <map>

using namespace std;

// ==================== 1. CHARACTER FREQUENCY MAP ====================

unordered_map<char, int> getCharFrequency(string s) {
    unordered_map<char, int> freq;
    for (char c : s) {
        freq[c]++;
    }
    return freq;
}

// Find first unique character
// LeetCode 387: First Unique Character in a String
int firstUniqChar(string s) {
    unordered_map<char, int> freq;
    
    // Count frequency
    for (char c : s) {
        freq[c]++;
    }
    
    // Find first unique
    for (int i = 0; i < s.length(); i++) {
        if (freq[s[i]] == 1) {
            return i;
        }
    }
    
    return -1;
}

// ==================== 2. GROUP ANAGRAMS ====================

// LeetCode 49: Group Anagrams
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

// Group anagrams without sorting (using frequency map as key)
vector<vector<string>> groupAnagramsOptimized(vector<string>& strs) {
    map<vector<int>, vector<string>> groups;
    
    for (string str : strs) {
        vector<int> freq(26, 0);
        for (char c : str) {
            freq[c - 'a']++;
        }
        groups[freq].push_back(str);
    }
    
    vector<vector<string>> result;
    for (auto& pair : groups) {
        result.push_back(pair.second);
    }
    
    return result;
}

// ==================== 3. LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS ====================

// LeetCode 3: Using hash map
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

// ==================== 4. VALID ANAGRAM ====================

// LeetCode 242: Valid Anagram
bool isAnagram(string s, string t) {
    if (s.length() != t.length()) return false;
    
    unordered_map<char, int> freq;
    
    for (char c : s) {
        freq[c]++;
    }
    
    for (char c : t) {
        freq[c]--;
        if (freq[c] < 0) return false;
    }
    
    return true;
}

// ==================== 5. FIND ALL ANAGRAMS IN A STRING ====================

// LeetCode 438: Find All Anagrams in a String
vector<int> findAnagrams(string s, string p) {
    vector<int> result;
    if (s.length() < p.length()) return result;
    
    vector<int> pFreq(26, 0);
    vector<int> windowFreq(26, 0);
    
    // Initialize frequency arrays
    for (int i = 0; i < p.length(); i++) {
        pFreq[p[i] - 'a']++;
        windowFreq[s[i] - 'a']++;
    }
    
    // Check first window
    if (pFreq == windowFreq) {
        result.push_back(0);
    }
    
    // Slide window
    for (int i = p.length(); i < s.length(); i++) {
        windowFreq[s[i - p.length()] - 'a']--;
        windowFreq[s[i] - 'a']++;
        
        if (pFreq == windowFreq) {
            result.push_back(i - p.length() + 1);
        }
    }
    
    return result;
}

// ==================== 6. WORD PATTERN ====================

// LeetCode 290: Word Pattern
bool wordPattern(string pattern, string s) {
    vector<string> words;
    string word = "";
    
    // Split string into words
    for (char c : s) {
        if (c == ' ') {
            if (!word.empty()) {
                words.push_back(word);
                word = "";
            }
        } else {
            word += c;
        }
    }
    if (!word.empty()) {
        words.push_back(word);
    }
    
    if (pattern.length() != words.size()) return false;
    
    unordered_map<char, string> charToWord;
    unordered_map<string, char> wordToChar;
    
    for (int i = 0; i < pattern.length(); i++) {
        char c = pattern[i];
        string w = words[i];
        
        if (charToWord.find(c) != charToWord.end()) {
            if (charToWord[c] != w) return false;
        } else {
            charToWord[c] = w;
        }
        
        if (wordToChar.find(w) != wordToChar.end()) {
            if (wordToChar[w] != c) return false;
        } else {
            wordToChar[w] = c;
        }
    }
    
    return true;
}

// ==================== 7. ISOMORPHIC STRINGS ====================

// LeetCode 205: Isomorphic Strings
bool isIsomorphic(string s, string t) {
    if (s.length() != t.length()) return false;
    
    unordered_map<char, char> sToT;
    unordered_map<char, char> tToS;
    
    for (int i = 0; i < s.length(); i++) {
        char sChar = s[i];
        char tChar = t[i];
        
        if (sToT.find(sChar) != sToT.end()) {
            if (sToT[sChar] != tChar) return false;
        } else {
            sToT[sChar] = tChar;
        }
        
        if (tToS.find(tChar) != tToS.end()) {
            if (tToS[tChar] != sChar) return false;
        } else {
            tToS[tChar] = sChar;
        }
    }
    
    return true;
}

// ==================== 8. LONGEST PALINDROME ====================

// LeetCode 409: Longest Palindrome
int longestPalindrome(string s) {
    unordered_map<char, int> freq;
    
    for (char c : s) {
        freq[c]++;
    }
    
    int length = 0;
    bool hasOdd = false;
    
    for (auto& pair : freq) {
        if (pair.second % 2 == 0) {
            length += pair.second;
        } else {
            length += pair.second - 1;
            hasOdd = true;
        }
    }
    
    return length + (hasOdd ? 1 : 0);
}

// ==================== 9. RANSOM NOTE ====================

// LeetCode 383: Ransom Note
bool canConstruct(string ransomNote, string magazine) {
    unordered_map<char, int> magazineFreq;
    
    for (char c : magazine) {
        magazineFreq[c]++;
    }
    
    for (char c : ransomNote) {
        if (magazineFreq.find(c) == magazineFreq.end() || 
            magazineFreq[c] == 0) {
            return false;
        }
        magazineFreq[c]--;
    }
    
    return true;
}

// ==================== 10. UNIQUE EMAIL ADDRESSES ====================

// LeetCode 929: Unique Email Addresses
int numUniqueEmails(vector<string>& emails) {
    unordered_set<string> uniqueEmails;
    
    for (string email : emails) {
        string processed = "";
        bool foundAt = false;
        bool foundPlus = false;
        
        for (char c : email) {
            if (c == '@') {
                foundAt = true;
                foundPlus = false;
                processed += c;
            } else if (foundAt) {
                processed += c;
            } else if (c == '+') {
                foundPlus = true;
            } else if (c == '.' && !foundPlus) {
                continue;
            } else if (!foundPlus) {
                processed += c;
            }
        }
        
        uniqueEmails.insert(processed);
    }
    
    return uniqueEmails.size();
}

// ==================== 11. JEWELS AND STONES ====================

// LeetCode 771: Jewels and Stones
int numJewelsInStones(string jewels, string stones) {
    unordered_set<char> jewelSet;
    
    for (char c : jewels) {
        jewelSet.insert(c);
    }
    
    int count = 0;
    for (char c : stones) {
        if (jewelSet.find(c) != jewelSet.end()) {
            count++;
        }
    }
    
    return count;
}

// ==================== 12. MINIMUM INDEX SUM OF TWO LISTS ====================

// LeetCode 599: Minimum Index Sum of Two Lists
vector<string> findRestaurant(vector<string>& list1, vector<string>& list2) {
    unordered_map<string, int> indexMap;
    
    for (int i = 0; i < list1.size(); i++) {
        indexMap[list1[i]] = i;
    }
    
    int minSum = INT_MAX;
    vector<string> result;
    
    for (int i = 0; i < list2.size(); i++) {
        if (indexMap.find(list2[i]) != indexMap.end()) {
            int sum = i + indexMap[list2[i]];
            if (sum < minSum) {
                minSum = sum;
                result.clear();
                result.push_back(list2[i]);
            } else if (sum == minSum) {
                result.push_back(list2[i]);
            }
        }
    }
    
    return result;
}

// ==================== 13. UNCOMMON WORDS FROM TWO SENTENCES ====================

// LeetCode 884: Uncommon Words from Two Sentences
vector<string> uncommonFromSentences(string s1, string s2) {
    unordered_map<string, int> freq;
    
    // Split and count words
    string word = "";
    for (char c : s1) {
        if (c == ' ') {
            if (!word.empty()) {
                freq[word]++;
                word = "";
            }
        } else {
            word += c;
        }
    }
    if (!word.empty()) {
        freq[word]++;
        word = "";
    }
    
    for (char c : s2) {
        if (c == ' ') {
            if (!word.empty()) {
                freq[word]++;
                word = "";
            }
        } else {
            word += c;
        }
    }
    if (!word.empty()) {
        freq[word]++;
    }
    
    vector<string> result;
    for (auto& pair : freq) {
        if (pair.second == 1) {
            result.push_back(pair.first);
        }
    }
    
    return result;
}

// ==================== MAIN FUNCTION (TESTING) ====================

int main() {
    cout << "=== Hash-Based Techniques ===\n\n";
    
    // Test first unique character
    string s1 = "leetcode";
    cout << "First unique char in '" << s1 << "' at index: " 
         << firstUniqChar(s1) << "\n";
    
    // Test group anagrams
    vector<string> strs = {"eat", "tea", "tan", "ate", "nat", "bat"};
    vector<vector<string>> groups = groupAnagrams(strs);
    cout << "\nGroup Anagrams:\n";
    for (auto& group : groups) {
        for (string s : group) {
            cout << s << " ";
        }
        cout << "\n";
    }
    
    // Test longest substring
    string s2 = "abcabcbb";
    cout << "\nLongest substring without repeating in '" << s2 << "': " 
         << lengthOfLongestSubstring(s2) << "\n";
    
    return 0;
}
