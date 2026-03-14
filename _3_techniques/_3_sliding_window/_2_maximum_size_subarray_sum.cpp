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

// with negatives
int longestSubarrayWithSumK(vector<int> a, long long k) {
   long long sum = 0;
   unordered_map<int, int> mp;
   int maxi = 0;
    for(int i=0;i<a.size();i++) {
     sum += a[i];

     if(sum == k) maxi = max(maxi, i+1);

     long long rem = sum - k;

     if(mp.find(rem) != mp.end()) maxi = max(maxi, i - mp[rem]);
     if(mp.find(sum) == mp.end()) mp[sum] = i;
    }

    return maxi;
}