<?php

function foo()
{
    global $color;

    include 'vars.php';

    echo "A $color $fruit";
}

foo();
echo "A $color $fruit";

/* vars.php is in the scope of foo() so     *
* $fruit is NOT available outside of this  *
* scope.  $color is because we declared it *
* as global.                               */
