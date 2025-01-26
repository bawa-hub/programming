<?php

// The first class callable syntax is introduced as of PHP 8.1.0, as a way of creating anonymous functions from callable

// Simple first class callable syntax
class Foo
{
    public function method()
    {
    }
    public static function staticmethod()
    {
    }
    public function __invoke()
    {
    }
}

$obj = new Foo();
$classStr = 'Foo';
$methodStr = 'method';
$staticmethodStr = 'staticmethod';


$f1 = strlen(...);
$f2 = $obj(...);  // invokable object
$f3 = $obj->method(...);
$f4 = $obj->$methodStr(...);
$f5 = Foo::staticmethod(...);
$f6 = $classStr::$staticmethodStr(...);

// traditional callable using string, array
$f7 = 'strlen'(...);
$f8 = [$obj, 'method'](...);
$f9 = [Foo::class, 'staticmethod'](...);

// The ... is part of the syntax, and not an omission. 

// Scope comparison of CallableExpr(...) and traditional callable
class Foo
{
    public function getPrivateMethod()
    {
        return [$this, 'privateMethod'];
    }

    private function privateMethod()
    {
        echo __METHOD__, "\n";
    }
}

$foo = new Foo;
$privateMethod = $foo->getPrivateMethod();
$privateMethod();
// Fatal error: Call to private method Foo::privateMethod() from global scope
// This is because call is performed outside from Foo and visibility will be checked from this point.

class Foo1
{
    public function getPrivateMethod()
    {
        // Uses the scope where the callable is acquired.
        return $this->privateMethod(...); // identical to Closure::fromCallable([$this, 'privateMethod']);
    }

    private function privateMethod()
    {
        echo __METHOD__, "\n";
    }
}

$foo1 = new Foo1;
$privateMethod = $foo1->getPrivateMethod();
$privateMethod();  // Foo1::privateMethod