import threading
import time

event = threading.Event()

def producer():
    print("Producer preparing data...")
    time.sleep(2)
    print("Producer: Data ready!")
    event.set()  # Signal the event

def consumer():
    print("Consumer waiting for data...")
    event.wait()  # Wait for the signal
    print("Consumer: Data received!")

threading.Thread(target=producer).start()
threading.Thread(target=consumer).start()
