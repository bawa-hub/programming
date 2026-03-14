// 02_encapsulation_access.cpp
// Encapsulation and access specifiers

#include <iostream>
#include <string>

class BankAccount {
private:
    std::string owner;
    double balance{};

public:
    BankAccount(const std::string& ownerName, double initialBalance)
        : owner(ownerName), balance(initialBalance) {}

    void deposit(double amount) {
        if (amount > 0) {
            balance += amount;
        }
    }

    bool withdraw(double amount) {
        if (amount > 0 && amount <= balance) {
            balance -= amount;
            return true;
        }
        return false;
    }

    double getBalance() const {
        return balance;
    }

    std::string getOwner() const {
        return owner;
    }
};

int main() {
    BankAccount acc("Vikram", 1000.0);
    acc.deposit(250.0);
    bool ok = acc.withdraw(500.0);

    std::cout << "Owner: " << acc.getOwner() << '\n';
    std::cout << "Withdraw success: " << std::boolalpha << ok << '\n';
    std::cout << "Balance: " << acc.getBalance() << '\n';

    // acc.balance = -1000.0; // ERROR if uncommented: balance is private

    return 0;
}

