<!-- It's possible to create instances of an object in a couple of ways:  -->

<?php
class Test
{
    static public function getNew()
    {
        return new static;
    }
}

class Child extends Test
{
}

$obj1 = new Test();
$obj2 = new $obj1;
var_dump($obj1 !== $obj2);

$obj3 = Test::getNew();
var_dump($obj3 instanceof Test);

$obj4 = Child::getNew();
var_dump($obj4 instanceof Child);
?>

<!-- 
    It is possible to access a member of a newly created object in a single expression: 
 -->

<?php
echo (new DateTime())->format('Y');
?>