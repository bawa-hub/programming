// https://practice.geeksforgeeks.org/problems/largest-subarray-with-0-sum/1

// brute force
int solve(int[] a, int n){
	int  max = 0;
	for(int i = 0; i < n; ++i){
		int sum = 0;
		for(int j = i; j < n; ++j){
			sum += a[j];
			if(sum == 0){
				max = max(max, j-i+1);
			}
		}
	}
	return max;
   }
// Time Complexity: O(N^2) as we have two loops for traversal
// Space Complexity: O(1) as we aren’t using any extra space

int maxLen(int A[], int n)
{
    // Your code here
    unordered_map<int, int> mpp;
    int maxi = 0;
    int sum = 0;
    for (int i = 0; i < n; i++)
    {
        sum += A[i];
        if (sum == 0)
        {
            maxi = i + 1;
        }
        else
        {
            if (mpp.find(sum) != mpp.end())
            {
                maxi = max(maxi, i - mpp[sum]);
            }
            else
            {
                mpp[sum] = i;
            }
        }
    }

    return maxi;
}
// Time Complexity: O(N), as we are traversing the array only once
// Space Complexity: O(N), in the worst case we would insert all array elements prefix sum into our hashmap
