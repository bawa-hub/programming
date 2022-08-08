<?php

// conflict resolution

/**
 * If two Traits insert a method with the same name, a fatal error is produced, 
 * if the conflict is not explicitly resolved. 
 */

trait A
{
    public function smallTalk()
    {
        echo 'a';
    }
    public function bigTalk()
    {
        echo 'A';
    }
}

trait B
{
    public function smallTalk()
    {
        echo 'b';
    }
    public function bigTalk()
    {
        echo 'B';
    }
}

class Talker
{
    use A, B {
        B::smallTalk insteadof A;
        A::bigTalk insteadof B;
    }
}

class Aliased_Talker
{
    use A, B {
        B::smallTalk insteadof A;
        A::bigTalk insteadof B;
        B::bigTalk as talk;
    }
}
