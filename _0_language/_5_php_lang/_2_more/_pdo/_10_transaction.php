<?php

// A transaction is a set of inter-dependent SQL statements that needs to execute in all-or-nothing mode
// A transaction is successful if all SQL statements executed successfully. 
// A failure of any statement will trigger the system to rollback to the original state to avoid data inconsistency

/**
 * PHP MySQL Transaction Demo
 */
class TransactionDemo
{

    const DB_HOST = 'localhost';
    const DB_NAME = 'pdo_learn';
    const DB_USER = 'root';
    const DB_PASSWORD = '';

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
            die($e->getMessage());
        }
    }

    /**
     * PDO instance
     * @var PDO 
     */
    private $pdo = null;

    /**
     * Transfer money between two accounts
     * @param int $from
     * @param int $to
     * @param float $amount
     * @return true on success or false on failure.
     */
    public function transfer($from, $to, $amount)
    {

        try {
            $this->pdo->beginTransaction();

            // get available amount of the transferer account
            $sql = 'SELECT amount FROM accounts WHERE id=:from';
            $stmt = $this->pdo->prepare($sql);
            $stmt->execute(array(":from" => $from));
            $availableAmount = (int) $stmt->fetchColumn();
            $stmt->closeCursor();

            if ($availableAmount < $amount) {
                echo 'Insufficient amount to transfer';
                return false;
            }
            // deduct from the transferred account
            $sql_update_from = 'UPDATE accounts
				SET amount = amount - :amount
				WHERE id = :from';
            $stmt = $this->pdo->prepare($sql_update_from);
            $stmt->execute(array(":from" => $from, ":amount" => $amount));
            $stmt->closeCursor();

            // add to the receiving account
            $sql_update_to = 'UPDATE accounts
                                SET amount = amount + :amount
                                WHERE id = :to';
            $stmt = $this->pdo->prepare($sql_update_to);
            $stmt->execute(array(":to" => $to, ":amount" => $amount));

            // commit the transaction
            $this->pdo->commit();

            echo 'The amount has been transferred successfully';

            return true;
        } catch (PDOException $e) {
            $this->pdo->rollBack();
            die($e->getMessage());
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
}

// test the transfer method
$obj = new TransactionDemo();

// transfer 30K from from account 1 to 2
// $obj->transfer(1, 2, 30000);

// transfer 5K from from account 1 to 2
$obj->transfer(1, 2, 5000);
