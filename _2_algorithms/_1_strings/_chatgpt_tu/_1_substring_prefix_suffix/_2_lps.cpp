// Given a string s, return the length of the longest prefix which is also a suffix, and not equal to the string itself.
// Input:  s = "level"
// Output: 1 ("l")

// Input:  s = "ababcab"
// Output: 2 ("ab")

#include <iostream>
using namespace std;

int longestPrefixSuffix(const string& s) {
    int n = s.length();
    for (int len = n - 1; len > 0; --len) {
        if (s.substr(0, len) == s.substr(n - len)) {
            return len;
        }
    }
    return 0;
}

int main() {
    string s = "ababcab";
    cout << "Length of LPS: " << longestPrefixSuffix(s) << endl;
    return 0;
}
