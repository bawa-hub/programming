num_string = '12'
num_integer = 23
print("Data type of num_string before Type Casting:",type(num_string))
# explicit type conversion
num_string = int(num_string)
print("Data type of num_string after Type Casting:",type(num_string))
num_sum = num_integer + num_string
print("Sum:",num_sum)
print("Data type of num_sum:",type(num_sum))
# Data type of num_string before Type Casting: <class 'str'>
# Data type of num_string after Type Casting: <class 'int'>
# Sum: 35
# Data type of num_sum: <class 'int'>

age = 30
print("my age is: " + str(age))
# my age is 30