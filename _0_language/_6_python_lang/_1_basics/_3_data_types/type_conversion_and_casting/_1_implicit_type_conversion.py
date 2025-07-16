# Python promotes the conversion of the lower data type (integer) to the higher data type (float) to avoid data loss
integer_number = 123
float_number = 1.23
new_number = integer_number + float_number
# display new value and resulting data type
print("Value:",new_number)
print("Data Type:",type(new_number))
# Value: 124.23
# Data Type: <class 'float'>