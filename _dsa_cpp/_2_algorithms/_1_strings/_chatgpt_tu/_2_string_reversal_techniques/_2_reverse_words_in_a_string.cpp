// Given a string s, reverse the order of words.
// A word is a sequence of non-space characters. Return the reversed sentence with no extra spaces.

// Input: "   the sky   is blue "
// Output: "blue is sky the"

#include <iostream>
#include <sstream>
#include <vector>
#include <algorithm>
using namespace std;

string trimAndReverseWords(string s) {
    stringstream ss(s);
    string word;
    vector<string> words;
    while (ss >> word) words.push_back(word);

    reverse(words.begin(), words.end());

    string result;
    for (string& w : words) result += w + " ";
    result.pop_back(); // remove trailing space
    return result;
}

int main() {
    string s = "   the sky   is blue ";
    cout << '"' << trimAndReverseWords(s) << '"' << endl;
    return 0;
}
// ⏱ Time: O(n)
// 📦 Space: O(n)