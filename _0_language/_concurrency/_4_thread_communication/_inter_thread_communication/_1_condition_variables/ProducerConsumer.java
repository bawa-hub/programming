package _4_thread_communication._inter_thread_communication._1_condition_variables;

class SharedData {
    private int counter = 0;

    // Method to produce data
    public synchronized void produce() throws InterruptedException {
        while (counter >= 1) {
            wait();  // Wait for consumer to consume
        }
        counter++;
        System.out.println("Produced: " + counter);
        notify();  // Notify the consumer to consume
    }

    // Method to consume data
    public synchronized void consume() throws InterruptedException {
        while (counter == 0) {
            wait();  // Wait for producer to produce
        }
        counter--;
        System.out.println("Consumed: " + counter);
        notify();  // Notify the producer to produce
    }
}

public class ProducerConsumer {
    public static void main(String[] args) {
        SharedData sharedData = new SharedData();

        // Producer thread
        Thread producer = new Thread(() -> {
            try {
                for (int i = 0; i < 10; i++) {
                    sharedData.produce();
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        // Consumer thread
        Thread consumer = new Thread(() -> {
            try {
                for (int i = 0; i < 10; i++) {
                    sharedData.consume();
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        producer.start();
        consumer.start();
    }
}

// In this example, the producer and consumer threads synchronize using the wait() and notify() methods to avoid producing or consuming too much at once.