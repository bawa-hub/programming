#  iterative
def my_atoi(s):
    sign, base, i = 1, 0, 0

    # Skip leading whitespaces
    while i < len(s) and s[i] == ' ':
        i += 1

    # Check for sign
    if i < len(s) and (s[i] == '-' or s[i] == '+'):
        sign = 1 - 2 * (s[i] == '-')
        i += 1

    # Process digits
    while i < len(s) and s[i].isdigit():
        digit = int(s[i])
        if base > (2**31 - 1) // 10 or (base == (2**31 - 1) // 10 and digit > 7):
            return 2**31 - 1 if sign == 1 else -2**31
        base = 10 * base + digit
        i += 1

    return base * sign

# Example usage:
s = "   -42"
result = my_atoi(s)
print(result)

