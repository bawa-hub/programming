<!-- Incorrect usage of default function arguments -->
<?php
function makeyogurt($type = "acidophilus", $flavour)
{
    return "Making a bowl of $type $flavour.\n";
}

echo makeyogurt("raspberry");   // won't work as expected

// Correct usage of default function arguments
function makeyogurt1($flavour, $type = "acidophilus")
{
    return "Making a bowl of $type $flavour.\n";
}

echo makeyogurt1("raspberry");   // works as expected