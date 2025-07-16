// Rotate a string s right by k characters.
// Input: s = "abcdef", k = 2
// Output: "efabcd"

#include <iostream>
#include <algorithm>
using namespace std;

string rotateRight(string s, int k) {
    int n = s.length();
    k = k % n;

    reverse(s.begin(), s.end());
    reverse(s.begin(), s.begin() + k);
    reverse(s.begin() + k, s.end());

    return s;
}

int main() {
    string s = "abcdef";
    int k = 2;
    cout << rotateRight(s, k) << endl;  // "efabcd"
    return 0;
}
