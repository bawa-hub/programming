// https://leetcode.com/problems/binary-subarrays-with-sum/
// https://leetcode.com/problems/binary-subarrays-with-sum/solutions/186683/c-java-python-sliding-window-o-1-space/ 

// sliding window
    int numSubarraysWithSum(vector<int>& A, int S) {
        return atMost(A, S) - atMost(A, S - 1); // for exact k, we do atmost(k) - atmost(k-1)
    }

    int atMost(vector<int>& nums, int sum) {
        if (sum < 0) return 0;
        int i=0,j=0,n=nums.size(),res=0;

        while(j<n) {
            sum-=nums[j];

            while(sum<0) {
                sum+=nums[i];
                i++;
            }

            res+=(j-i+1);
            j++;
        }
        return res;
    }
// Space O(1)
// Time O(N)

// using hashmap
    int numSubarraysWithSum(vector<int>& A, int S) {
        unordered_map<int, int> c({{0, 1}});
        int psum = 0, res = 0;
        for (int i : A) {
            psum += i;
            res += c[psum - S];
            c[psum]++;
        }
        return res;
    }
// Space O(N)
// Time O(N)