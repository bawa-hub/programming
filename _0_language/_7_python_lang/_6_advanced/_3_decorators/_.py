# https://www.programiz.com/python-programming/decorator


# a decorator is a design pattern that allows you to modify the functionality of a function by wrapping it in another function.
# The outer function is called the decorator, which takes the original function as an argument and returns a modified version of it.

# remember that everything in Python is an object, even functions are objects.

# any object which implements the special __call__() method is termed callable.
# So, in the most basic sense, a decorator is a callable that returns a callable.


def start_end_decorator(func):
    def wrapper():
        print("start")
        func()
        print("end")
    return wrapper    

@start_end_decorator
def printName():
    print("Bawa") 

# equivalent to 
# printName = start_end_decorator(printName) 

printName()  