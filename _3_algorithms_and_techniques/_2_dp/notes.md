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

# Recursive (required to base case) to iterative (base case to required) conversion

a). Declare base case
b). Express all states in for loop
c). copy the recurrence and write

# space optimized

if there is previous row and previous column, we can space optimized it
