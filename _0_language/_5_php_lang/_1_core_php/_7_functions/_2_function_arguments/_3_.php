<?php

// Use of default parameters in functions
function makecoffee($type = "cappuccino")
{
    return "Making a cup of $type.\n";
}
echo makecoffee();
echo makecoffee(null);
echo makecoffee("espresso");
// Making a cup of cappuccino.
// Making a cup of .
// Making a cup of espresso.