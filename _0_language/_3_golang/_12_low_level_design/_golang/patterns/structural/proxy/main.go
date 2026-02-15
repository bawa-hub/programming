package main

import "fmt"

type Image interface { Display() }

type RealImage struct{ filename string }

func (r RealImage) loadFromDisk() { fmt.Println("Loading", r.filename) }
func (r RealImage) Display()      { fmt.Println("Displaying", r.filename) }

type ProxyImage struct{ real *RealImage }

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{real: &RealImage{filename: filename}}
}

func (p *ProxyImage) Display() {
	p.real.loadFromDisk() // naive proxy (no caching), for demo
	p.real.Display()
}

func main() { NewProxyImage("photo.png").Display() }
