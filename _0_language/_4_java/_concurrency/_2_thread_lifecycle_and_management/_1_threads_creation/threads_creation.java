package _0_language._concurrency._2_thread_lifecycle_and_management._threads_creation;

import _0_language._concurrency._2_thread_lifecycle_and_management._threads_lifecycle.MyThread;

// In Java, threads can be created in two ways: by implementing the Runnable interface or extending the Thread class.

// Using Runnable interface
class MyRunnable implements Runnable {
    @Override
    public void run() {
        System.out.println("Thread is running!");
    }

    public static void main(String[] args) {
        MyRunnable myRunnable = new MyRunnable();
        Thread thread = new Thread(myRunnable);
        thread.start();  // Start the thread
    }
}

// Using Thread class
class MyThread extends Thread {
    @Override
    public void run() {
        System.out.println("Thread is running!");
    }

    public static void main(String[] args) {
        MyThread thread = new MyThread();
        thread.start();  // Start the thread
    }
}
