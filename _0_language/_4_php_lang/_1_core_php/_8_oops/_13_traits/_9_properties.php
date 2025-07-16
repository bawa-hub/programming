<?php

// Traits can also define properties. 
trait PropertiesTrait
{
    public $x = 1;
}

class PropertiesExample
{
    use PropertiesTrait;
}

$example = new PropertiesExample;
$example->x;

// If a trait defines a property then a class can not define a property with the same name unless it is compatible (same visibility and type, readonly modifier, and initial value), otherwise a fatal error is issued. 

// Conflict Resolution

// If a trait defines a property 
// then a class can not define a property with the same name 
// unless it is compatible (same visibility and initial value), 
// otherwise a fatal error is issued.

trait PropertiesTrait
{
    public $same = true;
    public $different1 = false;
    public bool $different2;
    public bool $different3;
}

class PropertiesExample
{
    use PropertiesTrait;
    public $same = true;
    public $different1 = true; // Fatal error
    public string $different2; // Fatal error
    readonly protected bool $different3; // Fatal error
}
