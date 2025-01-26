package _0_language._concurrency._3_synchronization_and_shared_resources._semaphores;

import java.util.concurrent.Semaphore;

public class SemaphoreExample {
    public static void main(String[] args) {
        Semaphore semaphore = new Semaphore(2); // Only 2 threads can access the resource at a time

        Runnable task = () -> {
            try {
                semaphore.acquire();
                System.out.println(Thread.currentThread().getName() + " is accessing the resource.");
                Thread.sleep(1000);
                semaphore.release();
                System.out.println(Thread.currentThread().getName() + " has released the resource.");
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        };

        for (int i = 0; i < 5; i++) {
            new Thread(task).start();
        }
    }
}
