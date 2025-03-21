class ImmutableObject:
    def __init__(self, value):
        self._value = value

    @property
    def value(self):
        return self._value

immutable = ImmutableObject(42)
print(immutable.value)  # Always thread-safe as value cannot change
