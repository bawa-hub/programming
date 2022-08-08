<?php

// Simple Class definition
class SimpleClass
{
    // property declaration
    public $var = 'a default value';

    // method declaration
    public function displayVar()
    {
        echo $this->var;
    }
}
// The pseudo-variable $this is available when a method is called from within an object context. 
// $this is the value of the calling object 

// Creating an instance
$instance = new SimpleClass();
// If there are no arguments to be passed to the class's constructor, 
// parentheses after the class name may be omitted.
$instance1 = new SimpleClass;

// This can also be done with a variable:
$className = 'SimpleClass';
$instance2 = new $className(); // new SimpleClass()

echo $instance->var;
