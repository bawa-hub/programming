import threading
import os
from multiprocessing import Process

# Thread task
def thread_task():
    print("This is running in a thread.")

# Process task
def process_task():
    print(f"Process ID: {os.getpid()}")

if __name__ == '__main__':
    # Example of Thread
    thread = threading.Thread(target=thread_task)
    thread.start()
    thread.join()

    # Example of Process
    process = Process(target=process_task)
    process.start()
    process.join()



# Why the Change is Needed:

#     The spawn method restarts the interpreter to execute the child process, and Python must be able to re-import the main module safely.
#     Without the if __name__ == '__main__': block, the child process would re-execute the entire script, leading to issues.
    
#     Additional Notes:

#     Always use if __name__ == '__main__': for scripts that involve multiprocessing when running on platforms like macOS and Windows.
#     The threading module does not have this requirement because threads share the same memory space as the main process.

