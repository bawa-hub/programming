<?php

/**
 * An int is a number of the set ℤ = {..., -2, -1, 0, 1, 2, ...}. 
 * 
 * ints can be specified in decimal (base 10), hexadecimal (base 16), octal (base 8) or binary (base 2) notation. The negation operator can be used to denote a negative int
 * 
 *  To use octal notation, precede the number with a 0 (zero). 
 * As of PHP 8.1.0, octal notation can also be preceded with 0o or 0O. 
 * To use hexadecimal notation precede the number with 0x. 
 * To use binary notation precede the number with 0b.
 * As of PHP 7.4.0, integer literals may contain underscores (_) between digits, for better readability of literals
 */

$a = 1234; // decimal number
$a = 0123; // octal number (equivalent to 83 decimal)
$a = 0o123; // octal number (as of PHP 8.1.0)
$a = 0x1A; // hexadecimal number (equivalent to 26 decimal)
$a = 0b11111111; // binary number (equivalent to 255 decimal)
$a = 1_234_567; // decimal number (as of PHP 7.4.0)
