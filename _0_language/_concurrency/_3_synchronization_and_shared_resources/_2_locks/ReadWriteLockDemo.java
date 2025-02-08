import java.util.concurrent.locks.ReentrantReadWriteLock;

class ReadWriteExample {
    private int value = 0;
    private final ReentrantReadWriteLock lock = new ReentrantReadWriteLock();

    public void write(int newValue) {
        lock.writeLock().lock();
        try {
            value = newValue;
            System.out.println(Thread.currentThread().getName() + " updated value to: " + value);
        } finally {
            lock.writeLock().unlock();
        }
    }

    public void read() {
        lock.readLock().lock();
        try {
            System.out.println(Thread.currentThread().getName() + " read value: " + value);
        } finally {
            lock.readLock().unlock();
        }
    }
}

public class ReadWriteLockDemo {
    public static void main(String[] args) {
        ReadWriteExample obj = new ReadWriteExample();

        Thread writer = new Thread(() -> obj.write(42), "Writer");
        Thread reader1 = new Thread(obj::read, "Reader-1");
        Thread reader2 = new Thread(obj::read, "Reader-2");

        writer.start();
        reader1.start();
        reader2.start();
    }
}

// Why use ReadWriteLock?

//     Improves performance by allowing concurrent readers.
//     Writers must wait until all readers have finished.