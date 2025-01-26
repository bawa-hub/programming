# sing an iterator method, we can loop through an object and return its elements.

# Technically, a Python iterator object must implement two special methods, __iter__() and __next__(),
# collectively called the iterator protocol.
# __iter__() returns the iterator object itself. If required, some initialization can be performed.
# __next__() must return the next item in the sequence. On reaching the end, and in subsequent calls, it must raise StopIteration

my_list = [4, 7, 0]
# create an iterator from the list
iterator = iter(my_list)
print(next(iterator))  # prints 4
print(next(iterator))  # prints 7
print(next(iterator))  # prints 0
# When we reach the end and there is no more data to be returned, we will get the StopIteration Exception.

# Using for Loop
my_list = [4, 7, 0]
for element in my_list:
    print(element)
# 4
# 7
# 0

# Building Custom Iterators


class PowTwo:
    """Class to implement an iterator
    of powers of two"""

    def __init__(self, max=0):
        self.max = max

    def __iter__(self):
        self.n = 0
        return self

    def __next__(self):
        if self.n <= self.max:
            result = 2 ** self.n
            self.n += 1
            return result
        else:
            raise StopIteration


numbers = PowTwo(3)
i = iter(numbers)
# Using next to get to the next iterator element
print(next(i))  # prints 1
print(next(i))  # prints 2
print(next(i))  # prints 4
print(next(i))  # prints 8
print(next(i))  # raises StopIteration exception


# Python Infinite Iterators
# An infinite iterator is an iterator that never ends, meaning that it will continue to produce elements indefinitely.

# from itertools import count
# create an infinite iterator that starts at 1 and increments by 1 each time
infinite_iterator = count(1)

# print the first 5 elements of the infinite iterator
for i in range(5):
    print(next(infinite_iterator))
