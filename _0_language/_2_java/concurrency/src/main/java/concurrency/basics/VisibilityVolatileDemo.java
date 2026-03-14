package concurrency.basics;

/**
 * Demonstrates the visibility problem and how volatile fixes it.
 *
 * How to use:
 * - Run as-is to see the "fixed" version using volatile.
 * - Then remove the volatile keyword and run again: the program may hang forever
 *   or take a very long time to exit because the worker thread may never see
 *   the updated value of running.
 */
public class VisibilityVolatileDemo {

    // Try toggling volatile on/off to observe behavior.
    private static volatile boolean running = true;

    public static void main(String[] args) throws InterruptedException {
        Thread worker = new Thread(() -> {
            System.out.println("Worker started, waiting for running=false...");
            long iterations = 0;
            while (running) {
                iterations++;
                // Optional: small hint to the JIT that the loop has side effects
                if (iterations % 10_000_000 == 0) {
                    Thread.yield();
                }
            }
            System.out.println("Worker stopped after iterations = " + iterations);
        }, "VisibilityWorker");

        worker.start();

        Thread.sleep(1000); // Give worker time to start and loop
        System.out.println("Main setting running=false");
        running = false;

        worker.join();
        System.out.println("Main exiting.");
    }
}

