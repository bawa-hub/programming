<?php

// In order to be able to directly modify array elements within the loop precede $value with &. 
// In that case the value will be assigned by reference. 

$arr = array(1, 2, 3, 4);
foreach ($arr as &$value) {
    $value = $value * 2;
    echo "$value<br>";
}
// $arr is now array(2, 4, 6, 8)
unset($value); // break the reference with the last element
var_dump($arr);
