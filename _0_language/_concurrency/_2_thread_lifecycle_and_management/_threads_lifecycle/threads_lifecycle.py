import threading
import time

def print_message():
    print("Thread is running!")
    time.sleep(2)
    print("Thread has finished!")

thread = threading.Thread(target=print_message)
print("State after creation:", thread.is_alive())

thread.start()  # Start the thread
print("State after start:", thread.is_alive())

thread.join()  # Wait for the thread to finish
print("State after termination:", thread.is_alive())
