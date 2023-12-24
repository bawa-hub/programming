#  https://leetcode.com/problems/find-the-duplicate-number/


def find_duplicate(arr):
    n = len(arr)
    freq = [0] * (n + 1)

    for num in arr:
        if freq[num] == 0:
            freq[num] += 1
        else:
            return num

    return 0

# Example usage:
arr = [2, 1, 1]
print("The duplicate element is", find_duplicate(arr))
