<?php

require_once '_0_dbconfig.php';

try {
    $pdo = new PDO("mysql:host=$host;dbname=$dbname", $username, $password);

    $sql = 'SELECT lastname,
                    firstname,
                    jobtitle
               FROM employees
              WHERE lastname LIKE ?';

    $q = $pdo->prepare($sql);
    $q->execute(['%son']);
    $q->setFetchMode(PDO::FETCH_ASSOC);

    while ($r = $q->fetch()) {
        echo sprintf('%s <br/>', $r['lastname']);
    }
} catch (PDOException $pe) {
    die("Could not connect to the database $dbname :" . $pe->getMessage());
}
