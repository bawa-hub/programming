// Given a string s and integer k, reverse the first k characters for every 2k characters from the start of the string.
// If there are fewer than k characters left, reverse all of them.
// If there are between k and 2k, reverse first k and leave rest.

// Input: s = "abcdefg", k = 2
// Output: "bacdfeg"

#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

string reverseStr(string s, int k) {
    int n = s.size();
    for (int i = 0; i < n; i += 2 * k) {
        int j = min(i + k, n);
        reverse(s.begin() + i, s.begin() + j);
    }
    return s;
}

int main() {
    string s = "abcdefg";
    int k = 2;
    cout << reverseStr(s, k) << endl;  // Output: "bacdfeg"
    return 0;
}
