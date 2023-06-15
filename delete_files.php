<?php

// @param  string  Target directory
// @param  string  Target file extension
// @return boolean True on success, False on failure
function unlink_recursive($dir_name, $ext)
{
    // Exit if there's no such directory
    if (!file_exists($dir_name)) {
        return false;
    }

    // Open the target directory
    $dir_handle = dir($dir_name);

    // Take entries in the directory one at a time
    while (false !== ($entry = $dir_handle->read())) {
        if ($entry == '.' || $entry == '..') {
            continue;
        }
        $abs_name = "$dir_name/$entry";

        if (is_file($abs_name) && preg_match("/^.+\.$ext$/", $entry)) {
            if (unlink($abs_name)) {
                continue;
            }
            return false;
        }

        // Recurse on the children if the current entry happens to be a "directory"
        if (is_dir($abs_name) || is_link($abs_name)) {
            unlink_recursive($abs_name, $ext);
        }
    }

    $dir_handle->close();
    return true;
}
