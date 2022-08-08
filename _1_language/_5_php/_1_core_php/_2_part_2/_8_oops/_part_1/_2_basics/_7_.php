<!-- It is not possible to extend multiple classes
a class can only inherit from one base class -->

<!-- inherited constants, methods, and properties can be overridden by 
redeclaring them with the same name defined in the parent class -->

<!-- if the parent class has defined a method or constant as final, 
they may not be overridden. -->

<!-- It is possible to access the overridden methods or static properties by referencing them with parent:: -->

<?php
class SimpleClass
{
    public $var = 'a default value';
    public function displayVar()
    {
        echo $this->var;
    }
}

// Simple Class Inheritance
class ExtendClass extends SimpleClass
{
    // Redefine the parent method
    function displayVar()
    {
        echo "Extending class\n";
        parent::displayVar();
    }
}

$extended = new ExtendClass();
$extended->displayVar();
