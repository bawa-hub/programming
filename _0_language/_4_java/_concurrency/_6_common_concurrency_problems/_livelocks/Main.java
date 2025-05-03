package _0_language._concurrency._6_common_concurrency_problems._livelocks;

class LivelockExample {
    static class Resource {
        private boolean inUse = false;

        public synchronized boolean use() {
            if (!inUse) {
                inUse = true;
                return true;
            }
            return false;
        }

        public synchronized void release() {
            inUse = false;
        }
    }

    public static void main(String[] args) {
        Resource resource = new Resource();

        Runnable thread1 = () -> {
            while (!resource.use()) {
                System.out.println("Thread 1: Waiting for resource...");
                try { Thread.sleep(50); } catch (InterruptedException e) {}
            }
            System.out.println("Thread 1: Acquired resource!");
        };

        Runnable thread2 = () -> {
            while (!resource.use()) {
                System.out.println("Thread 2: Waiting for resource...");
                try { Thread.sleep(50); } catch (InterruptedException e) {}
            }
            System.out.println("Thread 2: Acquired resource!");
        };

        new Thread(thread1).start();
        new Thread(thread2).start();
    }
}

