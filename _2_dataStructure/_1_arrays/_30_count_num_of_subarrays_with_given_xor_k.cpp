// https://practice.geeksforgeeks.org/problems/subsets-with-xor-value2023/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=subsets-with-xor-value

#include <bits/stdc++.h>
using namespace std;
class Solution
{
public:
    // brute force
    int solve(vector<int> &A, int B)
    {
        long long c = 0;
        for (int i = 0; i < A.size(); i++)
        {
            int current_xor = 0;
            for (int j = i; j < A.size(); j++)
            {
                current_xor ^= A[j];
                if (current_xor == B)
                    c++;
            }
        }
        return c;
    }
    //     Time Complexity: O(N2)
    // Space Complexity: O(1)

    // prefix xor and map
    int solve(vector<int> &A, int B)
    {
        unordered_map<int, int> visited;
        int cpx = 0;
        long long c = 0;
        for (int i = 0; i < A.size(); i++)
        {
            cpx ^= A[i];
            if (cpx == B)
                c++;
            int h = cpx ^ B;
            if (visited.find(h) != visited.end())
            {
                c = c + visited[h];
            }
            visited[cpx]++;
        }
        return c;
    }
    // Time Complexity: O(N)
    // Space Complexity: O(N)
};

int main()
{
    vector<int> A{4, 2, 2, 6, 4};
    Solution obj;
    int totalCount = obj.solve(A, 6);
    cout << "The total number of subarrays having a given XOR k is " << totalCount << endl;
}