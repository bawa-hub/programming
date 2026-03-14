package concurrency.advanced;

/**
 * Demonstrates a simple deadlock scenario and how to fix it.
 *
 * How to use:
 * - Run main() and observe that the program may hang (deadlock) in deadlockExample().
 * - Then comment out deadlockExample() and uncomment orderedLockingExample()
 *   to see how consistent lock ordering avoids deadlock.
 */
public class DeadlockDemo {

    private static final Object lockA = new Object();
    private static final Object lockB = new Object();

    public static void main(String[] args) throws InterruptedException {
        deadlockExample();
        // orderedLockingExample();
    }

    private static void deadlockExample() throws InterruptedException {
        Thread t1 = new Thread(() -> {
            synchronized (lockA) {
                System.out.println("T1 acquired lockA");
                sleep(200);
                System.out.println("T1 trying to acquire lockB");
                synchronized (lockB) {
                    System.out.println("T1 acquired lockB");
                }
            }
        }, "Deadlock-T1");

        Thread t2 = new Thread(() -> {
            synchronized (lockB) {
                System.out.println("T2 acquired lockB");
                sleep(200);
                System.out.println("T2 trying to acquire lockA");
                synchronized (lockA) {
                    System.out.println("T2 acquired lockA");
                }
            }
        }, "Deadlock-T2");

        t1.start();
        t2.start();

        t1.join(2000);
        t2.join(2000);

        System.out.println("If program has not exited, threads are likely deadlocked.");
    }

    private static void orderedLockingExample() throws InterruptedException {
        Thread t1 = new Thread(() -> acquireLocksInOrder("T1"), "Ordered-T1");
        Thread t2 = new Thread(() -> acquireLocksInOrder("T2"), "Ordered-T2");

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("Ordered locking example completed without deadlock.");
    }

    private static void acquireLocksInOrder(String name) {
        Object first = lockA;
        Object second = lockB;

        synchronized (first) {
            System.out.println(name + " acquired lockA");
            sleep(100);
            synchronized (second) {
                System.out.println(name + " acquired lockB");
            }
        }
    }

    private static void sleep(long millis) {
        try {
            Thread.sleep(millis);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }
}

