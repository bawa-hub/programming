/*
 * Basic String Operations
 * Essential operations every DSA candidate must master
 */

#include <iostream>
#include <string>
#include <algorithm>
#include <unordered_map>
#include <vector>
#include <cctype>

using namespace std;

// ==================== 1. STRING REVERSAL ====================

// Method 1: Using two pointers
string reverseStringTwoPointers(string s) {
    int left = 0, right = s.length() - 1;
    while (left < right) {
        swap(s[left], s[right]);
        left++;
        right--;
    }
    return s;
}

// Method 2: Using STL reverse
string reverseStringSTL(string s) {
    reverse(s.begin(), s.end());
    return s;
}

// Method 3: Recursive approach
void reverseRecursive(string& s, int left, int right) {
    if (left >= right) return;
    swap(s[left], s[right]);
    reverseRecursive(s, left + 1, right - 1);
}

// ==================== 2. PALINDROME CHECKING ====================

// Method 1: Two pointers
bool isPalindrome(string s) {
    int left = 0, right = s.length() - 1;
    while (left < right) {
        if (s[left] != s[right]) {
            return false;
        }
        left++;
        right--;
    }
    return true;
}

// Method 2: Case-insensitive palindrome (ignoring non-alphanumeric)
bool isPalindromeAdvanced(string s) {
    int left = 0, right = s.length() - 1;
    while (left < right) {
        // Skip non-alphanumeric characters
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

// Method 3: Recursive
bool isPalindromeRecursive(string s, int left, int right) {
    if (left >= right) return true;
    if (s[left] != s[right]) return false;
    return isPalindromeRecursive(s, left + 1, right - 1);
}

// ==================== 3. ANAGRAM DETECTION ====================

// Method 1: Sorting
bool isAnagramSorting(string s1, string s2) {
    if (s1.length() != s2.length()) return false;
    sort(s1.begin(), s1.end());
    sort(s2.begin(), s2.end());
    return s1 == s2;
}

// Method 2: Hash map (more efficient)
bool isAnagramHashMap(string s1, string s2) {
    if (s1.length() != s2.length()) return false;
    
    unordered_map<char, int> freq;
    
    // Count characters in s1
    for (char c : s1) {
        freq[c]++;
    }
    
    // Decrement for s2
    for (char c : s2) {
        freq[c]--;
        if (freq[c] < 0) return false;
    }
    
    return true;
}

// Method 3: Array (for lowercase letters only)
bool isAnagramArray(string s1, string s2) {
    if (s1.length() != s2.length()) return false;
    
    int count[26] = {0};
    
    for (int i = 0; i < s1.length(); i++) {
        count[s1[i] - 'a']++;
        count[s2[i] - 'a']--;
    }
    
    for (int i = 0; i < 26; i++) {
        if (count[i] != 0) return false;
    }
    
    return true;
}

// ==================== 4. SUBSTRING OPERATIONS ====================

// Check if string contains substring
bool containsSubstring(string text, string pattern) {
    return text.find(pattern) != string::npos;
}

// Find all occurrences of substring
vector<int> findAllOccurrences(string text, string pattern) {
    vector<int> positions;
    size_t pos = text.find(pattern);
    
    while (pos != string::npos) {
        positions.push_back(pos);
        pos = text.find(pattern, pos + 1);
    }
    
    return positions;
}

// Get substring between two delimiters
string substringBetween(string s, string start, string end) {
    size_t startPos = s.find(start);
    if (startPos == string::npos) return "";
    
    startPos += start.length();
    size_t endPos = s.find(end, startPos);
    
    if (endPos == string::npos) return "";
    
    return s.substr(startPos, endPos - startPos);
}

// ==================== 5. CHARACTER FREQUENCY ====================

// Count frequency of each character
unordered_map<char, int> getCharFrequency(string s) {
    unordered_map<char, int> freq;
    for (char c : s) {
        freq[c]++;
    }
    return freq;
}

// Find most frequent character
char mostFrequentChar(string s) {
    unordered_map<char, int> freq;
    int maxFreq = 0;
    char result = '\0';
    
    for (char c : s) {
        freq[c]++;
        if (freq[c] > maxFreq) {
            maxFreq = freq[c];
            result = c;
        }
    }
    
    return result;
}

// ==================== 6. STRING VALIDATION ====================

// Check if string contains only digits
bool isNumeric(string s) {
    for (char c : s) {
        if (!isdigit(c)) return false;
    }
    return true;
}

// Check if string contains only letters
bool isAlpha(string s) {
    for (char c : s) {
        if (!isalpha(c)) return false;
    }
    return true;
}

// Check if string is valid email (basic)
bool isValidEmail(string email) {
    size_t atPos = email.find('@');
    if (atPos == string::npos || atPos == 0) return false;
    
    size_t dotPos = email.find('.', atPos);
    if (dotPos == string::npos || dotPos == atPos + 1) return false;
    
    return dotPos < email.length() - 1;
}

// ==================== 7. STRING TRANSFORMATION ====================

// Remove all spaces
string removeSpaces(string s) {
    s.erase(remove(s.begin(), s.end(), ' '), s.end());
    return s;
}

// Remove duplicates (preserving order)
string removeDuplicates(string s) {
    unordered_map<char, bool> seen;
    string result = "";
    
    for (char c : s) {
        if (seen.find(c) == seen.end()) {
            seen[c] = true;
            result += c;
        }
    }
    
    return result;
}

// Convert to lowercase
string toLowerCase(string s) {
    transform(s.begin(), s.end(), s.begin(), ::tolower);
    return s;
}

// Convert to uppercase
string toUpperCase(string s) {
    transform(s.begin(), s.end(), s.begin(), ::toupper);
    return s;
}

// ==================== 8. STRING COMPARISON ====================

// Compare strings lexicographically
int compareStrings(string s1, string s2) {
    return s1.compare(s2);
    // Returns: <0 if s1 < s2, 0 if equal, >0 if s1 > s2
}

// Check if strings are equal (case-insensitive)
bool equalsIgnoreCase(string s1, string s2) {
    if (s1.length() != s2.length()) return false;
    
    for (int i = 0; i < s1.length(); i++) {
        if (tolower(s1[i]) != tolower(s2[i])) {
            return false;
        }
    }
    return true;
}

// ==================== MAIN FUNCTION (TESTING) ====================

int main() {
    cout << "=== Basic String Operations ===\n\n";
    
    // Test reversal
    string test = "hello";
    cout << "Reverse of '" << test << "': " << reverseStringTwoPointers(test) << "\n";
    
    // Test palindrome
    cout << "\nIs 'racecar' palindrome? " << (isPalindrome("racecar") ? "Yes" : "No") << "\n";
    cout << "Is 'A man a plan a canal Panama' palindrome? " 
         << (isPalindromeAdvanced("A man a plan a canal Panama") ? "Yes" : "No") << "\n";
    
    // Test anagram
    cout << "\nAre 'listen' and 'silent' anagrams? " 
         << (isAnagramHashMap("listen", "silent") ? "Yes" : "No") << "\n";
    
    // Test substring
    vector<int> positions = findAllOccurrences("ababab", "ab");
    cout << "\nOccurrences of 'ab' in 'ababab': ";
    for (int pos : positions) cout << pos << " ";
    cout << "\n";
    
    // Test frequency
    cout << "\nMost frequent char in 'hello': " << mostFrequentChar("hello") << "\n";
    
    return 0;
}
