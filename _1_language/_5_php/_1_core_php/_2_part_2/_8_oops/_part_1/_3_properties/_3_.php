<?php

// As of PHP 8.1.0, a property can be declared with the readonly modifier, which prevents modification of the property after initialization
class Test {
    public readonly string $prop;
 
    public function __construct(string $prop) {
        // Legal initialization.
        $this->prop = $prop;
    }
 }
 
 $test = new Test("foobar");
 // Legal read.
 var_dump($test->prop); // string(6) "foobar"
 
 // Illegal reassignment. It does not matter that the assigned value is the same.
 $test->prop = "foobar";
 // Error: Cannot modify readonly property Test::$prop