<?php

class InsertDataDemo
{

    const DB_HOST = 'localhost';
    const DB_NAME = 'pdo_learn';
    const DB_USER = 'root';
    const DB_PASSWORD = '';

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
        } catch (PDOException $pe) {
            die($pe->getMessage());
        }
    }

    /**
     * Insert a row into a table
     * @return
     */
    public function insert()
    {
        $sql = "INSERT INTO tasks (
                      subject,
                      description,
                      start_date,
                      end_date
                  )
                  VALUES (
                      'Learn PHP MySQL Insert Dat',
                      'PHP MySQL Insert data into a table',
                      '2013-01-01',
                      '2013-01-01'
                  )";

        return $this->pdo->exec($sql);
    }

    function insertSingleRow($subject, $description, $startDate, $endDate)
    {
        $task = array(
            ':subject' => $subject,
            ':description' => $description,
            ':start_date' => $startDate,
            ':end_date' => $endDate
        );

        $sql = 'INSERT INTO tasks (
                      subject,
                      description,
                      start_date,
                      end_date
                  )
                  VALUES (
                      :subject,
                      :description,
                      :start_date,
                      :end_date
                  );';

        $q = $this->pdo->prepare($sql);

        return $q->execute($task);
    }
}

$obj = new InsertDataDemo();
$obj->insert();
$obj->insertSingleRow(
    'MySQL PHP Insert Tutorial',
    'MySQL PHP Insert using prepared statement',
    '2013-01-01',
    '2013-01-02'
);
