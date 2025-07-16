<?php

// functions within functions

function foo()
{
    function bar()
    {
        echo "I don't exist until foo() is called.\n";
    }
}

/* We can't call bar() yet
   since it doesn't exist. */

foo();

/* Now we can call bar(),
   foo()'s processing has
   made it accessible. */

bar();

/**
 *  All functions and classes in PHP have the global scope - they can be called outside a function even if they were defined inside and vice versa.
 * PHP does not support function overloading, nor is it possible to undefine or redefine previously-declared functions. 
 */
