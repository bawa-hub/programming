<?php

// PHP supports passing arguments by value (the default), passing by reference,and default argument values. 
// Variable-length argument lists and Named Arguments are also supported.

//  Passing arrays to functions
function takes_array($input)
{
    echo "$input[0] + $input[1] = ", $input[0] + $input[1];
}

// Function Argument List with trailing Comma
function takes_many_args(
    $first_arg,
    $second_arg,
    $a_very_long_argument_name,
    $arg_with_default = 5,
    $again = 'a default string' // This trailing comma was not permitted before 8.0.0.
) {
    // ...
}

// Declaring optional arguments after mandatory arguments
function foo($a = [], $b)
{
} // Before
function foo1($a, $b)
{
}      // After

function bar(A $a = null, $b)
{
} // Still allowed
function bar1(?A $a, $b)
{
}       // Recommended
