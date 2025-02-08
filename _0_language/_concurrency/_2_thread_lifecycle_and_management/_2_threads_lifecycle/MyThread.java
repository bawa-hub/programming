package _0_language._concurrency._2_thread_lifecycle_and_management._2_threads_lifecycle;


// public class MyThread extends Thread {
//     @Override
//     public void run() {
//         System.out.println("Thread is running!");
//     }

//     public static void main(String[] args) {
//         MyThread thread = new MyThread();
//         System.out.println("State after creation: " + thread.getState());
        
//         thread.start();  // Start the thread
//         System.out.println("State after start: " + thread.getState());

//         try {
//             thread.join();  // Wait for the thread to finish
//         } catch (InterruptedException e) {
//             e.printStackTrace();
//         }

//         System.out.println("State after termination: " + thread.getState());
//     }
// }

class MyThread extends Thread {
    @Override
    public void run() {
        System.out.println("State in run(): " + Thread.currentThread().getState());

        try {
            // Move to TIMED_WAITING
            Thread.sleep(500);
            System.out.println("State after sleep: " + Thread.currentThread().getState());

            synchronized (this) {
                System.out.println("Thread going into WAITING state...");
                wait(); // Move to WAITING
                System.out.println("Thread resumed from WAITING...");
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) {
        MyThread thread = new MyThread();
        
        System.out.println("State after creation: " + thread.getState()); // NEW

        thread.start();
        System.out.println("State after start: " + thread.getState()); // RUNNABLE

        try {
            Thread.sleep(600); // Ensure `sleep()` has executed before checking
            System.out.println("State during sleep: " + thread.getState()); // TIMED_WAITING

            Thread.sleep(100); // Give time for thread to enter WAITING state
            System.out.println("State before notify (should be WAITING): " + thread.getState());

            synchronized (thread) {
                thread.notify(); // Wake up thread from WAITING
            }

            thread.join(); // Wait for thread to finish
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("State after termination: " + thread.getState()); // TERMINATED
    }
}

// Output:
// State after creation: NEW
// State after start: RUNNABLE
// State in run(): RUNNABLE
// State after sleep: RUNNABLE
// Thread going into WAITING state...
// State during sleep: WAITING
// State before notify (should be WAITING): WAITING
// Thread resumed from WAITING...
// State after termination: TERMINATED


// why State after start: comes before State in run() ?

/**
 * 
 * Explanation

    thread.start() is called in main()
        This does not immediately run the run() method.
        Instead, it moves the thread to the RUNNABLE state, meaning it's eligible for execution but waiting for the CPU.

    JVM schedules the new thread (MyThread)
        The main thread prints "State after start: RUNNABLE", because at this point, the new thread has not yet started running—it is just waiting in the RUNNABLE state.

    The CPU switches to the new thread (MyThread)
        Now, run() starts executing, and it prints "State in run(): RUNNABLE".

Key Point: Why This Happens?

    start() does not immediately call run().
    Thread scheduling is handled by the JVM and OS, which means there's a slight delay between calling start() and the actual execution of run().
    Meanwhile, the main thread continues execution without waiting, so it prints the state before run() begins.

Order of Execution

🔹 Thread main executes first:

    Calls start()
    Prints "State after start: RUNNABLE"
    Moves to the next statement

🔹 Thread MyThread starts running later:

    JVM schedules it
    Executes run()
    Prints "State in run(): RUNNABLE"

Can This Order Change?

Yes! Thread scheduling is not deterministic. In rare cases, if the new thread is scheduled immediately after start(), you might see "State in run(): RUNNABLE" before "State after start: RUNNABLE". However, in most cases, the main thread executes first.

 */