<?php

/**
 * when extending a class, the subclass inherits all of the public and protected methods, properties and constants from the parent class. 
 * Unless a class overrides those methods, they will retain their original functionality. 
 * 
 * Private methods of a parent class are not accessible to a child class. As a result, child classes may reimplement a private method themselves without regard for normal inheritance rules.
 * 
 * Prior to PHP 8.0.0, however, final and static restrictions were applied to private methods. As of PHP 8.0.0, the only private method restriction that is enforced is private final constructors, as that is a common way to "disable" the constructor when using static factory methods instead. 
 * 
 */

class Foo
{
    public function printItem($string)
    {
        echo 'Foo: ' . $string . PHP_EOL;
    }

    public function printPHP()
    {
        echo 'PHP is great.' . PHP_EOL;
    }
}

class Bar extends Foo
{
    public function printItem($string)
    {
        echo 'Bar: ' . $string . PHP_EOL;
    }
}

$foo = new Foo();
$bar = new Bar();
$foo->printItem('baz'); // Output: 'Foo: baz'
$foo->printPHP();       // Output: 'PHP is great' 
$bar->printItem('baz'); // Output: 'Bar: baz'
$bar->printPHP();       // Output: 'PHP is great'

// It is not allowed to override a read-write property with a readonly property or vice versa. 
class A
{
    public int $prop;
}
class B extends A
{
    // Illegal: read-write -> readonly
    public readonly int $prop;
}
