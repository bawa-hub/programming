<?php

require __DIR__ . '/vendor/autoload.php';

use \imonroe\ana\Ana;
use \bawa\composer_demo\Library;

echo time() . PHP_EOL;

echo Ana::sql_datetime(time()) . PHP_EOL;

$library = new Library;
$library->sayHi();
