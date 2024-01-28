// https://practice.geeksforgeeks.org/problems/max-sum-subarray-of-size-k5313/1

import java.util.*;

class Solution {
    public static long maximumSumSubarray(int K, ArrayList<Integer> Arr, int N) {
        int i = 0, j = 0;

        long sum = 0, maxi = Long.MIN_VALUE;

        while (j < N) {
            sum += Arr.get(j);
            if (j - i + 1 == K) {
                maxi = Math.max(maxi, sum);
                sum -= Arr.get(i);
                i++;
            }
            j++;
        }

        return maxi;
    }
}
