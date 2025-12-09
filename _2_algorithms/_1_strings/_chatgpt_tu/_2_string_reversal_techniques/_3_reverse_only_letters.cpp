// Given a string s, reverse only the letters, keeping non-letter characters in place.
// Input: "a-bC-dEf-ghIj"
// Output: "j-Ih-gfE-dCba"

#include <iostream>
#include <cctype>
using namespace std;

string reverseOnlyLetters(string s) {
    int i = 0, j = s.size() - 1;
    while (i < j) {
        if (!isalpha(s[i])) { i++; continue; }
        if (!isalpha(s[j])) { j--; continue; }
        swap(s[i++], s[j--]);
    }
    return s;
}

int main() {
    string s = "a-bC-dEf-ghIj";
    cout << reverseOnlyLetters(s) << endl;
    return 0;
}
