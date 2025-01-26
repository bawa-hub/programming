package _0_language._concurrency._2_thread_lifecycle_and_management._threads_pool;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class ThreadPoolExample {
    public static void main(String[] args) {
        ExecutorService executor = Executors.newFixedThreadPool(3); // Pool with 3 threads

        Runnable task = () -> {
            System.out.println(Thread.currentThread().getName() + " is executing a task");
        };

        for (int i = 0; i < 5; i++) {
            executor.execute(task);
        }

        executor.shutdown(); // Shutdown the pool after tasks are completed
    }
}

// The tasks will be executed by the three threads in the pool.