package main

import "fmt"

type Printer interface{ Print(doc string) }

type Scanner interface{ Scan(doc string) }

type Faxer interface{ Fax(doc string) }

type SimplePrinter struct{}
func (SimplePrinter) Print(doc string) { fmt.Println("Printing:", doc) }

type Workstation struct{}
func (Workstation) Print(doc string) { fmt.Println("Printing:", doc) }
func (Workstation) Scan(doc string)  { fmt.Println("Scanning:", doc) }

func main() {
	var p Printer = SimplePrinter{}
	p.Print("invoice")

	var ws interface{ Printer; Scanner } = Workstation{}
	ws.Print("pdf")
	ws.Scan("pdf")
}
