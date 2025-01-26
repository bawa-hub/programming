f = open("test.txt",'r',encoding = 'utf-8')
f.read(4)    # read the first 4 data
# 'This'
f.read(4)    # read the next 4 data
# ' is '
f.read()     # read in the rest till end of file
# 'my first file\nThis file\ncontains three lines\n'
f.read()  # further reading returns empty sting
# ''

f.tell()    # get the current file position
# 56
f.seek(0)   # bring file cursor to initial position
# 0
print(f.read())  # read the entire file
# This is my first file
# This file
# contains three lines


# We can read a file line-by-line using a for loop. This is both efficient and fast.
for line in f:
    print(line, end='')
# This is my first file
# This file
# contains three lines    


# we can use the readline() method to read individual lines of a file. 
# This method reads a file till the newline, including the newline character.

f.readline()
# 'This is my first file\n'
f.readline()
# 'This file\n'
f.readline()
# 'contains three lines\n'
f.readline()
# ''


# readlines() method returns a list of remaining lines of the entire file. 
f.readlines()
# ['This is my first file\n', 'This file\n', 'contains three lines\n']