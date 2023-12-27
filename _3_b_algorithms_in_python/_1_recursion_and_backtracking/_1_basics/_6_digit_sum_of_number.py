
# digit_sum(n) = digit_sum(n/10) + n%10;

def digit_sum(n):
    if n == 0:
        return 0
    return digit_sum(n // 10) + (n % 10)

# Example usage:
n = int(input("Enter a number: "))
result = digit_sum(n)
print(result)
