# In Python, threads can be created using the threading module.

import threading

def print_message():
    print("Thread is running!")

# Create a thread
thread = threading.Thread(target=print_message)
thread.start()  # Start the thread
