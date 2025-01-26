# key-value pairs, unordered, mutable

# Create dictionary
#  values can be of any data type and can repeat,
# keys must be of immutable type (string, number or tuple with immutable elements) and must be unique.
dict1 = {}
dict2 = {1: 'bawa', 2: 'vikram'}
dict3 = {'name': 'bawa', 1: [1, 2, 'bawa']}
dict4 = dict({1: 'rahul', 2: 'kumar'})
dict5 = dict([(1, 'apple'), (2, 'ball')])
print(dict1)
print(dict2)  # {1: 'bawa', 2: 'vikram'}
print(dict3)  # {'name': 'bawa', 1: [1, 2, 'bawa']}
print(dict4)  # {1: 'rahul', 2: 'kumar'}
print(dict5)  # {1: 'apple', 2: 'ball'}

# Access elements
my_dict = {'name': 'Jack', 'age': 26}
print(my_dict['name'])  # Jack
print(my_dict.get('age'))  # 26
print(my_dict.get('address'))  # None
print(my_dict['address'])  # KeyError

# Changing and adding Dictionary Elements
my_dict = {'name': 'Jack', 'age': 26}
my_dict['age'] = 27
print(my_dict)  # {'age': 27, 'name': 'Jack'}
my_dict['address'] = 'Downtown'
print(my_dict)  # {'address': 'Downtown', 'age': 27, 'name': 'Jack'}

# Removing elements from a dictionary
squares = {1: 1, 2: 4, 3: 9, 4: 16, 5: 25}
del squares[1]
print(squares)  # {2: 4, 3: 9, 4: 16, 5: 25}
print(squares.pop(4))  # 16
print(squares)  # {2: 4, 3: 9, 5: 25}
print(squares.popitem())  # (5, 25)
squares.clear()  # remove all items
print(squares)  # {}
del squares  # delete the dictionary itself
print(squares)  # Throws Error

# Membership Test for Dictionary Keys
squares = {1: 1, 3: 9, 5: 25, 7: 49, 9: 81}
print(1 in squares)  # True
print(2 not in squares)  # True
print(49 in squares)  # False

# Iterating through a Dictionary
squares = {1: 1, 3: 9, 5: 25, 7: 49, 9: 81}
for i in squares:
    print(squares[i])
