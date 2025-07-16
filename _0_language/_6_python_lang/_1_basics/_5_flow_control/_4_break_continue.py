# break statement terminates the loop containing it
# If the break statement is inside a nested loop, 
# the break statement will terminate the innermost loop.

# Use of break statement inside the loop
for val in "string":
    if val == "i":
        break
    print(val)
print("The end")
# s
# t
# r
# The end


# continue statement is used to skip the rest of the code inside a loop for the current iteration only.

# Program to show the use of continue statement inside loops
for val in "string":
    if val == "i":
        continue
    print(val)
print("The end")
# s
# t
# r
# n
# g
# The end