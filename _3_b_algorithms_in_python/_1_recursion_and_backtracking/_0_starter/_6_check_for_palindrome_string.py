#  iterative
def is_palindrome(s):
    left, right = 0, len(s) - 1
    while left < right:
        if not s[left].isalnum():
            left += 1
        elif not s[right].isalnum():
            right -= 1
        elif s[left].lower() != s[right].lower():
            return False
        else:
            left += 1
            right -= 1
    return True

# Example usage:
input_str = "A man, a plan, a canal, Panama"
result = is_palindrome(input_str)
print("Is palindrome:", result)

#  Time Complexity:  O(N)
#  Space Complexity: O(1)

# recursive

def is_palindrome_recursive(i, s):
    # Base Condition
    # If i exceeds half of the string, all elements are compared, return true.
    if i >= len(s) // 2:
        return True

    # If the start is not equal to the end, it's not a palindrome.
    if s[i] != s[len(s) - i - 1]:
        return False

    # If both characters are the same, increment i and check start+1 and end-1.
    return is_palindrome_recursive(i + 1, s)

# Example usage:
input_str = "madam"
result = is_palindrome_recursive(0, input_str)
print("Is palindrome:", result)

#  Time Complexity: O(N) { Precisely, O(N/2) as we compare the elements N/2 times and swap them}.
#  Space Complexity: O(1) { The elements of the given array are swapped in place so no extra space is required}.