package _0_language._concurrency._4_thread_communication._shared_memory_communicatioin;

class SharedMemoryExample {
    private int sharedValue = 0;

    public synchronized void increment() {
        sharedValue++;
    }

    public synchronized int getValue() {
        return sharedValue;
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        SharedMemoryExample sharedMemory = new SharedMemoryExample();

        Thread t1 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) sharedMemory.increment();
        });

        Thread t2 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) sharedMemory.increment();
        });

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("Final Value: " + sharedMemory.getValue());
    }
}

