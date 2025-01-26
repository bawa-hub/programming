<?php

interface Person
{
    public function giveIntro();
}

class Son implements Person
{
    public function giveIntro()
    {
        echo "I m son";
    }
}

class Husband implements Person
{
    public function giveIntro()
    {
        echo "I m husband";
    }
}

class Father implements Person
{
    public function giveIntro()
    {
        echo "I m father";
    }
}

$son = new Son();
$hus = new Husband();
$fath = new Father();

$persons = [$son, $hus, $fath];

foreach ($persons as $person) {
    $person->giveIntro();
    echo "<br>";
}
