temp = int(input("What is the temperature ? "))

if temp >=0 and temp <=30:
    print("good")
elif temp < 0 or temp > 30:
    print("bad")       
    
if not(temp > 30):
    print("Go outside")
# What is the temperature ? 20
# good 
# Go outside    