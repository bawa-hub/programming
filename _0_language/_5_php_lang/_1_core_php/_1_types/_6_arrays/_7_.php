<?php

// Array dereferencing
function getArray()
{
    return array(1, 2, 3);
}

$secondElement = getArray()[1];
// or
list(, $secondElement) = getArray();

var_dump($secondElement);
