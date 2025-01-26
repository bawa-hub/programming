<!-- 
    Class properties and methods live in separate "namespaces", 
    so it is possible to have a property and a method with the same name
 -->

<!--  Property access vs. method call -->

<?php
class Foo
{
    public $bar = 'property';

    public function bar()
    {
        return 'method';
    }
}

$obj = new Foo();
echo $obj->bar, PHP_EOL, $obj->bar(), PHP_EOL;
// property
// method

// Calling an anonymous function stored in a property
class Foo
{
    public $bar;

    public function __construct()
    {
        $this->bar = function () {
            return 42;
        };
    }
}

$obj = new Foo();

echo ($obj->bar)(), PHP_EOL; // 42
