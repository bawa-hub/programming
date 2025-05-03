package _0_language._concurrency._2_thread_lifecycle_and_management;

public class ThreadPriorityExample {
    public static void main(String[] args) {
        Thread highPriority = new Thread(() -> System.out.println("High-priority thread running"));
        Thread lowPriority = new Thread(() -> System.out.println("Low-priority thread running"));

        highPriority.setPriority(Thread.MAX_PRIORITY); // Set priority to 10
        lowPriority.setPriority(Thread.MIN_PRIORITY); // Set priority to 1

        lowPriority.start();
        highPriority.start();
    }
}

