import threading
import time

def print_message():
    for i in range(5):
        print(f"Thread running: {i}")
        time.sleep(1)  # Simulate some work

# Create and start the thread
thread = threading.Thread(target=print_message)
thread.start()

# Wait for the thread to finish
thread.join()

print("Main thread ends.")
