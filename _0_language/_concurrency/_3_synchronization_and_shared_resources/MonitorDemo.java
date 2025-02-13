package _3_synchronization_and_shared_resources;

class MonitorExample {
    private int count = 0;

    public synchronized void increment() {
        count++;
        System.out.println(Thread.currentThread().getName() + " - Count: " + count);
    }
}

public class MonitorDemo {
    public static void main(String[] args) {
        MonitorExample obj = new MonitorExample();

        Thread t1 = new Thread(obj::increment, "Thread-1");
        Thread t2 = new Thread(obj::increment, "Thread-2");

        t1.start();
        t2.start();
    }
}

// The synchronized block acts as a monitor, ensuring one thread at a time modifies count.