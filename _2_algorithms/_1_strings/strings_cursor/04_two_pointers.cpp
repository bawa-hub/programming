/*
 * Two Pointers Technique for Strings
 * Efficient solutions using two pointers approach
 */

#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <cctype>
#include <unordered_set>

using namespace std;

// ==================== 1. VALID PALINDROME ====================

// LeetCode 125: Valid Palindrome
bool isPalindrome(string s) {
    int left = 0, right = s.length() - 1;
    
    while (left < right) {
        // Skip non-alphanumeric characters
        while (left < right && !isalnum(s[left])) {
            left++;
        }
        while (left < right && !isalnum(s[right])) {
            right--;
        }
        
        if (tolower(s[left]) != tolower(s[right])) {
            return false;
        }
        
        left++;
        right--;
    }
    
    return true;
}

// ==================== 2. REVERSE WORDS IN STRING ====================

// LeetCode 151: Reverse Words in a String
string reverseWords(string s) {
    // Remove extra spaces
    string cleaned = "";
    int n = s.length();
    
    for (int i = 0; i < n; i++) {
        if (s[i] != ' ') {
            cleaned += s[i];
        } else if (!cleaned.empty() && cleaned.back() != ' ') {
            cleaned += ' ';
        }
    }
    
    // Remove trailing space
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

// ==================== 3. REMOVE DUPLICATES FROM SORTED ARRAY (STRING VERSION) ====================

string removeDuplicates(string s) {
    if (s.length() <= 1) return s;
    
    int writePos = 0;
    for (int readPos = 1; readPos < s.length(); readPos++) {
        if (s[readPos] != s[writePos]) {
            writePos++;
            s[writePos] = s[readPos];
        }
    }
    
    return s.substr(0, writePos + 1);
}

// Remove duplicates allowing at most 2 occurrences
string removeDuplicatesAtMostTwo(string s) {
    if (s.length() <= 2) return s;
    
    int writePos = 1;
    for (int readPos = 2; readPos < s.length(); readPos++) {
        if (s[readPos] != s[writePos] || s[readPos] != s[writePos - 1]) {
            writePos++;
            s[writePos] = s[readPos];
        }
    }
    
    return s.substr(0, writePos + 1);
}

// ==================== 4. MERGE SORTED STRINGS/ARRAYS ====================

string mergeSortedStrings(string s1, string s2) {
    string result = "";
    int i = 0, j = 0;
    
    while (i < s1.length() && j < s2.length()) {
        if (s1[i] <= s2[j]) {
            result += s1[i++];
        } else {
            result += s2[j++];
        }
    }
    
    while (i < s1.length()) {
        result += s1[i++];
    }
    
    while (j < s2.length()) {
        result += s2[j++];
    }
    
    return result;
}

// ==================== 5. STRING COMPRESSION ====================

// LeetCode 443: String Compression
int compress(vector<char>& chars) {
    int writePos = 0;
    int readPos = 0;
    int n = chars.size();
    
    while (readPos < n) {
        char currentChar = chars[readPos];
        int count = 0;
        
        // Count consecutive characters
        while (readPos < n && chars[readPos] == currentChar) {
            readPos++;
            count++;
        }
        
        // Write character
        chars[writePos++] = currentChar;
        
        // Write count if > 1
        if (count > 1) {
            string countStr = to_string(count);
            for (char c : countStr) {
                chars[writePos++] = c;
            }
        }
    }
    
    return writePos;
}

// ==================== 6. VALID PALINDROME II ====================

// LeetCode 680: Valid Palindrome II (can delete at most one character)
bool validPalindrome(string s) {
    int left = 0, right = s.length() - 1;
    
    while (left < right) {
        if (s[left] != s[right]) {
            // Try skipping left character
            int left1 = left + 1, right1 = right;
            bool valid1 = true;
            while (left1 < right1) {
                if (s[left1] != s[right1]) {
                    valid1 = false;
                    break;
                }
                left1++;
                right1--;
            }
            
            // Try skipping right character
            int left2 = left, right2 = right - 1;
            bool valid2 = true;
            while (left2 < right2) {
                if (s[left2] != s[right2]) {
                    valid2 = false;
                    break;
                }
                left2++;
                right2--;
            }
            
            return valid1 || valid2;
        }
        
        left++;
        right--;
    }
    
    return true;
}

// ==================== 7. REVERSE VOWELS OF A STRING ====================

// LeetCode 345: Reverse Vowels of a String
string reverseVowels(string s) {
    unordered_set<char> vowels = {'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'};
    
    int left = 0, right = s.length() - 1;
    
    while (left < right) {
        while (left < right && vowels.find(s[left]) == vowels.end()) {
            left++;
        }
        while (left < right && vowels.find(s[right]) == vowels.end()) {
            right--;
        }
        
        if (left < right) {
            swap(s[left], s[right]);
            left++;
            right--;
        }
    }
    
    return s;
}

// ==================== 8. IS SUBSEQUENCE ====================

// LeetCode 392: Is Subsequence
bool isSubsequence(string s, string t) {
    int i = 0, j = 0;
    
    while (i < s.length() && j < t.length()) {
        if (s[i] == t[j]) {
            i++;
        }
        j++;
    }
    
    return i == s.length();
}

// ==================== 9. PARTITION LABELS ====================

// LeetCode 763: Partition Labels
vector<int> partitionLabels(string s) {
    vector<int> result;
    
    // Store last occurrence of each character
    vector<int> lastOccurrence(26, -1);
    for (int i = 0; i < s.length(); i++) {
        lastOccurrence[s[i] - 'a'] = i;
    }
    
    int start = 0, end = 0;
    
    for (int i = 0; i < s.length(); i++) {
        end = max(end, lastOccurrence[s[i] - 'a']);
        
        if (i == end) {
            result.push_back(end - start + 1);
            start = i + 1;
        }
    }
    
    return result;
}

// ==================== 10. SORT COLORS (STRING VERSION - SORT CHARACTERS) ====================

// Sort string with three types of characters (like Dutch National Flag)
string sortThreeTypes(string s, char a, char b, char c) {
    int low = 0, mid = 0, high = s.length() - 1;
    
    while (mid <= high) {
        if (s[mid] == a) {
            swap(s[low], s[mid]);
            low++;
            mid++;
        } else if (s[mid] == b) {
            mid++;
        } else {
            swap(s[mid], s[high]);
            high--;
        }
    }
    
    return s;
}

// ==================== 11. CONTAINER WITH MOST WATER (STRING ANALOGY) ====================

// Maximum area between two characters (height-based)
int maxAreaBetweenChars(vector<int>& heights) {
    int left = 0, right = heights.size() - 1;
    int maxArea = 0;
    
    while (left < right) {
        int width = right - left;
        int height = min(heights[left], heights[right]);
        maxArea = max(maxArea, width * height);
        
        if (heights[left] < heights[right]) {
            left++;
        } else {
            right--;
        }
    }
    
    return maxArea;
}

// ==================== 12. TRAPPING RAIN WATER (STRING ANALOGY) ====================

int trap(vector<int>& height) {
    int n = height.size();
    if (n == 0) return 0;
    
    int left = 0, right = n - 1;
    int leftMax = 0, rightMax = 0;
    int water = 0;
    
    while (left < right) {
        if (height[left] < height[right]) {
            if (height[left] >= leftMax) {
                leftMax = height[left];
            } else {
                water += leftMax - height[left];
            }
            left++;
        } else {
            if (height[right] >= rightMax) {
                rightMax = height[right];
            } else {
                water += rightMax - height[right];
            }
            right--;
        }
    }
    
    return water;
}

// ==================== MAIN FUNCTION (TESTING) ====================

int main() {
    cout << "=== Two Pointers Technique ===\n\n";
    
    // Test valid palindrome
    string s1 = "A man, a plan, a canal: Panama";
    cout << "Is '" << s1 << "' palindrome? " 
         << (isPalindrome(s1) ? "Yes" : "No") << "\n";
    
    // Test reverse words
    string s2 = "  hello   world  ";
    cout << "Reverse words: '" << reverseWords(s2) << "'\n";
    
    // Test remove duplicates
    string s3 = "aabbcc";
    cout << "Remove duplicates: '" << removeDuplicates(s3) << "'\n";
    
    // Test valid palindrome II
    string s4 = "abca";
    cout << "Can '" << s4 << "' be palindrome? " 
         << (validPalindrome(s4) ? "Yes" : "No") << "\n";
    
    return 0;
}
