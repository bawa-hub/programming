<?php

/**
 * Declaring class properties or methods as static makes them accessible without needing an instantiation of the class. 
 * These can also be accessed statically within an instantiated class object.
 *  
 * Because static methods are callable without an instance of the object created, 
 * the pseudo-variable $this is not available inside methods declared as static. 
 * 
 * Static properties are accessed using the Scope Resolution Operator (::) 
 * and cannot be accessed through the object operator (->). 
 */

//  Static method example
class Foo
{
    public static $my_static = 'foo';

    public function staticValue()
    {
        return self::$my_static;
    }

    public static function aStaticMethod()
    {
        echo "I'm static method";
    }
}

class Bar extends Foo
{
    public function fooStatic()
    {
        return parent::$my_static;
    }
}

print Foo::$my_static;
echo "<br>";

$foo = new Foo();
print $foo->staticValue() . "\n";
print $foo->my_static . "\n";      // Undefined "Property" my_static 

print $foo::$my_static . "\n";
$classname = 'Foo';
print $classname::$my_static . "\n";

print Bar::$my_static . "\n";
$bar = new Bar();
print $bar->fooStatic() . "\n";

Foo::aStaticMethod();
echo "<br>";
$classname = 'Foo';
$classname::aStaticMethod();
echo "<br>";
