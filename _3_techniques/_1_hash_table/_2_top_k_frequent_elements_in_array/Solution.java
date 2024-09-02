// https://practice.geeksforgeeks.org/problems/top-k-frequent-elements-in-array/1
// https://leetcode.com/problems/top-k-frequent-elements/

import java.util.*;

class Solution {
    public List<Integer> topKFrequent(int[] nums, int k) {
        // Step 1: Build a frequency map
        Map<Integer, Integer> frequencyMap = new HashMap<>();
        for (int num : nums) {
            frequencyMap.put(num, frequencyMap.getOrDefault(num, 0) + 1);
        }

        // Step 2: Build a min-heap (priority queue)
        PriorityQueue<Map.Entry<Integer, Integer>> pq = new PriorityQueue<>(
            (a, b) -> a.getValue() - b.getValue()
        );

        for (Map.Entry<Integer, Integer> entry : frequencyMap.entrySet()) {
            pq.offer(entry);
            if (pq.size() > k) {
                pq.poll(); // Remove the element with the smallest frequency
            }
        }

        // Step 3: Extract the top k frequent elements
        List<Integer> res = new ArrayList<>();
        while (!pq.isEmpty()) {
            res.add(pq.poll().getKey());
        }

        // Optional: Reverse the list to return the elements in descending order of frequency
        Collections.reverse(res);

        return res;
    }

    public static void main(String[] args) {
        Solution sol = new Solution();
        int[] nums = {1, 1, 1, 2, 2, 3};
        int k = 2;
        List<Integer> topK = sol.topKFrequent(nums, k);
        System.out.println(topK); // Output: [1, 2]
    }
}
