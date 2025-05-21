package _2_dsa_java._1_arrays._2_2sum;
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

public class _4_2sum_2 {
    
}

class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int i = 0, j = numbers.length - 1;

        while (i < j) {
            int sum = numbers[i] + numbers[j];

            if (sum == target) {
                return new int[]{i + 1, j + 1}; // 1-based indexing
            } else if (sum < target) {
                i++;
            } else {
                j--;
            }
        }

        return new int[0]; // if no solution
    }
}
