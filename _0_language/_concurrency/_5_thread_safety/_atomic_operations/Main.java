package _0_language._concurrency._5_thread_safety._atomic_operations;

import java.util.concurrent.atomic.AtomicInteger;

class AtomicExample {
    private AtomicInteger count = new AtomicInteger(0);

    public void increment() {
        count.incrementAndGet();
    }

    public int getCount() {
        return count.get();
    }

    public static void main(String[] args) throws InterruptedException {
        AtomicExample atomicExample = new AtomicExample();

        Thread t1 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) atomicExample.increment();
        });

        Thread t2 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) atomicExample.increment();
        });

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("Final Count: " + atomicExample.getCount()); // Correct count
    }
}
