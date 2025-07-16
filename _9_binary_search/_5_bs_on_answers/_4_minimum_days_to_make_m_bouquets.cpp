// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/


// brute force
bool possible(vector<int> &arr, int day, int m, int k) {
    int n = arr.size(); //size of the array
    int cnt = 0;
    int noOfB = 0;
    // count the number of bouquets:
    for (int i = 0; i < n; i++) {
        if (arr[i] <= day) {
            cnt++;
        }
        else {
            noOfB += (cnt / k);
            cnt = 0;
        }
    }
    noOfB += (cnt / k);
    return noOfB >= m;
}
int roseGarden(vector<int> arr, int k, int m) {
    long long val = m * 1ll * k * 1ll;
    int n = arr.size(); //size of the array
    if (val > n) return -1; //impossible case.
    //find maximum and minimum:
    int mini = INT_MAX, maxi = INT_MIN;
    for (int i = 0; i < n; i++) {
        mini = min(mini, arr[i]);
        maxi = max(maxi, arr[i]);
    }

    for (int i = mini; i <= maxi; i++) {
        if (possible(arr, i, m, k))
            return i;
    }
    return -1;
}
// Time Complexity: O((max(arr[])-min(arr[])+1) * N), where {max(arr[]) -> maximum element of the array, min(arr[]) -> minimum element of the array, N = size of the array}.
// Reason: We are running a loop to check our answers that are in the range of [min(arr[]), max(arr[])]. For every possible answer, we will call the possible() function. Inside the possible() function, we are traversing the entire array, which results in O(N).
// Space Complexity: O(1) as we are not using any extra space to solve this problem.


class Solution
{
public:
    bool possible(vector<int> &arr, int day, int m, int k)
    {
        int cnt = 0;
        int noOfB = 0;

        for (int i = 0; i < arr.size(); i++)
        {
            if (arr[i] <= day)
                cnt++;
            else
            {
                noOfB += (cnt / k);
                cnt = 0;
            }
        }

        noOfB += (cnt / k);
        return noOfB >= m;
    }

    int minDays(vector<int> &bloomDay, int m, int k)
    {
        long long val = m * 1LL * k * 1LL;

        if (val > bloomDay.size())
            return -1;

        int mini = INT_MAX, maxi = INT_MIN;
        for (int i = 0; i < bloomDay.size(); i++)
        {
            mini = min(mini, bloomDay[i]);
            maxi = max(maxi, bloomDay[i]);
        }

        int l = mini, r = maxi;
        while (l <= r)
        {
            int mid = l + (r - l) / 2;

            if (possible(bloomDay, mid, m, k))
                r = mid - 1;
            else
                l = mid + 1;
        }
        return l;
    }
};

// Time Complexity: O(log(max(arr[])-min(arr[])+1) * N), where {max(arr[]) -> maximum element of the array, min(arr[]) -> minimum element of the array, N = size of the array}.
// Reason: We are applying binary search on our answers that are in the range of [min(arr[]), max(arr[])]. For every possible answer ‘mid’, we will call the possible() function. Inside the possible() function, we are traversing the entire array, which results in O(N).

// Space Complexity: O(1) as we are not using any extra space to solve this problem.