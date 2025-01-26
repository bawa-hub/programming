# Python closure is a nested function that allows us to access variables of the outer function even after the outer function is closed.

# Nested function in Python
def greet(name):
    # inner function
    def display_name():
        print("Hi", name)

    # call inner function
    display_name()


# call outer function
greet("John")
# Output: Hi John


# Closures
def greet():
    # variable defined outside the inner function
    name = "John"
    # return a nested anonymous function
    return lambda: "Hi " + name


# call the outer function
message = greet()
# call the inner function
print(message())
# Output: Hi John


def calculate():
    num = 1

    def inner_func():
        nonlocal num
        num += 2
        return num
    return inner_func


# call the outer function
odd = calculate()

# call the inner function
print(odd())
print(odd())
print(odd())

# call the outer function again
odd2 = calculate()
print(odd2())
# 3
# 5
# 7
# 3


# When to use closures?
# Closures can be used to avoid global values and provide data hiding


# All function objects have a __closure__ attribute that returns a tuple of cell objects if it is a closure function.
