<?php

/**
 * When an object is cloned, PHP will perform a shallow copy of all of the object's properties. 
 * Any properties that are references to other variables will remain references
 */

class Invoice
{
    private string $id;

    public function __construct()
    {
        $this->id = uniqid('invoice_');
        var_dump('__constructor');
    }

    public function __clone()
    {
        $this->id = uniqid('invoice_');
        var_dump('__clone');
    }

    public static function create()
    {
        return new static();
    }
}

$invoice1 = new Invoice();

// create new object
// $invoice2 = new $invoice1(); 
// var_dump($invoice1, $invoice2, Invoice::create());

// points to same object
// $invoice3 = $invoice1;
// var_dump($invoice1, $invoice3, $invoice1 === $invoice3);

// clone 
//  __constructor will not called instead __clone will called
$invoice4 = clone $invoice1;
var_dump($invoice1, $invoice4, $invoice1 === $invoice4);
