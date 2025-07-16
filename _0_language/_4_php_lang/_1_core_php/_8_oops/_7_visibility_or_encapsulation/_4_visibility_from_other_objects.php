<?php

/**
 * Objects of the same type will have access to each others private and protected members even though they are not the same instances. 
 * This is because the implementation specific details are already known when inside those objects. 
 */

//  Accessing private members of the same object type
class Test
{
    private $foo;

    public function __construct($foo)
    {
        $this->foo = $foo;
    }

    private function bar()
    {
        echo 'Accessed the private method.';
    }

    public function baz(Test $other)
    {
        // We can change the private property:
        $other->foo = 'hello';
        var_dump($other->foo);

        // We can also call the private method:
        $other->bar();
    }
}

$test = new Test('test');

$test->baz(new Test('other'));
