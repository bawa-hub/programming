<?php

function sumOfArray(&$arr)
{
    $sum = 0;
    for ($i = 0; $i < count($arr); $i++) $sum += $arr[$i];
    return $sum;
}

$a = [1, 2, 3, 4, 10];
echo sumOfArray($a);
