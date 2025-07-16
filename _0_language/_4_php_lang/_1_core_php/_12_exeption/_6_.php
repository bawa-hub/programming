<?php

// Omitting the caught variable (from >= php 8.0.0)

class SpecificException extends Exception
{
}

function test()
{
    throw new SpecificException('Oopsie');
}

try {
    test();
} catch (SpecificException) {
    print "A SpecificException was thrown, but we don't care about the details.";
}
