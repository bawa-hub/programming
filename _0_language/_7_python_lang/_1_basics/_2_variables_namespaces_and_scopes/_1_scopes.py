# we can declare variables in three different scopes: local scope, global, and nonlocal scope.

#  Local Variables
# We cannot access them outside the function.

def greet():
    message = 'Hello'
    print('Local', message)


greet()
print(message)
# Local Hello
# NameError: name 'message' is not defined


# Global Variables
# global variable can be accessed inside or outside of the function.

message = 'Hello'


def greet():
    # declare local variable
    print('Local', message)


greet()
print('Global', message)
# Local Hello
# Global Hello


# Nonlocal Variables
# nonlocal variables are used in nested functions whose local scope is not defined.
# This means that the variable can be neither in the local nor the global scope.

def outer():
    message = 'local'

    def inner():
        # declare nonlocal variable
        nonlocal message
        message = 'nonlocal'
        print("inner:", message)
    inner()
    print("outer:", message)


outer()
# inner: nonlocal
# outer: nonlocal

#  If we change the value of a nonlocal variable, the changes appear in the local variable.
