<?php

// Let's consider a scenario in which polymorphism is applied in a system that processes different types of notifications. Notifications can include emails, SMS messages, and push notifications. Here's a simplified example:

// Notification Interface
interface Notification
{
    public function send();
}

// EmailNotification class implementing the interface
class EmailNotification implements Notification
{
    private $to;
    private $subject;
    private $message;

    public function __construct($to, $subject, $message)
    {
        $this->to = $to;
        $this->subject = $subject;
        $this->message = $message;
    }

    public function send()
    {
        // Sending logic for email notification
        echo "Email sent to {$this->to} with subject '{$this->subject}'.\n";
    }
}

// SMSNotification class implementing the interface
class SMSNotification implements Notification
{
    private $phoneNumber;
    private $message;

    public function __construct($phoneNumber, $message)
    {
        $this->phoneNumber = $phoneNumber;
        $this->message = $message;
    }

    public function send()
    {
        // Sending logic for SMS notification
        echo "SMS sent to {$this->phoneNumber} with message '{$this->message}'.\n";
    }
}

// PushNotification class implementing the interface
class PushNotification implements Notification
{
    private $deviceToken;
    private $message;

    public function __construct($deviceToken, $message)
    {
        $this->deviceToken = $deviceToken;
        $this->message = $message;
    }

    public function send()
    {
        // Sending logic for push notification
        echo "Push notification sent to device {$this->deviceToken} with message '{$this->message}'.\n";
    }
}

// NotificationService class using polymorphism for sending notifications
class NotificationService
{
    public function sendNotification(Notification $notification)
    {
        $notification->send();
    }
}

// Client code using polymorphism to send different types of notifications
$notificationService = new NotificationService();

// Send email notification
$emailNotification = new EmailNotification("user@example.com", "Hello", "This is an email notification.");
$notificationService->sendNotification($emailNotification);

// Send SMS notification
$smsNotification = new SMSNotification("+1234567890", "This is an SMS notification.");
$notificationService->sendNotification($smsNotification);

// Send push notification
$pushNotification = new PushNotification("device_token", "This is a push notification.");
$notificationService->sendNotification($pushNotification);

// The Notification interface defines a common method send that all notification types must implement.
// Concrete classes (EmailNotification, SMSNotification, PushNotification) implement this interface with their specific sending logic.
// The NotificationService class uses polymorphism by allowing the dynamic sending of different types of notifications through the sendNotification method.

// This example demonstrates how polymorphism allows you to treat different types of notifications uniformly, making it easy to extend and maintain the notification system as new notification types are added. The NotificationService class doesn't need to know the specific type of each notification; it can send them using a common interface.