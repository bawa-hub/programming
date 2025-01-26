package _0_language._concurrency._6_common_concurrency_problems._starvation;

import java.util.concurrent.locks.ReentrantLock;

class StarvationExample {
    private final ReentrantLock lock = new ReentrantLock(true); // Fair lock

    public void accessResource(String threadName) {
        try {
            lock.lock();
            System.out.println(threadName + " is accessing the resource.");
            Thread.sleep(100);
        } catch (InterruptedException e) {
            e.printStackTrace();
        } finally {
            lock.unlock();
        }
    }

    public static void main(String[] args) {
        StarvationExample example = new StarvationExample();

        Runnable task = () -> {
            for (int i = 0; i < 5; i++) {
                example.accessResource(Thread.currentThread().getName());
            }
        };

        Thread t1 = new Thread(task, "Thread-1");
        Thread t2 = new Thread(task, "Thread-2");
        Thread t3 = new Thread(task, "Thread-3");

        t1.start();
        t2.start();
        t3.start();
    }
}
