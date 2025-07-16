<?php

// key is optional. If it is not specified, 
// PHP will use the increment of the largest previously used int key. 

$array = array("foo", "bar", "hello", "world");
var_dump($array);

// It is possible to specify the key only for some elements and 
// leave it out for others: 
$array1 = array(
    "a",
    "b",
    6 => "c",
    "d",
);
var_dump($array1);
