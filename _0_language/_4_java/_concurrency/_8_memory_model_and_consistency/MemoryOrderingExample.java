package _8_memory_model_and_consistency;

public class MemoryOrderingExample {
    private static volatile boolean flag = false;
    private static int count = 0;

    public static void main(String[] args) throws InterruptedException {
        Thread writer = new Thread(() -> {
            count = 1;
            flag = true;  // This happens-before the read of flag and count
        });

        Thread reader = new Thread(() -> {
            while (!flag) { /* Busy-wait */ }
            System.out.println("Count value: " + count);  // Should always print 1
        });

        writer.start();
        reader.start();
        writer.join();
        reader.join();
    }
}
// Explanation: The volatile keyword ensures that the update to flag is visible to other threads immediately. This prevents reordering of instructions and ensures the visibility of count.

