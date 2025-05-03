import threading

thread_local = threading.local()

def worker():
    thread_local.value = threading.current_thread().name
    print(f"{threading.current_thread().name} has value {thread_local.value}")

thread1 = threading.Thread(target=worker)
thread2 = threading.Thread(target=worker)

thread1.start()
thread2.start()
