package concurrency.advanced;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;

/**
 * Demonstrates basic CompletableFuture usage and composition.
 *
 * How to use:
 * - Run main() and follow the printed thread names and completion order.
 * - Modify delays and thread pool usage to see how it affects behavior.
 */
public class CompletableFutureDemo {

    public static void main(String[] args) throws ExecutionException, InterruptedException {
        simplePipeline();
        System.out.println("-----------------------------");
        combineTwoTasks();
    }

    private static void simplePipeline() throws ExecutionException, InterruptedException {
        System.out.println("Starting simple pipeline on thread " + Thread.currentThread().getName());

        CompletableFuture<Integer> future = CompletableFuture.supplyAsync(() -> {
            System.out.println("supplyAsync running on " + Thread.currentThread().getName());
            sleep(500);
            return 42;
        }).thenApply(result -> {
            System.out.println("thenApply on " + Thread.currentThread().getName());
            return result * 2;
        }).thenApply(result -> {
            System.out.println("second thenApply on " + Thread.currentThread().getName());
            return "Answer = " + result;
        });

        String finalResult = future.get();
        System.out.println("Final result: " + finalResult);
    }

    private static void combineTwoTasks() throws ExecutionException, InterruptedException {
        CompletableFuture<Integer> slow = CompletableFuture.supplyAsync(() -> {
            System.out.println("slow task on " + Thread.currentThread().getName());
            sleep(800);
            return 10;
        });

        CompletableFuture<Integer> fast = CompletableFuture.supplyAsync(() -> {
            System.out.println("fast task on " + Thread.currentThread().getName());
            sleep(300);
            return 5;
        });

        CompletableFuture<Integer> combined = slow.thenCombine(fast, (a, b) -> {
            System.out.println("Combining on " + Thread.currentThread().getName());
            return a + b;
        });

        System.out.println("Combined result = " + combined.get());
    }

    private static void sleep(long millis) {
        try {
            TimeUnit.MILLISECONDS.sleep(millis);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }
}

