# If two superclasses have the same method name and the derived class calls that method,
# Python uses the MRO to search for the right method to call.

#  the MRO specifies that methods should be inherited from the leftmost superclass first

class SuperClass1:
    def info(self):
        print("Super Class 1 method called")


class SuperClass2:
    def info(self):
        print("Super Class 2 method called")


class Derived(SuperClass1, SuperClass2):
    pass


d1 = Derived()
d1.info()

# Output: "Super Class 1 method called"
