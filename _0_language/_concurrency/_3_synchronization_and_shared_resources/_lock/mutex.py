import threading

counter = 0
lock = threading.Lock()

def increment():
    global counter
    with lock:  # Acquire the lock
        for _ in range(1000):
            counter += 1

threads = [threading.Thread(target=increment) for _ in range(2)]

for t in threads:
    t.start()

for t in threads:
    t.join()

print("Final Counter:", counter)
