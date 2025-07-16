
<?php
// This switch statement:

switch ($i) {
    case 0:
        echo "i equals 0";
        break;
    case 1:
        echo "i equals 1";
        break;
    case 2:
        echo "i equals 2";
        break;
}

// Is equivalent to:

if ($i == 0) {
    echo "i equals 0";
} elseif ($i == 1) {
    echo "i equals 1";
} elseif ($i == 2) {
    echo "i equals 2";
}



switch ($i) {
    case 0:
    case 1:
    case 2:
        echo "i is less than 3 but not negative";
        break;
    case 3:
        echo "i is 3";
}

// defuault case
switch ($i) {
    case 0:
        echo "i equals 0";
        break;
    case 1:
        echo "i equals 1";
        break;
    case 2:
        echo "i equals 2";
        break;
    default:
        echo "i is not equal to 0, 1 or 2";
}


// expression
$target = 1;
$start = 3;

switch ($target) {
    case $start - 1:
        print "A";
        break;
    case $start - 2:
        print "B";
        break;
    case $start - 3:
        print "C";
        break;
    case $start - 4:
        print "D";
        break;
}


// complex comparisons
$offset = 1;
$start = 3;

switch (true) {
    case $start - $offset === 1:
        print "A";
        break;
    case $start - $offset === 2:
        print "B";
        break;
    case $start - $offset === 3:
        print "C";
        break;
    case $start - $offset === 4:
        print "D";
        break;
}

// Prints "B"