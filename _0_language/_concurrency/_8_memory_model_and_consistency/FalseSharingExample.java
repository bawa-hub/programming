package _8_memory_model_and_consistency;

public class FalseSharingExample {
    private static class Data {
        volatile long a, b, c, d, e, f, g, h;
    }

    private static Data data = new Data();

    public static void main(String[] args) throws InterruptedException {
        Runnable task1 = () -> {
            for (int i = 0; i < 1000000; i++) {
                data.a = i;
            }
        };

        Runnable task2 = () -> {
            for (int i = 0; i < 1000000; i++) {
                data.b = i;
            }
        };

        // Both threads operate on variables a and b, which may cause false sharing
        Thread t1 = new Thread(task1);
        Thread t2 = new Thread(task2);

        t1.start();
        t2.start();
        t1.join();
        t2.join();
    }
}

// Explanation: Even though Thread 1 and Thread 2 operate on different variables, because the variables are on the same cache line, frequent cache invalidations can occur, resulting in performance bottlenecks.