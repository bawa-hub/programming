package main

import "fmt"

type Notifier interface { Send(msg string) }

type EmailNotifier struct{}
func (EmailNotifier) Send(msg string) { fmt.Println("Email:", msg) }

// Decorator

type SMSDecorator struct{ wrap Notifier }
func (d SMSDecorator) Send(msg string) {
	d.wrap.Send(msg)
	fmt.Println("SMS:", msg)
}

func main() {
	var n Notifier = EmailNotifier{}
	n.Send("hello")

	n = SMSDecorator{wrap: n}
	n.Send("hello again")
}
