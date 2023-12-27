def binary_search(array, x, low, high):
    while low <= high:
        mid = low + (high - low) // 2

        if array[mid] == x:
            return mid

        elif array[mid] < x:
            low = mid + 1

        else:
            high = mid - 1

    return -1

# Example usage:
arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
target = 7
result = binary_search(arr, target, 0, len(arr) - 1)

if result != -1:
    print(f"Element {target} is present at index {result}")
else:
    print(f"Element {target} is not present in the array")

def binary_search_recursive(a, l, r, num):
    if l <= r:
        mid = (l + r) // 2

        if a[mid] == num:
            return mid

        elif a[mid] > num:
            return binary_search_recursive(a, l, mid - 1, num)

        elif a[mid] < num:
            return binary_search_recursive(a, mid + 1, r, num)

    return -1

# Example usage:
arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
target = 7
result = binary_search_recursive(arr, 0, len(arr) - 1, target)

if result != -1:
    print(f"Element {target} is present at index {result}")
else:
    print(f"Element {target} is not present in the array")
