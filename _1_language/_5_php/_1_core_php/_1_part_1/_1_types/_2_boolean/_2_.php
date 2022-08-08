<?php

/**
 * When converting to bool, the following values are considered false:
 * 
 * 
    the boolean false itself
    the integer 0 (zero)
    the floats 0.0 and -0.0 (zero)
    the empty string, and the string "0"
    an array with zero elements
    the special type NULL (including unset variables)
    SimpleXML objects created from attributeless empty elements, 
    i.e. elements which have neither children nor attributes.

 * Every other value is considered true (including any resource and NAN). 
 */

var_dump((bool) "");        // bool(false)
var_dump((bool) "0");       // bool(false)
var_dump((bool) 1);         // bool(true)
var_dump((bool) -2);        // bool(true)
var_dump((bool) "foo");     // bool(true)
var_dump((bool) 2.3e5);     // bool(true)
var_dump((bool) array(12)); // bool(true)
var_dump((bool) array());   // bool(false)
var_dump((bool) "false");   // bool(true)