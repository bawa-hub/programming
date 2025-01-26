<?php
// defined at run time
define('IS_LOGGEDIN', true);
echo IS_LOGGEDIN, PHP_EOL;
echo defined('IS_LOGGEDIN'), PHP_EOL; // check constant is defined or not

// defined at compile time
const IS_PAID = true;
echo IS_PAID, PHP_EOL;

$paid = 'PAID';
define('STATUS_' . $paid, true); // dynamically created
echo STATUS_PAID, PHP_EOL;
