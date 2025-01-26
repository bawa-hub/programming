import os

# Get Current Directory
print(os.getcwd())
print(os.getcwdb()) # get it as bytes object.

# Changing Directory
os.chdir('/home/bawa/Learn/python_world')
print(os.getcwd())


# List Directories and Files
print(os.listdir())
print(os.listdir('/home/bawa'))

# Making a New Directory
os.mkdir('test')
print(os.listdir())

# Renaming a Directory or a File
os.rename('test','new_one')
print(os.listdir())

# Removing Directory or File
os.remove('test.txt') # remove file
os.rmdir('new_one') # remove dir
# Note: The rmdir() method can only remove empty directories.

# remove non-empty dir
import shutil
shutil.rmtree('test')