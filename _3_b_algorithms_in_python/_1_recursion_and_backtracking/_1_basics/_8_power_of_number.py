#  https://www.geeksforgeeks.org/write-a-c-program-to-calculate-powxn/
#  https://leetcode.com/problems/powx-n/

#  brute force using recursion
def power(x, n):
    if n == 0:
        return 1
    return power(x, n - 1) * x


#  brute force iterative
def power(x, n):
    ans = 1.0
    for i in range(n):
        ans *= x
    return ans
#  Time Complexity: O(N)
#  Space Complexity: O(1)

#  using binary exponentiation
def power(x, n):
    ans = 1.0
    nn = abs(n)

    while nn:
        if nn % 2:
            ans *= x
            nn -= 1
        else:
            x *= x
            nn //= 2

    if n < 0:
        ans = 1.0 / ans

    return ans

# Example usage:
result = power(2.0, 10)
print(result)

#  Time Complexity: O(log n)
#  Space Complexity: O(1)
