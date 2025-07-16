<?php

use App\Services\BarService;
use App\Services\FooService;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Cache;
use Illuminate\Support\Facades\Route;




Route::get('/', function (FooService $service) {
    // this will automaticall resolved this injected services and add to $resolved array in container.
    // dd(app());
    return "Hello world!";
});

Route::get('bar', function (BarService $service) {
    dd(app());
});

Route::get('app', 'AppController@index');
