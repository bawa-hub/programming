f = open("test.txt")    # open file in current directory
# f = open("C:/Python38/README.txt")  # specifying full path


# default is reading in text mode. 
# In this mode, we get strings when reading from the file.

# Mode
# 					Description
			
# r
# 					Opens a file for reading. (default)
			
# w
# 					Opens a file for writing. Creates a new file if it does not exist or truncates the file if it exists.
			
# x
# 					Opens a file for exclusive creation. If the file already exists, the operation fails.
			
# a
# 					Opens a file for appending at the end of the file without truncating it. Creates a new file if it does not exist.
			
# t
# 					Opens in text mode. (default)
			
# b
# 					Opens in binary mode.
			
# +
# 					Opens a file for updating (reading and writing)



# when working with files in text mode, it is highly recommended to specify the encoding type.
f = open("test.txt", mode='r', encoding='utf-8')