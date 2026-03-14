package concurrency.advanced;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentLinkedQueue;

/**
 * Demonstrates basic usage of concurrent collections.
 *
 * How to use:
 * - Run main() and inspect output to see concurrent writes and reads.
 * - Add more threads and operations to stress-test.
 */
public class ConcurrentCollectionsDemo {

    public static void main(String[] args) throws InterruptedException {
        concurrentHashMapExample();
        System.out.println("-----------------------------");
        concurrentQueueExample();
    }

    private static void concurrentHashMapExample() throws InterruptedException {
        Map<String, Integer> map = new ConcurrentHashMap<>();

        Runnable writer = () -> {
            for (int i = 0; i < 1000; i++) {
                map.put("key-" + i, i);
            }
        };

        Thread t1 = new Thread(writer, "Writer-1");
        Thread t2 = new Thread(writer, "Writer-2");

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("ConcurrentHashMap size = " + map.size());
        System.out.println("Sample value for key-500 = " + map.get("key-500"));
    }

    private static void concurrentQueueExample() throws InterruptedException {
        ConcurrentLinkedQueue<Integer> queue = new ConcurrentLinkedQueue<>();

        Thread producer = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                queue.offer(i);
            }
        }, "Producer");

        Thread consumer = new Thread(() -> {
            int polled = 0;
            while (polled < 1000) {
                Integer value = queue.poll();
                if (value != null) {
                    polled++;
                }
            }
            System.out.println("Consumer polled " + polled + " items");
        }, "Consumer");

        producer.start();
        consumer.start();

        producer.join();
        consumer.join();

        System.out.println("ConcurrentLinkedQueue size after = " + queue.size());
    }
}

