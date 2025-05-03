package _3_synchronization_and_shared_resources;

import java.util.concurrent.Semaphore;

class SharedResource {
    private final Semaphore semaphore = new Semaphore(2); // Only 2 threads can access at a time

    public void access() {
        try {
            semaphore.acquire(); // Acquire lock
            System.out.println(Thread.currentThread().getName() + " accessing resource...");
            Thread.sleep(1000); // Simulate work
        } catch (InterruptedException e) {
            e.printStackTrace();
        } finally {
            System.out.println(Thread.currentThread().getName() + " releasing resource.");
            semaphore.release(); // Release lock
        }
    }
}

public class SemaphoreDemo {
    public static void main(String[] args) {
        SharedResource resource = new SharedResource();

        for (int i = 0; i < 5; i++) {
            new Thread(resource::access, "Thread-" + i).start();
        }
    }
}

// Why use a semaphore?

//     Limits concurrent access to resources (e.g., database connections).