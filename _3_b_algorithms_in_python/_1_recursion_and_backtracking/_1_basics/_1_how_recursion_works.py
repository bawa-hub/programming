def print_recursive(num):
    # Base condition
    if num < 1:
        return

    # Self work
    print(num)

    # Recursive call
    print_recursive(num - 1)

# Example usage:
print_recursive(50)

