<?php

/***
 * 
 * The class keyword is also used for class name resolution. 
 * To obtain the fully qualified name of a class ClassName use ClassName::class. 
 * This is particularly useful with namespaced classes. 
 */

namespace NS {
    class ClassName
    {
    }

    echo ClassName::class;
}

/***
 * class name resolution using ::class is a compile time transformation. 
 * That means at the time the class name string is created no autoloading has happened yet. 
 * As a consequence, class names are expanded even if the class does not exist. 
 * No error is issued in that case. 
 */

//  Missing class name resolution
print Does\Not\Exist::class;
