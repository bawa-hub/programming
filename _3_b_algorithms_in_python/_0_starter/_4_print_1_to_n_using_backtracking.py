def f(i, n):
    if i < 1:
        return
    f(i - 1, n)
    print(i)

# Example usage:
n = 10
f(n, n)
