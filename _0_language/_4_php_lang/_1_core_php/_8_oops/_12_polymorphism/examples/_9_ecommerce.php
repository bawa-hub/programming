<?php

// Let's consider a simplified example of an e-commerce application that deals with various types of products. Products can include physical goods, digital goods, and services, each with its own specific behavior:

// Product Interface
interface Product
{
    public function getName();
    public function getPrice();
    public function getDescription();
}

// PhysicalProduct class implementing the interface
class PhysicalProduct implements Product
{
    private $name;
    private $price;
    private $weight;

    public function __construct($name, $price, $weight)
    {
        $this->name = $name;
        $this->price = $price;
        $this->weight = $weight;
    }

    public function getName()
    {
        return $this->name;
    }

    public function getPrice()
    {
        return $this->price;
    }

    public function getDescription()
    {
        return "Physical product weighing {$this->weight} kg";
    }
}

// DigitalProduct class implementing the interface
class DigitalProduct implements Product
{
    private $name;
    private $price;
    private $downloadLink;

    public function __construct($name, $price, $downloadLink)
    {
        $this->name = $name;
        $this->price = $price;
        $this->downloadLink = $downloadLink;
    }

    public function getName()
    {
        return $this->name;
    }

    public function getPrice()
    {
        return $this->price;
    }

    public function getDescription()
    {
        return "Digital product with download link: {$this->downloadLink}";
    }
}

// Service class implementing the interface
class Service implements Product
{
    private $name;
    private $price;
    private $duration;

    public function __construct($name, $price, $duration)
    {
        $this->name = $name;
        $this->price = $price;
        $this->duration = $duration;
    }

    public function getName()
    {
        return $this->name;
    }

    public function getPrice()
    {
        return $this->price;
    }

    public function getDescription()
    {
        return "Service with duration: {$this->duration} hours";
    }
}

// ShoppingCart class using polymorphism for managing products
class ShoppingCart
{
    private $products = [];

    public function addProduct(Product $product)
    {
        $this->products[] = $product;
    }

    public function displayProducts()
    {
        foreach ($this->products as $product) {
            echo "Product: {$product->getName()}, Price: {$product->getPrice()}, Description: {$product->getDescription()}\n";
        }
    }
}

// Client code using polymorphism to add and display different types of products
$cart = new ShoppingCart();

// Add products to the shopping cart
$physicalProduct = new PhysicalProduct("Laptop", 1200, 2.5);
$digitalProduct = new DigitalProduct("Software", 49.99, "https://example.com/download");
$service = new Service("Consulting", 100, 2);

$cart->addProduct($physicalProduct);
$cart->addProduct($digitalProduct);
$cart->addProduct($service);

// Display the products in the shopping cart
$cart->displayProducts();


// The Product interface defines common methods (getName, getPrice, getDescription) that all product types must implement.
// Concrete classes (PhysicalProduct, DigitalProduct, Service) implement this interface with their specific behavior for each method.
// The ShoppingCart class uses polymorphism by allowing the dynamic addition of different types of products through the addProduct method and displaying them through the displayProducts method.

// This example demonstrates how polymorphism allows you to treat different types of products uniformly, making it easy to extend and maintain the shopping cart system as new product types are added. The ShoppingCart class doesn't need to know the specific type of each product; it can manage and display them using a common interface