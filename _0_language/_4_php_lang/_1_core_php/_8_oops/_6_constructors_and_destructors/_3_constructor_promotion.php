<?php

/**
 * As of PHP 8.0.0, constructor parameters may also be promoted to correspond to an object property. It is very common for constructor parameters to be assigned to a property in the constructor but otherwise not operated upon.
 *  Constructor promotion provides a short-hand for that use case
 */

class Point
{
    public function __construct(protected int $x, protected int $y = 0)
    {
    }
}

// When a constructor argument includes a modifier, PHP will interpret it as both an object property and a constructor argument, and assign the argument value to the property. The constructor body may then be empty or may contain other statements. Any additional statements will be executed after the argument values have been assigned to the corresponding properties. 
