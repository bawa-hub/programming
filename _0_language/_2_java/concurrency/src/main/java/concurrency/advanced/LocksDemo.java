package concurrency.advanced;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

/**
 * Demonstrates using ReentrantLock vs synchronized.
 *
 * How to use:
 * - Run main() and observe that both counters reach the same final value.
 * - Experiment: remove lock.lock()/unlock() or synchronized and see races.
 */
public class LocksDemo {

    private static class LockBasedCounter {
        private final Lock lock = new ReentrantLock();
        private int value = 0;

        public void increment() {
            lock.lock();
            try {
                value++;
            } finally {
                lock.unlock();
            }
        }

        public int getValue() {
            lock.lock();
            try {
                return value;
            } finally {
                lock.unlock();
            }
        }
    }

    private static class SynchronizedCounter {
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
        LockBasedCounter lockCounter = new LockBasedCounter();
        SynchronizedCounter syncCounter = new SynchronizedCounter();

        Thread[] threads = new Thread[THREADS];
        for (int i = 0; i < THREADS; i++) {
            threads[i] = new Thread(() -> {
                for (int j = 0; j < INCREMENTS_PER_THREAD; j++) {
                    lockCounter.increment();
                    syncCounter.increment();
                }
            }, "Worker-" + i);
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
        System.out.println("Expected = " + expected);
        System.out.println("Lock-based counter      = " + lockCounter.getValue());
        System.out.println("Synchronized counter    = " + syncCounter.getValue());
        System.out.println("Time (ms)               = " + (end - start));
    }
}

