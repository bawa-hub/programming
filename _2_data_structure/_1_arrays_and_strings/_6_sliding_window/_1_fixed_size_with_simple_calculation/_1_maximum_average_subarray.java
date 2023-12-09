// https://leetcode.com/problems/maximum-average-subarray-i/

import java.util.*;

class Solution {
    public double findMaxAverage(ArrayList<Integer> nums, int k) {
        int i = 0, j = 0, n = nums.size();

        double sum = 0;
        double maxi = Double.NEGATIVE_INFINITY;

        while (j < n) {
            sum += nums.get(j);

            if (j - i + 1 == k) {
                double avg = sum / k;
                maxi = Math.max(maxi, avg);
                sum -= nums.get(i);
                i++;
            }

            j++;
        }

        return maxi;
    }
}
