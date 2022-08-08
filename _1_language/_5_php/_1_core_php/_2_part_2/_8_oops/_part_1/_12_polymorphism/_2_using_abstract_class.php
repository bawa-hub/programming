<?php

abstract class Person
{
    abstract public function greet();
}

class English extends Person
{
    public function greet()
    {
        return 'Hello!';
    }
}

class German extends Person
{
    public function greet()
    {
        return 'Hallo!';
    }
}

class French extends Person
{
    public function greet()
    {
        return 'Bonjour!';
    }
}

function greeting($people)
{
    foreach ($people as $person) {
        echo $person->greet() . '<br>';
    }
}

// Later, if you want to create a new class that extends the Person class, 
// you don’t need to change the greeting() function
class American extends Person
{
    public function greet()
    {
        return 'Hi!';
    }
}

$people = [
    new English(),
    new German(),
    new French(),
    new American()
];

greeting($people);
