def print_numbers(n):
    if n < 1:
        return
    print(n)
    print_numbers(n - 1)

# Example usage:
n = 10
print_numbers(n)
