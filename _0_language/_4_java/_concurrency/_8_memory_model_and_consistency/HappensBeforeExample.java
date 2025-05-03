package _8_memory_model_and_consistency;

public class HappensBeforeExample {
    private static volatile boolean flag = false;

    public static void main(String[] args) throws InterruptedException {
        Thread writer = new Thread(() -> {
            System.out.println("Thread 1: Writing to flag");
            flag = true;  // Write operation happens-before subsequent reads
        });

        Thread reader = new Thread(() -> {
            while (!flag) {  // Read operation will see the change made by writer
                // Busy-wait loop to check flag
            }
            System.out.println("Thread 2: Detected flag change!");
        });

        writer.start();
        reader.start();
        writer.join();
        reader.join();
    }
}


// Explanation: The volatile keyword ensures that the write to flag in Thread 1 happens-before the read in Thread 2, meaning Thread 2 will always see the updated value of flag once it's written.