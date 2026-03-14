package concurrency.basics;

/**
 * Demonstrates the three most common ways to create and run threads in Java.
 *
 * How to use:
 * - Run this class' main method.
 * - Observe the interleaving of output from different threads.
 * - Change thread names, add sleep, and see how scheduling behaves.
 */
public class ThreadCreationDemo {

    /**
     * 1) Extending Thread
     */
    static class MyThread extends Thread {
        public MyThread(String name) {
            super(name);
        }

        @Override
        public void run() {
            for (int i = 0; i < 5; i++) {
                System.out.println(Thread.currentThread().getName() + " (extends Thread) - i=" + i);
                sleepQuietly(100);
            }
        }
    }

    /**
     * 2) Implementing Runnable
     */
    static class MyRunnable implements Runnable {
        private final String name;

        MyRunnable(String name) {
            this.name = name;
        }

        @Override
        public void run() {
            for (int i = 0; i < 5; i++) {
                System.out.println(name + " (implements Runnable) - i=" + i
                        + ", actual thread=" + Thread.currentThread().getName());
                sleepQuietly(100);
            }
        }
    }

    private static void sleepQuietly(long millis) {
        try {
            Thread.sleep(millis);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            System.out.println("Thread interrupted: " + Thread.currentThread().getName());
        }
    }

    public static void main(String[] args) {
        System.out.println("Main thread: " + Thread.currentThread().getName());

        // 1) Using a class that extends Thread
        Thread t1 = new MyThread("ExtThread-1");

        // 2) Using a Runnable implementation
        Thread t2 = new Thread(new MyRunnable("Runnable-1"), "RunnableThread-1");

        // 3) Using a lambda (Runnable) – common in modern Java
        Thread t3 = new Thread(() -> {
            for (int i = 0; i < 5; i++) {
                System.out.println(Thread.currentThread().getName() + " (lambda Runnable) - i=" + i);
                sleepQuietly(100);
            }
        }, "LambdaThread-1");

        // Start threads
        t1.start();
        t2.start();
        t3.start();

        // Join threads to wait for completion before exiting main
        try {
            t1.join();
            t2.join();
            t3.join();
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }

        System.out.println("All threads finished. Main exiting.");
    }
}

