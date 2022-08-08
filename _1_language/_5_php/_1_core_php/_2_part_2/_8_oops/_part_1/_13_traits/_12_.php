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
