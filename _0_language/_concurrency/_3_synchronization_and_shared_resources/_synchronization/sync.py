import threading

# In Python, the threading module provides the Lock class to prevent multiple threads from accessing critical sections at the same time.

class Counter:
    def __init__(self):
        self.count = 0
        self.lock = threading.Lock()

    def increment(self):
        with self.lock:  # Acquire lock to ensure only one thread can increment at a time
            self.count += 1

    def get_count(self):
        return self.count

def increment_counter(counter):
    for _ in range(1000):
        counter.increment()

counter = Counter()

# Create two threads to increment the counter
thread1 = threading.Thread(target=increment_counter, args=(counter,))
thread2 = threading.Thread(target=increment_counter, args=(counter,))

thread1.start()
thread2.start()

thread1.join()
thread2.join()

print("Final count:", counter.get_count())  # Expected output: 2000


# We use a Lock to synchronize access to the increment() method, ensuring that only one thread can modify the count at a time.
