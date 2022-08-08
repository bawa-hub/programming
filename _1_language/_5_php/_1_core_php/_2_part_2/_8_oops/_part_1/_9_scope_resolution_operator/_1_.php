<?php

/**
 * Scope Resolution Operator(::) is a token that allows access to 
 * static, constant, and overridden properties or methods of a class. 
 * 
 */

//   from outside the class definition
class MyClass
{
    const CONST_VALUE = 'A constant value';

    protected function myFunc()
    {
        echo "MyClass::myFunc()\n";
    }
}

class OtherClass extends MyClass
{
    public static $my_static = 'static var';

    public static function doubleColon()
    {
        echo parent::CONST_VALUE . "\n";
        echo self::$my_static . "\n";
    }

    // Override parent's definition
    public function myFunc()
    {
        // But still call the parent function
        parent::myFunc();
        echo "OtherClass::myFunc()\n";
    }
}

$classname = 'MyClass';
echo $classname::CONST_VALUE;
echo "<br>";
echo MyClass::CONST_VALUE;
echo "<br>";

$classname = 'OtherClass';
$classname::doubleColon();
echo "<br>";
OtherClass::doubleColon();
echo "<br>";

$class = new OtherClass();
$class->myFunc();
