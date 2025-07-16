<?php

// function foo($arg_1, $arg_2, /* ..., */ $arg_n)
// {
//     echo "Example function.\n";
//     return $retval;
// }

// Functions need not be defined before they are referenced, 
// except when a function is conditionally defined. 

$makefoo = true;

/* We can't call foo() from here 
   since it doesn't exist yet,
   but we can call bar() */

bar();

if ($makefoo) {
  function foo()
  {
    echo "I don't exist until program execution reaches me.\n";
  }
}

/* Now we can safely call foo()
   since $makefoo evaluated to true */

if ($makefoo) foo();

function bar()
{
  echo "I exist immediately upon program start.\n";
}
