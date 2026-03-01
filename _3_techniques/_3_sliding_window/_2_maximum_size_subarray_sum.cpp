// https://www.codingninjas.com/studio/problems/longest-subarray-with-sum-k_6682399

// Note: this will not work for negative numbers

#include <vector>
using namespace std;


int longestSubarrayWithSumK(vector<int> a, long long k) {
    int i=0,j=0;
    long long sum = 0;
    int maxi = 0;

    while(j < a.size()) {
      sum += a[j];

          while (sum > k) {
                sum -= a[i];
                 i++;
            }

      if(sum == k) {
          maxi = max(maxi, j-i+1);
      }

      j++;
    }

    return maxi;
}