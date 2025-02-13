package _7_advanced_synchronization_primitives._2_latches;

import java.util.concurrent.CountDownLatch;

public class LatchExample {
    public static void main(String[] args) throws InterruptedException {
        final int NUM_THREADS = 3;
        CountDownLatch latch = new CountDownLatch(NUM_THREADS);

        Runnable task = () -> {
            try {
                System.out.println(Thread.currentThread().getName() + " is doing work.");
                Thread.sleep(1000); // Simulating work
                latch.countDown();
                System.out.println(Thread.currentThread().getName() + " completed work.");
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        };

        for (int i = 0; i < NUM_THREADS; i++) {
            new Thread(task).start();
        }

        latch.await(); // Main thread will wait until latch count reaches 0
        System.out.println("All threads have completed their work.");
    }
}

