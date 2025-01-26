package _0_language._concurrency._5_thread_safety._volatile_keyword;

class VolatileExample {
    private volatile boolean running = true;

    public void stop() {
        running = false;
    }

    public void run() {
        while (running) {
            System.out.println("Running...");
        }
        System.out.println("Stopped.");
    }

    public static void main(String[] args) throws InterruptedException {
        VolatileExample example = new VolatileExample();

        Thread thread = new Thread(example::run);
        thread.start();

        Thread.sleep(1000);
        example.stop();
    }
}

