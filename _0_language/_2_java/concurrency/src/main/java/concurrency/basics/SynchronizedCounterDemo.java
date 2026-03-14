package concurrency.basics;

/**
 * Demonstrates race conditions and how synchronized fixes them.
 *
 * How to use:
 * - Run main() as-is to see the "correct" result with synchronization.
 * - Then comment out the synchronized keyword in increment() and run again
 *   to observe lost updates and inconsistent final counts.
 */
public class SynchronizedCounterDemo {

    static class UnsafeCounter {
        private int value = 0;

        public void increment() {
            // Not synchronized: i++ is a read-modify-write sequence, NOT atomic.
            value++;
        }

        public int getValue() {
            return value;
        }
    }

    static class SafeCounter {
        private int value = 0;

        public synchronized void increment() {
            value++;
        }

        public synchronized int getValue() {
            return value;
        }
    }

    private static final int THREADS = 10;
    private static final int INCREMENTS_PER_THREAD = 100_000;

    public static void main(String[] args) throws InterruptedException {
        runUnsafeExample();
        System.out.println("-----------------------------");
        runSafeExample();
    }

    private static void runUnsafeExample() throws InterruptedException {
        UnsafeCounter counter = new UnsafeCounter();
        Thread[] threads = new Thread[THREADS];

        for (int i = 0; i < THREADS; i++) {
            threads[i] = new Thread(() -> {
                for (int j = 0; j < INCREMENTS_PER_THREAD; j++) {
                    counter.increment();
                }
            }, "Unsafe-" + i);
        }

        long start = System.currentTimeMillis();
        for (Thread t : threads) {
            t.start();
        }
        for (Thread t : threads) {
            t.join();
        }
        long end = System.currentTimeMillis();

        int expected = THREADS * INCREMENTS_PER_THREAD;
        System.out.println("UNSAFE COUNTER:");
        System.out.println("Expected = " + expected);
        System.out.println("Actual   = " + counter.getValue());
        System.out.println("Time (ms)= " + (end - start));
    }

    private static void runSafeExample() throws InterruptedException {
        SafeCounter counter = new SafeCounter();
        Thread[] threads = new Thread[THREADS];

        for (int i = 0; i < THREADS; i++) {
            threads[i] = new Thread(() -> {
                for (int j = 0; j < INCREMENTS_PER_THREAD; j++) {
                    counter.increment();
                }
            }, "Safe-" + i);
        }

        long start = System.currentTimeMillis();
        for (Thread t : threads) {
            t.start();
        }
        for (Thread t : threads) {
            t.join();
        }
        long end = System.currentTimeMillis();

        int expected = THREADS * INCREMENTS_PER_THREAD;
        System.out.println("SAFE COUNTER (synchronized):");
        System.out.println("Expected = " + expected);
        System.out.println("Actual   = " + counter.getValue());
        System.out.println("Time (ms)= " + (end - start));
    }
}

