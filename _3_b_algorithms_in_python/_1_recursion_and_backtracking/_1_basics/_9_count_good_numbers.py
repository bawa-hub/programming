#  https://leetcode.com/problems/count-good-numbers/

mod = 1000000007

class Solution:
    def power(self, x, n):
        if n == 0:
            return 1
        ans = self.power(x, n // 2)
        ans = (ans * ans) % mod
        if n % 2 == 1:
            ans = (ans * x) % mod
        return ans

    def count_good_numbers(self, n):
        numberOfOddPlaces = n // 2
        numberOfEvenPlaces = n // 2 + n % 2

        # from 0-9: there are 5 even num and 4 prime num
        return (self.power(5, numberOfEvenPlaces) * self.power(4, numberOfOddPlaces)) % mod

# Example usage:
sol = Solution()
result = sol.count_good_numbers(10)
print(result)


#  Time Complexity : O(logN).
#  Space Complexity : O(logN), Recursion stack space.