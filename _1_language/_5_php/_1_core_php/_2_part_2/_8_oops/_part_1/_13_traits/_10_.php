<?php

// static methods

trait StaticExample
{
    public static function doSomething()
    {
        return 'Doing something';
    }
}

class Example
{
    use StaticExample;
}

Example::doSomething();
