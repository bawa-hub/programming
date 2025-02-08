import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

class MutexExample {
    private int count = 0;
    private final Lock lock = new ReentrantLock();

    public void increment() {
        lock.lock();
        try {
            count++;
            System.out.println(Thread.currentThread().getName() + " - Count: " + count);
        } finally {
            lock.unlock();
        }
    }
}

public class MutexLockDemo {
    public static void main(String[] args) {
        MutexExample obj = new MutexExample();
        Runnable task = obj::increment;

        Thread t1 = new Thread(task, "Thread-1");
        Thread t2 = new Thread(task, "Thread-2");

        t1.start();
        t2.start();
    }
}
