# ordered, immutable, allow duplicate elements

# Different types of tuples

# Empty tuple
my_tuple = ()
print(my_tuple)

# Tuple having integers
my_tuple = (1, 2, 3)
my_tuple = 1,2,3 # paranthesis is optional
print(my_tuple)

# tuple with mixed datatypes
my_tuple = (1, "Hello", 3.4)
print(my_tuple)

# nested tuple
my_tuple = ("mouse", [8, 4, 6], (1, 2, 3))
print(my_tuple)

# Create a Python Tuple With one Element
var1 = ("Hello") # string
var2 = ("Hello",) # tuple,  trailing comma to indicate that it is a tuple
# Parentheses is optional
var3 = "hello",
print(type(var3))  # <class 'tuple'>

# Access Python Tuple Elements
# 1. Indexing
letters = ("p", "r", "o", "g", "r", "a", "m", "i", "z")
print(letters[0])   # prints "p" 
print(letters[5])   # prints "a"

# 2. Negative Indexing
letters = ('p', 'r', 'o', 'g', 'r', 'a', 'm', 'i', 'z')
print(letters[-1])   # prints 'z' 
print(letters[-3])   # prints 'm'

# 3. Slicing
my_tuple = ('p', 'r', 'o', 'g', 'r', 'a', 'm', 'i', 'z')
# elements 2nd to 4th index
print(my_tuple[1:4])  #  prints ('r', 'o', 'g')
# elements beginning to 2nd
print(my_tuple[:-7]) # prints ('p', 'r')
# elements 8th to end
print(my_tuple[7:]) # prints ('i', 'z')
# elements beginning to end
print(my_tuple[:]) # Prints ('p', 'r', 'o', 'g', 'r', 'a', 'm', 'i', 'z')

# Note: When we slice lists, the start index is inclusive but the end index is exclusive.

# Iterating through a Tuple
languages = ('Python', 'Swift', 'C++')
# iterating through the tuple
for language in languages:
    print(language)
    
# Check if an Item Exists
languages = ('Python', 'Swift', 'C++')
print('C' in languages)    # False
print('Python' in languages)    # True 



# Advantages of Tuple over List in Python
# We generally use tuples for heterogeneous (different) data types and lists for homogeneous (similar) data types.
# Since tuples are immutable, iterating through a tuple is faster than with a list. So there is a slight performance boost.
# Tuples that contain immutable elements can be used as a key for a dictionary. With lists, this is not possible.
# If you have data that doesn't change, implementing it as tuple will guarantee that it remains write-protected.