<?php 

// If the return is omitted the value null will be returned.
function square($num)
{
    return $num * $num;
}
echo square(4);   // outputs '16'.

// Returning an array to get multiple values
function small_numbers()
{
    return array (0, 1, 2);
}
list ($zero, $one, $two) = small_numbers();

// Returning a reference from a function
function &returns_reference()
{
    return $someref;
}

$newref =& returns_reference();
