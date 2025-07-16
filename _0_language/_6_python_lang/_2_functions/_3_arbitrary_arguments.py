# Sometimes, we do not know in advance the number of arguments that will be passed into a function.
# To handle this kind of situation, we can use arbitrary arguments in Python.

# use an asterisk (*) before the parameter name to denote this kind of argument


def find_sum(*numbers):
    result = 0

    for num in numbers:
        result = result + num

    print("Sum = ", result)


find_sum(1, 2, 3)
find_sum(4, 9)
# Sum =  6
# Sum =  13
