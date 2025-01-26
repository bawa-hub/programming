f = open("test.txt", encoding='utf-8')
# perform file operations
f.close()

# If an exception occurs when we are performing some operation with the file, the code exits without closing the file.

# A safer way is to use a try...finally block.
try:
    f = open("test.txt", encoding='utf-8')
    # perform file operations
finally:
    f.close()


# The best way to close a file is by using the with statement.
with open("test.txt", encoding='utf-8') as f:
    # perform file operations
