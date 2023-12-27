#  https://takeuforward.org/arrays/find-the-highest-lowest-frequency-element/

#  brute force
def count_freq(arr):
    n = len(arr)
    visited = [False] * n
    max_freq = 0
    min_freq = n
    max_ele = 0
    min_ele = 0

    for i in range(n):
        # Skip this element if already processed
        if visited[i]:
            continue

        # Count frequency
        count = 1
        for j in range(i + 1, n):
            if arr[i] == arr[j]:
                visited[j] = True
                count += 1

        if count > max_freq:
            max_ele = arr[i]
            max_freq = count
        if count < min_freq:
            min_ele = arr[i]
            min_freq = count

    print("The highest frequency element is:", max_ele)
    print("The lowest frequency element is:", min_ele)

# Example usage:
arr = [10, 5, 10, 15, 10, 5]
count_freq(arr)


#  Time Complexity: O(N*N), where N = size of the array. We are using the nested loop to find the frequency.
#  Space Complexity:  O(N), where N = size of the array. It is for the visited array we are using.

#  optimized (using hashmap)

def frequency(arr):
    freq_map = {}

    for num in arr:
        freq_map[num] = freq_map.get(num, 0) + 1

    max_freq = 0
    min_freq = len(arr)
    max_ele = 0
    min_ele = 0

    # Traverse through the map to find the elements.
    for element, count in freq_map.items():
        if count > max_freq:
            max_ele = element
            max_freq = count
        if count < min_freq:
            min_ele = element
            min_freq = count

    print("The highest frequency element is:", max_ele)
    print("The lowest frequency element is:", min_ele)

# Example usage:
arr = [10, 5, 10, 15, 10, 5]
frequency(arr)


#  Time Complexity: O(N), where N = size of the array. The insertion and retrieval operation in the map takes O(1) time.
# Space Complexity:  O(N), where N = size of the array. It is for the map we are using.