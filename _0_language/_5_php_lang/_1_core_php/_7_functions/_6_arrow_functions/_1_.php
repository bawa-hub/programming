<?php

// Both anonymous functions and arrow functions are implemented using the Closure class. 
// Arrow functions have the basic form fn (argument_list) => expr. 

// Arrow functions capture variables by value automatically
$y = 1;

$fn1 = fn ($x) => $x + $y;
// equivalent to using $y by value:
$fn2 = function ($x) use ($y) {
    return $x + $y;
};

var_export($fn1(3));
// 4

// Arrow functions capture variables by value automatically, even when nested
$z = 1;
$fn = fn ($x) => fn ($y) => $x * $y + $z;
// Outputs 51
var_export($fn(5)(10));

//  Examples of arrow functions
fn (array $x) => $x;
static fn (): int => $x;
fn ($x = 42) => $x;
fn (&$x) => $x;
fn & ($x) => $x;
fn ($x, ...$rest) => $rest;

// Values from the outer scope cannot be modified by arrow functions
$x = 1;
$fn = fn () => $x++; // Has no effect
$fn();
var_export($x);  // Outputs 1