# Create List
my_list = []  # empty list
my_list = list() # empty list
my_list = [1, "Hello", 3.4]  # list with mixed data types
my_list = ["mouse", [8, 4, 6], ['a']]  # nested list

# Length of List
languages = ['Python', 'Swift', 'C++']
len(languages)  # 3

# Access elements
my_list = ['p', 'r', 'o', 'b', 'e']
print(my_list[0])  # p
print(my_list[2])  # o
print(my_list[4])  # e
print(my_list[4.0])  # Error! Only integer can be used for indexing

# Negative indexing
print(my_list[-1])  # e
print(my_list[-5])  # p

# If the specified index does not exist in the list, Python throws the IndexError exception.

# List slicing
my_list = ['p', 'r', 'o', 'g', 'r', 'a', 'm', 'i', 'z']
print(my_list[2:5])  # ['o', 'g', 'r']
print(my_list[5:])  # ['a', 'm', 'i', 'z']
print(my_list[:])  # ['p', 'r', 'o', 'g', 'r', 'a', 'm', 'i', 'z']
#  When we slice lists, the start index is inclusive but the end index is exclusive

# Change List Items

# Python lists are mutable. Meaning lists are changeable.
# And, we can change items of a list by assigning new values using = operator
odd = [2, 4, 6, 8]
odd[0] = 1
print(odd)  # [1, 4, 6, 8]
odd[1:4] = [3, 5, 7]
print(odd)  # [1, 3, 5, 7]

# Add Elements to a Python List

# add one item to a list using the append() method
# or add several items using the extend() method.
odd = [1, 3, 5]
odd.append(7)
print(odd)  # [1, 3, 5, 7]
odd.extend([9, 11, 13])
print(odd)  # [1, 3, 5, 7, 9, 11, 13]

# Concatenating and repeating lists
odd = [1, 3, 5]
print(odd + [9, 7, 5])  # [1, 3, 5, 9, 7, 5]
print(["re"] * 3)  # ['re', 're', 're']

# Demonstration of list insert() method
odd = [1, 9]
odd.insert(1, 3)
print(odd)  # [1, 3, 9]
odd[2:2] = [5, 7]
print(odd)  # [1, 3, 5, 7, 9]

# Deleting list items
my_list = ['p', 'r', 'o', 'b', 'l', 'e', 'm']
del my_list[2]
print(my_list)  # ['p', 'r', 'b', 'l', 'e', 'm']
del my_list[1:5]
print(my_list)  # ['p', 'm']
del my_list  # delete the entire list
print(my_list)  # Error: List not defined

my_list = ['p', 'r', 'o', 'b', 'l', 'e', 'm']
my_list.remove('p')
print(my_list)  # ['r', 'o', 'b', 'l', 'e', 'm']
print(my_list.pop(1))  # 'o'
print(my_list)  # ['r', 'b', 'l', 'e', 'm']
print(my_list.pop())  # 'm'
print(my_list)  # ['r', 'b', 'l', 'e']
my_list.clear()
print(my_list)  # []


my_list = ['p', 'r', 'o', 'b', 'l', 'e', 'm']
my_list[2:3] = []
print(my_list)  # ['p', 'r', 'b', 'l', 'e', 'm']
my_list[2:5] = []
print(my_list)  # ['p', 'r', 'm']

# Iterating Through a List
for fruit in ['apple', 'banana', 'mango']:
    print("I like", fruit)
# I like apple
# I like banana
# I like mango

# List membership test
my_list = ['p', 'r', 'o', 'b', 'l', 'e', 'm']
print('p' in my_list)  # True
print('a' in my_list)  # False
print('c' not in my_list)  # True
