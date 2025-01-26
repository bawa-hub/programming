# def function_name(parameters):
# 	"""docstring"""
# 	statement(s)


def greet(name):
    """
    This function greets to
    the person passed in as
    a parameter
    """
    print("Hello, " + name + ". Good morning!")

greet('bawa')    



# Note: In python, the function definition should always be present before the function call. 
# Otherwise, we will get an error.

# The first string after the function header is called the docstring and is short for documentation string. 
# It is briefly used to explain what a function does.

print(greet.__doc__)