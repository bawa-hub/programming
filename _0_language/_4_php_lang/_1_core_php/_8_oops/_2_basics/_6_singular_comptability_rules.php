<?php

// Signature compatibility rules
// When overriding a method, its signature must be compatible with the parent method. 

// Compatible child methods
class Base
{
    public function foo(int $a)
    {
        echo "Valid\n";
    }
}

class Extend1 extends Base
{
    function foo(int $a = 5)
    {
        parent::foo($a);
    }
}

class Extend2 extends Base
{
    function foo(int $a, $b = 5)
    {
        parent::foo($a);
    }
}

$extended1 = new Extend1();
$extended1->foo();
$extended2 = new Extend2();
$extended2->foo(1);
// Valid
// Valid

// Fatal error when a child method removes a parameter
class Base
{
    public function foo(int $a = 5)
    {
        echo "Valid\n";
    }
}

class Extend extends Base
{
    function foo()
    {
        parent::foo(1);
    }
}

// Fatal error when a child method makes an optional parameter mandatory

class Base
{
    public function foo(int $a = 5)
    {
        echo "Valid\n";
    }
}

class Extend extends Base
{
    function foo(int $a)
    {
        parent::foo($a);
    }
}
