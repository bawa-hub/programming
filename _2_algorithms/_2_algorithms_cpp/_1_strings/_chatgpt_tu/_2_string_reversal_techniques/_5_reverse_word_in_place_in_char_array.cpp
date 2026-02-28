// You're given a vector<char> representing a sentence with words separated by single spaces (no leading/trailing spaces).
// Reverse the entire sentence word-by-word in-place, without using extra space.


// Input:  ['t','h','e',' ','s','k','y',' ','i','s',' ','b','l','u','e']
// Output: ['b','l','u','e',' ','i','s',' ','s','k','y',' ','t','h','e']

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void reverseRange(vector<char>& s, int i, int j) {
    while (i < j) swap(s[i++], s[j--]);
}

void reverseWordsInPlace(vector<char>& s) {
    int n = s.size();

    // Step 1: Reverse whole string
    reverseRange(s, 0, n - 1);

    // Step 2: Reverse each word
    int i = 0;
    while (i < n) {
        int j = i;
        while (j < n && s[j] != ' ') j++;
        reverseRange(s, i, j - 1);
        i = j + 1;
    }
}

int main() {
    vector<char> s = {'t','h','e',' ','s','k','y',' ','i','s',' ','b','l','u','e'};
    reverseWordsInPlace(s);
    for (char c : s) cout << c;
    cout << endl;
    return 0;
}
