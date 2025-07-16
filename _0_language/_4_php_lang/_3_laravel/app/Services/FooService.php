<?php

namespace App\Services;

class FooService
{
    protected $barService;

    public function __construct(BarService $barService)
    {
        $this->barService = $barService;
    }

    public function foo()
    {
        return $this->barService->bar();
    }
}
