from concurrent.futures import ThreadPoolExecutor

def task():
    print("Task executed by:", threading.current_thread().name)

if __name__ == "__main__":
    with ThreadPoolExecutor(max_workers=3) as executor:
        for _ in range(5):
            executor.submit(task)
