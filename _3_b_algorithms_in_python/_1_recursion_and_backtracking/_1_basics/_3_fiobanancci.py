from datetime import datetime

def fib(n):
#base condition
    if n == 0:
        return 0

#Self work
    if n == 1 or n == 2:
        return 1

#Recursion function
    else:
        return fib(n - 1) + fib(n - 2)

#Driver Code
if __name__ == "__main__":
#Get starting timepoint
    start = datetime.now()

#for loop to print the Fibonacci series.
    for i in range(45):
        print(fib(i), end=" ")

#Get ending timepoint
    stop = datetime.now()

    duration = stop - start

    print("\nTime taken by function:", duration)
