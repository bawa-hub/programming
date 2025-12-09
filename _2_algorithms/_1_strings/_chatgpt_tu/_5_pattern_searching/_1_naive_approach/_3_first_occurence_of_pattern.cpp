//  Return the index of the first occurrence of the pattern in the text. If the pattern is not found, return -1.
// Input: text = "hello world", pattern = "world"
// Output: 6

// Input: text = "hello world", pattern = "planet"
// Output: -1

#include <iostream>
using namespace std;

int firstOccurrence(const string &text, const string &pattern) {
    int n = text.size();
    int m = pattern.size();

    for (int i = 0; i <= n - m; i++) {
        int j = 0;
        while (j < m && text[i + j] == pattern[j]) {
            j++;
        }
        if (j == m) {
            return i;
        }
    }
    return -1;  // Not found
}

int main() {
    string text = "hello world";
    string pattern = "world";

    cout << "First occurrence at index: " << firstOccurrence(text, pattern) << endl;

    pattern = "planet";
    cout << "First occurrence at index: " << firstOccurrence(text, pattern) << endl;

    return 0;
}
