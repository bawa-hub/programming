<?php

// function fib($num)
// {
//     if ($num == 1) return 0;
//     if ($num == 2) return 1;
//     return fib($num - 1) + fib($num - 2);
// }

function fib($n)
{
    $f = array();
    $f[0] = 0;
    $f[1] = 1;
    for ($i = 2; $i <= $n; $i++) {
        $f[$i] = $f[$i - 1] + $f[$i - 2];
    }
    return $f[$n];
}

$time_start = microtime(true);
$time_end = microtime(true);

$execution_time = ($time_end - $time_start);
echo fib(1000000) . " ";
echo $execution_time * 1000;
