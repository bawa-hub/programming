package main

import "fmt"

type MultiFunctionPrinter interface {
	Print(doc string)
	Scan(doc string)
	Fax(doc string)
}

type OldPrinter struct{}

func (OldPrinter) Print(doc string) { fmt.Println("Printing:", doc) }
func (OldPrinter) Scan(doc string)  { fmt.Println("Scanning:", doc) }
func (OldPrinter) Fax(doc string)   { fmt.Println("Faxing:", doc) }

// SimplePrinter must implement methods it does not need.
type SimplePrinter struct{}

func (SimplePrinter) Print(doc string) { fmt.Println("Printing:", doc) }
func (SimplePrinter) Scan(doc string)  { fmt.Println("(unused) scanning:", doc) }
func (SimplePrinter) Fax(doc string)   { fmt.Println("(unused) faxing:", doc) }

func main() {
	var p MultiFunctionPrinter = SimplePrinter{}
	p.Print("report")
}
