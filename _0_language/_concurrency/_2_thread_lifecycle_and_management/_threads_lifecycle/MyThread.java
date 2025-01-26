package _0_language._concurrency._2_thread_lifecycle_and_management._threads_lifecycle;


public class MyThread extends Thread {
    @Override
    public void run() {
        System.out.println("Thread is running!");
    }

    public static void main(String[] args) {
        MyThread thread = new MyThread();
        System.out.println("State after creation: " + thread.getState());
        
        thread.start();  // Start the thread
        System.out.println("State after start: " + thread.getState());

        try {
            thread.join();  // Wait for the thread to finish
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("State after termination: " + thread.getState());
    }
}
