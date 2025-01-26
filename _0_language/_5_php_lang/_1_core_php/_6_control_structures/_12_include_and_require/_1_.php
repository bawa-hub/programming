<?php

// The include and require statements are identical, except upon failure:
// require will produce a fatal error (E_COMPILE_ERROR) and stop the script
// include will only produce a warning (E_WARNING) and the script will continue

// there is one big difference between include and require; 
// when a file is included with the include statement and PHP cannot find it, 
// the script will continue to execute:

// Use require when the file is required by the application.
// Use include when the file is not required 
// and application should continue when file is not found.

// include_once/require_once will inlcude only one time

// difference b/w include and include_once
// http://www.readmyviews.com/include-vs-include-once/


// echo "A $color $fruit"; // A

include 'vars.php';

echo "A $color $fruit"; // A green apple