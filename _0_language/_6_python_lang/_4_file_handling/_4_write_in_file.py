with open("test.txt",'w',encoding = 'utf-8') as f:
   f.write("my first file\n")
   f.write("This file\n\n")
   f.write("contains three lines\n")

# program will create a new file named test.txt in the current directory if it does not exist. 
# If it does exist, it is overwritten.
