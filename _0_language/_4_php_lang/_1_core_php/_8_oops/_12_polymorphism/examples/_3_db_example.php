<?php

// Interface for Database operations
interface DatabaseInterface
{
    public function connect();
    public function query($sql);
    public function disconnect();
}

// MySQL Database class implementing the interface
class MySQLDatabase implements DatabaseInterface
{
    public function connect()
    {
        echo "Connected to MySQL database.\n";
        // MySQL-specific connection logic
    }

    public function query($sql)
    {
        echo "Executing MySQL query: $sql\n";
        // MySQL-specific query execution logic
    }

    public function disconnect()
    {
        echo "Disconnected from MySQL database.\n";
        // MySQL-specific disconnection logic
    }
}

// PostgreSQL Database class implementing the interface
class PostgreSQLDatabase implements DatabaseInterface
{
    public function connect()
    {
        echo "Connected to PostgreSQL database.\n";
        // PostgreSQL-specific connection logic
    }

    public function query($sql)
    {
        echo "Executing PostgreSQL query: $sql\n";
        // PostgreSQL-specific query execution logic
    }

    public function disconnect()
    {
        echo "Disconnected from PostgreSQL database.\n";
        // PostgreSQL-specific disconnection logic
    }
}

// Client code using polymorphism to switch between databases
function performDatabaseOperations(DatabaseInterface $database, $sql)
{
    $database->connect();
    $database->query($sql);
    $database->disconnect();
}

// Example of using MySQL
$mysqlDatabase = new MySQLDatabase();
performDatabaseOperations($mysqlDatabase, "SELECT * FROM users");

// Example of using PostgreSQL
$postgresDatabase = new PostgreSQLDatabase();
performDatabaseOperations($postgresDatabase, "SELECT * FROM employees");


// In this example:

//     The DatabaseInterface defines a set of common methods (connect, query, and disconnect) that any database class must implement.
//     The MySQLDatabase and PostgreSQLDatabase classes implement the DatabaseInterface and provide their own specific implementations for each method.
//     The performDatabaseOperations function takes a DatabaseInterface as a parameter, allowing it to work with any class that implements the interface. 
//     This is the key to polymorphism.

// With this approach, you can easily switch between different database types by creating an instance of the desired database class and passing it to functions or classes that expect a DatabaseInterface object. This helps in keeping your code modular, extensible, and independent of the underlying database technology.