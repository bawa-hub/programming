// https://leetcode.com/problems/count-number-of-nice-subarrays/

    int numberOfSubarrays(vector<int>& A, int k) {
        return atMost(A, k) - atMost(A, k - 1);
    }


    int atMost(vector<int>& A, int k) {
        int i = 0,j=0,res=0,n = A.size();

        while(j<n) {
            if(A[j]%2!=0) k--;

            while(k<0) {
                if(A[i]%2!=0) k++;
                i++;
            }

            res+=(j-i+1);
            j++;
        }
        return res;
    }

//     Time O(N) for one pass
// Space O(1)

// note:
// To clarify, adding the length of the array does not directly give the count of subarrays. 
// Instead, it provides a way to estimate the upper bound on the count of subarrays. 
// Let me illustrate this with a proof.

// Let's say we have an array of length n. 
// To calculate the total number of subarrays, we can consider the following:

//     Length of Subarrays:
//     The shortest subarray contains one element, and the longest subarray spans the entire array. 
//     The length of the shortest subarray is 1, and the length of the longest subarray is n. 
//     All other subarrays have lengths between these extremes.
//     Counting Approach:
//     We can consider each element in the array and determine how many subarrays can be formed starting from that element. 
//     For example, if we start from the first element, we can form n subarrays (one for each possible length from 1 to n). 
//     If we start from the second element, we can form n-1 subarrays, and so on.

//     Total Count of Subarrays:
//     By summing up the counts of subarrays starting from each element in the array, we can estimate the total count of subarrays.

//     Let's calculate this sum:

//     Subarrays starting from the first element contribute n subarrays.
//     Subarrays starting from the second element contribute n - 1 subarrays.
//     Subarrays starting from the third element contribute n - 2 subarrays.
//     ...
//     Subarrays starting from the (n-1)th element contribute 2 subarrays.
//     Subarrays starting from the nth element contribute 1 subarray.

// Summing up these contributions: n+(n−1)+(n−2)+…+2+1
// This is the sum of the first n positive integers, which is known as the triangular number and is given by the formula n(n+1)/2​.
// Therefore, the total count of subarrays in an array of length n can be estimated as n(n+1)/2​.
// However, it's essential to note that this calculation provides an upper bound on the count of subarrays, as it considers all possible subarrays, including those that may not meet the given conditions (e.g., having k odd numbers). In practice, when dealing with constraints such as the number of odd elements in a subarray, the actual count of subarrays may be less than this upper bound.


// three pointer
    int numberOfSubarrays(vector<int>& A, int k) {
        int res = 0, i = 0, count = 0, n = A.size();
        for (int j = 0; j < n; j++) {
            if (A[j] & 1)
                --k, count = 0;
            while (k == 0)
                k += A[i++] & 1, ++count;
            res += count;
        }
        return res;
    }
//     Time O(N) for one pass
// Space O(1)