<?php

/**
 * An inherited member from a base class is overridden by a member inserted by a Trait
 * precedence order is that members from the current class override Trait methods, 
 * which in turn override inherited methods. 
 */

class Base
{
    public function sayHello()
    {
        echo 'Hello ';
    }
}

trait SayWorld
{
    public function sayHello()
    {
        parent::sayHello();
        echo 'World!';
    }
}

class MyHelloWorld extends Base
{
    use SayWorld;
}

$o = new MyHelloWorld();
$o->sayHello();
