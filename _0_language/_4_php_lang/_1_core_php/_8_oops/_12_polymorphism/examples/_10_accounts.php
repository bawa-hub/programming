<?php

// Account Interface
interface Account
{
    public function deposit($amount);
    public function withdraw($amount);
    public function getBalance();
    public function display();
}

// SavingsAccount class implementing the interface
class SavingsAccount implements Account
{
    private $balance;

    public function __construct($initialBalance)
    {
        $this->balance = $initialBalance;
    }

    public function deposit($amount)
    {
        $this->balance += $amount;
    }

    public function withdraw($amount)
    {
        if ($this->balance >= $amount) {
            $this->balance -= $amount;
        } else {
            echo "Insufficient funds.\n";
        }
    }

    public function getBalance()
    {
        return $this->balance;
    }

    public function display()
    {
        echo "Savings Account Balance: $this->balance\n";
    }
}

// CheckingAccount class implementing the interface
class CheckingAccount implements Account
{
    private $balance;
    private $overdraftLimit;

    public function __construct($initialBalance, $overdraftLimit)
    {
        $this->balance = $initialBalance;
        $this->overdraftLimit = $overdraftLimit;
    }

    public function deposit($amount)
    {
        $this->balance += $amount;
    }

    public function withdraw($amount)
    {
        if ($this->balance - $amount >= -$this->overdraftLimit) {
            $this->balance -= $amount;
        } else {
            echo "Transaction declined. Exceeds overdraft limit.\n";
        }
    }

    public function getBalance()
    {
        return $this->balance;
    }

    public function display()
    {
        echo "Checking Account Balance: $this->balance\n";
    }
}

// CreditCardAccount class implementing the interface
class CreditCardAccount implements Account
{
    private $balance;
    private $creditLimit;

    public function __construct($initialBalance, $creditLimit)
    {
        $this->balance = $initialBalance;
        $this->creditLimit = $creditLimit;
    }

    public function deposit($amount)
    {
        $this->balance += $amount;
    }

    public function withdraw($amount)
    {
        if ($this->balance + $this->creditLimit >= $amount) {
            $this->balance -= $amount;
        } else {
            echo "Transaction declined. Exceeds credit limit.\n";
        }
    }

    public function getBalance()
    {
        return $this->balance;
    }

    public function display()
    {
        echo "Credit Card Account Balance: $this->balance\n";
    }
}

// Bank class using polymorphism for managing accounts
class Bank
{
    private $accounts = [];

    public function addAccount(Account $account)
    {
        $this->accounts[] = $account;
    }

    public function displayAccounts()
    {
        foreach ($this->accounts as $account) {
            $account->display();
            echo "-------------------\n";
        }
    }
}

// Client code using polymorphism to add accounts and perform transactions
$bank = new Bank();

// Add accounts
$savingsAccount = new SavingsAccount(1000);
$checkingAccount = new CheckingAccount(2000, 500);
$creditCardAccount = new CreditCardAccount(0, 1000);

$bank->addAccount($savingsAccount);
$bank->addAccount($checkingAccount);
$bank->addAccount($creditCardAccount);

// Perform transactions
$savingsAccount->deposit(500);
$savingsAccount->withdraw(200);

$checkingAccount->withdraw(2500); // This will be declined due to overdraft limit

$creditCardAccount->withdraw(800);
$creditCardAccount->withdraw(1200); // This will be declined due to credit limit

// Display account balances
$bank->displayAccounts();


// The Account interface defines common methods (deposit, withdraw, getBalance, display) that all account types must implement.
// Concrete classes (SavingsAccount, CheckingAccount, CreditCardAccount) implement this interface with their specific transaction handling logic.
// The Bank class uses polymorphism by allowing the dynamic addition of different types of accounts through the addAccount method and displaying them through the displayAccounts method.

// This example demonstrates how polymorphism allows you to manage different types of accounts uniformly within a banking system. The Bank class doesn't need to know the specific type of each account; it can manage and display them using a common interface.