# a generator is a function that returns an iterator that produces a sequence of values when iterated over.
# Generators are useful when we want to produce a large sequence of values, but we don't want to store all of them in memory at once.

# Create Python Generator
def generator_name(arg):
    # statements
    yield something

# When the generator function is called, it does not execute the function body immediately.
# Instead, it returns a generator object that can be iterated over to produce the values.


def my_generator(n):
    value = 0
    while value < n:
        yield value
        value += 1


for value in my_generator(3):
    print(value)

# 0
# 1
# 2


# Python Generator Expression
squares_generator = (i * i for i in range(5))
for i in squares_generator:
    print(i)

# 0
# 1
# 4
# 9
# 16


# Use of Python Generators
# 1. Easy to Implement
class PowTwo:
    def __init__(self, max=0):
        self.n = 0
        self.max = max

    def __iter__(self):
        return self

    def __next__(self):
        if self.n > self.max:
            raise StopIteration

        result = 2 ** self.n
        self.n += 1
        return result


def PowTwoGen(max=0):
    n = 0
    while n < max:
        yield 2 ** n
        n += 1

# 2. Memory Efficient
# A normal function to return a sequence will create the entire sequence in memory before returning the result.
# This is an overkill, if the number of items in the sequence is very large.
# Generator implementation of such sequences is memory friendly and is preferred since it only produces one item at a time.

# 3. Represent Infinite Stream
# 4. Pipelining Generators


def fibonacci_numbers(nums):
    x, y = 0, 1
    for _ in range(nums):
        x, y = y, x+y
        yield x


def square(nums):
    for num in nums:
        yield num**2


print(sum(square(fibonacci_numbers(10))))

# Output: 4895
