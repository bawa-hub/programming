package _0_language._concurrency._3_synchronization_and_shared_resources._lock;

public class MutexExample {
    private int counter = 0;

    public synchronized void increment() {
        counter++;
    }

    public int getCounter() {
        return counter;
    }

    public static void main(String[] args) {
        MutexExample example = new MutexExample();

        Thread t1 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                example.increment();
            }
        });

        Thread t2 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                example.increment();
            }
        });

        t1.start();
        t2.start();

        try {
            t1.join();
            t2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("Final Counter: " + example.getCounter());
    }
}

// Output:

//     Without synchronization, the counter may have an incorrect value due to race conditions.
//     With synchronization, the correct result (2000) is guaranteed.