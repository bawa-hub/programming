# while loop in Python

# while test_expression:
#     Body of while


# Program to add natural numbers up to n, sum = 1+2+3+...+n
n = int(input("Enter n: "))
sum = 0
i = 1
while i <= n:
    sum = sum + i
    i = i+1    # update counter
# print the sum
print("The sum is", sum)


# While loop with else
counter = 0
while counter < 3:
    print("Inside loop")
    counter = counter + 1
else:
    print("Inside else")

# Inside loop
# Inside loop
# Inside loop
# Inside else    