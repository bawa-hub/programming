# the global keyword allows us to modify the variable outside of the current scope.
# It is used to create a global variable and make changes to the variable in a local context.


# Access and Modify Python Global Variable

# global variable
c = 1


def add():

    # increment c by 2
    c = c + 2
    print(c)


add()
# UnboundLocalError: local variable 'c' referenced before assignment
# we can only access the global variable but cannot modify it from inside the function.

c = 1


def add():

    # use of global keyword
    global c
    c = c + 2
    print(c)


add()
# Output: 3


# Global in Nested Functions

def outer_function():
    num = 20

    def inner_function():
        global num
        num = 25

    print("Before calling inner_function(): ", num)
    inner_function()
    print("After calling inner_function(): ", num)


outer_function()
print("Outside both function: ", num)
# Before calling inner_function():  20
# After calling inner_function():  20
# Outside both function:  25


# Rules of global Keyword
# When we create a variable inside a function, it is local by default.
# When we define a variable outside of a function, it is global by default. You don't have to use the global keyword.
# We use the global keyword to read and write a global variable inside a function
# Use of the global keyword outside a function has no effect.
