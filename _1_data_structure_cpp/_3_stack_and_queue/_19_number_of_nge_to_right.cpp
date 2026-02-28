// https://practice.geeksforgeeks.org/problems/number-of-nges-to-the-right/1
// https://www.geeksforgeeks.org/number-nges-right/

// C++ code for the above approach

#include <bits/stdc++.h>
using namespace std;

// brute force
int nextGreaterElements(vector<int>& a, int index)
{
	int count = 0, N = a.size();
	for (int i = index + 1; i < N; i++)
		if (a[i] > a[index])
			count++;

	return count;
}
// Time Complexity: O(NQ), and O(N) to answer a single query
// Auxiliary space: O(1) 

int main()
{

	vector<int> a = { 3, 4, 2, 7, 5, 8, 10, 6 };
	int Q = 2;
	vector<int> queries = { 0, 5 };
	for (int i = 0; i < Q; i++)
		cout << nextGreaterElements(a, queries[i]) << " ";
	return 0;
}
