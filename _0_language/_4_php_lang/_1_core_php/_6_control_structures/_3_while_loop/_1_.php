<?php

/* example 1 */

$i = 1;
while ($i <= 10) {
    echo $i++ . "<br>";  /* the printed value would be
                   $i before the increment
                   (post-increment) */
}

echo "<br>";
/* example 2 */

$i = 1;
while ($i <= 10) :
    echo $i . "<br>";
    $i++;
endwhile;
