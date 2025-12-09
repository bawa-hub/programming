//  Return true if the pattern is present in the text, otherwise false.
// Input: text = "abcdef", pattern = "cd"
// Output: true

// Input: text = "abcdef", pattern = "gh"
// Output: false

#include <iostream>
using namespace std;

bool patternExists(const string &text, const string &pattern) {
    int n = text.size();
    int m = pattern.size();

    for (int i = 0; i <= n - m; i++) {
        int j = 0;
        while (j < m && text[i + j] == pattern[j]) {
            j++;
        }
        if (j == m) {
            return true;
        }
    }
    return false;
}

int main() {
    string text = "abcdef";
    string pattern = "cd";

    cout << boolalpha << "Pattern exists: " << patternExists(text, pattern) << endl;

    pattern = "gh";
    cout << boolalpha << "Pattern exists: " << patternExists(text, pattern) << endl;

    return 0;
}
