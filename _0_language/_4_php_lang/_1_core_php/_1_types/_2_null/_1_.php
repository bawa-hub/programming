<?php

/**
 * The special null value represents a variable with no value. 
 * null is the only possible value of type null. 
 * 
 *  A variable is considered to be null if:

 * it has been assigned the constant null.
 * it has not been set to any value yet.
 * it has been unset().

 * There is only one value of type null, 
 * and that is the case-insensitive constant null. 
 */

$var = null;

// is_null — Finds whether a variable is null 
var_dump(is_null($not_exist), is_null($var));
