<?php

// Using ... to access variable arguments
function sum(...$numbers)
{
    $acc = 0;
    foreach ($numbers as $n) {
        $acc += $n;
    }
    return $acc;
}

echo sum(1, 2, 3, 4);
// 10

// Using ... to provide arguments
function add($a, $b)
{
    return $a + $b;
}

echo add(...[1, 2]) . "\n";

$a = [1, 2];
echo add(...$a);
// 3
// 3