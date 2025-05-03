package _0_language._concurrency._2_thread_lifecycle_and_management;

public class DaemonExample {
    public static void main(String[] args) {
        Thread daemonThread = new Thread(() -> {
            while (true) {
                System.out.println("Daemon thread is running...");
                try {
                    Thread.sleep(500);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        });

        daemonThread.setDaemon(true); // Set the thread as daemon
        daemonThread.start();

        System.out.println("Main thread is ending...");
        // The application will exit even if the daemon thread is still running
    }
}

// The Main thread is ending... message will be printed, and the program will terminate because the daemon thread does not keep the JVM alive.
