// https://practice.geeksforgeeks.org/problems/allocate-minimum-number-of-pages0937/1
// https://www.geeksforgeeks.org/allocate-minimum-number-pages/

using namespace std;

bool isPossible(int arr[], int n, int m, int curr_min)
{
	int studentsRequired = 1;
	int curr_sum = 0;

	for (int i = 0; i < n; i++)
	{

		if (arr[i] > curr_min)
			return false;

		if (curr_sum + arr[i] > curr_min)
		{
			studentsRequired++;

			curr_sum = arr[i];

			if (studentsRequired > m)
				return false;
		}
		else
			curr_sum += arr[i];
	}
	return true;
}

int findPages(int arr[], int n, int m)
{
	long long sum = 0;

	if (n < m)
		return -1;
	int mx = INT_MIN;

	for (int i = 0; i < n; i++)
	{
		sum += arr[i];
		mx = max(mx, arr[i]);
	}

	int start = mx, end = sum;
	int result = INT_MAX;

	while (start <= end)
	{
		int mid = (start + end) / 2;
		if (isPossible(arr, n, m, mid))
		{
			result = mid;
			end = mid - 1;
		}
		else
			start = mid + 1;
	}

	return result;
}

// Time Complexity: O(N*log (M - max)), where N is the number of different books , max denotes the maximum number of pages from all the books and M denotes the sum of number of pages of all different books
// Auxiliary Space: O(1)