#  https://leetcode.com/problems/longest-consecutive-sequence/description/

# brute force
def longest_consecutive(nums):
    if not nums:
        return 0

    nums.sort()

    ans = 1
    prev = nums[0]
    cur = 1

    for num in nums[1:]:
        if num == prev + 1:
            cur += 1
        elif num != prev:
            cur = 1
        prev = num
        ans = max(ans, cur)

    return ans

#     Time Complexity: We are first sorting the array which will take O(N * log(N)) time and then we are running a for loop which will take O(N) time. Hence, the overall time complexity will be O(N * log(N)).
#  Space Complexity: The space complexity for the above approach is O(1) because we are not using any auxiliary space

#     optimized
def longest_consecutive(nums):
    hash_set = set(nums)
    longest_streak = 0

    for num in nums:
        if num - 1 not in hash_set:
            current_num = num
            current_streak = 1

            while current_num + 1 in hash_set:
                current_num += 1
                current_streak += 1

            longest_streak = max(longest_streak, current_streak)

    return longest_streak


#  Time Complexity: The time complexity of the above approach is O(N) because we traverse each consecutive subsequence only once. (assuming HashSet takes O(1) to search)
#  Space Complexity: The space complexity of the above approach is O(N) because we are maintaining a HashSet.


# Example usage:
nums = [100, 4, 200, 1, 3, 2]
result = longest_consecutive(nums)
print(result)