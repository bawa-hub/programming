<?php

// If a string containing the name of a class is used with new, a new instance of that class will be created. If the class is in a namespace, its fully qualified name must be used when doing this. 
// If there are no arguments to be passed to the class's constructor, parentheses after the class name may be omitted. 

// Creating an instance
$instance = new SimpleClass();

// This can also be done with a variable:
$className = 'SimpleClass';
$instance = new $className(); // new SimpleClass()

// As of PHP 8.0.0, using new with arbitrary expressions is supported. This allows more complex instantiation if the expression produces a string. The expressions must be wrapped in parentheses. 

// Creating an instance using an arbitrary expression (as of php 8)
class ClassA extends \stdClass
{
}
class ClassB extends \stdClass
{
}
class ClassC extends ClassB
{
}
class ClassD extends ClassA
{
}

function getSomeClass(): string
{
    return 'ClassA';
}

var_dump(new (getSomeClass()));
var_dump(new ('Class' . 'B'));
var_dump(new ('Class' . 'C'));
var_dump(new (ClassD::class));
// object(ClassA)#1 (0) {
// }
// object(ClassB)#1 (0) {
// }
// object(ClassC)#1 (0) {
// }
// object(ClassD)#1 (0) {
// }


// In the class context, it is possible to create a new object by new self and new parent
// When assigning an already created instance of a class to a new variable, 
// the new variable will access the same instance as the object that was assigned. 
// This behaviour is the same when passing instances to a function. 
// A copy of an already created object can be made by cloning it.

// Object Assignment
class SimpleClass
{
    public $var = 'a default value';
    public function displayVar()
    {
        echo $this->var;
    }
}

$instance = new SimpleClass();
$assigned   =  $instance;
$reference  = &$instance;
$instance->var = '$assigned will have this value';
$instance = null; // $instance and $reference become null
var_dump($instance);
var_dump($reference);
var_dump($assigned);

// It's possible to create instances of an object in a couple of ways: 
// Creating new objects
class Test
{
    static public function getNew()
    {
        return new static;
    }
}

class Child extends Test
{
}

$obj1 = new Test();
$obj2 = new $obj1;
var_dump($obj1 !== $obj2);

$obj3 = Test::getNew();
var_dump($obj3 instanceof Test);

$obj4 = Child::getNew();
var_dump($obj4 instanceof Child);
// bool(true)
// bool(true)
// bool(true)

// It is possible to access a member of a newly created object in a single expression: 
echo (new DateTime())->format('Y');
