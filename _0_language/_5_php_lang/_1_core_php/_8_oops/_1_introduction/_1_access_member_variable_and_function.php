<!-- PHP 5 is very very flexible in accessing member variables and member functions -->

<?php

class Foo
{
    public $aMemberVar = 'aMemberVar Member Variable';
    public $aFuncName = 'aMemberFunc';

    function aMemberFunc()
    {
        print 'Inside `aMemberFunc()`';
    }
}

$foo = new Foo;
// You can access member variables in an object using another variable as name: 
$element = 'aMemberVar';
print $foo->$element;
// or use functions:
function getVarName()
{
    return 'aMemberVar';
}

print $foo->{getVarName()};

// Important Note: You must surround function name with { and } or PHP would think you are calling a member function of object "foo". 

// you can use a constant or literal as well: 
define('MY_CONSTANT', 'aMemberVar');
print $foo->{MY_CONSTANT}; // Prints "aMemberVar Member Variable"
print $foo->{'aMemberVar'}; // Prints "aMemberVar Member Variable" 

// You can use members of other objects as well: 
print $foo->{$otherObj->var};
print $foo->{$otherObj->func()};

// You can use mathods above to access member functions as well: 
print $foo->{'aMemberFunc'}(); // Prints "Inside `aMemberFunc()`"
print $foo->{$foo->aFuncName}(); // Prints "Inside `aMemberFunc()`" 
