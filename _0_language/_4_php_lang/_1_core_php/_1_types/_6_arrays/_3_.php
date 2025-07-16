<?php

// key can either be an int or a string. 
// The value can be of any type. 

/**
 * following key casts will occur:

 * Strings containing valid decimal ints, unless the number is preceded by a + sign, will be cast to the int type. 
 * E.g. the key "8" will actually be stored under 8. On the other hand "08" will not be cast, 
 * as it isn't a valid decimal integer.
 * 
 * Floats are also cast to ints, which means that the fractional part will be truncated. 
 * E.g. the key 8.7 will actually be stored under 8.
 * 
 * Bools are cast to ints, too, i.e. 
 * the key true will actually be stored under 1 and the key false under 0.
 * 
 * Null will be cast to the empty string, i.e. the key null will actually be stored under "".
 * 
 * Arrays and objects can not be used as keys. Doing so will result in a warning: Illegal offset type

 * If multiple elements in the array declaration use the same key, 
 * only the last one will be used as all others are overwritten. 
 */

$array = array(
    1    => "a",
    "1"  => "b",
    1.5  => "c",
    true => "d",
);
var_dump($array);

$array1 = array(
    "foo" => "bar",
    "bar" => "foo",
    100   => -100,
    -100  => 100,
);
var_dump($array1);
