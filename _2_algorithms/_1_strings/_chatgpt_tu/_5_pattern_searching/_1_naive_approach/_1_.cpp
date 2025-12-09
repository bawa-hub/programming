// for each index in the text, check if the substring starting there matches the pattern.
// If yes, record the index.
// Time Complexity: O(n*m), where n = length of text, m = length of pattern.

#include <iostream>
#include <vector>
using namespace std;

vector<int> naivePatternSearch(const string &text, const string &pattern)
{
    vector<int> result;
    int n = text.size();
    int m = pattern.size();

    for (int i = 0; i <= n - m; i++)
    {
        int j = 0;
        while (j < m && text[i + j] == pattern[j])
        {
            j++;
        }
        if (j == m)
        {
            result.push_back(i);
        }
    }

    return result;
}

int main()
{
    string text = "ABABDABACDABABCABAB";
    string pattern = "ABABCABAB";

    vector<int> occurrences = naivePatternSearch(text, pattern);

    cout << "Pattern found at indices: ";
    for (int index : occurrences)
    {
        cout << index << " ";
    }
    cout << endl;

    return 0;
}
