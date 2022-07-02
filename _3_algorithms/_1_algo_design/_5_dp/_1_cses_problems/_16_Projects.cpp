#include <bits/stdc++.h>
using namespace std;

#define mid(l, r) (l + (r - l) / 2);

struct Job
{
    long long start, finish, profit;
};

bool comparator(Job a, Job b)
{
    return a.finish < b.finish;
}

long long binarySearch(Job arr[], int i)
{
    long long low = 0, high = i - 1;
    while (low <= high)
    {
        long long mid = mid(low, high);
        if (arr[mid].finish < arr[i].start)
        {
            if (arr[mid + 1].finish < arr[i].start)
            {
                low = mid + 1;
            }
            else
            {
                return mid;
            }
        }
        else
        {
            high = mid - 1;
        }
    }
    return -1;
}

int main()
{
    long long n;
    cin >> n;
    struct Job arr[n];
    for (int i = 0; i <= n - 1; i++)
    {
        cin >> arr[i].start >> arr[i].finish >> arr[i].profit;
    }

    sort(arr, arr + n, comparator);

    vector<long long> dp(n, 0);

    dp[0] = arr[0].profit; // base case
    for (int i = 1; i <= n - 1; i++)
    {
        long long include = arr[i].profit; // including alone
        long long idx = binarySearch(arr, i);
        if (idx != -1)
        {
            include += dp[idx];
        }
        long long exclude = dp[i - 1];

        dp[i] = max(include, exclude);
    }

    cout << dp[n - 1];
}