<?php

/**
 * Within class methods non-static properties may be accessed by using -> (Object Operator): $this->property (where property is the name of the property).
 * Static properties are accessed by using the :: (Double Colon): self::$property
 * As of PHP 7.4.0, property definitions can include a Type declarations, with the exception of callable. 
 */

class User
{
    public int $id;
    public ?string $name;

    public function __construct(int $id, ?string $name)
    {
        $this->id = $id;
        $this->name = $name;
    }
}

$user = new User(1234, null);

var_dump($user->id);
var_dump($user->name);
