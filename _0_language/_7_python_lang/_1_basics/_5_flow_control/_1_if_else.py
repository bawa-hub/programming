#  if...elif...else Statement

# if test expression:
#     Body of if
# elif test expression:
#     Body of elif
# else: 
#     Body of else

# Nested if Example

num = float(input("Enter a number: "))
if num >= 0:
    if num == 0:
        print("Zero")
    else:
        print("Positive number")
else:
    print("Negative number")