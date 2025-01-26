# collections: Counter, namedtuple, OrderedDict, defaultdict, deque

from collections import Counter

a = "aaabbbccccdd"
myCounter = Counter(a)
print(myCounter)
# Counter({'c': 4, 'a': 3, 'b': 3, 'd': 2})