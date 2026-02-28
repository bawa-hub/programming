// Step-by-step Algorithm:
//     Compute hash of the pattern and the first window of text.
//     Slide the window over the text one by one:
//     Update hash of current window using rolling hash technique.
//     If current window hash == pattern hash, check characters one by one.
//     If they match, record the index.

#include <iostream>
#include <vector>
using namespace std;

#define d 256          // Number of characters in input alphabet
#define q 101          // A prime number for modulo

vector<int> rabinKarpSearch(const string &text, const string &pattern) {
    int n = text.size();
    int m = pattern.size();
    vector<int> result;

    if (m > n) return result;

    int p = 0;  // hash value for pattern
    int t = 0;  // hash value for text
    int h = 1;

    // The value of h would be "pow(d, m-1) % q"
    for (int i = 0; i < m - 1; i++)
        h = (h * d) % q;

    // Calculate initial hash values for pattern and first window of text
    for (int i = 0; i < m; i++) {
        p = (d * p + pattern[i]) % q;
        t = (d * t + text[i]) % q;
    }

    // Slide the pattern over text
    for (int i = 0; i <= n - m; i++) {
        // Check hash values of current window and pattern
        if (p == t) {
            // Check characters one by one
            int j = 0;
            for (; j < m; j++) {
                if (text[i + j] != pattern[j])
                    break;
            }
            if (j == m)
                result.push_back(i);
        }

        // Calculate hash for next window
        if (i < n - m) {
            t = (d * (t - text[i] * h) + text[i + m]) % q;

            // Make sure t is positive
            if (t < 0)
                t = t + q;
        }
    }

    return result;
}

int main() {
    string text = "GEEKS FOR GEEKS";
    string pattern = "GEEK";

    vector<int> matches = rabinKarpSearch(text, pattern);

    cout << "Pattern found at indices: ";
    for (int idx : matches) {
        cout << idx << " ";
    }
    cout << endl;

    return 0;
}

// Explanation:
//     d is the alphabet size (256 for ASCII).
//     q is a prime modulus to reduce hash collisions.
//     h helps to remove the leftmost character contribution from rolling hash.
//     We check character-by-character only when the hashes match.