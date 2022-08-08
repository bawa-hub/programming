<?php

/**
 * PHP MySQL Delete Data Demo
 */
class DeleteDataDemo
{

    const DB_HOST = 'localhost';
    const DB_NAME = 'pdo_learn';
    const DB_USER = 'root';
    const DB_PASSWORD = '';

    /**
     * PDO instance
     * @var PDO 
     */
    private $pdo = null;

    /**
     * Open a database connection to MySQL
     */
    public function __construct()
    {
        // open database connection
        $conStr = sprintf("mysql:host=%s;dbname=%s", self::DB_HOST, self::DB_NAME);
        try {
            $this->pdo = new PDO($conStr, self::DB_USER, self::DB_PASSWORD);
        } catch (PDOException $e) {
            die($e->getMessage());
        }
    }

    /**
     * Delete a task based on a specified task id
     * @param int $id
     * @return bool true on success or false on failure
     */
    public function delete($id)
    {

        $sql = 'DELETE FROM tasks
                WHERE task_id = :task_id';

        $q = $this->pdo->prepare($sql);

        return $q->execute([':task_id' => $id]);
    }

    /**
     * Delete all rows in the tasks table
     */
    public function deleteAll()
    {
        $sql = 'DELETE FROM tasks';
        return $this->pdo->exec($sql);
    }

    /**
     * Truncate the tasks table
     */
    public function truncateTable()
    {
        $sql = 'TRUNCATE TABLE tasks';
        return $this->pdo->exec($sql);
    }

    /**
     * close the database connection
     */
    public function __destruct()
    {
        $this->pdo = null;
    }
}

$obj = new DeleteDataDemo();
// delete id 2
$obj->delete(2);
