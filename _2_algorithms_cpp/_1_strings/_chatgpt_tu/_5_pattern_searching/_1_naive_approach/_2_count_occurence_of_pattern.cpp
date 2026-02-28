// Given a text and a pattern, count how many times the pattern occurs in the text using naive pattern searching.
// Input: text = "abababa", pattern = "aba"
// Output: 3

#include <iostream>
#include <vector>
using namespace std;

int countOccurrences(const string &text, const string &pattern) {
    int count = 0;
    int n = text.size();
    int m = pattern.size();

    for (int i = 0; i <= n - m; i++) {
        int j = 0;
        while (j < m && text[i + j] == pattern[j]) {
            j++;
        }
        if (j == m) {
            count++;
        }
    }
    return count;
}

int main() {
    string text = "abababa";
    string pattern = "aba";

    cout << "Occurrences of pattern: " << countOccurrences(text, pattern) << endl;

    return 0;
}
