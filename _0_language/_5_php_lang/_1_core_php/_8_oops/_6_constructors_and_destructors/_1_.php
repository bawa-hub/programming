<?php

/***
 *  Parent constructors are not called implicitly if the child class defines a constructor. 
 * In order to run a parent constructor, a call to parent::__construct() within the child constructor is required. 
 * If the child does not define a constructor 
 * then it may be inherited from the parent class just like a normal class method (if it was not declared as private). 
 */

class BaseClass
{
    function __construct()
    {
        print "In BaseClass constructor\n";
    }
}

class SubClass extends BaseClass
{
    function __construct()
    {
        parent::__construct();
        print "In SubClass constructor\n";
    }
}

class OtherSubClass extends BaseClass
{
    // inherits BaseClass's constructor
}

// In BaseClass constructor
$obj = new BaseClass();

// In BaseClass constructor
// In SubClass constructor
$obj = new SubClass();

// In BaseClass constructor
$obj = new OtherSubClass();
