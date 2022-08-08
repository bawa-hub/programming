<?php

/**
 * It is possible to define constants on a per-class basis remaining the same and unchangeable. 
 * The default visibility of class constants is public. 
 * Class constants can be redefined by a child class. 
 * As of PHP 8.1.0, class constants cannot be redefined by a child class if it is defined as final.
 * It's also possible for interfaces to have constants
 * It's possible to reference the class using a variable.
 *  class constants are allocated once per class, and not for each class instance. 
 */

class MyClass
{
    const CONSTANT = 'constant value';

    function showConstant()
    {
        echo  self::CONSTANT . "\n";
    }
}

echo MyClass::CONSTANT . "\n";

$classname = "MyClass";
echo $classname::CONSTANT . "\n";

$class = new MyClass();
$class->showConstant();

echo $class::CONSTANT . "\n";
