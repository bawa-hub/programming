// https://leetcode.com/problems/binary-subarrays-with-sum/
// https://leetcode.com/problems/binary-subarrays-with-sum/solutions/186683/c-java-python-sliding-window-o-1-space/ 

// similar problems
        // Number of Substrings Containing All Three Characters
        // Count Number of Nice Subarrays
        // Replace the Substring for Balanced String
        // Max Consecutive Ones III
        // Binary Subarrays With Sum
        // Subarrays with K Different Integers
        // Fruit Into Baskets
        // Shortest Subarray with Sum at Least K
        // Minimum Size Subarray Sum 

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
//     Space O(N)
// Time O(N)

// sliding window
    int numSubarraysWithSum(vector<int>& A, int S) {
        return atMost(A, S) - atMost(A, S - 1); // for exact k, we do atmost(k) - atmost(k-1)
    }

    int atMost(vector<int>& A, int S) {
        if (S < 0) return 0;
        int res = 0, i = 0, n = A.size();
        for (int j = 0; j < n; j++) {
            S -= A[j];
            while (S < 0)
                S += A[i++];
            res += j - i + 1;
        }
        return res;
    }
//     Space O(1)
// Time O(N)