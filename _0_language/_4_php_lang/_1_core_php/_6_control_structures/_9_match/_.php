<?php

// (PHP 8)
// The match expression branches evaluation based on an identity check of a value. 
// Similarly to a switch statement, a match expression has a subject expression that is compared against multiple alternatives. 
// Unlike switch, it will evaluate to a value much like ternary expressions. 
// Unlike switch, the comparison is an identity check (===) rather than a weak equality check (==). 
// Match expressions are available as of PHP 8.0.0. 

// Structure of a match expression
// $return_value = match (subject_expression) {
//     single_conditional_expression => return_expression,
//     conditional_expression1, conditional_expression2 => return_expression,
// };

// Basic match usage
$food = 'cake';
$return_value = match ($food) {
    'apple' => 'This food is an apple',
    'bar' => 'This food is a bar',
    'cake' => 'This food is a cake',
};
var_dump($return_value);

// The match expression is similar to a switch statement but has some key differences:

//     A match arm compares values strictly (===) instead of loosely as the switch statement does.
//     A match expression returns a value.
//     match arms do not fall-through to later cases the way switch statements do.
//     A match expression must be exhaustive.

// $result = match ($x) {
//     foo() => ...,
//     $this->bar() => ..., // $this->bar() isn't called if foo() === $x
//     $this->baz => beep(), // beep() isn't called unless $x === $this->baz
//     // etc.
// };

// match expression arms may contain multiple expressions separated by a comma. That is a logical OR, and is a short-hand for multiple match arms with the same right-hand side.
$result = match ($x) {
    // This match arm:
    $a, $b, $c => 5,
    // Is equivalent to these three match arms:
    $a => 5,
    $b => 5,
    $c => 5,
};

// A special case is the default pattern. This pattern matches anything that wasn't previously matched
$expressionResult = match ($condition) {
    1, 2 => foo(),
    3, 4 => bar(),
    default => baz(),
};

// Example of an unhandled match expression
$condition = 5;

try {
    match ($condition) {
        1, 2 => foo(),
        3, 4 => bar(),
    };
} catch (\UnhandledMatchError $e) {
    var_dump($e);
}



// Using match expressions to handle non identity checks
// It is possible to use a match expression to handle non-identity conditional cases by using true as the subject expression

// Using a generalized match expressions to branch on integer ranges
$age = 23;
$result = match (true) {
    $age >= 65 => 'senior',
    $age >= 25 => 'adult',
    $age >= 18 => 'young adult',
    default => 'kid',
};
var_dump($result);

// Using a generalized match expressions to branch on string content
$text = 'Bienvenue chez nous';
$result = match (true) {
    str_contains($text, 'Welcome') || str_contains($text, 'Hello') => 'en',
    str_contains($text, 'Bienvenue') || str_contains($text, 'Bonjour') => 'fr',
    // ...
};
var_dump($result);
