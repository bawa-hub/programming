<?php

/**
 * As of PHP 7.1.0, class constants may be defined as public, private, or protected. 
 * Constants declared without any explicit visibility keyword are defined as public. 
 */

class MyClass
{
    // Declare a public constant
    public const MY_PUBLIC = 'public';

    // Declare a protected constant
    protected const MY_PROTECTED = 'protected';

    // Declare a private constant
    private const MY_PRIVATE = 'private';

    public function foo()
    {
        echo self::MY_PUBLIC;
        echo self::MY_PROTECTED;
        echo self::MY_PRIVATE;
    }
}

$myclass = new MyClass();
MyClass::MY_PUBLIC; // Works
MyClass::MY_PROTECTED; // Fatal Error
MyClass::MY_PRIVATE; // Fatal Error
$myclass->foo(); // Public, Protected and Private work


/**
 * Define MyClass2
 */
class MyClass2 extends MyClass
{
    // This is public
    function foo2()
    {
        echo self::MY_PUBLIC;
        echo self::MY_PROTECTED;
        echo self::MY_PRIVATE; // Fatal Error
    }
}

$myclass2 = new MyClass2;
echo MyClass2::MY_PUBLIC; // Works
$myclass2->foo2(); // Public and Protected work, not Private