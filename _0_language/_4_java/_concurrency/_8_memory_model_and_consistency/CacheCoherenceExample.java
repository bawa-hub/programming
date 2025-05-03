package _8_memory_model_and_consistency;

public class CacheCoherenceExample {
    private static boolean flag = false;
    private static int count = 0;

    public static void main(String[] args) throws InterruptedException {
        Thread writer = new Thread(() -> {
            flag = true;
            count = 1;
        });

        Thread reader = new Thread(() -> {
            while (!flag) { /* Busy-wait */ }
            System.out.println("Count value: " + count);  // Expected: 1, but can be 0 due to stale cache.
        });

        writer.start();
        reader.start();
        writer.join();
        reader.join();
    }
}

// In this case, due to cache coherence issues, Thread 2 might see a stale value of count because the local cache in the processor might not immediately reflect the changes made by Thread 1 in shared memory.