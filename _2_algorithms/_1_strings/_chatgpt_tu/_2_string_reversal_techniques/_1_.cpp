#include <iostream>
#include <algorithm>
#include <sstream>
using namespace std;

// 1. Reverse Entire String
void reverseEntireString(string &s)
{
    int i = 0, j = s.size() - 1;
    while (i < j)
        swap(s[i++], s[j--]);
}

// 2. Reverse a Substring (i to j)
void reverseSubstring(string &s, int i, int j)
{
    while (i < j)
        swap(s[i++], s[j--]);
}

// 3. Reverse Words in a Sentence (Word Order Preserved)
string reverseEachWord(string s)
{
    stringstream ss(s);
    string word, result;
    while (ss >> word)
    {
        reverse(word.begin(), word.end());
        result += word + " ";
    }
    result.pop_back(); // Remove last space
    return result;
}
// Input: "I love C++" → Output: "I evol ++C"

// 4. Reverse Sentence Word Order
string reverseWordOrder(string s)
{
    stringstream ss(s);
    string word;
    vector<string> words;

    while (ss >> word)
        words.push_back(word);
    reverse(words.begin(), words.end());

    string result;
    for (auto &w : words)
        result += w + " ";
    result.pop_back();
    return result;
}
// "I love C++" → "C++ love I"

// 5. In-Place Reverse Words (No Extra Space, Hard)
// Steps:
//     Reverse whole string
//     Reverse each word
void reverseRange(string &s, int i, int j)
{
    while (i < j)
        swap(s[i++], s[j--]);
}

void reverseInPlaceWords(string &s)
{
    reverseRange(s, 0, s.size() - 1);

    int n = s.size(), i = 0;
    while (i < n)
    {
        if (s[i] == ' ')
        {
            i++;
            continue;
        }
        int j = i;
        while (j < n && s[j] != ' ')
            j++;
        reverseRange(s, i, j - 1);
        i = j;
    }
}

// 6. Handle Multiple Spaces Between Words

// 7. Reverse Only Vowels
bool isVowel(char c)
{
    return string("aeiouAEIOU").find(c) != string::npos;
}

void reverseVowels(string &s)
{
    int i = 0, j = s.length() - 1;
    while (i < j)
    {
        while (i < j && !isVowel(s[i]))
            i++;
        while (i < j && !isVowel(s[j]))
            j--;
        swap(s[i++], s[j--]);
    }
}

// 8. Rotation Trick (Double Reverse)

int main()
{
    string s = "hello";

    reverseEntireString(s);
    // reverse(s.begin(), s.end());
    cout << s << endl; // Output: "olleh"

    string s1 = "I love C++";
    cout << reverseEachWord(s1) << endl;
    cout << reverseWordOrder(s1) << endl;

    return 0;
}
