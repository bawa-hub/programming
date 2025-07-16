<?php

/**
 * By default, variables are always assigned by value
 *  when you assign an expression to a variable, 
 * the entire value of the original expression is copied into the destination variable
 * 
 * PHP also offers another way to assign values to variables: assign by reference
 */

$foo = 'Bob';              // Assign the value 'Bob' to $foo
$bar = &$foo;              // Reference $foo via $bar.
$bar = "My name is $bar";  // Alter $bar...
echo $bar, PHP_EOL;
echo $foo, PHP_EOL;
