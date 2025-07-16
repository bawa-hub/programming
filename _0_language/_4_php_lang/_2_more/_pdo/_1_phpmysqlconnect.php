<?php
require_once 'dbconfig.php';

try {
    $conn = new PDO("mysql:host=$host;dbname=$dbname", $username, $password);
    echo "Connected to $dbname at $host successfully.";
} catch (PDOException $pe) {
    die("Could not connect to the database $dbname :" . $pe->getMessage());
}
// When the script ends, PHP automatically closes the connection to the MySQL database server. 
// If you want to explicitly close the database connection, 
// you need to set the PDO object to null
$conn = null;
