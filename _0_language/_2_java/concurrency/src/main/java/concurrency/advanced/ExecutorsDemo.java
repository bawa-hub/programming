package concurrency.advanced;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

/**
 * Demonstrates using ExecutorService and thread pools instead of manually creating threads.
 *
 * How to use:
 * - Run main() and observe which thread names execute tasks.
 * - Change pool size, number of tasks, and see how it affects scheduling.
 */
public class ExecutorsDemo {

    public static void main(String[] args) throws InterruptedException, ExecutionException {
        simpleFixedThreadPoolExample();
        System.out.println("-----------------------------");
        callableExample();
    }

    private static void simpleFixedThreadPoolExample() throws InterruptedException {
        ExecutorService pool = Executors.newFixedThreadPool(3);

        Runnable task = () -> {
            String thread = Thread.currentThread().getName();
            System.out.println("Runnable running on " + thread);
            try {
                Thread.sleep(500);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        };

        for (int i = 0; i < 10; i++) {
            pool.submit(task);
        }

        pool.shutdown(); // No new tasks
        // In real code, use awaitTermination or try-with-resources style helpers.
        while (!pool.isTerminated()) {
            Thread.sleep(100);
        }

        System.out.println("Fixed thread pool example finished.");
    }

    private static void callableExample() throws InterruptedException, ExecutionException {
        ExecutorService pool = Executors.newFixedThreadPool(4);
        List<Callable<Integer>> tasks = new ArrayList<>();

        for (int i = 0; i < 8; i++) {
            final int id = i;
            tasks.add(() -> {
                String thread = Thread.currentThread().getName();
                System.out.println("Callable " + id + " running on " + thread);
                Thread.sleep(300);
                return id * id;
            });
        }

        List<Future<Integer>> futures = pool.invokeAll(tasks);

        for (Future<Integer> f : futures) {
            System.out.println("Result = " + f.get());
        }

        pool.shutdown();
        System.out.println("Callable example finished.");
    }
}

