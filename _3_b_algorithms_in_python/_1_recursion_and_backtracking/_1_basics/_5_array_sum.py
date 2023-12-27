
#  sum of elements upto index n-1 of array a of length n
def sum_of_array(n, a):
    if n < 0:
        return 0
    return sum_of_array(n - 1, a) + a[n]

# Example usage:
n = int(input("Enter the size of the array: "))
a = [int(input()) for _ in range(n)]
result = sum_of_array(n - 1, a)
print(result)
