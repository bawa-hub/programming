<?php

/**
 * special ::class constant allows for fully qualified class name resolution at compile time, 
 * this is useful for namespaced classes
 */

namespace foo {
    class bar
    {
    }

    echo bar::class; // foo\bar
}
