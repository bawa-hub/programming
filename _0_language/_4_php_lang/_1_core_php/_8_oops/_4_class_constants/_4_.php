<?php

// Class constant visibility modifiers, as of PHP 7.1.0
class Foo
{
    public const BAR = 'bar';
    private const BAZ = 'baz';
}
echo Foo::BAR, PHP_EOL;
echo Foo::BAZ, PHP_EOL;
