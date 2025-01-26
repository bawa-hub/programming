<?php

/**
 * PHP implements a way to reuse code called Traits. 
 * Traits are a mechanism for code reuse in single inheritance languages such as PHP
 * A Trait is similar to a class, 
 * but only intended to group functionality in a fine-grained and consistent way. 
 * It is not possible to instantiate a Trait on its own. 
 * It is an addition to traditional inheritance and enables horizontal composition of behavior; 
 * that is, the application of class members without requiring inheritance. 
 */

trait ezcReflectionReturnInfo
{
    function getReturnType()
    { /*1*/
    }
    function getReturnDescription()
    { /*2*/
    }
}

class ezcReflectionMethod extends ReflectionMethod
{
    use ezcReflectionReturnInfo;
    /* ... */
}

class ezcReflectionFunction extends ReflectionFunction
{
    use ezcReflectionReturnInfo;
    /* ... */
}
