// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/

class Solution {
public:
    int numOfSubarrays(vector<int>& arr, int k, int threshold) {

        int i=0,j=0,count=0,sum=0;
        int n = arr.size();

        while(j<n) {
            sum+=arr[j];

            if(j-i+1==k) {
               if(sum/k>=threshold) count++;
               sum-=arr[i];
               i++;
            }

            j++;
        }

        return count;
        
    }
};