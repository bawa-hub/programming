package _0_language._concurrency._3_synchronization_and_shared_resources._lock;

import java.util.concurrent.locks.ReentrantReadWriteLock;

public class ReadWriteLockExample {
    private int data = 0;
    private final ReentrantReadWriteLock lock = new ReentrantReadWriteLock();

    public void write(int value) {
        lock.writeLock().lock();
        try {
            data = value;
        } finally {
            lock.writeLock().unlock();
        }
    }

    public int read() {
        lock.readLock().lock();
        try {
            return data;
        } finally {
            lock.readLock().unlock();
        }
    }
}

