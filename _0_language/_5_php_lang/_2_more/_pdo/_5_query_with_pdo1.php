<?php

require_once '_0_dbconfig.php';

try {
    $pdo = new PDO("mysql:host=$host;dbname=$dbname", $username, $password);

    $sql = 'SELECT lastname,
                    firstname,
                    jobtitle
               FROM employees
              WHERE lastname LIKE :lname OR 
                    firstname LIKE :fname;';

    // prepare statement for execution
    $q = $pdo->prepare($sql);
    
    // pass values to the query and execute it
    $q->execute([':fname' => 'Le%',
        ':lname' => '%son']);
    
    $q->setFetchMode(PDO::FETCH_ASSOC);
    
    // print out the result set
    while ($r = $q->fetch()) {
        echo sprintf('%s <br/>', $r['lastname']);
    }
} catch (PDOException $e) {
    die("Could not connect to the database $dbname :" . $e->getMessage());
}