<?php

/**
 * As of PHP 8.0.0, the ::class constant may also be used on objects. 
 * This resolution happens at runtime, not compile time. 
 * Its effect is the same as calling get_class() on the object. 
 */

// Object name resolution
namespace NS {
    class ClassName
    {
    }
}
$c = new ClassName();
print $c::class;
