<?php

/**
 * Within class methods non-static properties may be accessed by using -> (Object Operator): $this->property (where property is the name of the property).
 * Static properties are accessed by using the :: (Double Colon): self::$property
 * As of PHP 7.4.0, property definitions can include a Type declarations, with the exception of callable. 
 */

//  Property declarations
class SimpleClass
{
    public $var1 = 'hello ' . 'world';
    public $var2 = <<<EOD
hello world
EOD;
    public $var3 = 1 + 2;
    // invalid property declarations:
    public $var4 = self::myStaticMethod();
    public $var5 = $myVar;

    // valid property declarations:
    public $var6 = myConstant;
    public $var7 = [true, false];

    public $var8 = <<<'EOD'
hello world
EOD;

    // Without visibility modifier:
    static $var9;
    readonly int $var10;
}

// Type declarations
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
// int(1234)
// NULL

// Typed properties must be initialized before accessing, otherwise an Error is thrown. 

class Shape
{
    public int $numberOfSides;
    public string $name;

    public function setNumberOfSides(int $numberOfSides): void
    {
        $this->numberOfSides = $numberOfSides;
    }

    public function setName(string $name): void
    {
        $this->name = $name;
    }

    public function getNumberOfSides(): int
    {
        return $this->numberOfSides;
    }

    public function getName(): string
    {
        return $this->name;
    }
}

$triangle = new Shape();
$triangle->setName("triangle");
$triangle->setNumberofSides(3);
var_dump($triangle->getName());
var_dump($triangle->getNumberOfSides());

$circle = new Shape();
$circle->setName("circle");
var_dump($circle->getName());
var_dump($circle->getNumberOfSides());
// string(8) "triangle"
// int(3)
// string(6) "circle"
// Fatal error: Uncaught Error: Typed property Shape::$numberOfSides must not be accessed before initialization
