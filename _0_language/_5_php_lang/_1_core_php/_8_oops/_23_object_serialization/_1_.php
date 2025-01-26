<?php

/**
 * serialize() returns a string containing a byte-stream representation of any value that can be stored in PHP. 
 * unserialize() can use this string to recreate the original variable values.
 * 
 * Using serialize to save an object will save all variables in an object. 
 * The methods in an object will not be saved, only the name of the class
 * 
 * In order to be able to unserialize() an object, the class of that object needs to be defined. 
 * That is, if you have an object of class A and serialize this, you'll get a string that refers to class A and contains all values of variables contained in it. 
 * If you want to be able to unserialize this in another file, an object of class A, the definition of class A must be present in that file first. 
 * This can be done for example by storing the class definition of class A in an include file and including this file or making use of the spl_autoload_register() function. 
 * 
 */

// A.php:

class A
{
    public $one = 1;

    public function show_one()
    {
        echo $this->one;
    }
}

// page1.php:

include "A.php";

$a = new A;
$s = serialize($a);
// store $s somewhere where page2.php can find it.
file_put_contents('store', $s);

// page2.php:

// this is needed for the unserialize to work properly.
include "A.php";

$s = file_get_contents('store');
$a = unserialize($s);

// now use the function show_one() of the $a object.  
$a->show_one();
