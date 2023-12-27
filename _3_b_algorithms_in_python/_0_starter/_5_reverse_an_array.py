def print_array(ans):
    print("The reversed array is:- ")
    for num in ans:
        print(num, end=" ")

#  using extra array
def reverse_array(arr):
    n = len(arr)
    ans = [0] * n
    for i in range(n - 1, -1, -1):
        ans[n - i - 1] = arr[i]
    print_array(ans)

def print_array(ans):
    print("The reversed array is:- ")
    for num in ans:
        print(num, end=" ")


#  Time Complexity: O(n), single-pass for reversing array.
# Space Complexity: O(n), for the extra array used.

#  Space-optimized iterative method
def reverse_array(arr):
    p1, p2 = 0, len(arr) - 1
    while p1 < p2:
        arr[p1], arr[p2] = arr[p2], arr[p1]
        p1 += 1
        p2 -= 1
    print_array(arr)

def print_array(arr):
    print("The reversed array is:- ")
    for num in arr:
        print(num, end=" ")

#  Time Complexity: O(n), single-pass involved.
#  Space Complexity: O(1)

#   Recursive method
def reverse_array(arr, start, end):
    if start < end:
        arr[start], arr[end] = arr[end], arr[start]
        reverse_array(arr, start + 1, end - 1)
#  Time Complexity: O(n)
#  Space Complexity: O(1)

#  Using library function
def reverse_array(arr):
    arr.reverse()

