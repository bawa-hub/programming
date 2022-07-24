<?php

//  allow access to only one instance of a particular class (during runtime)

final class Product
{

    private static $instance;
    public $mix;

    public static function getInstance()
    {
        if (!(self::$instance instanceof self)) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    private function __construct()
    {
    }

    private function __clone()
    {
    }
}
$firstProduct = Product::getInstance();
$secondProduct = Product::getInstance();

$firstProduct->mix = 'test';
$secondProduct->mix = 'example';

var_dump($firstProduct);
var_dump($secondProduct);
print_r($firstProduct->mix);
echo PHP_EOL;
print_r($secondProduct->mix);
