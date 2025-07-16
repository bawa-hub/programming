<?php

/**
 *  goto operator can be used to jump to another section in the program
 * target point is specified by a case-sensitive label followed by a colon
 * target label must be within the same file and context,
 * meaning that you cannot jump out of a function or method, 
 * nor can you jump into one
 * also cannot jump into any sort of loop or switch structure
 * may jump out of these, and a common use is to use a goto in place of a multi-level break
 */

goto a;
echo 'Foo';

a:
echo 'Bar';
