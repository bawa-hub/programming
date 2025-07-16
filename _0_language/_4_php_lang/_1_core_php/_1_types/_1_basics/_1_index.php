<?php

/**  PHP supports ten primitive types. */

/** Four scalar types: */

// 1. bool
$bool = true;
echo gettype($bool);
// 2. int
$int = 1;
echo gettype($int);
echo is_int($int);
// 3. float (floating-point number, aka double)
$float = 1.11;
echo gettype($float);
// 4. string
$string = "Hello world";
echo gettype($string);

/** Four compound types: */

// 5. array
$arr = ['monu', 'sonu', 'gonu'];
echo gettype($arr);
// 6. object
class Test
{
}
$test = new Test();
echo gettype($test);
// 7. callable
// 8. iterable


/** Two special types: */

// 9. resource
// 10. NULL

/**
 * 
 * var_dump() function returns the data type and value of an expression
 * To get a human-readable representation of a type for debugging, use the gettype() function
 * To check for a certain type, do not use gettype(), but rather the is_type functions
 * To forcibly convert a variable to a certain type, either cast the variable or use the settype() function on it.
 */
