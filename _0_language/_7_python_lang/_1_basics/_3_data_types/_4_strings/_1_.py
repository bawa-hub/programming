name = "Python"
print(name)  # Python
message = "I love Python."
print(message)  # I love Python.

# String Length
greet = "Hello"
len(greet)  # 5

# Access String Characters
greet = 'hello'
print(greet[1])  # "e"
print(greet[-4])  # "e"
print(greet[1:4])  # "ell"

# In Python, strings are immutable. That means the characters of a string cannot be changed
message = 'Hola Amigos'
message[0] = 'H'
print(message)  # TypeError: 'str' object does not support item assignment
message = 'Hello Friends'
prints(message)  # prints "Hello Friends"

# Multiline String
message = """
Never gonna give you up
Never gonna let you down
"""
print(message)
# Never gonna give you up
# Never gonna let you down

# Compare Two Strings
str1 = "Hello, world!"
str2 = "I love Python."
str3 = "Hello, world!"
print(str1 == str2)  # False
print(str1 == str3)  # True

# Join Two or More Strings
greet = "Hello, "
name = "Jack"
result = greet + name
print(result)  # Hello, Jack

# Iterate Through String
greet = 'Hello'
for letter in greet:
    print(letter)
# H
# e
# l
# l
# o

# Membership Test
print('a' in 'program')  # True
print('at' not in 'battle')  # False

# String Formatting (f-Strings)
name = 'Cathy'
country = 'UK'
print(f'{name} is from {country}')  # Cathy is from UK

# Escape Sequences
example = "He said, \"What's there?\""  # escape double quotes
example = 'He said, "What\'s there?"'  # escape single quotes
print(example)  # He said, "What's there?"


# slicing (indexing[] or slice())
# [start:stop:step]
name = "Bawa Vikram"
fname = name[0:4]
# fname = name[:4] # by default start with 0
lname = name[5:11]
# lname = name[5:] # means upto last index
print(fname)
# Bawa
print(lname)
# Vikram
funny_name = name[0:11:2]
print(funny_name)
# Bw irm
reversed_name = name[::-1]
print(reversed_name)
# markiV awaB