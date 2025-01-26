<?php

// There is no int division operator in PHP, 
// to achieve this use the intdiv() function.
// value can be cast to an int to round it towards zero, 
// or the round() function provides finer control over rounding. 

var_dump(25 / 7);         // float(3.5714285714286)
var_dump((int) (25 / 7)); // int(3)
var_dump(round(25 / 7));  // float(4)