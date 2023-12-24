#  https://practice.geeksforgeeks.org/problems/top-k-frequent-elements-in-array/1
#  https://leetcode.com/problems/top-k-frequent-elements/

import heapq
from typing import List

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # Create a frequency map
        freq_map = {}
        for num in nums:
            freq_map[num] = freq_map.get(num, 0) + 1

        # Use a min heap
        min_heap = []
        for key, value in freq_map.items():
            heapq.heappush(min_heap, (value, key))
            if len(min_heap) > k:
                heapq.heappop(min_heap)

        # Extract elements from the heap
        result = []
        while min_heap:
            result.append(heapq.heappop(min_heap)[1])

        # Reverse the result to get the elements in descending order of frequency
        return result[::-1]

# Example usage:
nums = [1, 1, 1, 2, 2, 3]
k = 2
solution = Solution()
result = solution.topKFrequent(nums, k)
print(result)
