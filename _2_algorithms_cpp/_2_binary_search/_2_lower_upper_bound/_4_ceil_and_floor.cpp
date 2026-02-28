// https://practice.geeksforgeeks.org/problems/ceil-the-floor2802/1

// valid in sorted array
// ceil = smallest num in array >= x [same as lower_bound]
// floor = largest num in array <= x

// same as lower_bound
int ceil(int arr[], int n, int x)
{
    int l = 0, r = n - 1;
    int ans = -1;

    while (l <= r)
    {
        int mid = l + (r - l) / 2;

        if (arr[mid] >= x)
        {
            ans = arr[mid];
            r = mid - 1;
        }
        else
        {
            l = mid + 1;
        }
    }

    return ans;
}

int floor(int arr[], int n, int x)
{
    int l = 0, r = n - 1;
    int ans = -1;

    while (l <= r)
    {
        int mid = l + (r - l) / 2;

        if (arr[mid] <= x)
        {
            ans = arr[mid];
            l = mid + 1;
        }
        else
        {
            r = mid - 1;
        }
    }

    return ans;
}