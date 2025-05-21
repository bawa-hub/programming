package _2_dsa_java._1_arrays._2_2sum;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// https://leetcode.com/problems/two-sum/

public class _1_2sum {

    // brute force
    public int[] twoSumBruteForce(int[] nums, int target) {
        for (int i = 0; i < nums.length; ++i) {
            for (int j = i + 1; j < nums.length; ++j) {
                if (nums[i] + nums[j] == target) {
                    return new int[] { i, j };
                }
            }
        }
        return new int[0]; // if no solution
    }

    // TC: O(n^2)
    // SC: O(1)

    // two pointer
    public int[] twoSumTwoPointer(int[] nums, int target) {
        int[] store = Arrays.copyOf(nums, nums.length);
        Arrays.sort(store);

        int left = 0, right = store.length - 1;
        int n1 = 0, n2 = 0;
        while (left < right) {
            int sum = store[left] + store[right];
            if (sum == target) {
                n1 = store[left];
                n2 = store[right];
                break;
            } else if (sum > target) {
                right--;
            } else {
                left++;
            }
        }

        List<Integer> result = new ArrayList<>();
        for (int i = 0; i < nums.length; ++i) {
            if (nums[i] == n1 || nums[i] == n2) {
                result.add(i);
                if (result.size() == 2)
                    break;
            }
        }

        return new int[] { result.get(0), result.get(1) };
    }

    // TC: O(nlogn)
    // SC: O(n)

    // hashing
    public int[] twoSumHashing(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; ++i) {
            int complement = target - nums[i];
            if (map.containsKey(complement)) {
                return new int[] { i, map.get(complement) };
            }
            map.put(nums[i], i);
        }
        return new int[0]; // if no solution
    }

    // TC: O(n)
    // SC: O(n)

}
