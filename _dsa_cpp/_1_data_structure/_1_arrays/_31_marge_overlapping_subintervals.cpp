// https://leetcode.com/problems/merge-intervals/

#include <bits/stdc++.h>
using namespace std;

// brute force
vector<pair<int, int>> merge(vector<pair<int, int>> &arr)
{

    int n = arr.size();
    sort(arr.begin(), arr.end());
    vector<pair<int, int>> ans;

    for (int i = 0; i < n; i++)
    {
        int start = arr[i].first, end = arr[i].second;

        // since the intervals already lies
        // in the data structure present we continue
        if (!ans.empty())
        {
            if (start <= ans.back().second)
            {
                continue;
            }
        }

        for (int j = i + 1; j < n; j++)
        {
            if (arr[j].first <= end)
            {
                end = max(end, arr[j].second);
            }
        }

        end = max(end, arr[i].second);

        ans.push_back({start, end});
    }

    return ans;
}

int main()
{
    vector<pair<int, int>> arr;
    arr = {{1, 3}, {2, 4}, {2, 6}, {8, 9}, {8, 10}, {9, 11}, {15, 18}, {16, 17}};
    vector<pair<int, int>> ans = merge(arr);

    cout << "Merged Overlapping Intervals are " << endl;

    for (auto it : ans)
    {
        cout << it.first << " " << it.second << "\n";
    }
}
// Time Complexity: O(NlogN)+O(N*N). O(NlogN) for sorting the array, and O(N*N) because we are checking to the right for each index which is a nested loop.
// Space Complexity: O(N), as we are using a separate data structure.

// optimal
vector<vector<int>> merge(vector<vector<int>> &intervals)
{

    sort(intervals.begin(), intervals.end());
    vector<vector<int>> merged;

    for (int i = 0; i < intervals.size(); i++)
    {
        if (merged.empty() || merged.back()[1] < intervals[i][0])
        {
            vector<int> v = {
                intervals[i][0],
                intervals[i][1]};

            merged.push_back(v);
        }
        else
        {
            merged.back()[1] = max(merged.back()[1], intervals[i][1]);
        }
    }

    return merged;
}

int main()
{
    vector<vector<int>> arr;
    arr = {{1, 3}, {2, 4}, {2, 6}, {8, 9}, {8, 10}, {9, 11}, {15, 18}, {16, 17}};
    vector<vector<int>> ans = merge(arr);

    cout << "Merged Overlapping Intervals are " << endl;

    for (auto it : ans)
    {
        cout << it[0] << " " << it[1] << "\n";
    }
}
// Time Complexity: O(NlogN) + O(N). O(NlogN) for sorting and O(N) for traversing through the array.
// Space Complexity: O(N) to return the answer of the merged intervals.