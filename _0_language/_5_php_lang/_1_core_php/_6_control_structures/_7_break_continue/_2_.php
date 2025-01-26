<?php

// continue is used within looping structures to skip the rest of the current loop iteration 
// and continue execution at the condition evaluation and 
// then the beginning of the next iteration. 

// continue accepts an optional numeric argument which tells it how many levels of enclosing loops it should skip to the end of. 
// The default value is 1, thus skipping to the end of the current loop. 

$arr = array('one', 'two', 'three', 'four', 'stop', 'five');
foreach ($arr as $key => $value) {
    if (!($key % 2)) { // skip even members
        continue;
    }
    echo "$value<br>";
}

echo "<br>";

$i = 0;
while ($i++ < 5) {
    echo "Outer<br />\n";
    while (1) {
        echo "Middle<br />\n";
        while (1) {
            echo "Inner<br />\n";
            continue 3;
        }
        echo "This never gets output.<br />\n";
    }
    echo "Neither does this.<br />\n";
}
