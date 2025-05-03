package _7_advanced_synchronization_primitives._4_semaphore;

import java.util.concurrent.Semaphore;

public class SemaphoreExample {
    public static void main(String[] args) throws InterruptedException {
        Semaphore semaphore = new Semaphore(2); // Allows 2 threads to access at a time

        Runnable task = () -> {
            try {
                semaphore.acquire();
                System.out.println(Thread.currentThread().getName() + " acquired the semaphore.");
                Thread.sleep(1000); // Simulate work
            } catch (InterruptedException e) {
                e.printStackTrace();
            } finally {
                semaphore.release();
                System.out.println(Thread.currentThread().getName() + " released the semaphore.");
            }
        };

        for (int i = 0; i < 5; i++) {
            new Thread(task).start();
        }
    }
}

