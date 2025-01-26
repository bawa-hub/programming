<?php

/**
 * Iterable is a pseudo-type introduced in PHP 7.1. 
 * It accepts any array or object implementing the Traversable interface
 * 
 * Both of these types are iterable using foreach and 
 * can be used with yield from within a generator. 
 */

//  Iterable can be used as a parameter type to indicate that a function requires a set of values
// If a value is not an array or instance of Traversable, a TypeError will be thrown.
function foo(iterable $iterable)
{
    foreach ($iterable as $value) {
        // ...
    }
}

// Parameters declared as iterable may use null or an array as a default value. 
function bar(iterable $iterable = [])
{
    // ...
}

// Iterable can also be used as a return type to indicate a function will return an iterable value. 
// If the returned value is not an array or instance of Traversable, a TypeError will be thrown. 
function baz(): iterable
{
    return [1, 2, 3];
}

// Functions declaring iterable as a return type may also be generators. 
function gen(): iterable
{
    yield 1;
    yield 2;
    yield 3;
}
