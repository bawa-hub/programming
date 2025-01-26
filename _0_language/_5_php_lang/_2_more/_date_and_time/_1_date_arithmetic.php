<?php

// Declare two dates
$start_date = strtotime("2021-12-29");
$end_date = strtotime(date('Y-m-d'));

// Get the difference and divide into
// total no. seconds 60/60/24 to get
// number of days
echo "Difference between two dates: "
    . ($end_date - $start_date) / 60 / 60 / 24;
