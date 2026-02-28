// https://practice.geeksforgeeks.org/problems/check-if-an-array-is-sorted0701/1

class Solution
{
public:
    bool arraySortedOrNot(int arr[], int n)
    {
        for (int i = 1; i < n; i++)
        {
            if (arr[i] < arr[i - 1])
                return false;
        }
        return true;
    }
};