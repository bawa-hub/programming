# for loop in Python

# for val in sequence:
#     loop body

# Program to find the sum of all numbers stored in a list
numbers = [6, 5, 3, 8, 4, 2, 5, 4, 11]
sum = 0
for val in numbers:
    sum = sum+val
print("The sum is", sum) # 48



# The range() function
print(range(10)) # range(0, 10)
print(list(range(10))) # [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
print(list(range(2, 8))) # [2, 3, 4, 5, 6, 7]
print(list(range(2, 20, 3))) # [2, 5, 8, 11, 14, 17]


# Program to iterate through a list using indexing

genre = ['pop', 'rock', 'jazz']
# iterate over the list using index
for i in range(len(genre)):
    print("I like", genre[i])
# I like pop
# I like rock
# ​I like jazz    


# for loop with else
digits = [0, 1, 5]
for i in digits:
    print(i)
else:
    print("No items left.")
# 0
# 1
# 5
# No items left.

#  break keyword can be used to stop a for loop. In such cases, the else part is ignored.
# a for loop's else part runs if no break occurs.


# program to display student's marks from record
student_name = 'Soyuj'
marks = {'James': 90, 'Jules': 55, 'Arthur': 77}
for student in marks:
    if student == student_name:
        print(marks[student])
        break
else:
    print('No entry with that name found.')
# No entry with that name found.    