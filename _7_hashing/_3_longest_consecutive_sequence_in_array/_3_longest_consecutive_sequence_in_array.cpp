// https://leetcode.com/problems/longest-consecutive-sequence/description/

#include <bits/stdc++.h>
using namespace std;

// brute force
int longestConsecutive(vector<int> nums)
{
    if (nums.size() == 0)
    {
        return 0;
    }

    sort(nums.begin(), nums.end());

    int ans = 1;
    int prev = nums[0];
    int cur = 1;

    for (int i = 1; i < nums.size(); i++)
    {
        if (nums[i] == prev + 1)
        {
            cur++;
        }
        else if (nums[i] != prev)
        {
            cur = 1;
        }
        prev = nums[i];
        ans = max(ans, cur);
    }
    return ans;
}
//     Time Complexity: We are first sorting the array which will take O(N * log(N)) time and then we are running a for loop which will take O(N) time. Hence, the overall time complexity will be O(N * log(N)).
// Space Complexity: The space complexity for the above approach is O(1) because we are not using any auxiliary space

//    optimized
int longestConsecutive(vector<int> &nums)
{
    set<int> hashSet;
    for (int num : nums)
    {
        hashSet.insert(num);
    }

    int longestStreak = 0;

    for (int num : nums)
    {
        if (!hashSet.count(num - 1))
        {
            int currentNum = num;
            int currentStreak = 1;

            while (hashSet.count(currentNum + 1))
            {
                currentNum += 1;
                currentStreak += 1;
            }

            longestStreak = max(longestStreak, currentStreak);
        }
    }

    return longestStreak;
}
// Time Complexity: The time complexity of the above approach is O(N) because we traverse each consecutive subsequence only once. (assuming HashSet takes O(1) to search)
// Space Complexity: The space complexity of the above approach is O(N) because we are maintaining a HashSet.

int main()
{
    vector<int> arr{100, 200, 1, 2, 3, 4};
    int lon = longestConsecutive(arr);
    cout << "The longest consecutive sequence is " << lon << endl;
}