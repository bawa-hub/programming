<?php

class Base
{
    public function foo(int $a = 5)
    {
        echo "Valid\n";
    }
}

// Fatal error when a child method removes a parameter
class Extend extends Base
{
    function foo()
    {
        parent::foo(1);
    }
}

// Fatal error when a child method makes an optional parameter mandatory

class Extend1 extends Base
{
    function foo(int $a)
    {
        parent::foo($a);
    }
}
