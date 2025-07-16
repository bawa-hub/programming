<?php

// Using the as syntax, one can also adjust the visibility of the method in the exhibiting class. 

trait HelloWorld
{
    public function sayHello()
    {
        echo 'Hello World!';
    }
}

// Change visibility of sayHello
class MyClass1
{
    use HelloWorld {
        sayHello as protected;
    }
}

// Alias method with changed visibility
// sayHello visibility not changed
class MyClass2
{
    use HelloWorld {
        sayHello as private myPrivateHello;
    }
}
