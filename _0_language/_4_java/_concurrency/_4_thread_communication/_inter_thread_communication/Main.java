package _0_language._concurrency._4_thread_communication._inter_thread_communication;

class Message {
    private String content;
    private boolean hasMessage = false;

    public synchronized void write(String message) throws InterruptedException {
        while (hasMessage) {
            wait(); // Wait until the message is read
        }
        this.content = message;
        hasMessage = true;
        notifyAll(); // Notify waiting threads
    }

    public synchronized String read() throws InterruptedException {
        while (!hasMessage) {
            wait(); // Wait until a message is written
        }
        hasMessage = false;
        notifyAll(); // Notify waiting threads
        return content;
    }
}

public class Main {
    public static void main(String[] args) {
        Message message = new Message();

        Thread writer = new Thread(() -> {
            try {
                message.write("Hello from Writer!");
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        Thread reader = new Thread(() -> {
            try {
                System.out.println("Reader received: " + message.read());
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        writer.start();
        reader.start();
    }
}

