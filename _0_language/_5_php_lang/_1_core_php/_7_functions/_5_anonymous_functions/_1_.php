<?php 

// Anonymous functions are implemented using the Closure class. 

echo preg_replace_callback('~-([a-z])~', function ($match) {
    return strtoupper($match[1]);
}, 'hello-world');
// outputs helloWorld

// Closures can also be used as the values of variables; PHP automatically converts such expressions into instances of the Closure internal class


//  Anonymous function variable assignment example
$greet = function($name)
{
    printf("Hello %s\r\n", $name);
};

$greet('World');
$greet('PHP');

// https://www.php.net/manual/en/functions.anonymous.php