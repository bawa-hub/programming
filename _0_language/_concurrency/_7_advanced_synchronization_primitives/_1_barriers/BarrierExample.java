package _7_advanced_synchronization_primitives._1_barriers;

import java.util.concurrent.CyclicBarrier;

public class BarrierExample {
    private static final int NUM_THREADS = 3;
    private static CyclicBarrier barrier = new CyclicBarrier(NUM_THREADS, () -> {
        System.out.println("All threads reached the barrier, proceeding.");
    });

    public static void main(String[] args) {
        Runnable task = () -> {
            try {
                System.out.println(Thread.currentThread().getName() + " is performing work.");
                Thread.sleep(1000); // Simulating work
                System.out.println(Thread.currentThread().getName() + " reached the barrier.");
                barrier.await();
            } catch (Exception e) {
                e.printStackTrace();
            }
        };

        for (int i = 0; i < NUM_THREADS; i++) {
            new Thread(task).start();
        }
    }
}
