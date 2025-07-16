<?php

/**
 * Scope Resolution Operator(::) is a token that allows access to 
 * static, constant, and overridden properties or methods of a class. 
 * 
 */

//  :: from outside the class definition
class MyClass
{
    const CONST_VALUE = 'A constant value';
}

$classname = 'MyClass';
echo $classname::CONST_VALUE;

echo MyClass::CONST_VALUE;

// :: from inside the class definition
class OtherClass extends MyClass
{
    public static $my_static = 'static var';

    public static function doubleColon()
    {
        echo parent::CONST_VALUE . "\n";
        echo self::$my_static . "\n";
    }
}

$classname = 'OtherClass';
$classname::doubleColon();

OtherClass::doubleColon();


// Calling a parent's method
class MyClass
{
    protected function myFunc()
    {
        echo "MyClass::myFunc()\n";
    }
}

class OtherClass extends MyClass
{
    // Override parent's definition
    public function myFunc()
    {
        // But still call the parent function
        parent::myFunc();
        echo "OtherClass::myFunc()\n";
    }
}

$class = new OtherClass();
$class->myFunc();

// A class constant, class property (static), and class function (static) can all share the same name and be accessed using the double-colon
class A
{

    public static $B = '1'; # Static class variable.

    const B = '2'; # Class constant.

    public static function B()
    { # Static class function.
        return '3';
    }
}

echo A::$B . A::B . A::B(); # Outputs: 123