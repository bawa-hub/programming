<?php

// The final keyword prevents child classes from overriding a method or constant by prefixing the definition with final. If the class itself is being defined final then it cannot be extended. 

// Final methods example

class BaseClass
{
    public function test()
    {
        echo "BaseClass::test() called\n";
    }

    final public function moreTesting()
    {
        echo "BaseClass::moreTesting() called\n";
    }
}

class ChildClass extends BaseClass
{
    public function moreTesting()
    {
        echo "ChildClass::moreTesting() called\n";
    }
}
// Results in Fatal error: Cannot override final method BaseClass::moreTesting()

//  Final class example
final class BaseClass
{
    public function test()
    {
        echo "BaseClass::test() called\n";
    }

    // As the class is already final, the final keyword is redundant
    final public function moreTesting()
    {
        echo "BaseClass::moreTesting() called\n";
    }
}

class ChildClass extends BaseClass
{
}
// Results in Fatal error: Class ChildClass may not inherit from final class (BaseClass)


//  Final constants example as of PHP 8.1.0
class Foo
{
    final public const X = "foo";
}

class Bar extends Foo
{
    public const X = "bar";
}

// Fatal error: Bar::X cannot override final constant Foo::X

// Properties cannot be declared final: only classes, methods, and constants (as of PHP 8.1.0) may be declared as final. As of PHP 8.0.0, private methods may not be declared final except for the constructor. 