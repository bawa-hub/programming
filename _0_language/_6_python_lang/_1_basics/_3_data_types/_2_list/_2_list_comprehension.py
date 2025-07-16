# List comprehension is an elegant and concise way to create a new list from an existing list.
# https://www.programiz.com/python-programming/list-comprehension

pow2 = [2 ** x for x in range(10)]
print(pow2)  # [1, 2, 4, 8, 16, 32, 64, 128, 256, 512]
# equivalent to
# pow2 = []
# for x in range(10):
#    pow2.append(2 ** x)

pow2 = [2 ** x for x in range(10) if x > 5]
print(pow2)  # [64, 128, 256, 512]

odd = [x for x in range(20) if x % 2 == 1]
print(odd)  # [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]

l = [x+y for x in ['Python ', 'C '] for y in ['Language', 'Programming']]
# ['Python Language', 'Python Programming', 'C Language', 'C Programming']
print(l)
