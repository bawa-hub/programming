// lower_bound
int lower_bound(int arr[], int low,
                int high, int X)
{
    if (low > high)
    {
        return low;
    }
    int mid = low + (high - low) / 2;

    if (arr[mid] >= X)
    {
        return lower_bound(arr, low,
                           mid - 1, X);
    }

    return lower_bound(arr, mid + 1,
                       high, X);
}

// upper_bound
int upper_bound(int arr[], int low,
                int high, int X)
{
    if (low > high)
        return low;
    int mid = low + (high - low) / 2;

    if (arr[mid] <= X)
    {
        return upper_bound(arr, mid + 1,
                           high, X);
    }

    return upper_bound(arr, low,
                       mid - 1, X);
}
