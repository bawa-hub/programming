<?php

/**
 * PHP MySQL Create Table Demo
 */
class CreateTableDemo
{

    /**
     * database host
     */
    const DB_HOST = 'localhost';

    /**
     * database name
     */
    const DB_NAME = 'pdo_learn';

    /**
     * database user
     */
    const DB_USER = 'root';
    /*
     * database password
     */
    const DB_PASSWORD = '';

    /**
     *
     * @var type 
     */
    private $pdo = null;

    /**
     * Open the database connection
     */
    public function __construct()
    {
        // open database connection
        $conStr = sprintf("mysql:host=%s;dbname=%s", self::DB_HOST, self::DB_NAME);
        try {
            $this->pdo = new PDO($conStr, self::DB_USER, self::DB_PASSWORD);
        } catch (PDOException $e) {
            echo $e->getMessage();
        }
    }

    /**
     * close the database connection
     */
    public function __destruct()
    {
        // close the database connection
        $this->pdo = null;
    }

    /**
     * create the tasks table
     * @return boolean returns true on success or false on failure
     */
    public function createTaskTable()
    {
        $sql = <<<EOSQL
            CREATE TABLE IF NOT EXISTS tasks (
                task_id     INT AUTO_INCREMENT PRIMARY KEY,
                subject     VARCHAR (255)        DEFAULT NULL,
                start_date  DATE                 DEFAULT NULL,
                end_date    DATE                 DEFAULT NULL,
                description VARCHAR (400)        DEFAULT NULL
            );
EOSQL;
        return $this->pdo->exec($sql);
    }
}

// create the tasks table
$obj = new CreateTableDemo();
$obj->createTaskTable();
