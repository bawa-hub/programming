def print_string(name, n):
    if n == 0:
        return
    print(name)
    print_string(name, n - 1)

# Example usage:
name = "Bawa"
n = 10
print_string(name, n)


