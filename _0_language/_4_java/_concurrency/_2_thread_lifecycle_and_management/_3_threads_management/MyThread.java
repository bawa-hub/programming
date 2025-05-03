package _0_language._concurrency._2_thread_lifecycle_and_management._threads_management;


// In Java, you can manage threads by:
//     Starting a thread using the start() method.
//     Joining a thread using join(), which makes the calling thread wait until the specified thread finishes.
//     Interrupting a thread using interrupt(), which signals a thread to stop execution.

class MyThread extends Thread {
    @Override
    public void run() {
        for (int i = 0; i < 5; i++) {
            System.out.println("Thread running: " + i);
            try {
                Thread.sleep(500);  // Simulate work by sleeping
            } catch (InterruptedException e) {
                System.out.println("Thread interrupted");
            }
        }
    }

    public static void main(String[] args) {
        MyThread thread = new MyThread();
        
        thread.start();  // Start the thread
        
        try {
            thread.join();  // Wait for the thread to finish
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        
        System.out.println("Main thread ends.");
    }
}

// In this example, the main thread waits for MyThread to complete using join(). The main thread will print its message only after MyThread terminates.