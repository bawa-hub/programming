using namespace std;

int binarySearch(int arr[], int start,
                 int end, int x)
{
    bool isAsc = arr[start] < arr[end];
    if (end >= start)
    {
        int middle = start + (end - start) / 2;

        if (arr[middle] == x)
            return middle;

        if (isAsc == true)
        {
            if (arr[middle] > x)
                return binarySearch(
                    arr, start,
                    middle - 1, x);

            return binarySearch(arr, middle + 1,
                                end, x);
        }
        else
        {
            if (arr[middle] < x)
                return binarySearch(arr, start,
                                    middle - 1, x);

            return binarySearch(arr, middle + 1,
                                end, x);
        }
    }

    return -1;
}
