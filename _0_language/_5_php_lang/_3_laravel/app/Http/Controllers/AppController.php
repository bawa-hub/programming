<?php

namespace App\Http\Controllers;

use App\Services\FooService;
use Illuminate\Http\Request;

class AppController extends Controller
{
    public function index(FooService $fooService)
    {
        echo $fooService->foo();
    }
}
