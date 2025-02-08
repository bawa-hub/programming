package _0_language._concurrency._3_synchronization_and_shared_resources._synchronization;

// In Java, you can use the synchronized keyword to ensure that only one thread can execute a method or block of code at a time.

class Counter {
    private int count = 0;

    // Synchronize method to ensure only one thread can increment the count
    public synchronized void increment() {
        count++;
    }

    public int getCount() {
        return count;
    }
}

public class Main {
    public static void main(String[] args) {
        Counter counter = new Counter();

        // Create two threads that increment the counter
        Thread t1 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                counter.increment();
            }
        });

        Thread t2 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                counter.increment();
            }
        });

        t1.start();
        t2.start();

        try {
            t1.join();
            t2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("Final count: " + counter.getCount());  // Expected output: 2000
    }
}


// In this example:
    // We use synchronized on the increment() method to prevent two threads from simultaneously modifying the count variable, ensuring that the final count is correct.
    // We use synchronized to lock the method and allow only one thread at a time.