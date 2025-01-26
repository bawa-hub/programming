<?php

// Traits can, as of PHP 8.2.0, also define constants.

trait ConstantsTrait
{
    public const FLAG_MUTABLE = 1;
    final public const FLAG_IMMUTABLE = 5;
}

class ConstantsExample
{
    use ConstantsTrait;
}

$example = new ConstantsExample;
echo $example::FLAG_MUTABLE; // 1

// If a trait defines a constant then a class can not define a constant with the same name unless it is compatible (same visibility, initial value, and finality), otherwise a fatal error is issued.

// Conflict Resolution
trait ConstantsTrait
{
    public const FLAG_MUTABLE = 1;
    final public const FLAG_IMMUTABLE = 5;
}

class ConstantsExample
{
    use ConstantsTrait;
    public const FLAG_IMMUTABLE = 5; // Fatal error
}
