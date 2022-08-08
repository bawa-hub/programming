<?php

require("Employee.php");

use employee\Employee;
use const employee\MAX;
use function employee\getData;

$emp = new Employee();
$name = $emp->getName();
var_dump($name);
echo MAX;
echo "<br>";
echo getData();
