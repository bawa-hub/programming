# how to identify dp problem generically ?

count the total number of ways
min/max
trying out all possible ways -> think recursion -> then dp

# memoization

store the value of subproblems in some map/table

# Steps to follow for recurrence relation in DP problems:

a). Try to represent the problem (all states) in terms of index. [(i,j) in case of matrix] -> write base case
b). Do all possible stuffs on that index a/c to problem statement -> write recrurrrence relation
c). Sum of all stuffs -> if ques says count all ways
    minimum( of all stuffs ) -> if ques says find min
d). memoize it -> look at changing parameter and make dp array (1d if one param and 2d if 2 param)

# Recursive (required to base case) to iterative (base case to required) conversion
base case to up == bottom up approach
a). Copy base case
b). Express all states in for loop 
note: if 2 states then nested for loop and accordingly.
c). copy the recurrence and write

Note:
if in top-down approach, base case hits when -1 then apply shifting of index in bottom-up

# space optimized
note: when there is a concept of prev and current think space optimization.
if there is previous variable or row and previous column, we can space optimized it
