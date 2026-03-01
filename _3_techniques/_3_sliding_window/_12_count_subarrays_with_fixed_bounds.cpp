// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/solutions/2708654/c-easy-explanation-w-example-100-faster-sliding-window/



class Solution {
public:
    long long countSubarrays(vector<int>& nums, int minK, int maxK) 
    {
        int  n = nums.size();
        long long cnt = 0, mini = -1, maxi = -1; 
        int i = 0, j = 0;                        

        while(j < n)                             
        {
            //subarray of no use now, move window
            if(nums[j] < minK || nums[j] > maxK) 
            {
                 //set new mini and maxi of new window
                mini = maxi = -1;  

                //slide the window            
                i = j+1;                         
            }

            //update the index of recently observed minK
            if (nums[j] == minK) mini = j;       

            //update the index of recently observed maxK
			if (nums[j] == maxK) maxi = j;      
             
            //if 2nd part of max is -ve means we don't have minK and maxK in window now so max(0, -ve) = 0, no increment
            cnt += max(0LL, min(mini, maxi) - i + 1);  // think of this as j-i+1

            j++;                                 //keep increasing the window towards right
        }
		return cnt;
    }
};


long long countSubarrays(vector<int>& A, int minK, int maxK) {
        long res = 0, jbad = -1, jmin = -1, jmax = -1, n = A.size();
        for (int i = 0; i < n; ++i) {
            if (A[i] < minK || A[i] > maxK) jbad = i;
            if (A[i] == minK) jmin = i;
            if (A[i] == maxK) jmax = i;
            res += max(0L, min(jmin, jmax) - jbad);
        }
        return res;
    }