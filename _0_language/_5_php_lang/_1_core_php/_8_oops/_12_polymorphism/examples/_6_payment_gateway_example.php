<?php

// PaymentGateway Interface
interface PaymentGateway
{
    public function processPayment($amount);
}

// PayPal PaymentGateway class implementing the interface
class PayPalPaymentGateway implements PaymentGateway
{
    public function processPayment($amount)
    {
        // PayPal-specific payment processing logic
        echo "Processed payment of $$amount via PayPal.\n";
    }
}

// Stripe PaymentGateway class implementing the interface
class StripePaymentGateway implements PaymentGateway
{
    public function processPayment($amount)
    {
        // Stripe-specific payment processing logic
        echo "Processed payment of $$amount via Stripe.\n";
    }
}

// CreditCard PaymentGateway class implementing the interface
class CreditCardPaymentGateway implements PaymentGateway
{
    public function processPayment($amount)
    {
        // Credit card-specific payment processing logic
        echo "Processed payment of $$amount via Credit Card.\n";
    }
}

// Order class using polymorphism for payment processing
class Order
{
    private $paymentGateway;

    public function setPaymentGateway(PaymentGateway $gateway)
    {
        $this->paymentGateway = $gateway;
    }

    public function processPayment($amount)
    {
        $this->paymentGateway->processPayment($amount);
    }
}

// Client code using polymorphism to switch between payment gateways
$order = new Order();

// PayPal payment
$paypalGateway = new PayPalPaymentGateway();
$order->setPaymentGateway($paypalGateway);
$order->processPayment(100);

// Stripe payment
$stripeGateway = new StripePaymentGateway();
$order->setPaymentGateway($stripeGateway);
$order->processPayment(150);

// Credit card payment
$creditCardGateway = new CreditCardPaymentGateway();
$order->setPaymentGateway($creditCardGateway);
$order->processPayment(200);


// The PaymentGateway interface defines a common method processPayment that all payment gateways must implement.
// Concrete classes (PayPalPaymentGateway, StripePaymentGateway, CreditCardPaymentGateway) implement this interface with their specific payment processing logic.
// The Order class uses polymorphism by allowing the dynamic switching of payment gateways at runtime through the setPaymentGateway method