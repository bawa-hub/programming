# https://www.programiz.com/python-programming/modules
# https://www.programiz.com/python-programming/package

# https://docs.python.org/3/py-modindex.html

# see all the built in modules

# using help
help('modules')

# using dir
import sys
modules = dir(sys.modules)
print(modules)

import message
import message as m
from message import bye,hello
from message import *