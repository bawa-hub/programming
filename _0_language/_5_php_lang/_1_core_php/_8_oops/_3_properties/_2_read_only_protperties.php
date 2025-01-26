<?php

// As of PHP 8.1.0, a property can be declared with the readonly modifier, which prevents modification of the property after initialization
class Test
{
    public readonly string $prop;

    public function __construct(string $prop)
    {
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

//  The readonly modifier can only be applied to typed properties. A readonly property without type constraints can be created using the Mixed type
// Readonly static properties are not supported. 

// Illegal initialization of readonly properties
class Test1
{
    public readonly string $prop;
}

$test1 = new Test1;
// Illegal initialization outside of private scope.
$test1->prop = "foobar";
// Error: Cannot initialize readonly property Test1::$prop from global scope