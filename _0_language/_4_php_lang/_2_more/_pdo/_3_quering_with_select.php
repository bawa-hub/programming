<?php
require_once '_0_dbconfig.php';

try {
    $pdo = new PDO("mysql:host=$host;dbname=$dbname", $username, $password);

    $sql = 'SELECT lastname,
                    firstname,
                    jobtitle
               FROM employees
              ORDER BY lastname';

    $q = $pdo->query($sql);
    $q->setFetchMode(PDO::FETCH_ASSOC);
} catch (PDOException $e) {
    die("Could not connect to the database $dbname :" . $e->getMessage());
}
?>
<!DOCTYPE html>
<html>

<head>
    <title>PHP MySQL Query Data Demo</title>
    <link href="css/bootstrap.min.css" rel="stylesheet">
    <link href="css/style.css" rel="stylesheet">
</head>

<body>
    <div id="container">
        <h1>Employees</h1>
        <table class="table table-bordered table-condensed">
            <thead>
                <tr>
                    <th>First Name</th>
                    <th>Last Name</th>
                    <th>Job Title</th>
                </tr>
            </thead>
            <tbody>
                <?php while ($row = $q->fetch()) : ?>
                    <tr>
                        <td><?php echo htmlspecialchars($row['lastname']) ?></td>
                        <td><?php echo htmlspecialchars($row['firstname']); ?></td>
                        <td><?php echo htmlspecialchars($row['jobtitle']); ?></td>
                    </tr>
                <?php endwhile; ?>
            </tbody>
        </table>
</body>
</div>

</html>