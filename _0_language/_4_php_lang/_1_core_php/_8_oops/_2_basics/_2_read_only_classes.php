<?php

/***
 * Marking a class as readonly will add the readonly modifier to every declared property, 
 * and prevent the creation of dynamic properties. Moreover, it is impossible to add support for them by using the AllowDynamicProperties attribute. 
 * Attempting to do so will trigger a compile-time error. 
 */

#[\AllowDynamicProperties]
readonly class Foo
{
}

// Fatal error: Cannot apply #[AllowDynamicProperties] to readonly class Foo

// neither untyped, nor static properties can be marked with the readonly modifier, readonly classes cannot declare them either:
readonly class Foo
{
    public $bar;
}

// Fatal error: Readonly property Foo::$bar must have type
readonly class Foo
{
    public static int $bar;
}

// Fatal error: Readonly class Foo cannot declare static properties

// A readonly class can be extended if, and only if, the child class is also a readonly class. 
