<?php

class SimpleClass
{
    public $var = 'a default value';
    public function displayVar()
    {
        echo $this->var;
    }
}

// In the class context, it is possible to create a new object by new self and new parent
// When assigning an already created instance of a class to a new variable, 
// the new variable will access the same instance as the object that was assigned. 
// This behaviour is the same when passing instances to a function. 
// A copy of an already created object can be made by cloning it.

$instance = new SimpleClass();
$assigned   =  $instance;
$reference  = &$instance;
$instance->var = '$assigned will have this value';
$instance = null; // $instance and $reference become null
var_dump($instance);
var_dump($reference);
var_dump($assigned);
