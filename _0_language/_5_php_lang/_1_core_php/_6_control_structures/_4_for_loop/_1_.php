<?php
/* example 1 */

for ($i = 1; $i <= 10; $i++) {
    echo $i . "<br>";
}
echo "<br>";

/* example 2 */

for ($i = 1;; $i++) {
    if ($i > 10) {
        break;
    }
    echo $i;
}
echo "<br>";

/* example 3 */

$i = 1;
for (;;) {
    if ($i > 10) {
        break;
    }
    echo $i;
    $i++;
}
echo "<br>";

/* example 4 */

for ($i = 1, $j = 0; $i <= 10; $j += $i, print $i, $i++);
