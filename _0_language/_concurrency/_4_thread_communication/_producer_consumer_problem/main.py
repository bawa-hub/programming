import threading
import queue
import time

buffer = queue.Queue(maxsize=5)

def producer():
    for i in range(10):
        buffer.put(i)
        print(f"Produced: {i}")
        time.sleep(0.5)

def consumer():
    while True:
        item = buffer.get()
        print(f"Consumed: {item}")
        buffer.task_done()
        time.sleep(1)

threading.Thread(target=producer).start()
threading.Thread(target=consumer, daemon=True).start()

time.sleep(10)  # Allow time for processing
