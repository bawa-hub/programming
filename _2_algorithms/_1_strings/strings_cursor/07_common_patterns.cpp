/*
 * Common Interview Patterns for Strings
 * Frequently asked patterns in coding interviews
 */

#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <stack>
#include <queue>
#include <algorithm>
#include <sstream>
#include <climits>

using namespace std;

// ==================== 1. STRING ENCODING/DECODING ====================

// LeetCode 271: Encode and Decode Strings
class Codec {
public:
    // Encodes a list of strings to a single string
    string encode(vector<string>& strs) {
        string encoded = "";
        for (string str : strs) {
            encoded += to_string(str.length()) + "#" + str;
        }
        return encoded;
    }
    
    // Decodes a single string to a list of strings
    vector<string> decode(string s) {
        vector<string> result;
        int i = 0;
        
        while (i < s.length()) {
            // Find the delimiter
            int j = i;
            while (s[j] != '#') {
                j++;
            }
            
            // Extract length
            int length = stoi(s.substr(i, j - i));
            
            // Extract string
            result.push_back(s.substr(j + 1, length));
            
            // Move to next string
            i = j + 1 + length;
        }
        
        return result;
    }
};

// ==================== 2. STRING INTERLEAVING ====================

// LeetCode 97: Interleaving String
bool isInterleave(string s1, string s2, string s3) {
    int m = s1.length();
    int n = s2.length();
    
    if (m + n != s3.length()) return false;
    
    vector<vector<bool>> dp(m + 1, vector<bool>(n + 1, false));
    dp[0][0] = true;
    
    // Initialize first row
    for (int j = 1; j <= n; j++) {
        dp[0][j] = dp[0][j - 1] && s2[j - 1] == s3[j - 1];
    }
    
    // Initialize first column
    for (int i = 1; i <= m; i++) {
        dp[i][0] = dp[i - 1][0] && s1[i - 1] == s3[i - 1];
    }
    
    // Fill DP table
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            dp[i][j] = (dp[i - 1][j] && s1[i - 1] == s3[i + j - 1]) ||
                       (dp[i][j - 1] && s2[j - 1] == s3[i + j - 1]);
        }
    }
    
    return dp[m][n];
}

// ==================== 3. VALID PARENTHESES ====================

// LeetCode 20: Valid Parentheses
bool isValidParentheses(string s) {
    stack<char> st;
    unordered_map<char, char> mapping = {
        {')', '('},
        {'}', '{'},
        {']', '['}
    };
    
    for (char c : s) {
        if (mapping.find(c) != mapping.end()) {
            // Closing bracket
            if (st.empty() || st.top() != mapping[c]) {
                return false;
            }
            st.pop();
        } else {
            // Opening bracket
            st.push(c);
        }
    }
    
    return st.empty();
}

// ==================== 4. REMOVE INVALID PARENTHESES ====================

// LeetCode 301: Remove Invalid Parentheses
vector<string> removeInvalidParentheses(string s) {
    vector<string> result;
    unordered_set<string> visited;
    queue<string> q;
    
    q.push(s);
    visited.insert(s);
    bool found = false;
    
    while (!q.empty()) {
        string current = q.front();
        q.pop();
        
        if (isValidParentheses(current)) {
            result.push_back(current);
            found = true;
        }
        
        if (found) continue;
        
        // Generate all possible strings by removing one character
        for (int i = 0; i < current.length(); i++) {
            if (current[i] != '(' && current[i] != ')') continue;
            
            string next = current.substr(0, i) + current.substr(i + 1);
            if (visited.find(next) == visited.end()) {
                visited.insert(next);
                q.push(next);
            }
        }
    }
    
    return result;
}

// ==================== 5. STRING TO INTEGER (ATOI) ====================

// LeetCode 8: String to Integer (atoi)
int myAtoi(string s) {
    int i = 0;
    int n = s.length();
    
    // Skip whitespace
    while (i < n && s[i] == ' ') {
        i++;
    }
    
    if (i == n) return 0;
    
    // Check sign
    int sign = 1;
    if (s[i] == '+' || s[i] == '-') {
        sign = (s[i] == '-') ? -1 : 1;
        i++;
    }
    
    // Convert digits
    long long result = 0;
    while (i < n && isdigit(s[i])) {
        result = result * 10 + (s[i] - '0');
        
        // Check overflow
        if (sign * result > INT_MAX) return INT_MAX;
        if (sign * result < INT_MIN) return INT_MIN;
        
        i++;
    }
    
    return sign * result;
}

// ==================== 6. ROMAN TO INTEGER ====================

// LeetCode 13: Roman to Integer
int romanToInt(string s) {
    unordered_map<char, int> values = {
        {'I', 1},
        {'V', 5},
        {'X', 10},
        {'L', 50},
        {'C', 100},
        {'D', 500},
        {'M', 1000}
    };
    
    int result = 0;
    for (int i = 0; i < s.length(); i++) {
        if (i < s.length() - 1 && values[s[i]] < values[s[i + 1]]) {
            result -= values[s[i]];
        } else {
            result += values[s[i]];
        }
    }
    
    return result;
}

// ==================== 7. INTEGER TO ROMAN ====================

// LeetCode 12: Integer to Roman
string intToRoman(int num) {
    vector<pair<int, string>> values = {
        {1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
        {100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
        {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}
    };
    
    string result = "";
    for (auto& p : values) {
        while (num >= p.first) {
            result += p.second;
            num -= p.first;
        }
    }
    
    return result;
}

// ==================== 8. ZIGZAG CONVERSION ====================

// LeetCode 6: Zigzag Conversion
string convert(string s, int numRows) {
    if (numRows == 1) return s;
    
    vector<string> rows(min(numRows, (int)s.length()));
    int currentRow = 0;
    bool goingDown = false;
    
    for (char c : s) {
        rows[currentRow] += c;
        
        if (currentRow == 0 || currentRow == numRows - 1) {
            goingDown = !goingDown;
        }
        
        currentRow += goingDown ? 1 : -1;
    }
    
    string result = "";
    for (string row : rows) {
        result += row;
    }
    
    return result;
}

// ==================== 9. LONGEST VALID PARENTHESES ====================

// LeetCode 32: Longest Valid Parentheses
int longestValidParentheses(string s) {
    stack<int> st;
    st.push(-1);  // Base for calculation
    int maxLen = 0;
    
    for (int i = 0; i < s.length(); i++) {
        if (s[i] == '(') {
            st.push(i);
        } else {
            st.pop();
            if (st.empty()) {
                st.push(i);
            } else {
                maxLen = max(maxLen, i - st.top());
            }
        }
    }
    
    return maxLen;
}

// ==================== 10. SIMPLIFY PATH ====================

// LeetCode 71: Simplify Path
string simplifyPath(string path) {
    vector<string> stack;
    stringstream ss(path);
    string token;
    
    while (getline(ss, token, '/')) {
        if (token == "" || token == ".") {
            continue;
        } else if (token == "..") {
            if (!stack.empty()) {
                stack.pop_back();
            }
        } else {
            stack.push_back(token);
        }
    }
    
    string result = "";
    for (string dir : stack) {
        result += "/" + dir;
    }
    
    return result.empty() ? "/" : result;
}

// ==================== 11. BASIC CALCULATOR ====================

// LeetCode 224: Basic Calculator
int calculate(string s) {
    stack<int> st;
    int result = 0;
    int number = 0;
    int sign = 1;
    
    for (char c : s) {
        if (isdigit(c)) {
            number = number * 10 + (c - '0');
        } else if (c == '+') {
            result += sign * number;
            number = 0;
            sign = 1;
        } else if (c == '-') {
            result += sign * number;
            number = 0;
            sign = -1;
        } else if (c == '(') {
            st.push(result);
            st.push(sign);
            result = 0;
            sign = 1;
        } else if (c == ')') {
            result += sign * number;
            number = 0;
            result *= st.top(); st.pop();
            result += st.top(); st.pop();
        }
    }
    
    if (number != 0) {
        result += sign * number;
    }
    
    return result;
}

// ==================== 12. REVERSE STRING II ====================

// LeetCode 541: Reverse String II
string reverseStr(string s, int k) {
    int n = s.length();
    for (int i = 0; i < n; i += 2 * k) {
        int left = i;
        int right = min(i + k - 1, n - 1);
        
        while (left < right) {
            swap(s[left], s[right]);
            left++;
            right--;
        }
    }
    
    return s;
}

// ==================== 13. REPEATED SUBSTRING PATTERN ====================

// LeetCode 459: Repeated Substring Pattern
bool repeatedSubstringPattern(string s) {
    int n = s.length();
    
    for (int len = 1; len <= n / 2; len++) {
        if (n % len != 0) continue;
        
        string pattern = s.substr(0, len);
        bool valid = true;
        
        for (int i = len; i < n; i += len) {
            if (s.substr(i, len) != pattern) {
                valid = false;
                break;
            }
        }
        
        if (valid) return true;
    }
    
    return false;
}

// ==================== MAIN FUNCTION (TESTING) ====================

int main() {
    cout << "=== Common Interview Patterns ===\n\n";
    
    // Test encoding/decoding
    Codec codec;
    vector<string> strs = {"hello", "world", "test"};
    string encoded = codec.encode(strs);
    vector<string> decoded = codec.decode(encoded);
    cout << "Encoded: " << encoded << "\n";
    cout << "Decoded: ";
    for (string s : decoded) cout << s << " ";
    cout << "\n\n";
    
    // Test valid parentheses
    cout << "Is '()[]{}' valid? " << (isValid("()[]{}") ? "Yes" : "No") << "\n";
    
    // Test atoi
    cout << "atoi('   -42'): " << myAtoi("   -42") << "\n";
    
    // Test roman to int
    cout << "Roman 'MCMXCIV' to int: " << romanToInt("MCMXCIV") << "\n";
    
    return 0;
}
