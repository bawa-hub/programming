<?php

/**
 * Classes defined as abstract cannot be instantiated, 
 * and any class that contains at least one abstract method must also be abstract.
 * 
 * Methods defined as abstract simply declare the method's signature; 
 * they cannot define the implementation.
 * 
 * When inheriting from an abstract class, 
 * all methods marked abstract in the parent's class declaration must be defined by the child class, 
 * and follow the usual inheritance and signature compatibility rules. 
 */

abstract class AbstractClass
{
    // Force Extending class to define this method
    abstract protected function getValue();
    abstract protected function prefixValue($prefix);

    // Common method
    public function printOut()
    {
        print $this->getValue() . "\n";
    }
}

class ConcreteClass1 extends AbstractClass
{
    protected function getValue()
    {
        return "ConcreteClass1";
    }

    public function prefixValue($prefix)
    {
        return "{$prefix}ConcreteClass1";
    }
}

class ConcreteClass2 extends AbstractClass
{
    public function getValue()
    {
        return "ConcreteClass2";
    }

    public function prefixValue($prefix)
    {
        return "{$prefix}ConcreteClass2";
    }
}

$class1 = new ConcreteClass1;
$class1->printOut();
echo $class1->prefixValue('FOO_') . "\n";

$class2 = new ConcreteClass2;
$class2->printOut();
echo $class2->prefixValue('FOO_') . "\n";
