package _3_synchronization_and_shared_resources;

class BankAccount {
    private int balance = 1000;

    public synchronized void withdraw(int amount) {
        if (balance >= amount) {
            System.out.println(Thread.currentThread().getName() + " is withdrawing " + amount);
            balance -= amount;
            System.out.println(Thread.currentThread().getName() + " completed withdrawal. Balance: " + balance);
        } else {
            System.out.println("Insufficient funds for " + Thread.currentThread().getName());
        }
    }
}

public class CriticalSectionExample {
    public static void main(String[] args) {
        BankAccount account = new BankAccount();

        Thread t1 = new Thread(() -> account.withdraw(700), "Thread-1");
        Thread t2 = new Thread(() -> account.withdraw(500), "Thread-2");

        t1.start();
        t2.start();
    }
}

// Problem without Synchronization: If both threads withdraw money at the same time, the balance might become negative due to race conditions.
// Solution: We use synchronized to lock the method and allow only one thread at a time.