
<?php

// Alternate Precedence Order Example

trait HelloWorld
{
    public function sayHello()
    {
        echo 'Hello World!';
    }
}

class TheWorldIsNotEnough
{
    use HelloWorld;
    public function sayHello()
    {
        echo 'Hello Universe!';
    }
}

$o = new TheWorldIsNotEnough();
$o->sayHello();
?>
