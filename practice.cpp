#include <iostream>
#include <vector>
#include <string>
#include <map>
#include <queue>

using namespace std;

// string smallestString(int N, vector<string> S)
// {
//     map<string, int> substrings;
//     for (int i = 0; i < S.size(); i++)
//     {
//         string s = S[i];
//         for (int i = 0; i < s.length(); i++)
//         {
//             for (int j = i; j < s.length(); j++)
//             {
//                 substrings[s.substr(i, j - i + 1)]++;
//             }
//         }
//     }

//     queue<string> q;
//     for (char c = 'a'; c <= 'z'; c++)
//     {
//         string str(1, c);
//         q.push(str);
//     }
//     while (!q.empty())
//     {
//         string curr = q.front();
//         q.pop();
//         if (substrings.find(curr) == substrings.end())
//         {
//             return curr;
//         }

//         for (int i = 0; i < 26; ++i)
//         {
//             curr.push_back(i + 'a');
//             q.push(curr);
//             curr.pop_back();
//         }
//     }
//     return " ";
// }

// int solve(int N, int H, vector<int> piles)
// {
//     int ans = -1;
//     int low = 1, high;

//     high = *max_element(piles.begin(),
//                         piles.end());

//     while (low <= high)

//     {
//         int K = low + (high - low) / 2;

//         int time = 0;

//         for (int ai : piles)
//         {

//             time += (ai + K - 1) / K;
//         }

//         if (time <= H)
//         {
//             ans = K;
//             high = K - 1;
//         }
//         else
//         {
//             low = K + 1;
//         }
//     }

//     return ans;
// }

int main()
{
    vector<string> S(3);
    S.push_back('abd');
    S.push_back('eg');
    S.push_back('acd');
    cout << "hi";
}