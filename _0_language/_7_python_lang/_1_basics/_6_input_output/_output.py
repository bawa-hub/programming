# Output
# Syntax of print()
# print(object= separator= end= file= flush=)


    # object - value(s) to be printed
    # sep (optional) - allows us to separate multiple objects inside print().
    # end (optional) - allows us to add add specific values like new line "\n", tab "\t"
    # file (optional) - where the values are printed. It's default value is sys.stdout (screen)
    # flush (optional) - boolean specifying if the output is flushed or buffered. Default: False

print('Good Morning!', end= ' ')
print('It is rainy today')
# Good Morning! It is rainy today

print('New Year', 2023, 'See you soon!', sep= '. ')
# New Year. 2023. See you soon!

# Concatenated Strings
name = 'bawa'
print('Programiz is ' + 'awesome ' + name)
# Programiz is awesome bawa

# Output formatting
x = 5
y = 10
print('The value of x is {} and y is {}'.format(x,y))