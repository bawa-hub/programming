<?php

// PHP possesses a destructor concept similar to that of other object-oriented languages, such as C++. The destructor method will be called as soon as there are no other references to a particular object, or in any order during the shutdown sequence.

class MyDestructableClass
{
    function __construct()
    {
        print "In constructor\n";
    }

    function __destruct()
    {
        print "Destroying " . __CLASS__ . "\n";
    }
}

$obj = new MyDestructableClass();
