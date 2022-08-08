<?php

// If a trait defines a property 
// then a class can not define a property with the same name 
// unless it is compatible (same visibility and initial value), 
// otherwise a fatal error is issued. 

trait PropertiesTrait
{
    public $same = true;
    public $different = false;
}

class PropertiesExample
{
    use PropertiesTrait;
    public $same = true;
    public $different = true; // Fatal error
}
