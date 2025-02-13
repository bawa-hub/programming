package _3_synchronization_and_shared_resources;

import java.util.concurrent.atomic.AtomicBoolean;

class Spinlock {
    private final AtomicBoolean locked = new AtomicBoolean(false);

    public void lock() {
        while (!locked.compareAndSet(false, true)) {
            // Busy-wait (spinning)
        }
    }

    public void unlock() {
        locked.set(false);
    }
}

// Use Case: Low-latency scenarios where waiting threads should not sleep.