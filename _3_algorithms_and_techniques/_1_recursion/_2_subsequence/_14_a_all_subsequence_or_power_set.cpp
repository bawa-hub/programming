// https://leetcode.com/problems/subsets/
// https://practice.geeksforgeeks.org/problems/power-set4302/1

// subsequences is contiguous/non-contiguous sequences, that follows the order

// print all the subsequence or power set of the given array
#include <bits/stdc++.h>
using namespace std;

void f(int idx, vector<int> &ds, int arr[], int n)
{
    if (idx == n)
    {
        cout << "{ ";
        for (auto it : ds)
        {
            cout << it << " ";
        }
        cout << "}";
        cout << endl;
        return;
    }

    // not pick, or not take condition, this element is not added to your subsequence
    f(idx + 1, ds, arr, n);

    // take or pick the particular index into the subsequence
    ds.push_back(arr[idx]);
    f(idx + 1, ds, arr, n);
    ds.pop_back();
}
// Time complexity - O(n*2^n)
// Space complexity - O(n)

int main()
{
    int arr[] = {3, 1, 2};
    int n = 3;
    vector<int> ds;
    f(0, ds, arr, n);

    return 0;
}

// print all possible subsequences of the String
// void solve(int i, string s, string &f) {
// 	if (i == s.length()) {
// 		cout << f << " ";
// 		return;
// 	}
// 	//picking
// 	f = f + s[i];
// 	solve(i + 1, s,  f);
// 	//poping out while backtracking
// 	f.pop_back();
// 	solve(i + 1, s,  f);
// }
// int main() {
// 	string s = "abc";
// 	string f = "";
// 	cout<<"All possible subsequences are: "<<endl;
// 	solve(0, s, f);
// }