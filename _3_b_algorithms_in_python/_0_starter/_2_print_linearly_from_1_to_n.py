def print_numbers(i, n):
    if i > n:
        return
    print(i)
    print_numbers(i + 1, n)

# Example usage:
n = 10
i = 1
print_numbers(i, n)
