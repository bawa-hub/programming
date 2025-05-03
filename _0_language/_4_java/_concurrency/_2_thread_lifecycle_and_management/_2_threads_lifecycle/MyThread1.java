package _0_language._concurrency._2_thread_lifecycle_and_management._threads_lifecycle;

class MyThread1 extends Thread {
    @Override
    public void run() {
        try {
            Thread.sleep(1000);  // Simulate some work
            System.out.println("Thread is running!");
        } catch (InterruptedException e) {
            System.out.println("Thread interrupted.");
        }
    }

    public static void main(String[] args) {
        MyThread1 thread = new MyThread();
        
        System.out.println("State after creation: " + thread.getState()); // NEW
        
        thread.start();  // Starts the thread
        
        System.out.println("State after calling start(): " + thread.getState()); // RUNNABLE or TIMED_WAITING (depending on scheduler)
        
        try {
            thread.join();  // Waits for the thread to finish execution
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        
        System.out.println("State after completion: " + thread.getState()); // TERMINATED
    }
}


// Java Thread States:

//     NEW: Thread created but not started.
//     RUNNABLE: Thread is eligible for running, waiting for CPU time.
//     TIMED_WAITING: Thread is in a waiting state but has a specified timeout (like Thread.sleep()).
//     WAITING: Thread is waiting indefinitely (like Object.wait()).
//     TERMINATED: Thread has finished execution.