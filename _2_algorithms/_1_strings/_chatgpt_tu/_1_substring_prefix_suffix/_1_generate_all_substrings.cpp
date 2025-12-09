// Goal: Write a function to print all possible substrings of a string.

#include <iostream>
using namespace std;

// Generate All Substrings
void printAllSubstrings(const string& s) {
    int n = s.length();
    for (int i = 0; i < n; ++i) {          
        for (int j = i; j < n; ++j) {     
            cout << s.substr(i, j - i + 1) << endl;
        }
    }
}
// Total substrings = n * (n + 1) / 2

// Generate All Prefixes
void printAllPrefixes(const string& s) {
    for (int i = 1; i <= s.length(); ++i) {
        cout << s.substr(0, i) << endl;
    }
}
// Input: "abcd"
// Output: "a", "ab", "abc", "abcd"

// Generate All Suffixes
void printAllSuffixes(const string& s) {
    for (int i = 0; i < s.length(); ++i) {
        cout << s.substr(i) << endl;
    }
}
// Input: "abcd"
// Output: "abcd", "bcd", "cd", "d"

// Given a string s, count all substrings where the first and last characters are the same.
int countSubstringsSameStartEnd(const string& s) {
    unordered_map<char, int> freq;
    for (char c : s) {
        freq[c]++;
    }

    int count = 0;
    for (auto& [ch, f] : freq) {
        count += f; // Single character substrings
        count += (f * (f - 1)) / 2; // Combinations of substrings
    }
    return count;
}
// We can count the frequency of each character and use this formula:
//     For each char c appearing f times:
//     Count of valid substrings = f + (f choose 2)
//     f substrings of length 1 + all pairs (f * (f - 1) / 2)

int main() {
    string str = "abc";
    printAllSubstrings(str);
    return 0;
}

