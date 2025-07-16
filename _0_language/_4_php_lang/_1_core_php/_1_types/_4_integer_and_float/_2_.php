<?php

/**
 *  size of an int is platform-dependent, 
 * although a maximum value of about two billion is the usual value (that's 32 bits signed)
 * 
 *  64-bit platforms usually have a maximum value of about 9E18. 
 * PHP does not support unsigned ints
 * 
 * int size can be determined using the constant PHP_INT_SIZE, 
 * maximum value using the constant PHP_INT_MAX, 
 * and minimum value using the constant PHP_INT_MIN. 
 * 
 * If PHP encounters a number beyond the bounds of the int type, 
 * it will be interpreted as a float instead. 
 * Also, an operation which results in a number beyond the bounds of the int type will return a float instead. 
 */

//  Integer overflow on a 32-bit system
$large_number = 2147483647;
var_dump($large_number);                     // int(2147483647)

$large_number = 2147483648;
var_dump($large_number);                     // float(2147483648)

$million = 1000000;
$large_number =  50000 * $million;
var_dump($large_number);                     // float(50000000000)