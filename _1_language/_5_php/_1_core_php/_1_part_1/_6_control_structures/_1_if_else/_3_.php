<?php
$v = 1;

$r = (1 == $v) ? 'Yes' : 'No'; // $r is set to 'Yes'
$r = (3 == $v) ? 'Yes' : 'No'; // $r is set to 'No'

echo (1 == $v) ? 'Yes' : 'No'; // 'Yes' will be printed

// and since PHP 5.3
$v = 'My Value';
$r = ($v) ?: 'No Value'; // $r is set to 'My Value' because $v is evaluated to TRUE

$v = '';
echo ($v) ?: 'No Value'; // 'No Value' will be printed because $v is evaluated to FALSE
